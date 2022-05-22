//关闭订单
package request

import "encoding/xml"

type UnifiedTradeRefundQueryRequest struct {
	XMLName xml.Name `xml:"xml"`
	CommonParam
	TransactionId string `xml:"transaction_id,omitempty"`
	OutRefundNo   string `xml:"out_refund_no,omitempty"`
	RefundId      string `xml:"refund_id,omitempty"`
}

func (r UnifiedTradeRefundQueryRequest) ServiceName() string {
	return "unified.trade.refundquery"
}

func (r UnifiedTradeRefundQueryRequest) Full() Request {
	if r.SignType == "" {
		r.SignType = "MD5"
	}
	r.Service = r.ServiceName()
	return r
}

func (r UnifiedTradeRefundQueryRequest) Encode(sign string) []byte {
	r.Sign = sign
	content, _ := xml.Marshal(r)
	return content
}
