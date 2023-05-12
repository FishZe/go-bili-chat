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
	// 修改Json解析器
	bili.SetJsonCoder(&Json{})
	// 修改客户端优先级模式
	bili.SetClientPriorityMode(bili.DelayClientPriority) // 优先使用延迟低的
	// bili.SetClientPriorityMode(bili.NoCDNClientPriority) // 优先使用无CDN
	// 新建一个命令处理器
	h := bili.GetNewHandler()
	// 注册一个处理，将该直播间的弹幕消息绑定到这个函数
	RoomId := 23015128
	// 弹幕
	h.AddOption(handle.CmdDanmuMsg, RoomId, func(event handle.MsgEvent) {
		fmt.Printf("[%v][弹幕] %v (%v): %v\n", event.RoomId, event.DanMuMsg.Data.Sender.Name, event.DanMuMsg.Data.Medal.MedalName, event.DanMuMsg.Data.Content)
	})
	// 超级留言
	h.AddOption(handle.CmdSuperChatMessage, RoomId, func(event handle.MsgEvent) {
		fmt.Printf("[%v][超级留言] %v: %v\n", event.RoomId, event.SuperChatMessage.Data.UserInfo.Uname, event.SuperChatMessage.Data.Message)
	})
	// 礼物
	h.AddOption(handle.CmdSendGift, RoomId, func(event handle.MsgEvent) {
		fmt.Printf("[%v][礼物] %v: %v\n", event.RoomId, event.SendGift.Data.Name, event.SendGift.Data.GiftName)
	})
	// 大航海
	h.AddOption(handle.CmdUserToastMsg, RoomId, func(event handle.MsgEvent) {
		fmt.Printf("[%v][大航海] %v: %v\n", event.RoomId, event.UserToastMsg.Data.Username, event.UserToastMsg.Data.RoleName)
	})
	// 进入直播间
	h.AddOption(handle.CmdInteractWord, RoomId, func(event handle.MsgEvent) {
		fmt.Printf("[%v][进入直播间] %v (%v) \n", event.RoomId, event.InteractWord.Data.Name, event.InteractWord.Data.FansMedal.MedalName)
	})
	// 心跳
	heart := h.AddOption(handle.CmdHeartBeatReply, RoomId, func(event handle.MsgEvent) {
		fmt.Printf("[%v][心跳包回复] %d \n", event.RoomId, event.HeartBeatReply.Sum)
	})
	// 连接到直播间
	_ = h.AddRoom(RoomId)
	// 删掉一个直播间
	// _ = h.DelRoom(RoomId)
	// 删掉一个处理
	h.DelOption(heart)
	// 获取连接个数
	fmt.Println("已连接到: ", h.CountRoom(), "个直播间")
	// 启动处理器
	h.Run()
	// 如果连接个数过多, 建议把 h.Run() 改为 go h.Run()并移动到添加直播间前
}
