package main

import (
	"fmt"
	"github.com/FishZe/go_bilichat_core/client"
	"github.com/FishZe/go_bilichat_core/handler"
	"time"
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
	h.Handler.CmdChan = make(chan map[string]string, 100)
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

func main() {
	h := GetNewHandler()
	go h.Run()
	living := []int{23946255, 1519992, 25243084, 21762782, 156789, 23074404, 14052636, 24399819, 21768293, 4375197, 23170704, 25033029, 22309191, 25158669, 14629550, 21345362, 4101456, 14963868, 22696954, 24970821, 14795432, 12032317, 24188245, 680462, 24593085, 3822389, 395113, 4400273, 23233451, 24483629, 76326, 23749574, 58593, 1287123, 725364, 22259479, 22644333, 21564812, 22228014, 596082, 1868279, 22422551, 6204892, 11771685, 1140818, 23588360, 11268906, 22543433, 470799}
	for _, v := range living {
		h.AddOption(handler.CmdDanmuMsg, v, func(event handler.MsgEvent) {
			fmt.Printf("[%v] %v: %v\n", event.RoomId, event.DanMuMsg.Data.Sender.Name, event.DanMuMsg.Data.Content)
		})
		h.AddRoom(LiveRoom{RoomId: v})
	}
	for {
		time.Sleep(10 * time.Second)
	}
}
