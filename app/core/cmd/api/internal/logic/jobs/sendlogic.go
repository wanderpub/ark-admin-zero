package jobs

import (
	"context"
	"fmt"
	"time"

	"ark-admin-zero/app/core/cmd/api/internal/svc"
	"ark-admin-zero/app/core/cmd/api/internal/types"
	"ark-admin-zero/app/core/model"

	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/jsonx"
	"github.com/zeromicro/go-zero/core/logx"
)

type SendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendLogic {
	return &SendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendLogic) Send(req *types.SendRequest) (resp *types.Response, err error) {
	var sendModel = &model.SendTaskModel{
		Code:              req.Code,
		MessageTemplateId: req.MessageTemplateId,
		Time:              req.Time,
		MessageParamList: model.MessageParam{
			Receiver:  req.MessageParam.Receiver,
			Variables: model.ContentModel(req.MessageParam.Variables),
			Extra:     req.MessageParam.Extra,
		},
	}
	json, _ := jsonx.MarshalToString(sendModel)
	fmt.Println(json)
	// Create a task with typename and payload.
	payload, err := jsonx.Marshal(sendModel)
	if err != nil {
		return &types.Response{
			Message: err.Error(),
		}, err
	}
	typeName := fmt.Sprintf("task:%s", req.Code)
	t := asynq.NewTask(typeName, payload)

	//1.获取当前时区
	location, _ := time.LoadLocation("Asia/Shanghai")
	var processTime = time.Now().In(location)
	if req.Time > 0 {
		if req.Time > time.Now().In(location).Unix() {
			processTime = time.Unix(req.Time, 0).In(location)
		}
	}
	fmt.Println(processTime.Unix())
	fmt.Println(processTime.Format("2006-01-02 15:04:05"))
	info, err := l.svcCtx.AsynqClient.Enqueue(t, asynq.ProcessAt(processTime))
	// res, _ := jsonx.MarshalToString(info)
	return &types.Response{
		Message: info.ID,
	}, err
}
