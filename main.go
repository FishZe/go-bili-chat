package Go_BiliChat

import (
	"github.com/FishZe/Go-BiliChat/client"
	"github.com/FishZe/Go-BiliChat/handler"
	"sync"
)

type Handler struct {
	Handler handler.Handler
	rooms   sync.Map
}

type LiveRoom struct {
	RoomId int
	Client client.Client
}

func GetNewHandler() Handler {
	h := Handler{}
	h.Handler.DoFunc = make(map[string]map[int][]func(event handler.MsgEvent), 0)
	h.Handler.CmdChan = make(chan map[string]interface{}, 100000)
	return h
}

func (h *Handler) AddOption(Cmd string, RoomId int, Do func(event handler.MsgEvent)) {
	h.Handler.AddOption(Cmd, RoomId, Do)
}

func (h *Handler) AddRoom(roomId int) {
	if _, ok := h.rooms.Load(roomId); ok {
		return
	}
	room := LiveRoom{}
	room.RoomId = roomId
	room.Client.RoomId = room.RoomId
	room.Client.BiliChat(h.Handler.CmdChan)
	h.rooms.Store(room.RoomId, room)
}

func (h *Handler) DelRoom(RoomId int) {
	if _, ok := h.rooms.Load(RoomId); !ok {
		return
	}
	h.Handler.DelRoomOption(RoomId)
	if c, ok := h.rooms.Load(RoomId); ok {
		cl := c.(LiveRoom).Client
		cl.Close()
		h.rooms.Delete(RoomId)
	}
}

func (h *Handler) Run() {
	h.Handler.CmdHandler()
}
