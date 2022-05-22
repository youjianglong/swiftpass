package response

import "encoding/xml"

type UnifiedTradeRefundQueryResponse struct {
	XMLName xml.Name `xml:"xml"`
	CommonParams
	DeviceInfo       string `xml:"device_info"`
	TransactionId    string `xml:"transaction_id"`
	OutTradeNo       string `xml:"out_trade_no"`
	RefundCount      string `xml:"refund_count"`
	OutRefundNo0     string `xml:"out_refund_no_0"`
	RefundId0        string `xml:"refund_id_0"`
	RefundChannel0   string `xml:"refund_channel_0"`
	RefundFee0       string `xml:"refund_fee_0"`
	CouponRefundFee0 string `xml:"coupon_refund_fee_0"`
	RefundTime0      string `xml:"refund_time_0"`
	RefundStatus0    string `xml:"refund_status_0"`
	OutRefundNo1     string `xml:"out_refund_no_1"`
	RefundId1        string `xml:"refund_id_1"`
	RefundChannel1   string `xml:"refund_channel_1"`
	RefundFee1       string `xml:"refund_fee_1"`
	CouponRefundFee1 string `xml:"coupon_refund_fee_1"`
	RefundTime1      string `xml:"refund_time_1"`
	RefundStatus1    string `xml:"refund_status_1"`
}
