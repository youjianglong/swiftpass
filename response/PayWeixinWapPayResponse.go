package response

import "encoding/xml"

type PayWeixinWapPayResponse struct {
	XMLName xml.Name `xml:"xml"`
	CommonParams
	AppId   string `xml:"appid"`
	PayInfo string `xml:"pay_info"`
}
