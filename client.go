package swiftpass

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/xml"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/youjianglong/swiftpass/request"
	"github.com/youjianglong/swiftpass/response"
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

var (
	ErrNoRsaKey      = errors.New("not rsa key")
	ErrRsaPrivateKey = errors.New("failed to load rsa private key")
	ErrRsaPublicKey  = errors.New("failed to load rsa public key")
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
func (swiftClient *SwiftClient) LoadRsaKeyFile(priveteKey, publicKey string) error {
	readPrivateByte, readErr := ioutil.ReadFile(priveteKey)
	if readErr != nil || len(readPrivateByte) == 0 {
		return ErrRsaPrivateKey
	}
	swiftClient.PrivateKey = readPrivateByte

	readPublicByte, readErr := ioutil.ReadFile(publicKey)
	if readErr != nil || len(readPublicByte) == 0 {
		return ErrRsaPublicKey
	}
	swiftClient.PublicKey = readPublicByte

	return nil
}

//加载rsa秘钥文件
func (swiftClient *SwiftClient) LoadRsaKeyByte(priveteKey, publicKey []byte) {
	swiftClient.PrivateKey = priveteKey
	swiftClient.PublicKey = publicKey
}

//生成md5签名
func (swiftClient *SwiftClient) GenerateSign(req request.Request) string {
	params := XmlObject2Map(req)
	delete(params, "sign")
	paramStr := BuildQueryString(params)
	paramStr += "&key=" + swiftClient.Key
	return strings.ToUpper(Md5String(paramStr))
}

//生成ras签名
func (swiftClient *SwiftClient) GenerateSignRsa(req request.Request) (string, error) {
	if len(swiftClient.PrivateKey) == 0 || len(swiftClient.PublicKey) == 0 {
		return "", ErrNoRsaKey
	}

	params := XmlObject2Map(req)
	delete(params, "sign")
	data := []byte(BuildQueryString(params))

	h := sha256.New()
	h.Write(data)
	hashed := h.Sum(nil)
	privateKey, err := DecodePrivateKey(swiftClient.PrivateKey)
	newData, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(newData), nil
}

//提交支付
func (c *SwiftClient) Execute(req request.Request, swpResp response.Response) error {
	var sign string
	var err error

	req = req.Full()
	if c.SignType == "" || c.SignType == SignTypeMd5 {
		sign = c.GenerateSign(req)
	} else {
		sign, err = c.GenerateSignRsa(req)
	}

	if err != nil {
		return err
	}

	data := req.Encode(sign)
	c.RequestXml = string(data)

	httpReq, _ := http.NewRequest("POST", c.GetWay, strings.NewReader(string(data)))
	httpReq.Header.Set("Accept", "text/xml,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	httpReq.Header.Set("Charset", "UTF-8")

	res, err := c.client.Do(httpReq)
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
func (c *SwiftClient) Verify(body io.Reader) bool {
	params := Xml2Map(body)
	sign := params["sign"]
	if sign == "" {
		return false
	}

	delete(params, "sign")

	paramStr := BuildQueryString(params)
	signType := params["sign_type"]
	if signType == SignTypeMd5 { //MD5签名
		return sign == strings.ToUpper(Md5String(paramStr+"&key="+c.Key))
	}
	key, err := DecodePublicKey(c.PublicKey)
	if err != nil {
		return false
	}
	h := sha256.New()
	h.Write([]byte(paramStr))
	hashed := h.Sum(nil)
	sig, _ := base64.StdEncoding.DecodeString(sign)

	return rsa.VerifyPKCS1v15(key, crypto.SHA256, hashed, sig) == nil
}
