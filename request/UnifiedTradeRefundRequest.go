//关闭订单
package request

import "encoding/xml"

type UnifiedTradeRefundRequest struct {
	XMLName xml.Name `xml:"xml"`
	CommonParam
	TransactionId string `xml:"transaction_id"`
	OutRefundNo   string `xml:"out_refund_no"`
	TotalFee      string `xml:"total_fee"`
	RefundFee     string `xml:"refund_fee"`
	RefundChannel string `xml:"refund_channel"`
}

func (r UnifiedTradeRefundRequest) ServiceName() string {
	return "unified.trade.refund"
}

func (r UnifiedTradeRefundRequest) Full() Request {
	if r.SignType == "" {
		r.SignType = "MD5"
	}
	r.Service = r.ServiceName()
	return r
}

func (r UnifiedTradeRefundRequest) Encode(sign string) []byte {
	r.Sign = sign
	content, _ := xml.Marshal(r)
	return content
}
