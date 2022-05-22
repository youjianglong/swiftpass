package swiftpass

import (
	"encoding/xml"
)

// Notify 订单通知 https://open.swiftpass.cn/openapi/doc?index_1=171&index_2=35&chapter_1=2225&chapter_2=2231
type Notify struct {
	XMLName          xml.Name `xml:"xml"`
	Version          string   `xml:"version,omitempty"`            //版本号，version默认值是2.0
	Charset          string   `xml:"charset,omitempty"`            //字符集，可选值 UTF-8，默认为 UTF-8
	Status           string   `xml:"status"`                       //返回状态码，0表示成功，非0表示失败此字段是通信标识，非交易标识，交易是否成功需要查看 result_code 来判断
	Message          string   `xml:"message,omitempty"`            //返回信息，如非空，为错误原因签名失败参数格式校验错误
	ResultCode       string   `xml:"result_code,omitempty"`        //业务结果，0表示成功非0表示失败
	MchId            string   `xml:"mch_id,omitempty"`             //门店号，由平台分配
	DeviceInfo       string   `xml:"device_info,omitempty"`        //终端设备号
	NonceStr         string   `xml:"nonce_str,omitempty"`          //随机字符串，不长于 32 位
	ErrCode          string   `xml:"err_code,omitempty"`           //参考错误码
	ErrMsg           string   `xml:"err_msg,omitempty"`            //结果信息描述
	SignType         string   `xml:"sign_type,omitempty"`          //签名类型，取值：MD5、RSA_1_256、RSA_1_1，默认：MD5
	Sign             string   `xml:"sign,omitempty"`               //签名，MD5/RSA_1_256/RSA_1_1签名结果，详见“安全规范”
	Openid           string   `xml:"openid,omitempty"`             //用户在商户 appid 下的唯一标识
	TradeType        string   `xml:"trade_type,omitempty"`         //交易类型
	IsSubscribe      string   `xml:"is_subscribe,omitempty"`       //用户是否关注公众账号，Y-关注，N-未关注，仅在公众账号类型支付有效
	PayResult        string   `xml:"pay_result,omitempty"`         //支付结果：0—成功；其它—失败
	PayInfo          string   `xml:"pay_info,omitempty"`           //支付结果信息，支付成功时为空
	TransactionId    string   `xml:"transaction_id,omitempty"`     //平台交易号
	OutTransactionId string   `xml:"out_transaction_id,omitempty"` //第三方订单号
	ThirdOrderNo     string   `xml:"third_order_no,omitempty"`     //第三方商户单号
	SubIsSubscribe   string   `xml:"sub_is_subscribe,omitempty"`   //用户是否关注子公众账号，Y-关注，N-未关注，仅在公众账号类型支付有效
	SubAppid         string   `xml:"sub_appid,omitempty"`          //子商户appid
	SubOpenid        string   `xml:"sub_openid,omitempty"`         //用户在商户 sub_appid 下的唯一标识
	OutTradeNo       string   `xml:"out_trade_no,omitempty"`       //商户系统内部的定单号，32个字符内、可包含字母
	TotalFee         int      `xml:"total_fee,omitempty"`          //总金额，以分为单位，不允许包含任何字、符号
	CashFee          int      `xml:"cash_fee,omitempty"`           //现金支付金额订单现金支付金额，详见支付金额
	CouponFee        int      `xml:"coupon_fee,omitempty"`         //现金券支付金额<=订单总金额， 订单总金额-现金券金额为现金支付金额
	PromotionDetail  string   `xml:"promotion_detail,omitempty"`   //优惠详情
	FeeType          string   `xml:"fee_type,omitempty"`           //货币类型，符合 ISO 4217 标准的三位字母代码，默认人民币：CNY
	Attach           string   `xml:"attach,omitempty"`             //商家数据包，原样返回
	BankType         string   `xml:"bank_type,omitempty"`          //银行类型
	BankBillno       string   `xml:"bank_billno,omitempty"`        //银行订单号，非银行卡支付则为空
	TimeEnd          string   `xml:"time_end,omitempty"`           //支付完成时间，格式为yyyyMMddHHmmss，如2009年12月27日9点10分10秒表示为20091227091010。时区为GMT+8 beijing。该时间取自平台服务器
	Mdiscount        int      `xml:"mdiscount,omitempty"`          //免充值优惠金额
}
