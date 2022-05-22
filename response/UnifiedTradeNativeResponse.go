package response

import "encoding/xml"

type UnifiedTradeNativeResponse struct {
	XMLName xml.Name `xml:"xml"`
	CommonParams
	CodeUrl    string `xml:"code_url"`
	CodeImgUrl string `xml:"code_img_url"`
	UUID       string `xml:"uuid,omitempty"`
}
