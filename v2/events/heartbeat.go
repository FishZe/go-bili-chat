package events

import "encoding/binary"

type HeartBeatEvent struct {
	*BLiveEvent
	Hot uint32
}

type HeartBeatHandler func(event *HeartBeatEvent)

func (d HeartBeatHandler) Cmd() string {
	return CmdHeartBeatReply
}

func (d HeartBeatHandler) On(event *BLiveEvent) {
	d(&HeartBeatEvent{
		BLiveEvent: event,
		Hot:        binary.BigEndian.Uint32(event.RawMessage),
	})
}
