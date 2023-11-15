package types

import (
	"time"
)

// GPTModel 1 openai大模型 2 讯飞星火大模型
type GPTModel int

const (
	_ GPTModel = iota
	OpenAIModel
	SparkModel
)

type GPTModelArg struct {
	Spark   *SparkArg   `json:"spark"`
	Openapi *OpenapiArg `json:"openapi"`
}
type SparkArg struct {
	WssDomain string `json:"wss_domain"`
	AppID     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
	AppKey    string `json:"app_key"`
}
type OpenapiArg struct {
	APIKey            string        `json:"api_key"`
	AutoPass          bool          `json:"auto_pass"`
	SessionTimeout    time.Duration `json:"session_timeout"`
	MaxTokens         uint          `json:"max_tokens"`
	Model             string        `json:"model"`
	Temperature       float64       `json:"temperature"`
	ReplyPrefix       string        `json:"reply_prefix"`
	SessionClearToken string        `json:"session_clear_token"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
