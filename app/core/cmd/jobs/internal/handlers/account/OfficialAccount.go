package account

type OfficialAccount struct {
	AppID          string            `json:"app_id"`           // appid
	AppSecret      string            `json:"app_secret"`       // appsecret
	Token          string            `json:"token"`            // token
	EncodingAESKey string            `json:"encoding_aes_key"` // EncodingAESKey
	OpenId         string            `json:"openId"`
	TemplateId     string            `json:"templateId"`
	Url            string            `json:"url"`
	MiniProgramId  string            `json:"miniProgramId"`
	Path           string            `json:"path"`
	Map            map[string]string `json:"map"`
}
