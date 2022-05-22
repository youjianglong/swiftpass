//关闭订单
package request

import "encoding/xml"

type UnifiedTradeCloseRequest struct {
	XMLName xml.Name `xml:"xml"`
	CommonParam
}

func (pwnr UnifiedTradeCloseRequest) ServiceName() string {
	return "unified.trade.close"
}

func (pwnr UnifiedTradeCloseRequest) DecodeToXml(sign string) []byte {
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
