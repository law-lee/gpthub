# gpthub
a hub for various gpt model

# 配置
支持文件和环境变量读取

本地config.json文件示例
```
{
  "spark": {
    "wss_domain": "wss://aichat.xf-yun.com/v1/chat",
    "app_id": "",
    "app_secret":"",
    "app_key":""
  },
  "openapi": {
    "api_key": "",
    "auto_pass": true,
    "session_timeout": 60,
    "max_tokens": 1024,
    "model": "text-davinci-003",
    "temperature": 1,
    "reply_prefix": "来自机器人回复：",
    "session_clear_token": "清空会话"
  }
}
```