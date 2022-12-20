package go_bilichat_core

import (
	"fmt"
	"github.com/FishZe/go_bilichat_core/Handler"
	"github.com/FishZe/go_bilichat_core/client"
)

const (
	CmdDanmuMsg                  = "DANMU_MSG"
	CmdSuperChatMessage          = "SUPER_CHAT_MESSAGE"
	CmdSuperChatMessageJpn       = "SUPER_CHAT_MESSAGE_JPN"
	CmdWatchedChange             = "WATCHED_CHANGE"
	CmdSendGift                  = "SEND_GIFT"
	CmdOnlineRankCount           = "ONLINE_RANK_COUNT"
	CmdOnlineRankV2              = "ONLINE_RANK_V2"
	CmdOnlineRankTop3            = "ONLINE_RANK_TOP3"
	CmdLikeInfoV3Click           = "LIKE_INFO_V3_CLICK"
	CmdInteractWord              = "INTERACT_WORD"
	CmdStopLiveRoomList          = "STOP_LIVE_ROOM_LIST"
	CmdLikeInfoV3Update          = "LIKE_INFO_V3_UPDATE"
	CmdHotRankChange             = "HOT_RANK_CHANGED"
	CmdNoticeMsg                 = "NOTICE_MSG"
	CmdRoomRealTimeMessageUpdate = "ROOM_REAL_TIME_MESSAGE_UPDATE"
	CmdWidgetBanner              = "WIDGET_BANNER"
	CmdHotRankChangedV2          = "HOT_RANK_CHANGED_V2"
	CmdGuardHonorThousand        = "GUARD_HONOR_THOUSAND"
	CmdLive                      = "LIVE"
	CmdRoomChange                = "ROOM_CHANGE"
	CmdRoomBlockMsg              = "ROOM_BLOCK_MSG"
	CmdFullScreenSpecialEffect   = "FULL_SCREEN_SPECIAL_EFFECT"
	CmdCommonNoticeDanmaku       = "COMMON_NOTICE_DANMAKU"
	CmdTradingScore              = "TRADING_SCORE"
	CmdPreparing                 = "PREPARING"
	CmdGuardBuy                  = "GUARD_BUY"
	CmdGiftStarProcess           = "GIFT_STAR_PROCESS"
	CmdRoomSkinMsg               = "ROOM_SKIN_MSG"
	CmdEnterEffect               = "ENTER_EFFECT"
)

type option struct {
	Cmd    string
	RoomId int
	DoFunc func(event Handler.MsgEvent)
}

type Handle struct {
	Options []option
	cmdChan chan map[string]interface{}
}

type LiveRoom struct {
	RoomId int
	client client.Client
}

func GetNewHandler() Handle {
	return Handle{cmdChan: make(chan map[string]interface{}, 1)}
}

func (live *LiveRoom) New(roomId int) {
	live.client.RoomId = roomId
}

func (handle *Handle) New(Cmd string, RoomId int, DoFunc func(event Handler.MsgEvent)) {
	handle.Options = append(handle.Options, option{Cmd: Cmd, RoomId: RoomId, DoFunc: DoFunc})
}

func (handle *Handle) Binding(room LiveRoom) error {
	room.client.RoomId = room.RoomId
	if room.client.RoomId == 0 {
		return fmt.Errorf("room id is 0")
	}
	err := room.client.BiliChat(handle.cmdChan)
	if err != nil {
		return err
	}
	return nil
}

func (handle *Handle) Run() {
	var options []Handler.Options
	for _, option := range handle.Options {
		options = append(options, Handler.Options{RoomId: []int{option.RoomId}, Cmd: option.Cmd, DoFunc: option.DoFunc})
	}
	h := Handler.Handler{Options: options, CmdChan: handle.cmdChan}
	go h.CmdHandler()
}
