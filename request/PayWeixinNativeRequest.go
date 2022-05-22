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

func (r PayWeixinNativeRequest) ServiceName() string {
	return "pay.weixin.native"
}

func (r PayWeixinNativeRequest) Full() Request {
	if r.SignType == "" {
		r.SignType = "MD5"
	}
	r.Service = r.ServiceName()
	return r
}

func (r PayWeixinNativeRequest) Encode(sign string) []byte {
	r.Sign = sign
	content, _ := xml.Marshal(r)
	return content
}
