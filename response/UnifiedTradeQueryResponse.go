package response

import "encoding/xml"

type UnifiedTradeQueryResponse struct {
	XMLName xml.Name `xml:"xml"`
	CommonParams
	AppId            string `xml:"appid"`
	OpenId           string `xml:"openid"`
	DeviceInfo       string `xml:"device_info"`
	TradeState       string `xml:"trade_state"`
	TradeType        string `xml:"trade_type"`
	IsSubscribe      string `xml:"is_subscribe"`
	TransactionId    string `xml:"transaction_id"`
	OutTransactionId string `xml:"out_transaction_id"`
	TotalFee         string `xml:"total_fee"`
	CouponFee        string `xml:"coupon_fee"`
	FeeType          string `xml:"fee_type"`
	Attach           string `xml:"attach"`
	BankType         string `xml:"bank_type"`
	OutTradeNo       string `xml:"out_trade_no"`
	BankBillNo       string `xml:"bank_billno"`
	TimeEnd          string `xml:"time_end"`
}
