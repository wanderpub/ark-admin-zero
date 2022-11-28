package handlers

import (
	"ark-admin-zero/app/core/cmd/jobs/internal/handlers/account"
	"ark-admin-zero/app/core/cmd/jobs/internal/svc"
	"ark-admin-zero/app/core/model"
	"ark-admin-zero/common/utils"
	"context"

	"github.com/pkg/errors"
	"github.com/wanghuiyt/ding"
	"github.com/zeromicro/go-zero/core/jsonx"
)

const SendAll = "@all"

type dingDingRobotHandler struct {
	svcCtx *svc.ServiceContext
}

func NewDingDingRobotHandler(svcCtx *svc.ServiceContext) IHandler {
	return dingDingRobotHandler{
		svcCtx: svcCtx,
	}
}

func (h dingDingRobotHandler) DoHandler(ctx context.Context, sendTask model.SendTaskModel) (err error) {
	var acc account.DingDingRobotAccount
	err = jsonx.UnmarshalFromString(sendTask.TaskInfo.SendAccountConfig, &acc)
	if err != nil {
		return errors.Wrap(err, "dingDingRobotHandler get account err")
	}
	var at []string
	d := ding.Webhook{
		AccessToken: acc.AccessToken,
		Secret:      acc.Secret,
		EnableAt:    true,
	}

	if utils.ArrayStringIn(sendTask.TaskInfo.Receiver, SendAll) {
		d.AtAll = true
	} else {
		at = sendTask.TaskInfo.Receiver
	}

	err = d.SendMessage(sendTask.MessageParamList.Variables.Content, at...)
	if err != nil {
		return errors.Wrap(err, "dingDingRobotHandler SendMessage err")
	}
	return nil
}
