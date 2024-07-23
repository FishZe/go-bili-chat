package go_bili_chat

import (
	"errors"
	"github.com/FishZe/go-bili-chat/v2/events"
	"net/http"
	"sync"

	"github.com/FishZe/go-bili-chat/v2/client"
	"github.com/FishZe/go-bili-chat/v2/handler"
	log "github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
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

func SetHeader(header http.Header) {
	client.Header = header
}

func SetHeaderUA(ua string) {
	client.Header.Set("User-Agent", ua)
}

func SetBuvid(Buvid string) {
	client.Buvid = Buvid
}

func SetHeaderCookie(cookie string) {
	client.Header.Set("Cookie", cookie)
}

func SetUID(uid int64) {
	client.UID = uid
}

func GetNewHandler() *Handler {
	h := Handler{}
	h.Handler.DoFunc = make(handler.CmdTable, 0)
	h.Handler.CmdChan = make(chan *events.BLiveEvent, 1024)
	h.Handler.FuncPath = make(map[*events.BLiveEventHandler]handler.Path)
	return &h
}

func (h *Handler) AddOption(RoomId int, Do events.BLiveEventHandler) *events.BLiveEventHandler {
	if RoomId <= 10000 && RoomId != 0 {
		RealRoomId, err := client.GetRealRoomId(RoomId)
		if err != nil {
			log.Error(err)
			return nil
		} else if RealRoomId == 0 {
			log.Error(GetRoomFailed)
			return nil
		}
		log.Debug(RoomId, " is short roomid, the real roomid is: ", RealRoomId)
		RoomId = RealRoomId
	}
	return h.Handler.AddOption(RoomId, Do)
}

func (h *Handler) DelOption(f *events.BLiveEventHandler) {
	h.Handler.DelOption(f)
}

func (h *Handler) ExistRoom(RoomId int) bool {
	if _, ok := h.rooms.Load(RoomId); ok {
		return true
	}
	return false
}

func (h *Handler) AddRawRoom(RawRoom client.WsAuthBody) error {
	if _, ok := h.rooms.Load(RawRoom.Roomid); ok {
		return RoomAlreadyExist
	}
	room := LiveRoom{}
	room.RoomId = RawRoom.Roomid
	room.Client.RoomInfo = RawRoom
	room.Client.BiliChat(h.Handler.CmdChan)
	h.rooms.Store(room.RoomId, room)
	return nil
}

func (h *Handler) AddRoom(RoomId int) error {
	if RoomId <= 10000 {
		RealRoomId, err := client.GetRealRoomId(RoomId)
		if err != nil {
			return err
		} else if RealRoomId == 0 {
			return GetRoomFailed
		}
		log.Debug(RoomId, " is short roomid, the real roomid is: ", RealRoomId)
		RoomId = RealRoomId
	}
	return h.AddRawRoom(client.WsAuthBody{
		Roomid:   RoomId,
		Protover: 3,
		UID:      client.UID,
		Buvid:    client.Buvid,
		Type:     2,
		Platform: "web",
		Key:      "",
	})
}

func (h *Handler) DelRoom(RoomId int) error {
	if RoomId <= 10000 {
		RealRoomId, err := client.GetRealRoomId(RoomId)
		if err != nil {
			return err
		} else if RealRoomId == 0 {
			return GetRoomFailed
		}
		log.Debug(RoomId, " is short roomid, the real roomid is: ", RealRoomId)
		RoomId = RealRoomId
	}
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
