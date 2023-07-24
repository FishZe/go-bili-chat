package handler

import (
	log "github.com/sirupsen/logrus"
	"reflect"
	"strconv"
	"sync"
)

var cmdFieldId sync.Map
var cmdSetFunc sync.Map

func init() {
	t := reflect.TypeOf(MsgEvent{})
	for i := 0; i < t.NumField(); i++ {
		cmdFieldId.Store(t.Field(i).Name, i)
	}
	cmdSetFunc.Store(CmdDanmuMsg, struct{}{})
	cmdSetFunc.Store(CmdNoticeMsg, struct{}{})
	cmdSetFunc.Store(CmdSuperChatMessage, struct{}{})
	cmdSetFunc.Store(CmdHeartBeatReply, struct{}{})
}

// SetDanMuMsg 设置弹幕消息
// 该消息为list结构, 部分字段含义未知, 因此目前只有部分内容
// TODO: 完善更多字段
func (*Handler) SetDanMuMsg(msg map[string]interface{}) (m MsgEvent) {
	danMu := DanMuMsg{}
	danMu.Cmd = CmdDanmuMsg
	danMuMsg := make(map[string]interface{}, 0)
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &danMuMsg); err == nil {
		danMu.Data.Content = danMuMsg["info"].([]interface{})[1].(string)
		danMu.Data.SendTimeStamp = int(danMuMsg["info"].([]interface{})[9].(map[string]interface{})["ts"].(float64))
		danMu.Data.SenderEnterRoomTimeStamp = int(danMuMsg["info"].([]interface{})[0].([]interface{})[4].(float64))
		danMu.Data.SendMillionTimeStamp = int64(danMuMsg["info"].([]interface{})[0].([]interface{})[5].(float64))
		danMu.Data.Sender.Uid = int64(danMuMsg["info"].([]interface{})[2].([]interface{})[0].(float64))
		danMu.Data.Sender.Name = danMuMsg["info"].([]interface{})[2].([]interface{})[1].(string)
		// 部分情况下, 弹幕发送者未佩戴牌子, 需要判断
		if len(danMuMsg["info"].([]interface{})[3].([]interface{})) != 0 {
			danMu.Data.Medal.GuardLevel = int(danMuMsg["info"].([]interface{})[3].([]interface{})[0].(float64))
			danMu.Data.Medal.MedalName = danMuMsg["info"].([]interface{})[3].([]interface{})[1].(string)
			danMu.Data.Medal.TargetID = int(danMuMsg["info"].([]interface{})[3].([]interface{})[11].(float64))
			danMu.Data.Medal.AnchorRoomId = int(danMuMsg["info"].([]interface{})[3].([]interface{})[3].(float64))
		}
		m = MsgEvent{Cmd: CmdDanmuMsg, DanMuMsg: &danMu, RoomId: msg["RoomId"].(int)}
	}
	return
}

// SetNoticeMsg 可能为系统消息
// TODO: 尝试优化
func (*Handler) SetNoticeMsg(msg map[string]interface{}) (m MsgEvent) {
	noticeMsg := NoticeMsg{}
	notice := make(map[string]interface{}, 0)
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &notice); err == nil {
		// 这个字段很奇怪, 类型不确定, 需要特判
		switch notice["real_roomid"].(type) {
		case float64:
			notice["real_roomid"] = strconv.FormatFloat(notice["real_roomid"].(float64), 'f', -1, 64)
		case int:
			notice["real_roomid"] = strconv.Itoa(notice["real_roomid"].(int))
		}
		if dataJson, err := JsonCoder.Marshal(notice); err == nil {
			if err = JsonCoder.Unmarshal(dataJson, &noticeMsg); err == nil {
				m = MsgEvent{Cmd: CmdNoticeMsg, NoticeMsg: &noticeMsg, RoomId: msg["RoomId"].(int)}
			}
		}
	}
	return
}

// SetSuperChatMessage 超级留言
// TODO: 尝试优化
func (*Handler) SetSuperChatMessage(msg map[string]interface{}) (m MsgEvent) {
	superChatMsg := SuperChatMessage{}
	superChatMsg.Cmd = CmdSuperChatMessage
	sc := make(map[string]interface{}, 0)
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &sc); err == nil {
		// id 和 uid 类型不确定, 需要特判
		switch sc["data"].(map[string]interface{})["id"].(type) {
		case float64:
			sc["data"].(map[string]interface{})["id"] = strconv.FormatFloat(sc["data"].(map[string]interface{})["id"].(float64), 'f', -1, 64)
		case int:
			sc["data"].(map[string]interface{})["id"] = strconv.Itoa(sc["data"].(map[string]interface{})["id"].(int))
		}
		switch sc["data"].(map[string]interface{})["uid"].(type) {
		case float64:
			sc["data"].(map[string]interface{})["uid"] = strconv.FormatFloat(sc["data"].(map[string]interface{})["uid"].(float64), 'f', -1, 64)
		case int:
			sc["data"].(map[string]interface{})["uid"] = strconv.Itoa(sc["data"].(map[string]interface{})["uid"].(int))
		}
		if dataJson, err := JsonCoder.Marshal(sc["data"]); err == nil {
			if err = JsonCoder.Unmarshal(dataJson, &superChatMsg.Data); err == nil {
				m = MsgEvent{Cmd: CmdSuperChatMessage, SuperChatMessage: &superChatMsg, RoomId: msg["RoomId"].(int)}
			}
		}
	}
	return
}

func (*Handler) SetHeartBeatReply(msg map[string]interface{}) (m MsgEvent) {
	heartBeatReply := HeartBeatReply{Sum: msg["msg"].(int)}
	m = MsgEvent{Cmd: CmdHeartBeatReply, RoomId: msg["RoomId"].(int), HeartBeatReply: &heartBeatReply}
	return
}

// DefaultCmd 默认处理函数
func (*Handler) DefaultCmd(msg map[string]interface{}) (m MsgEvent) {
	cmd := msg["cmd"].(string)
	m = MsgEvent{Cmd: cmd, RoomId: msg["RoomId"].(int)}
	if id, ok := cmdFieldId.Load(CmdName[cmd]); !ok {
		return
	} else {
		v := reflect.ValueOf(&m).Elem()
		x := v.Field(id.(int)).Interface().(cmdInterface).New()
		if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), x); err == nil {
			v.Field(id.(int)).Set(reflect.ValueOf(x))
		} else {
			log.Warnf("cmd: %s, msg: %s, err: %s", cmd, msg["msg"].(string), err.Error())
		}
	}
	return
}
