package request

import "encoding/xml"

type PayWeixinJsPayRequest struct {
	XMLName xml.Name `xml:"xml"`
	CommonParam
	IsRaw       string `xml:"is_raw"`
	IsMinipg    string `xml:"is_minipg"`
	Body        string `xml:"body"`
	SubOpenid   string `xml:"sub_openid"`
	SubAppId    string `xml:"sub_appid"`
	Attach      string `xml:"attach"`
	TotalFee    string `xml:"total_fee"`
	NotifyUrl   string `xml:"notify_url"`
	TimeStart   string `xml:"time_start"`
	TimeExpire  string `xml:"time_expire"`
	GoodsTag    string `xml:"goods_tag"`
	ProductId   string `xml:"product_id"`
	MchCreateIp string `xml:"mch_create_ip"`
}

func (pwnr PayWeixinJsPayRequest) ServiceName() string {
	return "pay.weixin.jspay"
}

func (pwnr PayWeixinJsPayRequest) DecodeToXml(sign string) []byte {
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
