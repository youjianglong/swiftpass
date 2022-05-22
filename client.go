package swiftpass

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"

	"swiftpass/request"
	"swiftpass/response"
)

const (
	SWIFTPASS_GETWAY = "https://pay.hstypay.com/v2/pay/gateway"
)

const (
	RESULT_STATUS_SUCCESS = "0"
	RESULT_CODE_SUCCESS   = "0"
)

const (
	TRADE_STATE_SUCCESS = "SUCCESS"
	TRADE_STATE_REFUND  = "REFUND"
	TRADE_STATE_NOTPAY  = "NOTPAY"
	TRADE_STATE_CLOSED  = "CLOSED"
	TRADE_STATE_REVERSE = "REVERSE"
	TRADE_STATE_REVOKED = "REVOKED"
)

const (
	SignTypeMd5 = "MD5"
	SignTypeRsa = "RSA"
)

type SwiftClient struct {
	Key         string
	GetWay      string
	SignType    string
	RequestXml  string
	ResponseXml string
	PrivateKey  []byte
	PublicKey   []byte

	client *http.Client
}

func NewDefaultClient(key string) *SwiftClient {
	var client SwiftClient
	client.GetWay = SWIFTPASS_GETWAY
	client.Key = key
	client.SignType = SignTypeMd5
	client.client = &http.Client{
		Transport: &http.Transport{
			MaxIdleConns: 64,
		},
	}
	return &client
}

//加载rsa秘钥文件
func (swiftClient *SwiftClient) LoadRsaKeyFile(priveteKey, publicKey string) {
	readPrivateByte, readErr := ioutil.ReadFile(priveteKey)
	if readErr != nil || len(readPrivateByte) == 0 {
		panic("private key file read fail")
	}
	swiftClient.PrivateKey = readPrivateByte
	readPublicByte, readErr := ioutil.ReadFile(publicKey)
	if readErr != nil || len(readPublicByte) == 0 {
		panic("private key file read fail")
	}
	swiftClient.PublicKey = readPublicByte
}

//加载rsa秘钥文件
func (swiftClient *SwiftClient) LoadRsaKeyByte(priveteKey, publicKey []byte) {
	swiftClient.PrivateKey = priveteKey
	swiftClient.PublicKey = publicKey
}

//加载私钥
func loadPrivateKey(content []byte) *rsa.PrivateKey {
	block, _ := pem.Decode(content)
	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	return key
}

//加载公钥
func loadPublicKey(content []byte) *rsa.PublicKey {
	block, _ := pem.Decode(content)
	key, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}
	return key
}

//生成md5签名
func (swiftClient *SwiftClient) GenerateSign(req request.Request) string {
	xmlData := req.DecodeToXml("")
	params := Xml2Map(bytes.NewReader(xmlData))
	paramStr := buildQuery(params)
	paramStr += "&key=" + swiftClient.Key
	sign := strings.ToUpper(Md5String(paramStr))
	return sign
}

//生成ras签名
func (swiftClient *SwiftClient) GenerateSignRsa(req request.Request) string {
	if len(swiftClient.PrivateKey) == 0 || len(swiftClient.PublicKey) == 0 {
		panic("please load key file!")
	}
	xmlData := req.DecodeToXml("")
	params := Xml2Map(bytes.NewReader(xmlData))
	data := []byte(buildQuery(params))
	h := sha256.New()
	h.Write(data)
	hashed := h.Sum(nil)
	privateKey := loadPrivateKey(swiftClient.PrivateKey)
	newData, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(newData)
}

//提交支付
func (c *SwiftClient) Execute(swpReq request.Request, swpResp response.Response) error {
	sign := ""
	if c.SignType == "" || c.SignType == SignTypeMd5 {
		sign = c.GenerateSign(swpReq)
	} else {
		sign = c.GenerateSignRsa(swpReq)
	}

	data := swpReq.DecodeToXml(sign)
	c.RequestXml = string(data)

	req, _ := http.NewRequest("POST", c.GetWay, strings.NewReader(string(data)))
	req.Header.Set("Accept", "text/xml,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set("Charset", "UTF-8")

	res, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	content, rerr := ioutil.ReadAll(res.Body)
	if rerr != nil {
		return rerr
	}
	c.ResponseXml = string(content)
	encodeErr := xml.Unmarshal(content, swpResp)
	if encodeErr != nil {
		return encodeErr
	}
	return nil
}

//验证通知签名
func (c *SwiftClient) Verify(n Notify) bool {
	if n.Sign == "" {
		return false
	}
	xmlData, _ := xml.Marshal(n)
	params := Xml2Map(bytes.NewReader(xmlData))
	delete(params, "sign")
	paramStr := buildQuery(params)

	if n.SignType == SignTypeMd5 { //MD5签名
		return n.Sign == strings.ToUpper(Md5String(paramStr+"&key="+c.Key))
	}
	key := loadPublicKey(c.PublicKey)
	h := sha256.New()
	h.Write([]byte(paramStr))
	hashed := h.Sum(nil)
	sig, _ := base64.StdEncoding.DecodeString(n.Sign)

	return rsa.VerifyPKCS1v15(key, crypto.SHA256, hashed, sig) == nil
}

//对字典key进行排序
func sortKeys(m map[string]string) []string {
	keys := make([]string, 0)
	for key, _ := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

//生成url请求串
func buildQuery(params map[string]string) string {
	signParams := map[string]string{}
	for k, v := range params {
		if v == "" {
			continue
		}
		signParams[k] = v
	}

	keys := sortKeys(signParams)

	queryStr := ""
	for _, key := range keys {
		queryStr += key + "=" + signParams[key] + "&"
	}
	return queryStr[:len(queryStr)-1]
}
