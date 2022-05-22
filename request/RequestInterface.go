package request

//威富通的支付接口
type Request interface {
	ServiceName() string
	Full() Request
	Encode(sign string) []byte
}

type CommonParam struct {
	Service        string `xml:"service,omitempty"`          //接口类型
	Version        string `xml:"version,omitempty"`          //版本号，version默认值是2.0
	Charset        string `xml:"charset,omitempty"`          //可选值 UTF-8 ，默认为 UTF-8
	SignType       string `xml:"sign_type,omitempty"`        //签名类型，取值：MD5、RSA_1_256、RSA_1_1，默认：MD5
	SignAgentNo    string `xml:"sign_agentno,omitempty"`     //由平台分配。传入了此参数时，数据的签名使用的将是服务商的signKey
	MchId          string `xml:"mch_id,omitempty"`           //门店编号，由平台分配
	OutTradeNo     string `xml:"out_trade_no,omitempty"`     //商户系统内部的订单号 ,32个字符内、 可包含字母,确保在商户系统唯一
	DeviceInfo     string `xml:"device_info,omitempty"`      //终端设备号
	OpUserId       string `xml:"op_user_id,omitempty"`       //操作员帐号,默认为门店编号
	NonceStr       string `xml:"nonce_str,omitempty"`        //随机字符串，不长于32位
	LimitCreditPay string `xml:"limit_credit_pay,omitempty"` //限定用户使用时能否使用信用卡，值为1，禁用信用卡；值为0或者不传此参数则不禁用
	Sign           string `xml:"sign,omitempty"`             //MD5/RSA_1_256/RSA_1_1签名结果，详见“安全规范”
}
