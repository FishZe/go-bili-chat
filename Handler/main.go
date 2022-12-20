package Handler

import (
	"fmt"
	"time"
)

type Handler struct {
	CmdChan chan map[string]interface{}
}

func (handler *Handler) CmdHandler() {
	msgHandlerFunc := map[string]func(map[string]interface{}){
		CmdDanmuMsg:                  SetDanMuMsg,
		CmdInteractWord:              SetInteractWord,
		CmdOnlineRankCount:           SetOnlineRankCount,
		CmdWatchedChange:             SetWatchedChange,
		CmdNoticeMsg:                 SetNoticeMsg,
		CmdSuperChatMessage:          SetSuperChatMessage,
		CmdSuperChatMessageJpn:       SetSuperChatMessage,
		CmdSendGift:                  SetSendGift,
		CmdOnlineRankV2:              SetOnlineRankV2,
		CmdOnlineRankTop3:            SetOnlineRankTop3,
		CmdLikeInfoV3Click:           SetLikeInfoV3Click,
		CmdStopLiveRoomList:          SetStopLiveRoomList,
		CmdLikeInfoV3Update:          SetLikeInfoV3Update,
		CmdHotRankChange:             SetHotRankChange,
		CmdRoomRealTimeMessageUpdate: SetRoomRealTimeMessageUpdate,
		CmdWidgetBanner:              SetWidgetBanner,
		CmdHotRankChangedV2:          SetHotRankChangedV2,
	}
	for {
		select {
		case msg, ok := <-handler.CmdChan:
			if ok {
				if _, ok := msgHandlerFunc[msg["cmd"].(string)]; ok {
					msgHandlerFunc[msg["cmd"].(string)](msg)
				} else {
					fmt.Println(msg["cmd"].(string))
				}
			}
		default:
			time.Sleep(10 * time.Microsecond)
			continue
		}

	}
}
