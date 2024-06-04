package events

import (
	"github.com/FishZe/go-bili-chat/v2/serializer"
	log "github.com/sirupsen/logrus"
)

type UserToastEvent struct {
	*BLiveEvent
	UserToastMsg
}

type UserToastHandler func(event *UserToastEvent)

func (d UserToastHandler) Cmd() string {
	return CmdUserToastMsg
}

func (d UserToastHandler) On(event *BLiveEvent) {
	var msg UserToastMsg
	err := serializer.JsonCoder.Unmarshal(event.RawMessage, &msg)
	if err != nil {
		log.Warnf("failed to deserialize %s message %s", d.Cmd(), err)
		return
	}
	d(&UserToastEvent{
		BLiveEvent:   event,
		UserToastMsg: msg,
	})
}
