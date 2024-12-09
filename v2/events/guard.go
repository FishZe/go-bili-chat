package events

import (
	"github.com/FishZe/go-bili-chat/v2/serializer"
	log "github.com/sirupsen/logrus"
)

type GuardBuyEvent struct {
	*BLiveEvent
	GuardBuyMsg
}

type GuardBuyHandler func(event *GuardBuyEvent)

func (d GuardBuyHandler) Cmd() string {
	return CmdUserToastMsg
}

func (d GuardBuyHandler) On(event *BLiveEvent) {
	var msg GuardBuyMsg
	err := serializer.JsonCoder.Unmarshal(event.RawMessage, &msg)
	if err != nil {
		log.Warnf("failed to deserialize %s message %s", d.Cmd(), err)
		return
	}
	d(&GuardBuyEvent{
		BLiveEvent:  event,
		GuardBuyMsg: msg,
	})
}
