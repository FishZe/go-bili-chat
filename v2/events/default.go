package events

type DefaultEventHandler struct {
	cmd     string
	handler func(event *BLiveEvent)
}

func NewDefaultEventHandler(cmd string, handler func(event *BLiveEvent)) *DefaultEventHandler {
	return &DefaultEventHandler{cmd: cmd, handler: handler}
}

func (d DefaultEventHandler) Cmd() string {
	return CmdInteractWord
}

func (d DefaultEventHandler) On(event *BLiveEvent) {
	d.handler(event)
}
