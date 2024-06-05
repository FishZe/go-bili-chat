package events

import (
	"github.com/FishZe/go-bili-chat/v2/serializer"
	log "github.com/sirupsen/logrus"
)

type SendGiftMsg struct {
	*BLiveEvent
	SendGift
}

type SendGiftHandler func(event *SendGiftMsg)

func (d SendGiftHandler) Cmd() string {
	return CmdSendGift
}

func (d SendGiftHandler) On(event *BLiveEvent) {
	var msg SendGift
	err := serializer.JsonCoder.Unmarshal(event.RawMessage, &msg)
	if err != nil {
		log.Warnf("failed to deserialize %s message %s", d.Cmd(), err)
		return
	}
	d(&SendGiftMsg{
		BLiveEvent: event,
		SendGift:   msg,
	})
}
