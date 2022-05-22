package response

type Response interface {
	DecodeToMap()
}

type CommonParams struct {
	Version    string `xml:"version"`
	Charset    string `xml:"charsex"`
	SignType   string `xml:"sign_type"`
	Status     string `xml:"status"`
	Message    string `xml:"message"`
	Code       string `xml:"code"`
	ResultCode string `xml:"result_code"`
	MchId      string `xml:"mch_id"`
	NonceStr   string `xml:"nonce_str"`
	ErrCode    string `xml:"err_code"`
	ErrMsg     string `xml:"err_msg"`
	Sign       string `xml:"sign"`
}

func (p CommonParams) DecodeToMap() {
}
