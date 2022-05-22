package request

import "encoding/xml"

type PayWeixinNativeRequest struct {
	XMLName xml.Name `xml:"xml"`
	CommonParam
	Body        string `xml:"body,omitempty"`
	Attach      string `xml:"attach,omitempty"`
	TotalFee    string `xml:"total_fee,omitempty"`
	NotifyUrl   string `xml:"notify_url,omitempty"`
	TimeStart   string `xml:"time_start,omitempty"`
	TimeExpire  string `xml:"time_expire,omitempty"`
	GoodsTag    string `xml:"goods_tag,omitempty"`
	ProductId   string `xml:"product_id,omitempty"`
	MchCreateIp string `xml:"mch_create_ip,omitempty"`
}

func (pwnr PayWeixinNativeRequest) ServiceName() string {
	return "pay.weixin.native"
}

func (pwnr PayWeixinNativeRequest) DecodeToXml(sign string) []byte {
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
