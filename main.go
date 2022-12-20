package main

import (
	"Go-BiliChat-Core/Handler"
	"Go-BiliChat-Core/client"
	"fmt"
)

func main() {
	c := client.Client{RoomId: 23805029}
	msgHandler, err := c.BiliChat()
	if err != nil {
		fmt.Println(err)
	}
	h := Handler.Handler{CmdChan: msgHandler.CmdChan}
	h.CmdHandler()
}
