package go_bilichat_core

import (
	"github.com/FishZe/go_bilichat_core/client"
	"github.com/FishZe/go_bilichat_core/handler"
)

type Handler struct {
	Handler handler.Handler
	rooms   map[int]LiveRoom
}

type LiveRoom struct {
	RoomId int
	Client client.Client
}

func GetNewHandler() Handler {
	h := Handler{}
	h.rooms = make(map[int]LiveRoom)
	h.Handler.DoFunc = make(map[string]map[int][]func(event handler.MsgEvent), 0)
	h.Handler.CmdChan = make(chan map[string]interface{}, 100000)
	return h
}

func (h *Handler) AddOption(Cmd string, RoomId int, Do func(event handler.MsgEvent)) {
	h.Handler.AddOption(Cmd, RoomId, Do)
}

func (h *Handler) AddRoom(roomId int) {
	if _, ok := h.rooms[roomId]; ok {
		return
	}
	room := LiveRoom{}
	room.RoomId = roomId
	room.Client.RoomId = room.RoomId
	room.Client.BiliChat(h.Handler.CmdChan)
	h.rooms[room.RoomId] = room
}

func (h *Handler) DelRoom(RoomId int) {
	if _, ok := h.rooms[RoomId]; !ok {
		return
	}
	h.Handler.DelRoomOption(RoomId)
	c := h.rooms[RoomId].Client
	c.Close()
	delete(h.rooms, RoomId)
}

func (h *Handler) Run() {
	h.Handler.CmdHandler()
}
