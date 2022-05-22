package request

import "encoding/xml"

type PayAlipayJsPayRequest struct {
	XMLName xml.Name `xml:"xml"`
	CommonParam
	Body                 string `xml:"body"`
	Attach               string `xml:"attach,omitempty"`
	TotalFee             string `xml:"total_fee,omitempty"`
	MchCreateIp          string `xml:"mch_create_ip,omitempty"`
	NotifyUrl            string `xml:"notify_url,omitempty"`
	LimitCreditPay       string `xml:"limit_credit_pay,omitempty"`
	TimeStart            string `xml:"time_start,omitempty"`
	TimeExpire           string `xml:"time_expire,omitempty"`
	QrCodeTimeoutExpress string `xml:"qr_code_timeout_express,omitempty"`
	GoodsTag             string `xml:"goods_tag,omitempty"`
	ProductId            string `xml:"product_id,omitempty"`
	BuyerLogonId         string `xml:"buyer_logon_id,omitempty"`
	BuyerId              string `xml:"buyer_id,omitempty"`
}

func (r PayAlipayJsPayRequest) ServiceName() string {
	return "pay.alipay.jspay"
}

func (r PayAlipayJsPayRequest) Full() Request {
	if r.SignType == "" {
		r.SignType = "MD5"
	}
	r.Service = r.ServiceName()
	return r
}

func (r PayAlipayJsPayRequest) Encode(sign string) []byte {
	r.Sign = sign
	content, _ := xml.Marshal(r)
	return content
}
