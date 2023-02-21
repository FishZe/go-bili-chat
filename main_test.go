package Go_BiliChat

import (
	"fmt"
	"github.com/FishZe/Go-BiliChat/handler"
	"testing"
)

func TestMain(m *testing.M) {
	h := GetNewHandler()
	h.AddOption(handler.CmdDanmuMsg, 80397, func(event handler.MsgEvent) {
		fmt.Printf("[%v] %v: %v\n", event.RoomId, event.DanMuMsg.Data.Sender.Name, event.DanMuMsg.Data.Content)
	}, "danmuHandler")
	h.AddRoom(80397)
	h.Run()
}
