// Code generated by goctl. DO NOT EDIT.
package types

type Request struct {
	Username     string      `json:"username"`
	Password     string      `json:"password,optional"`
	Captcha      *CaptchaReq `json:"captcha,optional"`
	Phone        string      `json:"phone,optional"`
	Promotion    string      `json:"promotion,optional"`
	Code         string      `json:"code,optional"`
	Country      string      `json:"country,optional"`
	SuperPartner string      `json:"superPartner,optional"`
	Ip           string      `json:"ip,optional"`
}

type CaptchaReq struct {
	Server string `json:"server"`
	Token  string `json:"token"`
}

type Response struct {
	Message string `json:"message"`
}

type CodeRequest struct {
	Phone   string `json:"phone,optional"`
	Country string `json:"country,optional"`
}

type CodeResponse struct{}
