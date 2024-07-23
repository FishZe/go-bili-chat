package events

import (
	"github.com/tidwall/gjson"
)

type DanMuEvent BLiveEvent

type DanMuEventHandler func(event *DanMuEvent)

func (d DanMuEventHandler) Cmd() string {
	return CmdDanmuMsg
}

func (d DanMuEventHandler) On(event *BLiveEvent) {
	d((*DanMuEvent)(event))
}

func (d *DanMuEvent) GetContent() string {
	return gjson.GetBytes(d.RawMessage, "info.1").String()
}

func (d *DanMuEvent) GetUID() int64 {
	return gjson.GetBytes(d.RawMessage, "info.2.0").Int()
}

func (d *DanMuEvent) GetUserName() string {
	return gjson.GetBytes(d.RawMessage, "info.2.1").String()
}

func (d *DanMuEvent) GetFace() string {
	return gjson.GetBytes(d.RawMessage, "info.0.15.user.base.face").String()
}

func (d *DanMuEvent) GetMedal() *FansMedal {
	medalInfo := gjson.GetBytes(d.RawMessage, "info.3")
	if !medalInfo.IsArray() {
		return nil
	}
	return &FansMedal{
		MedalLevel:   int(medalInfo.Get("0").Int()),
		MedalName:    medalInfo.Get("1").String(),
		AnchorUname:  medalInfo.Get("2").String(),
		AnchorRoomid: int(medalInfo.Get("3").Int()),
		TargetId:     medalInfo.Get("12").Int(),
	}
}

func (d *DanMuEvent) GetGuardLevel() int {
	return int(gjson.GetBytes(d.RawMessage, "info.3.10").Int())
}

func (d *DanMuEvent) GetTimestamp() int {
	return int(gjson.GetBytes(d.RawMessage, "info.9.st").Int())
}

func (d *DanMuEvent) GetMilliTimestamp() int64 {
	return gjson.GetBytes(d.RawMessage, "info.0.4").Int()
}
