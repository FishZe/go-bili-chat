package go_bili_chat

import (
	"fmt"
	"github.com/FishZe/go-bili-chat/handler"
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestMain(m *testing.M) {
	// 可选: 修改日志等级 请删除import的注释
	ChangeLogLevel(log.DebugLevel)
	// 新建一个命令处理器
	h := GetNewHandler()
	// 注册一个处理，将该直播间的弹幕消息绑定到这个函数
	// 弹幕
	roomId := 21545805
	h.AddOption(handler.CmdDanmuMsg, roomId, func(event handler.MsgEvent) {
		fmt.Printf("[%v][弹幕] %v (%v): %v\n", event.RoomId, event.DanMuMsg.Data.Sender.Name, event.DanMuMsg.Data.Medal.MedalName, event.DanMuMsg.Data.Content)
	})
	// 超级留言
	h.AddOption(handler.CmdSuperChatMessage, roomId, func(event handler.MsgEvent) {
		fmt.Printf("[%v][超级留言] %v: %v\n", event.RoomId, event.SuperChatMessage.Data.UserInfo.Uname, event.SuperChatMessage.Data.Message)
	})
	// 礼物
	h.AddOption(handler.CmdSendGift, roomId, func(event handler.MsgEvent) {
		fmt.Printf("[%v][礼物] %v: %v\n", event.RoomId, event.SendGift.Data.Name, event.SendGift.Data.GiftName)
	})
	// 大航海
	h.AddOption(handler.CmdUserToastMsg, roomId, func(event handler.MsgEvent) {
		fmt.Printf("[%v][大航海] %v: %v\n", event.RoomId, event.UserToastMsg.Data.Username, event.UserToastMsg.Data.RoleName)
	})
	h.AddOption(handler.CmdInteractWord, roomId, func(event handler.MsgEvent) {
		fmt.Printf("[%v][进入直播间] %v (%v) \n", event.RoomId, event.InteractWord.Data.Name, event.InteractWord.Data.FansMedal.MedalName)
	})
	// 连接到直播间
	_ = h.AddRoom(roomId)
	// 启动处理器
	h.Run()
}
