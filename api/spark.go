package api

import (
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/law-lee/gpthub/pkg"
	"github.com/law-lee/gpthub/types"
)

// Spark 讯飞星火大模型
type Spark struct {
	Model *types.SparkArg
}

func (s *Spark) Send(req string) (string, error) {
	cfg := s.Model
	conn, err := pkg.NewWebsocketConn(cfg)
	if err != nil {
		return "", fmt.Errorf("new websocket conn: %v", err)
	}
	data := pkg.GenParams1(cfg.AppID, req)
	err = conn.WriteJSON(data)
	if err != nil {
		return "", fmt.Errorf("websocket write json: %v", err)
	}
	var answer = ""
	//获取返回的数据
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			return "", fmt.Errorf("websocket read message: %v", err)
		}

		var data map[string]interface{}
		err1 := json.Unmarshal(msg, &data)
		if err1 != nil {
			return "", fmt.Errorf("websocket parsing JSON: %v", err)
		}
		//fmt.Println(string(msg))
		//解析数据
		payload := data["payload"].(map[string]interface{})
		choices := payload["choices"].(map[string]interface{})
		header := data["header"].(map[string]interface{})
		code := header["code"].(float64)

		if code != 0 {
			logrus.Errorf("[Spark] close conn error: %s", data["payload"])
		}
		status := choices["status"].(float64)
		//fmt.Println(status)
		text := choices["text"].([]interface{})
		content := text[0].(map[string]interface{})["content"].(string)
		if status != 2 {
			answer += content
		} else {
			//fmt.Println("收到最终结果")
			answer += content
			//usage := payload["usage"].(map[string]interface{})
			//temp := usage["text"].(map[string]interface{})
			//totalTokens := temp["total_tokens"].(float64)
			//fmt.Println("total_tokens:", totalTokens)
			if err = conn.Close(); err != nil {
				logrus.Errorf("[Spark] close conn error: %v", err)
			}
			break
		}

	}
	//输出返回结果
	return answer, nil
}
