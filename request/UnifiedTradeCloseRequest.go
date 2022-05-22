//关闭订单
package request

import "encoding/xml"

type UnifiedTradeCloseRequest struct {
	XMLName xml.Name `xml:"xml"`
	CommonParam
}

func (r UnifiedTradeCloseRequest) ServiceName() string {
	return "unified.trade.close"
}

func (r UnifiedTradeCloseRequest) Full() Request {
	if r.SignType == "" {
		r.SignType = "MD5"
	}
	r.Service = r.ServiceName()
	return r
}

func (r UnifiedTradeCloseRequest) Encode(sign string) []byte {
	r.Sign = sign
	content, _ := xml.Marshal(r)
	return content
}
