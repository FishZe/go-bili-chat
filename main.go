package go_bili_chat

import (
	"errors"
	"github.com/FishZe/go-bili-chat/client"
	"github.com/FishZe/go-bili-chat/handler"
	log "github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
	"sync"
)

const DefaultClientPriority = 1 << 0
const DelayClientPriority = 1 << 1
const NoCDNClientPriority = 1 << 2

var ClientPriorityMode = DefaultClientPriority
var GetRoomFailed = errors.New("get room failed")
var RoomAlreadyExist = errors.New("room already exist")
var RoomNotExist = errors.New("room not exist")

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
	SetClientPriorityMode(ClientPriorityMode)
}

func ChangeLogLevel(level log.Level) {
	log.SetLevel(level)
}

func SetClientPriorityMode(mode int) {
	ClientPriorityMode = mode
	client.ChangeSequenceMode(mode)
}

func GetNewHandler() *Handler {
	h := Handler{}
	h.Handler.DoFunc = make(map[string]map[int][]func(event handler.MsgEvent), 0)
	h.Handler.CmdChan = make(chan map[string]interface{}, 10)
	h.Handler.Init()
	return &h
}

func (h *Handler) AddOption(Cmd string, RoomId int, Do func(event handler.MsgEvent), funcName ...string) {
	h.Handler.AddOption(Cmd, RoomId, Do, funcName...)
}

func (h *Handler) AddRoom(roomId int) error {
	if _, ok := h.rooms.Load(roomId); ok {
		return RoomAlreadyExist
	}
	room := LiveRoom{}
	if roomId <= 10000 {
		RealRoomId, err := client.GetRealRoomId(roomId)
		if err != nil {
			return err
		} else if RealRoomId == 0 {
			return GetRoomFailed
		}
		log.Debug(roomId, " is short roomid, the real roomid is: ", RealRoomId)
		roomId = RealRoomId
	}
	room.RoomId = roomId
	room.Client.RoomId = room.RoomId
	room.Client.BiliChat(h.Handler.CmdChan)
	h.rooms.Store(room.RoomId, room)
	return nil
}

func (h *Handler) DelRoom(RoomId int) error {
	if _, ok := h.rooms.Load(RoomId); !ok {
		return RoomNotExist
	}
	h.Handler.DelRoomOption(RoomId)
	if c, ok := h.rooms.Load(RoomId); ok {
		cl := c.(LiveRoom).Client
		cl.Close()
		h.rooms.Delete(RoomId)
	}
	return nil
}

func (h *Handler) CountRoom() int {
	count := 0
	h.rooms.Range(func(key, value interface{}) bool {
		count++
		return true
	})
	return count
}

func (h *Handler) Run() {
	h.Handler.CmdHandler()
}
