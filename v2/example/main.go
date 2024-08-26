package main

import (
	"encoding/json"
	"fmt"
	bili "github.com/FishZe/go-bili-chat/v2"
	"github.com/FishZe/go-bili-chat/v2/events"
	"github.com/FishZe/go-bili-chat/v2/utils"
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

type MyEventHandler func(data string)

func (m MyEventHandler) On(event *events.BLiveEvent) {
	//解析事件
	m(fmt.Sprintf("[%v][自定义解析]弹幕数据的长度为: %d", event.RoomId, len(event.RawMessage)))
}

func (m MyEventHandler) Cmd() string {
	//指定事件ID
	return events.CmdDanmuMsg
}

func init() {
	// 可选: 修改日志等级
	// bili.ChangeLogLevel(log.DebugLevel)
	// 修改Json解析器
	// bili.SetJsonCoder(&Json{})

	/*
		修改客户端优先级模式
		bili.SetClientPriorityMode(bili.DelayClientPriority) // 优先使用延迟低的
		bili.SetClientPriorityMode(bili.NoCDNClientPriority) // 优先使用无CDN
	*/

	// 修改连接时认证包的UID, 默认为1
	// bili.SetUID(208259)

	/*
		修改连接时的header
		header := http.Header{
			"User-Agent": []string{"bilibili-live-tools"},
		}
		bili.SetHeader(header)
	*/

}

func main() {
	// 新建一个命令处理器
	h := bili.GetNewHandler()
	// 注册一个处理，将该直播间的弹幕消息绑定到这个函数
	RoomId := 22344968

	bili.SetUID(45700000)
	bili.SetHeaderCookie("DedeUserID__ckMd5=a08xxx; SESSDATA=6f867xxx; bili_jct=ab4af4xxx; DedeUserID=34xxxx")
	// 弹幕
	h.AddOption(RoomId, events.DanMuEventHandler(func(event *events.DanMuEvent) {
		fmt.Printf("[%v][弹幕] %v (%v): %v\n", event.RoomId, event.GetUserName(), event.GetMedal().MedalName, event.GetContent())
	}))
	//// 超级留言
	h.AddOption(RoomId, events.SuperChatHandler(func(event *events.SuperChatEvent) {
		fmt.Printf("[%v][超级留言] %v: %v\n", event.RoomId, event.Data.UserInfo.Uname, event.Data.Message)
	}))
	//// 礼物
	h.AddOption(RoomId, events.SendGiftHandler(func(event *events.SendGiftMsg) {
		fmt.Printf("[%v][礼物] %v: %v\n", event.RoomId, event.Data.Name, event.Data.GiftName)
	}))
	//// 大航海
	h.AddOption(RoomId, events.GuardBuyHandler(func(event *events.GuardBuyEvent) {
		fmt.Printf("[%v][大航海] %v: %v\n", event.RoomId, event.Data.Username, utils.GetGuardName(event.Data.GuardLevel))
	}))
	//// 进入直播间
	h.AddOption(RoomId, events.InteractWordHandler(func(event *events.InteractWordEvent) {
		fmt.Printf("[%v][进入直播间] %v (%v) \n", event.RoomId, event.Data.Uname, event.Data.FansMedal.MedalName)
	}))
	//// 心跳
	heart := h.AddOption(RoomId, events.HeartBeatHandler(func(event *events.HeartBeatEvent) {
		fmt.Printf("[%v][心跳包回复] %d \n", event.RoomId, event.Hot)
	}))
	//// 原始事件
	h.AddOption(RoomId, events.NewDefaultEventHandler(events.CmdDanmuMsg, func(event *events.BLiveEvent) {
		fmt.Printf("[%v][原始弹幕] %s\n", event.RoomId, event.RawMessage)
	}))
	//// 自定义解析器
	h.AddOption(RoomId, MyEventHandler(func(data string) {
		fmt.Println(data)
	}))
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
	// 并循环启动多个处理器并行处理，否则可能消息阻塞导致丢失部分消息
}
