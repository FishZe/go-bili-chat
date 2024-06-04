package events

import (
	"github.com/FishZe/go-bili-chat/v2/serializer"
	log "github.com/sirupsen/logrus"
)

type SuperChatEvent struct {
	*BLiveEvent
	SuperChatMessage
}

type SuperChatHandler func(event *SuperChatEvent)

func (d SuperChatHandler) Cmd() string {
	return CmdSuperChatMessage
}

func (d SuperChatHandler) On(event *BLiveEvent) {
	var msg SuperChatMessage
	err := serializer.JsonCoder.Unmarshal(event.RawMessage, &msg)
	if err != nil {
		log.Warnf("failed to deserialize %s message: %s", d.Cmd(), err)
		return
	}
	d(&SuperChatEvent{
		BLiveEvent:       event,
		SuperChatMessage: msg,
	})
}
