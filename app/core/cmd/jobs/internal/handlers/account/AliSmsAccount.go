package account

type AliSmsAccount struct {
	SecretId   string `json:"secretId"`
	SecretKey  string `json:"secretKey"`
	Region     string `json:"region"`     //地域
	Url        string `json:"url"`        //url地址
	SignName   string `json:"signName"`   //签名
	TemplateId string `json:"templateId"` //模版ID
}
