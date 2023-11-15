package main

import (
	"fmt"

	"github.com/law-lee/gpthub/service"
)

func main() {
	text := "你是谁，可以干什么"
	gpt, err := service.NewModel(2)
	if err != nil {
		fmt.Printf("new gpt model err: %v", err)
		return
	}
	r, err := gpt.Send(text)
	if err != nil {
		fmt.Printf("request gpt err: %v", err)
	}
	fmt.Println(r)
}
