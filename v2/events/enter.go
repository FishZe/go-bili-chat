package events

import (
	"github.com/FishZe/go-bili-chat/v2/serializer"
	log "github.com/sirupsen/logrus"
)

type InteractWordEvent struct {
	*BLiveEvent
	InteractWord
}

type InteractWordHandler func(event *InteractWordEvent)

func (d InteractWordHandler) Cmd() string {
	return CmdInteractWord
}

func (d InteractWordHandler) On(event *BLiveEvent) {
	var msg InteractWord
	err := serializer.JsonCoder.Unmarshal(event.RawMessage, &msg)
	if err != nil {
		log.Warnf("failed to deserialize %s message %s", d.Cmd(), err)
		return
	}
	d(&InteractWordEvent{
		BLiveEvent:   event,
		InteractWord: msg,
	})
}
