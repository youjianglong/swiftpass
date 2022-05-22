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

func (pwnr PayWeixinWapPayRequest) ServiceName() string {
	return "pay.weixin.wappay"
}

func (pwnr PayWeixinWapPayRequest) DecodeToXml(sign string) []byte {
	pwnr.Sign = sign
	if pwnr.SignType == "" {
		pwnr.SignType = "MD5"
	}
	pwnr.Service = pwnr.ServiceName()
	xmlByte, decodeError := xml.Marshal(pwnr)
	if decodeError == nil {
		return xmlByte
	} else {
		return nil
	}
}
