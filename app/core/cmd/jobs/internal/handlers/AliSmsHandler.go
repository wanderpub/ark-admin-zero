package handlers

import (
	"ark-admin-zero/app/core/cmd/jobs/internal/handlers/account"
	"ark-admin-zero/app/core/cmd/jobs/internal/svc"
	"ark-admin-zero/app/core/model"
	"context"
	"fmt"
	"strings"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	smsapi "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/jsonx"
)

type aliSmsHandler struct {
	svcCtx *svc.ServiceContext
}

func NewAliSmsHandler(svcCtx *svc.ServiceContext) IHandler {
	return aliSmsHandler{
		svcCtx: svcCtx,
	}
}

func (h aliSmsHandler) DoHandler(ctx context.Context, sendTask model.SendTaskModel) (err error) {
	var acc account.AliSmsAccount
	err = jsonx.UnmarshalFromString(sendTask.TaskInfo.SendAccountConfig, &acc)
	if err != nil {
		return errors.Wrap(err, "alismsHandler get account err")
	}
	openConfig := &openapi.Config{
		AccessKeyId:     &acc.SecretId,
		AccessKeySecret: &acc.SecretKey,
	}
	// 访问的域名
	openConfig.Endpoint = tea.String(acc.Url)
	client, _ := smsapi.NewClient(openConfig)

	// err = jsonx.UnmarshalFromString(sendTask.TaskInfo.Config, &acc)
	// if err != nil {
	// 	return errors.Wrap(err, "alismsHandler get account err")
	// }
	request := &smsapi.SendSmsRequest{}
	request.SetSignName(sendTask.MessageParamList.Variables.SignName)
	request.SetTemplateCode(sendTask.MessageParamList.Variables.TemplateId)

	templateParam, _ := jsonx.MarshalToString(sendTask.MessageParamList.Variables.Map)
	request.SetTemplateParam(templateParam)
	//收短信手机号，多个用","分隔
	request.SetPhoneNumbers(strings.Join(sendTask.TaskInfo.Receiver, ","))
	//发送短信
	response, err := client.SendSms(request)
	if err != nil {
		return errors.Wrap(err, "alismsHandler DialAndSend err")
	}
	fmt.Printf("%v", response.Body)
	return nil
}
