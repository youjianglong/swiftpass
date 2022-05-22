package response

import "encoding/xml"

type PayWeixinJsPayResponse struct {
	XMLName xml.Name `xml:"xml"`
	CommonParams
	DeviceInfo string `xml:"device_info"`
	CodeUrl    string `xml:"code_url"`
	CodeImgUrl string `xml:"code_img_url"`
	AppId      string `xml:"appid"`
	TokenId    string `xml:"token_id"`
	PayInfo    string `xml:"pay_info"`
}

type WeixinJsPayPayInfo struct {
	AppId       string `json:"appId"`
	TimeStamp   string `json:"timeStamp"`
	Status      string `json:"status"`
	SignType    string `json:"signType"`
	Package     string `json:"package"`
	CallbackUrl string `json:"callback_url"`
	NonceStr    string `json:"nonceStr"`
	PaySign     string `json:"paySign"`
}
