package events

type BLiveEventHandler interface {
	On(event *BLiveEvent)
	Cmd() string
}

type BLiveEvent struct {
	Cmd        string
	RoomId     int
	RawMessage []byte
}
