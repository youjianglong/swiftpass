package request

import (
	"encoding/xml"
)

//统一扫码支付下单
type UnifiedTradeNativeRequest struct {
	XMLName xml.Name `xml:"xml"`
	CommonParam
	Body              string `xml:"body"`                           //商品描述
	ProfitShareInfos  string `xml:"profit_share_infos,omitempty"`   //分账信息，https://open.swiftpass.cn/openapi/wiki?index=271&chapter=71
	HbFqNum           string `xml:"hb_fq_num,omitempty"`            //只支持传"3"|"6"|"12"，只适用于支付宝支付
	HbFqSellerPercent string `xml:"hb_fq_seller_percent,omitempty"` //只支持传"0"|"100"，商家承担手续费传"100"，用户承担传"0"，在有hb_fq_num字段时默认为“0 ”
	Attach            string `xml:"attach,omitempty"`               //商户附加信息，可做扩展参数
	SubAppid          string `xml:"sub_appid,omitempty"`            //微信公众平台基本配置中的AppID(应用ID)，传入后支付成功可返回对应公众号下的用户openid
	TotalFee          int    `xml:"total_fee"`                      //总金额，以分为单位，不允许包含任何字、符号
	NeedReceipt       bool   `xml:"need_receipt,omitempty"`         //需要和微信公众平台的发票功能联合，传入true时，微信支付成功消息和支付详情页将出现开票入口[新增need_receipt【适用于微信】]
	MchCreateIp       string `xml:"mch_create_ip"`                  //上传商户真实的发起交易的终端出网IP
	NotifyUrl         string `xml:"notify_url"`                     //接收平台通知的URL，需给绝对路径，255字符内格式如:http://wap.tenpay.com/tenpay.asp，确保平台能通过互联网访问该地址，如不需要接收通知，请传：http://127.0.0.1
	CallbackUrl       string `xml:"callback_url,omitempty"`         //前端页面跳转的URL（包括支付成功和关闭时都会跳到这个地址,商户需自行处理逻辑），需给绝对路径，255字符内格式如:http://wap.tenpay.com/callback.asp注:该地址只作为前端页面的一个跳转，须使用notify_url通知结果作为支付最终结果。
	TimeStart         string `xml:"time_start,omitempty"`           //订单生成时间，格式为yyyyMMddHHmmss，如2009年12月25日9点10分10秒表示为20091225091010。时区为GMT+8 beijing。该时间取自商户服务器。注：订单生成时间与超时时间需要同时传入才会生效。
	TimeExpire        string `xml:"time_expire,omitempty"`          //订单失效时间，格式为yyyyMMddHHmmss，如2009年12月27日9点10分10秒表示为20091227091010。时区为GMT+8 beijing。该时间取自商户服务器。注：订单生成时间与超时时间需要同时传入才会生效。
	GoodsTag          string `xml:"goods_tag,omitempty"`            //商品标记，微信平台配置的商品标记，用于优惠券或者满减使用
}

//分账信息
type ProfitShareInfo struct {
	TransIn string `xml:"trans_in"`
	TransNo string `xml:"trans_no"`
	Amount  int    `xml:"amount"`
	Desc    string `xml:"desc"`
}

func (r UnifiedTradeNativeRequest) ServiceName() string {
	return "unified.trade.native"
}

func (r UnifiedTradeNativeRequest) Full() Request {
	if r.SignType == "" {
		r.SignType = "MD5"
	}
	r.Service = r.ServiceName()
	return r
}

func (r UnifiedTradeNativeRequest) Encode(sign string) []byte {
	r.Sign = sign
	content, _ := xml.Marshal(r)
	return content
}
