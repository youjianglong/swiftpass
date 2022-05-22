package response

import "encoding/xml"

type UnifiedTradeRefundResponse struct {
	XMLName xml.Name `xml:"xml"`
	CommonParams
	DeviceInfo      string `xml:"device_info"`
	TransactionId   string `xml:"transaction_id"`
	OutTradeNo      string `xml:"out_trade_no"`
	OutRefundNo     string `xml:"out_refund_no"`
	RefundId        string `xml:"refund_id"`
	RefundChannel   string `xml:"refund_channel"`
	RefundFee       string `xml:"refund_fee"`
	CouponRefundFee string `xml:"coupon_refund_fee"`
}
