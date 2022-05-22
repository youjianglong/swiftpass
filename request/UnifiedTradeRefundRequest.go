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

func (pwnr UnifiedTradeRefundRequest) ServiceName() string {
	return "unified.trade.refund"
}

func (pwnr UnifiedTradeRefundRequest) DecodeToXml(sign string) []byte {
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
