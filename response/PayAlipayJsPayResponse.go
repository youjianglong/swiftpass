package response

import "encoding/xml"

type PayAlipayJsPayResponse struct {
	XMLName xml.Name `xml:"xml"`
	CommonParams
	DeviceInfo string `xml:"device_info"`
	PayInfo    string `xml:"pay_info"`
	PayUrl     string `xml:"pay_url"`
}

type AlipayJsPayPayInfo struct {
	TradeNo string `json:"tradeNO"`
	Status  string `json:"status"`
}
