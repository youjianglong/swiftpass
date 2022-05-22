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

func (pwnr UnifiedTradeRefundQueryRequest) ServiceName() string {
	return "unified.trade.refundquery"
}

func (pwnr UnifiedTradeRefundQueryRequest) DecodeToXml(sign string) []byte {
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
