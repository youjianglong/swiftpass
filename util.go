package swiftpass

import (
	"crypto/md5"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"encoding/xml"
	"fmt"
	"io"
	"reflect"
	"sort"
	"strconv"
	"strings"
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

//对字典key进行排序
func SortKeys(m map[string]string) []string {
	keys := make([]string, 0)
	for key, _ := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

//生成url请求串
func BuildQueryString(params map[string]string) string {
	signParams := map[string]string{}
	for k, v := range params {
		if v == "" {
			continue
		}
		signParams[k] = v
	}

	keys := SortKeys(signParams)

	queryStr := ""
	for _, key := range keys {
		queryStr += key + "=" + signParams[key] + "&"
	}
	return queryStr[:len(queryStr)-1]
}

//解码私钥
func DecodePrivateKey(content []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(content)
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

//解码公钥
func DecodePublicKey(content []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(content)
	return x509.ParsePKCS1PublicKey(block.Bytes)
}

func XmlObject2Map(obj interface{}) map[string]string {
	valueOf := reflect.ValueOf(obj)
	kind := valueOf.Kind()
	for kind == reflect.Ptr {
		valueOf = valueOf.Elem()
		kind = valueOf.Kind()
	}
	if kind != reflect.Struct {
		return nil
	}

	typeOf := valueOf.Type()
	m := make(map[string]string)
	num := valueOf.NumField()

	for i := 0; i < num; i++ {
		fieldType := typeOf.Field(i)
		if !fieldType.IsExported() || fieldType.Name == "XMLName" {
			continue
		}

		tag := fieldType.Tag.Get("xml")
		if tag == "-" {
			continue
		}

		if fieldType.Anonymous {
			subMap := XmlObject2Map(valueOf.Field(i).Interface())
			if subMap != nil {
				for k, v := range subMap {
					m[k] = v
				}
			}
			continue
		}

		tags := strings.Split(tag, ",")

		key := strings.TrimSpace(tags[0])
		if key == "" {
			key = fieldType.Name
		}

		var omitempty bool
		if len(tags) > 1 && strings.Contains(tag, "omitempty") {
			omitempty = true
		}

		fieldValue := valueOf.Field(i)
		var val string
		switch fieldValue.Kind() {
		case reflect.String:
			val = fieldValue.String()
		case reflect.Float64, reflect.Float32:
			val = strconv.FormatFloat(fieldValue.Float(), 'f', -1, 64)
		case reflect.Bool:
			if fieldValue.Bool() {
				val = "true"
			} else {
				if omitempty {
					val = ""
				} else {
					val = "false"
				}
			}
		default:
			val = fmt.Sprintf("%v", fieldValue.Interface())
		}
		if val == "" && omitempty {
			continue
		}
		m[key] = val
	}

	return m
}
