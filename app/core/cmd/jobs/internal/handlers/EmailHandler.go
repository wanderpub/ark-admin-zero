package handlers

import (
	"ark-admin-zero/app/core/cmd/jobs/internal/handlers/account"
	"ark-admin-zero/app/core/cmd/jobs/internal/svc"
	"ark-admin-zero/app/core/model"
	"context"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/jsonx"
	"gopkg.in/gomail.v2"
)

type emailHandler struct {
	svcCtx *svc.ServiceContext
}

func NewEmailHandler(svcCtx *svc.ServiceContext) IHandler {
	return emailHandler{
		svcCtx: svcCtx,
	}
}

func (h emailHandler) DoHandler(ctx context.Context, sendTask model.SendTaskModel) (err error) {
	var acc account.EmailAccount
	err = jsonx.UnmarshalFromString(sendTask.TaskInfo.SendAccountConfig, &acc)
	if err != nil {
		return errors.Wrap(err, "emailHandler get account err")
	}
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(acc.Username, "官方"))
	m.SetHeader("To", sendTask.TaskInfo.Receiver...) //主送
	m.SetHeader("Subject", sendTask.MessageParamList.Variables.Title)
	//发送html格式邮件。
	m.SetBody("text/html", sendTask.MessageParamList.Variables.Content)
	d := gomail.NewDialer(acc.Host, acc.Port, acc.Username, acc.Password)
	if err := d.DialAndSend(m); err != nil {
		return errors.Wrap(err, "emailHandler DialAndSend err")
	}
	return nil
}
