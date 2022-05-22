package request

import "encoding/xml"

type PayAlipayNativeRequest struct {
	XMLName xml.Name `xml:"xml"`
	CommonParam
	Body                 string `xml:"body"`
	Attach               string `xml:"attach"`
	TotalFee             string `xml:"total_fee"`
	NotifyUrl            string `xml:"notify_url"`
	LimitCreditPay       string `xml:"limit_credit_pay"`
	TimeStart            string `xml:"time_start"`
	TimeExpire           string `xml:"time_expire"`
	QrCodeTimeoutExpress string `xml:"qr_code_timeout_express"`
	GoodsTag             string `xml:"goods_tag"`
	ProductId            string `xml:"product_id"`
	MchCreateIp          string `xml:"mch_create_ip"`
}

func (r PayAlipayNativeRequest) ServiceName() string {
	return "pay.alipay.native"
}

func (r PayAlipayNativeRequest) Full() Request {
	if r.SignType == "" {
		r.SignType = "MD5"
	}
	r.Service = r.ServiceName()
	return r
}

func (r PayAlipayNativeRequest) Encode(sign string) []byte {
	r.Sign = sign
	content, _ := xml.Marshal(r)
	return content
}
