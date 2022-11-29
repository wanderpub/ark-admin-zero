package listen

import (
	"ark-admin-zero/app/core/cmd/jobs/internal/handlers"
	"ark-admin-zero/app/core/cmd/jobs/internal/svc"
	"ark-admin-zero/app/core/model"
	"ark-admin-zero/common/utils"
	"context"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/hibiken/asynq"
	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/limit"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/time/rate"
)

// A list of task types.
const (
	//邮件
	TypeSend string = "task:send"
	//邮件
	TypeEmail string = "task:email"
	//短信
	TypeSms string = "task:sms"
	//阿里短信
	TypeAliSms string = "task:alisms"
	//腾讯短信
	TypeTencentSms string = "task:tencentsms"
	//公众号模版消息
	TypeOffiaccount string = "task:offiaccount"
	//小程序消息
	TypeMiniprogram string = "task:miniprogram"
	//企业微信
	TypeEnterpriseWechat string = "task:enterprisewechat"
	//钉钉群机器人
	TypeDingRobot string = "task:dingrobot"
	//钉钉工作通知
	TypeDingWork string = "task:dingwork"
	//腾讯推送
	TypeTencentPush string = "task:tencentpush"
)

// Rate is 10 events/sec and permits burst of at most 30 events.
var limiter = rate.NewLimiter(100, 300)

// 消费消息
func (l *AsynqTask) handleSendProcess(ctx context.Context, t *asynq.Task) error {
	if !limiter.Allow() { //限速
		return &RateLimitError{
			RetryIn: time.Duration(rand.Intn(10)) * time.Second,
		}
	}
	json, _ := jsonx.MarshalToString(t.Payload())
	fmt.Println(json)
	var p model.SendTaskModel
	if err := jsonx.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}
	fmt.Println("接收人:", p.MessageParamList.Receiver)
	messageTemplate, err := l.svcCtx.TaskTemplateModel.FindOne(ctx, p.MessageTemplateId)
	if err != nil {
		return errorx.Wrapf(err, "查询模板异常 err:%v 模板id:%d", err, p.MessageTemplateId)
	}
	sendAccount, err := l.svcCtx.TaskAccountModel.FindOne(ctx, messageTemplate.SendAccount)
	if err != nil {
		return errorx.Wrapf(err, "查询模板消息帐号异常 err:%v 帐号id:%d", err, messageTemplate.SendAccount)
	}
	p.TaskInfo = model.TaskInfo{
		MessageTemplateId: messageTemplate.Id,
		IdType:            messageTemplate.IdType,
		SendChannel:       messageTemplate.SendChannel,
		TemplateType:      messageTemplate.TemplateType,
		MsgType:           messageTemplate.MsgType,
		ShieldType:        messageTemplate.ShieldType,
		SendAccount:       messageTemplate.SendAccount,
		Config:            messageTemplate.Config,
		BusinessId:        utils.GenerateBusinessId(messageTemplate.Id, messageTemplate.TemplateType),
		Receiver:          utils.ArrayStringUniq(strings.Split(p.MessageParamList.Receiver, ",")),
		SendAccountConfig: sendAccount.Config,
	}
	// logx.Infof("消息内容 %v, now is %v\n", p, time.Now())
	if messageTemplate.DeduplicationConfig != "" {
		var limit model.PeriodLimit
		err = jsonx.UnmarshalFromString(messageTemplate.DeduplicationConfig, &limit)
		if err != nil {
			logx.Errorw("PeriodLimit json err", logx.Field("info", p), logx.Field("err", err))
		} else {
			//滑动窗口去重
			receiver, _ := periodLimit(*l.svcCtx, &p, &limit)
			p.TaskInfo.Receiver = receiver
		}
	}
	//发送消息
	if len(p.TaskInfo.Receiver) > 0 {
		h := handlers.GetHandler(p.TaskInfo.SendChannel)
		logx.Infof("消息内容 %v, now is %v\n", p, time.Now())
		err := h.DoHandler(ctx, p)
		if err != nil {
			logx.Errorw("DoHandler err", logx.Field("info", p), logx.Field("err", err))
		}
	} else {
		fmt.Println("没有要发送的对象，正常返回。")
	}
	return nil
}

// 滑动窗口过滤去重
func periodLimit(svcCtx svc.ServiceContext, taskInfo *model.SendTaskModel, plimit *model.PeriodLimit) (returnReceiver []string, err error) {
	returnReceiver = make([]string, 0)
	key := fmt.Sprintf("periodlimit_%d", taskInfo.MessageTemplateId)
	l := limit.NewPeriodLimit(plimit.Seconds, plimit.Quota, svcCtx.RedisClient, key)
	for _, receiver := range taskInfo.TaskInfo.Receiver {
		key := periodLimitKey(taskInfo, receiver)
		code, err := l.Take(key)
		if err != nil {
			logx.Errorw("slideWindowLimit Take ", logx.Field("err", err))
			continue
		}
		//表示到了上限 直接过滤掉
		if code != limit.OverQuota {
			fmt.Println("没有达到上限")
			returnReceiver = append(returnReceiver, receiver)
		}
	}
	return returnReceiver, nil
}

// 生成KEY，用于滑动窗口去重
func periodLimitKey(taskInfo *model.SendTaskModel, receiver string) string {
	str, _ := jsonx.Marshal(taskInfo.MessageParamList.Variables)
	return utils.MD5(cast.ToString(taskInfo.MessageTemplateId) + receiver + string(str))
}

// 以下为限速使用
type RateLimitError struct {
	RetryIn time.Duration
}

func (e *RateLimitError) Error() string {
	return fmt.Sprintf("rate limited (retry in  %v)", e.RetryIn)
}

func IsRateLimitError(err error) bool {
	_, ok := err.(*RateLimitError)
	return ok
}

func retryDelay(n int, err error, task *asynq.Task) time.Duration {
	var ratelimitErr *RateLimitError
	if errors.As(err, &ratelimitErr) {
		return ratelimitErr.RetryIn
	}
	return asynq.DefaultRetryDelayFunc(n, err, task)
}
