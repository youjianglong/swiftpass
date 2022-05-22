package response

import "encoding/xml"

type PayAlipayNativeResponse struct {
	XMLName xml.Name `xml:"xml"`
	CommonParams
	DeviceInfo string `xml:"device_info"`
	CodeUrl    string `xml:"code_url"`
	CodeImgUrl string `xml:"code_img_url"`
}
