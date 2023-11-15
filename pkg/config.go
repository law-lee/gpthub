package pkg

import (
	"encoding/json"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/law-lee/gpthub/types"
)

var config *types.GPTModelArg
var once sync.Once

// LoadConfig 加载配置
func LoadConfig() *types.GPTModelArg {
	once.Do(func() {
		config = &types.GPTModelArg{
			Spark:   &types.SparkArg{},
			Openapi: &types.OpenapiArg{},
		}
		// 判断配置文件是否存在，存在直接JSON读取
		_, err := os.Stat("config.json")
		if err == nil {
			f, err := os.Open("config.json")
			if err != nil {
				logrus.Fatalf("open config err: %v", err)
				return
			}
			defer f.Close()
			encoder := json.NewDecoder(f)
			err = encoder.Decode(config)
			if err != nil {
				logrus.Fatalf("decode config err: %v", err)
				return
			}
		}
		// 有环境变量使用环境变量
		// openapi
		ApiKey := os.Getenv("APIKEY")
		AutoPass := os.Getenv("AUTO_PASS")
		SessionTimeout := os.Getenv("SESSION_TIMEOUT")
		Model := os.Getenv("MODEL")
		MaxTokens := os.Getenv("MAX_TOKENS")
		Temperature := os.Getenv("TEMPREATURE")
		ReplyPrefix := os.Getenv("REPLY_PREFIX")
		SessionClearToken := os.Getenv("SESSION_CLEAR_TOKEN")
		// spark api
		WssDomain := os.Getenv("WSSDOMAIN")
		AppID := os.Getenv("APPID")
		AppSecret := os.Getenv("APPSECRET")
		AppKey := os.Getenv("APPKEY")
		if WssDomain != "" {
			config.Spark.WssDomain = WssDomain
		}
		if AppID != "" {
			config.Spark.AppID = AppID
		}
		if AppSecret != "" {
			config.Spark.AppSecret = AppSecret
		}
		if AppKey != "" {
			config.Spark.AppKey = AppKey
		}
		if ApiKey != "" {
			config.Openapi.APIKey = ApiKey
		}
		if AutoPass != "" {
			config.Openapi.AutoPass = true
		}
		if SessionTimeout != "" {
			duration, _ := time.ParseDuration(SessionTimeout)
			config.Openapi.SessionTimeout = duration
		}
		if Model != "" {
			config.Openapi.Model = Model
		}
		if MaxTokens != "" {
			max, _ := strconv.Atoi(MaxTokens)
			config.Openapi.MaxTokens = uint(max)
		}
		if Temperature != "" {
			temp, _ := strconv.ParseFloat(Temperature, 64)
			config.Openapi.Temperature = temp
		}
		if ReplyPrefix != "" {
			config.Openapi.ReplyPrefix = ReplyPrefix
		}
		if SessionClearToken != "" {
			config.Openapi.SessionClearToken = SessionClearToken
		}
	})

	return config
}
