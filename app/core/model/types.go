package model

type TaskInfo struct {
	MessageTemplateId int64    `json:"messageTemplateId"`
	BusinessId        int64    `json:"businessId"`
	Receiver          []string `json:"receiver"` //先去重
	IdType            int64    `json:"idType"`
	SendChannel       int64    `json:"sendChannel"`
	TemplateType      int64    `json:"templateType"`
	MsgType           int64    `json:"msgType"`
	ShieldType        int64    `json:"shieldType"`
	SendAccount       int64    `json:"sendAccount"`
	Config            string   `json:"config"`
	SendAccountConfig string   `json:"sendAccountConfig"`
}

type SendTaskModel struct {
	Code              string       `json:"code"`
	MessageTemplateId int64        `json:"message_template_id"`
	Time              int64        `json:"time"`
	MessageParamList  MessageParam `json:"message_params"`
	TaskInfo          TaskInfo     `json:"task_info"`
}

type MessageParam struct {
	Receiver  string                 `json:"receiver"`       //接收者 多个用,逗号号分隔开
	Variables ContentModel           `json:"variables"`      //可选 消息内容中的可变部分(占位符替换)
	Extra     map[string]interface{} `json:"extra,optional"` //可选 扩展参数
}

type ContentModel struct {
	Map        map[string]string `json:"map,optional"`         //消息数据key/value形式
	Array      []string          `json:"array,optional"`       //消息数据数组形式
	TemplateId string            `json:"template_id,optional"` // 发送消息的模版ID
	Url        string            `json:"url,optional"`         // 消息的URL地址
	Title      string            `json:"title,optional"`       //标题
	Content    string            `json:"content,optional"`     //内容
	MediaId    string            `json:"media_id,optional"`    //媒体ID
	SendType   string            `json:"send_type,optional"`   //类型
	SignName   string            `json:"sign_name,optional"`   //签名
	AppID      string            `json:"app_id,optional"`      //appid
	ID         string            `json:"id,optional"`          //外部编号
}

type PeriodLimit struct {
	Quota   int `json:"quota"`   //配额
	Seconds int `json:"seconds"` //秒
}
