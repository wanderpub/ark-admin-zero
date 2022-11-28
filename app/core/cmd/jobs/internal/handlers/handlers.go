package handlers

import (
	"ark-admin-zero/app/core/cmd/jobs/internal/svc"
	"ark-admin-zero/common/enums/channelType"
	"sync"
)

var (
	once          sync.Once
	handlerHolder map[int64]IHandler
)

// SetUp 初始化所有handler
func SetUp(svcCtx *svc.ServiceContext) {
	once.Do(func() {
		handlerHolder = map[int64]IHandler{
			channelType.Email:              NewEmailHandler(svcCtx),
			channelType.AliSms:             NewAliSmsHandler(svcCtx),
			channelType.TencentSms:         NewTencentSmsHandler(svcCtx),
			channelType.OfficialAccounts:   NewOfficialAccountHandler(svcCtx),
			channelType.DingDing:           NewDingDingRobotHandler(svcCtx),
			channelType.EnterpriseWeChat:   NewEnterpriseWeChatHandler(svcCtx),
			channelType.DingDingWorkNotice: NewDingDingWorkNoticeHandler(svcCtx),
		}
	})
}

func GetHandler(sendChannel int64) IHandler {
	return handlerHolder[sendChannel]
}
