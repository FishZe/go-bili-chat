package main

import (
	"encoding/json"
	"fmt"
	bili "github.com/FishZe/go-bili-chat"
	handle "github.com/FishZe/go-bili-chat/handler"
	log "github.com/sirupsen/logrus"
)

type Json struct{}

func (j *Json) Unmarshal(data []byte, v interface{}) error {
	log.Println(string(data))
	return json.Unmarshal(data, v)
}

func (j *Json) Marshal(v interface{}) ([]byte, error) {
	log.Println(v)
	return json.Marshal(v)
}

func main() {
	// 可选: 修改日志等级 请删除import的注释
	// bili.ChangeLogLevel(log.DebugLevel)
	bili.SetJsonCoder(&Json{})
	bili.SetClientPriorityMode(bili.DelayClientPriority) // 优先使用延迟低的
	// bili.SetClientPriorityMode(bili.NoCDNClientPriority) // 优先使用无CDN
	// 新建一个命令处理器
	h := bili.GetNewHandler()
	// 注册一个处理，将该直播间的弹幕消息绑定到这个函数
	// 弹幕
	h.AddOption(handle.CmdDanmuMsg, 21545805, func(event handle.MsgEvent) {
		fmt.Printf("[%v][弹幕] %v (%v): %v\n", event.RoomId, event.DanMuMsg.Data.Sender.Name, event.DanMuMsg.Data.Medal.MedalName, event.DanMuMsg.Data.Content)
	}, "弹幕")
	// 超级留言
	h.AddOption(handle.CmdSuperChatMessage, 21545805, func(event handle.MsgEvent) {
		fmt.Printf("[%v][超级留言] %v: %v\n", event.RoomId, event.SuperChatMessage.Data.UserInfo.Uname, event.SuperChatMessage.Data.Message)
	}, "超级留言")
	// 礼物
	h.AddOption(handle.CmdSendGift, 21545805, func(event handle.MsgEvent) {
		fmt.Printf("[%v][礼物] %v: %v\n", event.RoomId, event.SendGift.Data.Name, event.SendGift.Data.GiftName)
	}, "礼物")
	// 大航海
	h.AddOption(handle.CmdUserToastMsg, 21545805, func(event handle.MsgEvent) {
		fmt.Printf("[%v][大航海] %v: %v\n", event.RoomId, event.UserToastMsg.Data.Username, event.UserToastMsg.Data.RoleName)
	}, "大航海")
	h.AddOption(handle.CmdInteractWord, 21545805, func(event handle.MsgEvent) {
		fmt.Printf("[%v][进入直播间] %v (%v) \n", event.RoomId, event.InteractWord.Data.Name, event.InteractWord.Data.FansMedal.MedalName)
	}, "进入直播间")
	// 连接到直播间
	_ = h.AddRoom(21545805)
	// 启动处理器
	h.Run()
}
