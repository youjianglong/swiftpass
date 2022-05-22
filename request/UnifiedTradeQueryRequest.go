package request

import "encoding/xml"

type UnifiedTradeQueryRequest struct {
	XMLName xml.Name `xml:"xml"`
	CommonParam
	TransactionId string `xml:"transaction_id"`
}

func (pwnr UnifiedTradeQueryRequest) ServiceName() string {
	return "unified.trade.query"
}

func (pwnr UnifiedTradeQueryRequest) DecodeToXml(sign string) []byte {
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
