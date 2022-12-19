package main

import (
	"Go-BiliChat-Core/client"
	"fmt"
	"time"
)

func main() {
	c := client.Client{RoomId: 5169315}
	msgHandler, err := c.BiliChat()
	if err != nil {
		fmt.Println(err)
	}
	for {
		select {
		case _, ok := <-msgHandler.CmdChan:
			if ok {
				continue
			}
		default:
			time.Sleep(10 * time.Microsecond)
			continue
		}

	}
}
