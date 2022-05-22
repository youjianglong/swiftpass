package swiftpass

import (
	"swiftpass/pkg/swiftpass/request"
	"swiftpass/pkg/swiftpass/response"
	"testing"
	"time"
)

func TestClent(t *testing.T) {
	client := NewDefaultClient("ZZx7owPBgtBYFT6vrtAocJIdDimLmJAV")
	req := request.UnifiedTradeNativeRequest{}
	req.MchId = "1030051268"
	req.OutTradeNo = "test_" + time.Now().Format("060102150405")
	req.Body = "一份面条"
	req.TotalFee = 600
	req.MchCreateIp = "127.0.0.1"
	req.NotifyUrl = "http://127.0.0.1"
	req.NonceStr = time.Now().Format("060102150405")

	resp := response.UnifiedTradeNativeResponse{}
	err := client.Execute(req, &resp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(client.RequestXml)
	t.Log(client.ResponseXml)
	t.Log(resp)
}
