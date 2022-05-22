package swiftpass

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/xml"
	"io"
)

func Md5String(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func Xml2Map(reader io.Reader) map[string]string {
	val := XMLStringMap{}
	_ = xml.NewDecoder(reader).Decode(&val)
	return val
}
