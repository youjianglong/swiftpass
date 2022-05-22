package response

import "encoding/xml"

type UnifiedTradeCloseResponse struct {
	XMLName xml.Name `xml:"xml"`
	CommonParams
}
