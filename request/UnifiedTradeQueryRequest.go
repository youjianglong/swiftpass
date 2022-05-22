package request

import "encoding/xml"

type UnifiedTradeQueryRequest struct {
	XMLName xml.Name `xml:"xml"`
	CommonParam
	TransactionId string `xml:"transaction_id"`
}

func (r UnifiedTradeQueryRequest) ServiceName() string {
	return "unified.trade.query"
}

func (r UnifiedTradeQueryRequest) Full() Request {
	if r.SignType == "" {
		r.SignType = "MD5"
	}
	r.Service = r.ServiceName()
	return r
}

func (r UnifiedTradeQueryRequest) Encode(sign string) []byte {
	r.Sign = sign
	content, _ := xml.Marshal(r)
	return content
}
