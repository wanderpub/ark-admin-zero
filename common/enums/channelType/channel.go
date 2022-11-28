package channelType

const (
	Im                 = 10
	Push               = 20
	Sms                = 30
	AliSms             = 31
	TencentSms         = 32
	Email              = 40
	OfficialAccounts   = 50
	MiniProgram        = 60
	EnterpriseWeChat   = 70
	DingDing           = 80
	DingDingWorkNotice = 90
)

var (
	TypeText = map[int]string{
		Im:                 "IM(站内信)",
		Push:               "push(通知栏)",
		Sms:                "sms(短信)",
		AliSms:             "阿里云(短信)",
		TencentSms:         "腾讯云(短信)",
		Email:              "email(邮件)",
		OfficialAccounts:   "OfficialAccounts(服务号)",
		MiniProgram:        "miniProgram(小程序)",
		EnterpriseWeChat:   "EnterpriseWeChat(企业微信)",
		DingDing:           "dingDingRobot(钉钉机器人)",
		DingDingWorkNotice: "dingDingWorkNotice(钉钉工作通知)",
	}
	TypeCodeEn = map[int]string{
		Im:                 "im",
		Push:               "push",
		Sms:                "sms",
		AliSms:             "ali_sms",
		TencentSms:         "tencent_sms",
		Email:              "email",
		OfficialAccounts:   "official_accounts",
		MiniProgram:        "mini_program",
		EnterpriseWeChat:   "enterprise_we_chat",
		DingDing:           "ding_ding_robot",
		DingDingWorkNotice: "ding_ding_work_notice",
	}
)
