package handlers

import (
	"ark-admin-zero/app/core/cmd/jobs/internal/handlers/account"
	"ark-admin-zero/app/core/cmd/jobs/internal/svc"
	"ark-admin-zero/app/core/model"
	"context"
	"fmt"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/jsonx"
)

type tencentSmsHandler struct {
	svcCtx *svc.ServiceContext
}

func NewTencentSmsHandler(svcCtx *svc.ServiceContext) IHandler {
	return tencentSmsHandler{
		svcCtx: svcCtx,
	}
}

func (h tencentSmsHandler) DoHandler(ctx context.Context, sendTask model.SendTaskModel) (err error) {
	var acc account.TencentSmsAccount
	err = jsonx.UnmarshalFromString(sendTask.TaskInfo.SendAccountConfig, &acc)
	if err != nil {
		return errorx.Wrap(err, "alismsHandler get account err")
	}
	// 实例化一个认证对象，入参需要传入腾讯云账户secretId，secretKey,此处还需注意密钥对的保密
	// 密钥可前往https://console.cloud.tencent.com/cam/capi网站进行获取
	credential := common.NewCredential(
		acc.SecretId,
		acc.SecretKey,
	)
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = acc.Url
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := sms.NewClient(credential, acc.Region, cpf)

	// err = jsonx.UnmarshalFromString(sendTask.TaskInfo.Config, &acc)
	// if err != nil {
	// 	return errorx.Wrap(err, "alismsHandler get account err")
	// }
	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := sms.NewSendSmsRequest()
	request.SmsSdkAppId = common.StringPtr(sendTask.MessageParamList.Variables.AppID)
	request.SignName = common.StringPtr(sendTask.MessageParamList.Variables.SignName)
	request.TemplateId = common.StringPtr(sendTask.MessageParamList.Variables.TemplateId)

	request.PhoneNumberSet = common.StringPtrs(sendTask.TaskInfo.Receiver)
	request.TemplateParamSet = common.StringPtrs(sendTask.MessageParamList.Variables.Array)
	fmt.Println(request)
	// 返回的resp是一个SendSmsResponse的实例，与请求对象对应
	response, err := client.SendSms(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		panic(err)
	}
	// 输出json格式的字符串回包
	fmt.Printf("%s", response.ToJsonString())
	return nil
}
