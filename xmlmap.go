package swiftpass

import (
	"encoding/xml"
	"io"
)

type XMLStringMap map[string]string

type xmlMapEntry struct {
	XMLName xml.Name
	Value   string `xml:",chardata"`
}

func (m XMLStringMap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if len(m) == 0 {
		return nil
	}

	start.Name.Local = "xml"

	err := e.EncodeToken(start)
	if err != nil {
		return err
	}

	for k, v := range m {
		e.Encode(xmlMapEntry{XMLName: xml.Name{Local: k}, Value: v})
	}

	return e.EncodeToken(start.End())
}

func (m XMLStringMap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		var e xmlMapEntry

		err := d.Decode(&e)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		m[e.XMLName.Local] = e.Value
	}
	return nil
}

func (m XMLStringMap) ToXML() ([]byte, error) {
	return xml.Marshal(m)
}

func UnmarshalXMLToMap(data []byte) (XMLStringMap, error) {
	m := XMLStringMap{}
	err := xml.Unmarshal(data, &m)
	return m, err
}
