package go_bili_chat

import (
	"github.com/FishZe/go-bili-chat/client"
	"github.com/FishZe/go-bili-chat/handler"
	log "github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"sync"
)

const DefaultClientSequence = 1
const DelayClientSequence = 2

var ClientSequenceMode = DefaultClientSequence

type Handler struct {
	Handler handler.Handler
	rooms   sync.Map
}

type LiveRoom struct {
	RoomId int
	Client client.Client
}

func init() {
	log.SetFormatter(&easy.Formatter{
		TimestampFormat: "01-02 15:04:05",
		LogFormat:       "[bili-live][%time%][%lvl%]: %msg% \n",
	})
	ChangeLogLevel(log.ErrorLevel)
	SetJsonCoder(&DefaultJson{})
	SetClientSequenceMode(DefaultClientSequence)
}

func ChangeLogLevel(level log.Level) {
	log.SetLevel(level)
}

func SetClientSequenceMode(mode int) {
	ClientSequenceMode = mode
	client.ChangeSequenceMode(mode)
}

func GetNewHandler() Handler {
	h := Handler{}
	h.Handler.DoFunc = make(map[string]map[int][]func(event handler.MsgEvent), 0)
	h.Handler.CmdChan = make(chan map[string]interface{}, 10)
	h.Handler.Init()
	return h
}

func (h *Handler) AddOption(Cmd string, RoomId int, Do func(event handler.MsgEvent), funcName ...string) {
	h.Handler.AddOption(Cmd, RoomId, Do, funcName...)
}

func (h *Handler) AddRoom(roomId int) {
	if _, ok := h.rooms.Load(roomId); ok {
		return
	}
	room := LiveRoom{}
	if roomId <= 10000 {
		RealroomId, err := client.GetRealRoomId(roomId)
		if err != nil {
			log.Error(err)
			return
		}
		log.Info(roomId, " is short roomid, the real roomid is: ", RealroomId)
		roomId = RealroomId
	}
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
