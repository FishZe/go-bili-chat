package handler

import (
	"encoding/json"
	"strconv"
)

// SetDanMuMsg 设置弹幕消息
// 该消息为list结构, 部分字段含义未知, 因此目前只有部分内容
func (_ *Handler) SetDanMuMsg(msg map[string]interface{}) (m MsgEvent) {
	danMu := DanMuMsg{}
	danMu.Cmd = CmdDanmuMsg
	danMuMsg := make(map[string]interface{}, 0)
	if err := json.Unmarshal([]byte(msg["msg"].(string)), &danMuMsg); err == nil {
		danMu.Data.Content = danMuMsg["info"].([]interface{})[1].(string)
		danMu.Data.SendTimeStamp = int(danMuMsg["info"].([]interface{})[9].(map[string]interface{})["ts"].(float64))
		danMu.Data.SenderEnterRoomTimeStamp = int(danMuMsg["info"].([]interface{})[0].([]interface{})[4].(float64))
		danMu.Data.SendMillionTimeStamp = int64(danMuMsg["info"].([]interface{})[0].([]interface{})[5].(float64))
		danMu.Data.Sender.Uid = int64(danMuMsg["info"].([]interface{})[2].([]interface{})[0].(float64))
		danMu.Data.Sender.Name = danMuMsg["info"].([]interface{})[2].([]interface{})[1].(string)
		// 部分情况下, 弹幕发送者未佩戴牌子, 需要判断
		if len(danMuMsg["info"].([]interface{})[3].([]interface{})) != 0 {
			danMu.Data.Medal.GuardLevel = int(danMuMsg["info"].([]interface{})[3].([]interface{})[0].(float64))
			danMu.Data.Medal.MedalName = danMuMsg["info"].([]interface{})[3].([]interface{})[1].(string)
			danMu.Data.Medal.TargetID = int(danMuMsg["info"].([]interface{})[3].([]interface{})[11].(float64))
			danMu.Data.Medal.AnchorRoomId = int(danMuMsg["info"].([]interface{})[3].([]interface{})[3].(float64))
		}
		m = MsgEvent{Cmd: CmdDanmuMsg, DanMuMsg: &danMu, RoomId: msg["RoomId"].(int)}
	}
	return
}

// SetInteractWord 设置欢迎消息
func (_ *Handler) SetInteractWord(msg map[string]interface{}) (m MsgEvent) {
	interactMsg := InteractWord{}
	if err := json.Unmarshal([]byte(msg["msg"].(string)), &interactMsg); err == nil {
		m = MsgEvent{Cmd: CmdInteractWord, InteractWord: &interactMsg, RoomId: msg["RoomId"].(int)}
	}
	return
}

// SetOnlineRankCount 暂时未知
func (_ *Handler) SetOnlineRankCount(msg map[string]interface{}) (m MsgEvent) {
	onlineRankCount := OnlineRankCount{}
	if err := json.Unmarshal([]byte(msg["msg"].(string)), &onlineRankCount); err == nil {
		m = MsgEvent{Cmd: CmdOnlineRankCount, OnlineRankCount: &onlineRankCount, RoomId: msg["RoomId"].(int)}
	}
	return
}

// SetWatchedChange 暂时未知
func (_ *Handler) SetWatchedChange(msg map[string]interface{}) (m MsgEvent) {
	watchedChange := WatchedChange{}
	if err := json.Unmarshal([]byte(msg["msg"].(string)), &watchedChange); err == nil {
		m = MsgEvent{Cmd: CmdWatchedChange, WatchedChange: &watchedChange, RoomId: msg["RoomId"].(int)}
	}
	return
}

// SetNoticeMsg 可能为系统消息
// TODO: 尝试优化
func (_ *Handler) SetNoticeMsg(msg map[string]interface{}) (m MsgEvent) {
	noticeMsg := NoticeMsg{}
	notice := make(map[string]interface{}, 0)
	if err := json.Unmarshal([]byte(msg["msg"].(string)), &notice); err == nil {
		// 这个字段很奇怪, 类型不确定, 需要特判
		switch notice["real_roomid"].(type) {
		case float64:
			notice["real_roomid"] = strconv.FormatFloat(notice["real_roomid"].(float64), 'f', -1, 64)
		case int:
			notice["real_roomid"] = strconv.Itoa(notice["real_roomid"].(int))
		}
		if dataJson, err := json.Marshal(notice); err == nil {
			if err = json.Unmarshal(dataJson, &noticeMsg); err == nil {
				m = MsgEvent{Cmd: CmdNoticeMsg, NoticeMsg: &noticeMsg, RoomId: msg["RoomId"].(int)}
			}
		}
	}
	return
}

// SetSuperChatMessage 超级留言
// TODO: 尝试优化
func (_ *Handler) SetSuperChatMessage(msg map[string]interface{}) (m MsgEvent) {
	superChatMsg := SuperChatMessage{}
	superChatMsg.Cmd = CmdSuperChatMessage
	sc := make(map[string]interface{}, 0)
	if err := json.Unmarshal([]byte(msg["msg"].(string)), &sc); err == nil {
		// id 和 uid 类型不确定, 需要特判
		switch sc["data"].(map[string]interface{})["id"].(type) {
		case float64:
			sc["data"].(map[string]interface{})["id"] = strconv.FormatFloat(sc["data"].(map[string]interface{})["id"].(float64), 'f', -1, 64)
		case int:
			sc["data"].(map[string]interface{})["id"] = strconv.Itoa(sc["data"].(map[string]interface{})["id"].(int))
		}
		switch sc["data"].(map[string]interface{})["uid"].(type) {
		case float64:
			sc["data"].(map[string]interface{})["uid"] = strconv.FormatFloat(sc["data"].(map[string]interface{})["uid"].(float64), 'f', -1, 64)
		case int:
			sc["data"].(map[string]interface{})["uid"] = strconv.Itoa(sc["data"].(map[string]interface{})["uid"].(int))
		}
		if dataJson, err := json.Marshal(sc["data"]); err == nil {
			if err = json.Unmarshal(dataJson, &superChatMsg.Data); err == nil {
				m = MsgEvent{Cmd: CmdSuperChatMessage, SuperChatMessage: &superChatMsg, RoomId: msg["RoomId"].(int)}
			}
		}
	}
	return
}

// SetSendGift 赠送礼物
func (_ *Handler) SetSendGift(msg map[string]interface{}) (m MsgEvent) {
	sendGift := SendGift{}
	if err := json.Unmarshal([]byte(msg["msg"].(string)), &sendGift); err == nil {
		m = MsgEvent{Cmd: CmdSendGift, SendGift: &sendGift, RoomId: msg["RoomId"].(int)}
	}
	return
}

// SetOnlineRankV2 未知
func (_ *Handler) SetOnlineRankV2(msg map[string]interface{}) (m MsgEvent) {
	onlineRankV2 := OnlineRankV2{}
	if err := json.Unmarshal([]byte(msg["msg"].(string)), &onlineRankV2); err == nil {
		m = MsgEvent{Cmd: CmdOnlineRankV2, OnlineRankV2: &onlineRankV2, RoomId: msg["RoomId"].(int)}
	}
	return
}

// SetOnlineRankTop3 未知
func (_ *Handler) SetOnlineRankTop3(msg map[string]interface{}) (m MsgEvent) {
	onlineRankTop3 := OnlineRankTop3{}
	if err := json.Unmarshal([]byte(msg["msg"].(string)), &onlineRankTop3); err == nil {
		m = MsgEvent{Cmd: CmdOnlineRankTop3, OnlineRankTop3: &onlineRankTop3, RoomId: msg["RoomId"].(int)}
	}
	return
}

// SetLikeInfoV3Click 可能为点赞
func (_ *Handler) SetLikeInfoV3Click(msg map[string]interface{}) (m MsgEvent) {
	likeInfoV3Click := LikeInfoV3Click{}
	if err := json.Unmarshal([]byte(msg["msg"].(string)), &likeInfoV3Click); err == nil {
		m = MsgEvent{Cmd: CmdLikeInfoV3Click, LikeInfoV3Click: &likeInfoV3Click, RoomId: msg["RoomId"].(int)}
	}
	return
}

// SetStopLiveRoomList 未知
func (_ *Handler) SetStopLiveRoomList(msg map[string]interface{}) (m MsgEvent) {
	stopLiveRoomList := StopLiveRoomList{}
	if err := json.Unmarshal([]byte(msg["msg"].(string)), &stopLiveRoomList); err == nil {
		m = MsgEvent{Cmd: CmdStopLiveRoomList, StopLiveRoomList: &stopLiveRoomList, RoomId: msg["RoomId"].(int)}
	}
	return
}

// SetLikeInfoV3Update 未知
func (_ *Handler) SetLikeInfoV3Update(msg map[string]interface{}) (m MsgEvent) {
	likeInfoV3Update := LikeInfoV3Update{}
	if err := json.Unmarshal([]byte(msg["msg"].(string)), &likeInfoV3Update); err == nil {
		m = MsgEvent{Cmd: CmdLikeInfoV3Update, LikeInfoV3Update: &likeInfoV3Update, RoomId: msg["RoomId"].(int)}
	}
	return
}

// SetHotRankChange 未知
func (_ *Handler) SetHotRankChange(msg map[string]interface{}) (m MsgEvent) {
	hotRankChange := HotRankChange{}
	if err := json.Unmarshal([]byte(msg["msg"].(string)), &hotRankChange); err == nil {
		m = MsgEvent{Cmd: CmdHotRankChange, HotRankChange: &hotRankChange, RoomId: msg["RoomId"].(int)}
	}
	return
}

// SetRoomRealTimeMessageUpdate 未知
func (_ *Handler) SetRoomRealTimeMessageUpdate(msg map[string]interface{}) (m MsgEvent) {
	roomRealTimeMessageUpdate := RoomRealTimeMessageUpdate{}
	if err := json.Unmarshal([]byte(msg["msg"].(string)), &roomRealTimeMessageUpdate); err == nil {
		m = MsgEvent{Cmd: CmdRoomRealTimeMessageUpdate, RoomRealTimeMessageUpdate: &roomRealTimeMessageUpdate, RoomId: msg["RoomId"].(int)}

	}
	return
}

// SetWidgetBanner 未知
func (_ *Handler) SetWidgetBanner(msg map[string]interface{}) (m MsgEvent) {
	widgetBanner := WidgetBanner{}
	if err := json.Unmarshal([]byte(msg["msg"].(string)), &widgetBanner); err == nil {
		m = MsgEvent{Cmd: CmdWidgetBanner, WidgetBanner: &widgetBanner, RoomId: msg["RoomId"].(int)}
	}
	return
}

// SetHotRankChangedV2 未知
func (_ *Handler) SetHotRankChangedV2(msg map[string]interface{}) (m MsgEvent) {
	hotRankChangedV2 := HotRankChangedV2{}
	if err := json.Unmarshal([]byte(msg["msg"].(string)), &hotRankChangedV2); err == nil {
		m = MsgEvent{Cmd: CmdHotRankChangedV2, HotRankChangedV2: &hotRankChangedV2, RoomId: msg["RoomId"].(int)}
	}
	return
}

// SetGuardHonorThousand 未知
func (_ *Handler) SetGuardHonorThousand(msg map[string]interface{}) (m MsgEvent) {
	guardHonorThousand := GuardHonorThousand{}
	if err := json.Unmarshal([]byte(msg["msg"].(string)), &guardHonorThousand); err == nil {
		m = MsgEvent{Cmd: CmdGuardHonorThousand, GuardHonorThousand: &guardHonorThousand, RoomId: msg["RoomId"].(int)}
	}
	return
}

// SetLive 开始直播
func (_ *Handler) SetLive(msg map[string]interface{}) (m MsgEvent) {
	live := Live{}
	if err := json.Unmarshal([]byte(msg["msg"].(string)), &live); err == nil {
		m = MsgEvent{Cmd: CmdLive, Live: &live, RoomId: msg["RoomId"].(int)}
	}
	return
}

// SetRoomChange 未知
func (_ *Handler) SetRoomChange(msg map[string]interface{}) (m MsgEvent) {
	roomChange := RoomChange{}
	if err := json.Unmarshal([]byte(msg["msg"].(string)), &roomChange); err == nil {
		m = MsgEvent{Cmd: CmdRoomChange, RoomChange: &roomChange, RoomId: msg["RoomId"].(int)}
	}
	return
}

// SetRoomBlockMsg 未知
func (_ *Handler) SetRoomBlockMsg(msg map[string]interface{}) (m MsgEvent) {
	roomBlockMsg := RoomBlockMsg{}
	if err := json.Unmarshal([]byte(msg["msg"].(string)), &roomBlockMsg); err == nil {
		m = MsgEvent{Cmd: CmdRoomBlockMsg, RoomBlockMsg: &roomBlockMsg, RoomId: msg["RoomId"].(int)}
	}
	return
}

// SetFullScreenSpecialEffect 可能为礼物特效
func (_ *Handler) SetFullScreenSpecialEffect(msg map[string]interface{}) (m MsgEvent) {
	fullScreenSpecialEffect := FullScreenSpecialEffect{}
	if err := json.Unmarshal([]byte(msg["msg"].(string)), &fullScreenSpecialEffect); err == nil {
		m = MsgEvent{Cmd: CmdFullScreenSpecialEffect, FullScreenSpecialEffect: &fullScreenSpecialEffect, RoomId: msg["RoomId"].(int)}
	}
	return
}

// SetCommonNoticeDanmaku 未知
func (_ *Handler) SetCommonNoticeDanmaku(msg map[string]interface{}) (m MsgEvent) {
	commonNoticeDanmaku := CommonNoticeDanmaku{}
	if err := json.Unmarshal([]byte(msg["msg"].(string)), &commonNoticeDanmaku); err == nil {
		m = MsgEvent{Cmd: CmdCommonNoticeDanmaku, CommonNoticeDanmaku: &commonNoticeDanmaku, RoomId: msg["RoomId"].(int)}
	}
	return
}

// SetTradingScore 未知
func (_ *Handler) SetTradingScore(msg map[string]interface{}) (m MsgEvent) {
	tradingScore := TradingScore{}
	if err := json.Unmarshal([]byte(msg["msg"].(string)), &tradingScore); err == nil {
		m = MsgEvent{Cmd: CmdTradingScore, TradingScore: &tradingScore, RoomId: msg["RoomId"].(int)}
	}
	return
}

// SetPreparing 开始准备
func (_ *Handler) SetPreparing(msg map[string]interface{}) (m MsgEvent) {
	preparing := Preparing{}
	preparing.Cmd = CmdPreparing
	tmp := make(map[string]interface{})
	if err := json.Unmarshal([]byte(msg["msg"].(string)), &tmp); err == nil {
		preparing.RoomId = msg["RoomId"].(string)
		m = MsgEvent{Cmd: CmdPreparing, Preparing: &preparing, RoomId: msg["RoomId"].(int)}
	}
	return
}

// SetGuardBuy 大航海购买
func (_ *Handler) SetGuardBuy(msg map[string]interface{}) (m MsgEvent) {
	guardBuy := GuardBuy{}
	if err := json.Unmarshal([]byte(msg["msg"].(string)), &guardBuy); err == nil {
		m = MsgEvent{Cmd: CmdGuardBuy, GuardBuy: &guardBuy, RoomId: msg["RoomId"].(int)}
	}
	return
}

// SetGiftStarProcess 未知
func (_ *Handler) SetGiftStarProcess(msg map[string]interface{}) (m MsgEvent) {
	giftStarProcess := GiftStarProcess{}
	if err := json.Unmarshal([]byte(msg["msg"].(string)), &giftStarProcess); err == nil {
		m = MsgEvent{Cmd: CmdGiftStarProcess, GiftStarProcess: &giftStarProcess, RoomId: msg["RoomId"].(int)}
	}
	return
}

// SetRoomSkinMsg 未知
func (_ *Handler) SetRoomSkinMsg(msg map[string]interface{}) (m MsgEvent) {
	roomSkinMsg := RoomSkinMsg{}
	if err := json.Unmarshal([]byte(msg["msg"].(string)), &roomSkinMsg); err == nil {
		m = MsgEvent{Cmd: CmdRoomSkinMsg, RoomSkinMsg: &roomSkinMsg, RoomId: msg["RoomId"].(int)}
	}
	return
}

// SetEntryEffect 未知
func (_ *Handler) SetEntryEffect(msg map[string]interface{}) (m MsgEvent) {
	enterEffect := EntryEffect{}
	if err := json.Unmarshal([]byte(msg["msg"].(string)), &enterEffect); err == nil {
		m = MsgEvent{Cmd: CmdEntryEffect, EntryEffect: &enterEffect, RoomId: msg["RoomId"].(int)}
	}
	return
}

// SetUserToastMsg 上舰长
func (_ *Handler) SetUserToastMsg(msg map[string]interface{}) (m MsgEvent) {
	userToastMsg := UserToastMsg{}
	if err := json.Unmarshal([]byte(msg["msg"].(string)), &userToastMsg); err == nil {
		m = MsgEvent{Cmd: CmdUserToastMsg, UserToastMsg: &userToastMsg, RoomId: msg["RoomId"].(int)}
	}
	return
}
