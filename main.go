package go_bilichat_core

import (
	"github.com/FishZe/go_bilichat_core/client"
	"github.com/FishZe/go_bilichat_core/handler"
)

type Handler struct {
	Handler handler.Handler
}

type LiveRoom struct {
	RoomId int
	Client client.Client
}

func GetNewHandler() Handler {
	h := Handler{}
	h.Handler.DoFunc = make(map[string]map[int][]func(event handler.MsgEvent), 0)
	h.Handler.CmdChan = make(chan map[string]interface{}, 100)
	return h
}

func (h *Handler) AddOption(Cmd string, RoomId int, Do func(event handler.MsgEvent)) {
	h.Handler.AddOption(Cmd, RoomId, Do)
}

func (h *Handler) AddRoom(room LiveRoom) {
	room.Client.RoomId = room.RoomId
	room.Client.BiliChat(h.Handler.CmdChan)
}

func (h *Handler) Run() {
	h.Handler.CmdHandler()
}
