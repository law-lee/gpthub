package pkg

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/gorilla/websocket"

	"github.com/law-lee/gpthub/types"
)

func NewWebsocketConn(spark *types.SparkArg) (*websocket.Conn, error) {
	d := websocket.Dialer{
		HandshakeTimeout: 5 * time.Second,
	}
	//握手并建立websocket 连接
	urlStr, err := assembleAuthUrl1(spark.WssDomain, spark.AppKey, spark.AppSecret)
	if err != nil {
		return nil, err
	}
	conn, _, err := d.Dial(urlStr, nil)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// assembleAuthUrl1 创建鉴权url  apikey 即 hmac username
func assembleAuthUrl1(hosturl string, apiKey, apiSecret string) (string, error) {
	if hosturl == "" || apiKey == "" || apiSecret == "" {
		return "", fmt.Errorf("model args is empty.")
	}
	ul, err := url.Parse(hosturl)
	if err != nil {
		return "", fmt.Errorf("url parse err: %v", err)
	}
	//签名时间
	date := time.Now().UTC().Format(time.RFC1123)
	//date = "Tue, 28 May 2019 09:10:42 MST"
	//参与签名的字段 host ,date, request-line
	signString := []string{"host: " + ul.Host, "date: " + date, "GET " + ul.Path + " HTTP/1.1"}
	//拼接签名字符串
	sign := strings.Join(signString, "\n")
	// fmt.Println(sign)
	//签名结果
	sha := HmacWithShaTobase64("hmac-sha256", sign, apiSecret)
	// fmt.Println(sha)
	//构建请求参数 此时不需要urlencoding
	authUrl := fmt.Sprintf("hmac username=\"%s\", algorithm=\"%s\", headers=\"%s\", signature=\"%s\"", apiKey,
		"hmac-sha256", "host date request-line", sha)
	//将请求参数使用base64编码
	authorization := base64.StdEncoding.EncodeToString([]byte(authUrl))

	v := url.Values{}
	v.Add("host", ul.Host)
	v.Add("date", date)
	v.Add("authorization", authorization)
	//将编码后的字符串url encode后添加到url后面
	callurl := hosturl + "?" + v.Encode()
	return callurl, nil
}

func HmacWithShaTobase64(algorithm, data, key string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(data))
	encodeData := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(encodeData)
}

// GenParams1 生成参数
func GenParams1(appid, question string) map[string]interface{} { // 根据实际情况修改返回的数据结构和字段名

	messages := []types.Message{
		{Role: "user", Content: question},
	}

	data := map[string]interface{}{ // 根据实际情况修改返回的数据结构和字段名
		"header": map[string]interface{}{ // 根据实际情况修改返回的数据结构和字段名
			"app_id": appid, // 根据实际情况修改返回的数据结构和字段名
		},
		"parameter": map[string]interface{}{ // 根据实际情况修改返回的数据结构和字段名
			"chat": map[string]interface{}{ // 根据实际情况修改返回的数据结构和字段名
				"domain":      "general",    // 根据实际情况修改返回的数据结构和字段名
				"temperature": float64(0.8), // 根据实际情况修改返回的数据结构和字段名
				"top_k":       int64(6),     // 根据实际情况修改返回的数据结构和字段名
				"max_tokens":  int64(2048),  // 根据实际情况修改返回的数据结构和字段名
				"auditing":    "default",    // 根据实际情况修改返回的数据结构和字段名
			},
		},
		"payload": map[string]interface{}{ // 根据实际情况修改返回的数据结构和字段名
			"message": map[string]interface{}{ // 根据实际情况修改返回的数据结构和字段名
				"text": messages, // 根据实际情况修改返回的数据结构和字段名
			},
		},
	}
	return data // 根据实际情况修改返回的数据结构和字段名
}
