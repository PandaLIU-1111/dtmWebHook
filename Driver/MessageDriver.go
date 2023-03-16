package Driver

import (
	"crypto/hmac"
	"crypto/sha256"
	"dtmWebHook.com/m/v2/Config"
	"encoding/base64"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/url"
	"strconv"
	"time"
)

type MessageDriver interface {
	Send(config Config.NotifyConfig, content string)
}

type DingTalkMessageDriver string

func (m *DingTalkMessageDriver) Send(config Config.NotifyConfig, content string) {
	// Create a Resty Client
	client := resty.New()

	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	stringToSign := timestamp + "\n" + config.Secret
	byteSecret := []byte(config.Secret)
	h := hmac.New(sha256.New, byteSecret)
	h.Write([]byte(stringToSign))
	signData := h.Sum(nil)
	sign := url.QueryEscape(base64.StdEncoding.EncodeToString(signData))

	resp, err := client.R().
		SetQueryParams(map[string]string{
			"access_token": config.Token,
			"sign":         sign,
			"timestamp":    timestamp,
		}).
		SetBody(map[string]interface{}{
			"msgtype": "markdown",
			"markdown": map[string]string{
				"title": "DTM消息",
				"text":  content,
			},
		}).
		SetHeader("Content-Type", "Application/json").
		EnableTrace().
		Post("https://oapi.dingtalk.com/robot/send")

	// Explore response object
	fmt.Println("Response Info:")
	fmt.Println("  URL      :", resp.Request.URL)
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()
}
