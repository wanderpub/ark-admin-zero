package handlers

import (
	"ark-admin-zero/app/core/cmd/jobs/internal/handlers/account"
	"ark-admin-zero/app/core/cmd/jobs/internal/svc"
	"ark-admin-zero/app/core/model"
	"context"
	"strings"

	"github.com/pkg/errors"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/message"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
)

const colorSep = "|" //以|分割颜色

type officialAccountHandler struct {
	svcCtx *svc.ServiceContext
}

func NewOfficialAccountHandler(svcCtx *svc.ServiceContext) IHandler {
	return officialAccountHandler{
		svcCtx: svcCtx,
	}
}

func (h officialAccountHandler) DoHandler(ctx context.Context, sendTask model.SendTaskModel) (err error) {
	var acc account.OfficialAccount
	err = jsonx.UnmarshalFromString(sendTask.TaskInfo.SendAccountConfig, &acc)
	if err != nil {
		return errors.Wrap(err, "officialAccountHandler get account err")
	}

	wc := wechat.NewWechat()
	cacheImpl := cache.NewRedis(ctx, &cache.RedisOpts{
		Host:     h.svcCtx.Config.Redis.Host,
		Password: h.svcCtx.Config.Redis.Pass,
	})
	cfg := &offConfig.Config{
		AppID:          acc.AppID,
		AppSecret:      acc.AppSecret,
		Token:          acc.Token,
		EncodingAESKey: acc.EncodingAESKey,
		Cache:          cacheImpl,
	}
	subscribe := wc.GetOfficialAccount(cfg).GetTemplate()
	templateId := sendTask.MessageParamList.Variables.TemplateId
	url := sendTask.MessageParamList.Variables.Url
	params := make(map[string]*message.TemplateDataItem, len(sendTask.MessageParamList.Variables.Map))

	for key, val := range sendTask.MessageParamList.Variables.Map {
		color := ""
		value := ""
		arr := strings.Split(val, colorSep)
		if len(arr) == 1 {
			value = arr[0]
		}
		if len(arr) == 2 {
			value = arr[0]
			color = arr[1]
		}
		params[key] = &message.TemplateDataItem{Value: value, Color: color}
	}
	var msgIds []int64
	//如果需要实现跳转小程序 需要在getRealWxMpTemplateId里面返回对应的数据进行操作
	for _, receiver := range sendTask.TaskInfo.Receiver {
		msgID, err := subscribe.Send(&message.TemplateMessage{
			ToUser:     receiver,
			TemplateID: templateId,
			URL:        url,
			Data:       params,
		})
		if err != nil {
			logx.Errorw("officialAccountHandler send msg",
				logx.Field("err", err),
				logx.Field("receiver", receiver),
				logx.Field("templateId", templateId))
			continue
		}
		msgIds = append(msgIds, msgID)
	}
	logx.Infow("officialAccountHandler send success", logx.Field("msgIds", msgIds))
	return nil
}
