package request

import "encoding/xml"

type PayWeixinWapPayRequest struct {
	XMLName xml.Name `xml:"xml"`
	CommonParam
	GroupNo     string `xml:"groupno"`
	Body        string `xml:"body"`
	Attach      string `xml:"attach"`
	TotalFee    string `xml:"total_fee"`
	NotifyUrl   string `xml:"notify_url"`
	CallbackUrl string `xml:"callback_url"`
	TimeStart   string `xml:"time_start"`
	TimeExpire  string `xml:"time_expire"`
	MchAppName  string `xml:"mch_app_name"`
	MchAppId    string `xml:"mch_app_id	"`
	GoodsTag    string `xml:"goods_tag"`
	MchCreateIp string `xml:"mch_create_ip"`
}

func (r PayWeixinWapPayRequest) ServiceName() string {
	return "pay.weixin.wappay"
}

func (r PayWeixinWapPayRequest) Full() Request {
	if r.SignType == "" {
		r.SignType = "MD5"
	}
	r.Service = r.ServiceName()
	return r
}

func (r PayWeixinWapPayRequest) Encode(sign string) []byte {
	r.Sign = sign
	content, _ := xml.Marshal(r)
	return content
}
