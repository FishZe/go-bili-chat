package main

import (
	"fmt"
	bili "github.com/FishZe/go-bili-chat"
	handle "github.com/FishZe/go-bili-chat/handler"
)

func main() {
	// 可选: 修改日志等级 请删除import的注释
	// bili.ChangeLogLevel(log.DebugLevel)
	// 新建一个命令处理器
	h := bili.GetNewHandler()
	// 注册一个处理，将该直播间的弹幕消息绑定到这个函数
	// 弹幕
	h.AddOption(handle.CmdDanmuMsg, 0, func(event handle.MsgEvent) {
		fmt.Printf("[%v][弹幕] %v (%v): %v\n", event.RoomId, event.DanMuMsg.Data.Sender.Name, event.DanMuMsg.Data.Medal.MedalName, event.DanMuMsg.Data.Content)
	})
	// 超级留言
	h.AddOption(handle.CmdSuperChatMessage, 0, func(event handle.MsgEvent) {
		fmt.Printf("[%v][超级留言] %v: %v\n", event.RoomId, event.SuperChatMessage.Data.UserInfo.Uname, event.SuperChatMessage.Data.Message)
	})
	// 礼物
	h.AddOption(handle.CmdSendGift, 0, func(event handle.MsgEvent) {
		fmt.Printf("[%v][礼物] %v: %v\n", event.RoomId, event.SendGift.Data.Name, event.SendGift.Data.GiftName)
	})
	// 大航海
	h.AddOption(handle.CmdUserToastMsg, 0, func(event handle.MsgEvent) {
		fmt.Printf("[%v][大航海] %v: %v\n", event.RoomId, event.UserToastMsg.Data.Username, event.UserToastMsg.Data.RoleName)
	})
	h.AddOption(handle.CmdInteractWord, 0, func(event handle.MsgEvent) {
		fmt.Printf("[%v][进入直播间] %v (%v) \n", event.RoomId, event.InteractWord.Data.Name, event.InteractWord.Data.FansMedal.MedalName)
	})
	// 连接到直播间
	h.AddRoom(22637261)
	// 启动处理器
	h.Run()
}
