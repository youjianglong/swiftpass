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

func (pwnr PayAlipayNativeRequest) ServiceName() string {
	return "pay.alipay.native"
}

func (pwnr PayAlipayNativeRequest) DecodeToXml(sign string) []byte {
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
