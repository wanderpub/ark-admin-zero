// Code generated by goctl. DO NOT EDIT.
package types

type LoginReq struct {
	CaptchaId  string `json:"captchaId"`
	VerifyCode string `json:"verifyCode"`
	Account    string `json:"account"`
	Password   string `json:"password"`
}

type LoginResp struct {
	Token string `json:"token"`
}

type InfoResp struct {
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
}

type ProfileResp struct {
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Gender   int64  `json:"gender"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	Remark   string `json:"remark"`
	Avatar   string `json:"avatar"`
}

type Menu struct {
	Id           int64  `json:"id"`
	ParentId     int64  `json:"parentId"`
	Name         string `json:"name"`
	Router       string `json:"router"`
	Type         int64  `json:"type"`
	Icon         string `json:"icon"`
	OrderNum     int64  `json:"orderNum"`
	ViewPath     string `json:"viewPath"`
	IsShow       int64  `json:"isShow"`
	ActiveRouter string `json:"activeRouter"`
}

type PermMenuResp struct {
	Menus []Menu   `json:"menus"`
	Perms []string `json:"perms"`
}

type PasswordReq struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

type CaptchaResp struct {
	CaptchaId  string `json:"captchaId"`
	VerifyCode string `json:"verifyCode"`
}
