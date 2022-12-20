package Handler

import (
	"time"
)

type Handler struct {
	CmdChan chan map[string]interface{}
	Options []Options
}

func (handler *Handler) CmdHandler() {
	msgHandlerFunc := map[string]func(map[string]interface{}) MsgEvent{
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
		CmdGuardHonorThousand:        SetGuardHonorThousand,
		CmdLive:                      SetLive,
		CmdRoomChange:                SetRoomChange,
		CmdRoomBlockMsg:              SetRoomBlockMsg,
		CmdFullScreenSpecialEffect:   SetFullScreenSpecialEffect,
		CmdCommonNoticeDanmaku:       SetCommonNoticeDanmaku,
		CmdTradingScore:              SetTradingScore,
		CmdPreparing:                 SetPreparing,
		CmdGuardBuy:                  SetGuardBuy,
		CmdGiftStarProcess:           SetGiftStarProcess,
		CmdRoomSkinMsg:               SetRoomSkinMsg,
		CmdEnterEffect:               SetEnterEffect,
	}
	doFuncs := make(map[string]map[int][]func(event MsgEvent), 0)
	for _, v := range handler.Options {
		if v.RoomId == nil || len(v.RoomId) == 0 || v.Cmd == "" || v.DoFunc == nil {
			continue
		}
		// 命令未绑定
		if _, ok := doFuncs[v.Cmd]; !ok {
			doFuncs[v.Cmd] = make(map[int][]func(event MsgEvent))
		}
		for _, roomId := range v.RoomId {
			// 房间未绑定
			if _, ok := doFuncs[v.Cmd][roomId]; !ok {
				doFuncs[v.Cmd][roomId] = make([]func(event MsgEvent), 0)
			}
			doFuncs[v.Cmd][roomId] = append(doFuncs[v.Cmd][roomId], v.DoFunc)
		}
	}
	for {
		select {
		case msg, ok := <-handler.CmdChan:
			if ok {
				if _, ok := msgHandlerFunc[msg["cmd"].(string)]; ok {
					msgEvent := msgHandlerFunc[msg["cmd"].(string)](msg)
					if msgEvent.Cmd == "" || msgEvent.RoomId == 0 {
						continue
					}
					if _, ok := doFuncs[msg["cmd"].(string)]; ok {
						if _, ok := doFuncs[msg["cmd"].(string)][msgEvent.RoomId]; ok {
							for _, v := range doFuncs[msg["cmd"].(string)][msgEvent.RoomId] {
								go v(msgEvent)
							}
						}
					}
				}
			}
		default:
			time.Sleep(10 * time.Microsecond)
			continue
		}

	}
}
