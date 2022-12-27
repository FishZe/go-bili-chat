package handler

import (
	"encoding/json"
	"log"
	"strconv"
)

func (_ *Handler) SetDanMuMsg(msg map[string]interface{}) MsgEvent {
	danMu := DanMuMsg{}
	danMu.Cmd = CmdDanmuMsg
	danMuMsg := make(map[string]interface{}, 0)
	err := json.Unmarshal([]byte(msg["msg"].(string)), &danMuMsg)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	danMu.Data.Content = danMuMsg["info"].([]interface{})[1].(string)
	danMu.Data.SendTimeStamp = int(danMuMsg["info"].([]interface{})[9].(map[string]interface{})["ts"].(float64))
	danMu.Data.SenderEnterRoomTimeStamp = int(danMuMsg["info"].([]interface{})[0].([]interface{})[4].(float64))
	danMu.Data.SendMillionTimeStamp = int64(danMuMsg["info"].([]interface{})[0].([]interface{})[5].(float64))
	danMu.Data.Sender.Uid = int64(danMuMsg["info"].([]interface{})[2].([]interface{})[0].(float64))
	danMu.Data.Sender.Name = danMuMsg["info"].([]interface{})[2].([]interface{})[1].(string)
	if len(danMuMsg["info"].([]interface{})[3].([]interface{})) != 0 {
		danMu.Data.Medal.GuardLevel = int(danMuMsg["info"].([]interface{})[3].([]interface{})[0].(float64))
		danMu.Data.Medal.MedalName = danMuMsg["info"].([]interface{})[3].([]interface{})[1].(string)
		danMu.Data.Medal.TargetID = int(danMuMsg["info"].([]interface{})[3].([]interface{})[11].(float64))
		danMu.Data.Medal.AnchorRoomId = int(danMuMsg["info"].([]interface{})[3].([]interface{})[3].(float64))
	}
	return MsgEvent{Cmd: CmdDanmuMsg, DanMuMsg: &danMu, RoomId: msg["RoomId"].(int)}
}

func (_ *Handler) SetInteractWord(msg map[string]interface{}) MsgEvent {
	interactMsg := InteractWord{}
	err := json.Unmarshal([]byte(msg["msg"].(string)), &interactMsg)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	return MsgEvent{Cmd: CmdInteractWord, InteractWord: &interactMsg, RoomId: msg["RoomId"].(int)}
}

func (_ *Handler) SetOnlineRankCount(msg map[string]interface{}) MsgEvent {
	onlineRankCount := OnlineRankCount{}
	err := json.Unmarshal([]byte(msg["msg"].(string)), &onlineRankCount)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}

	return MsgEvent{Cmd: CmdOnlineRankCount, OnlineRankCount: &onlineRankCount, RoomId: msg["RoomId"].(int)}
}

func (_ *Handler) SetWatchedChange(msg map[string]interface{}) MsgEvent {
	watchedChange := WatchedChange{}
	err := json.Unmarshal([]byte(msg["msg"].(string)), &watchedChange)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}

	return MsgEvent{Cmd: CmdWatchedChange, WatchedChange: &watchedChange, RoomId: msg["RoomId"].(int)}
}

func (_ *Handler) SetNoticeMsg(msg map[string]interface{}) MsgEvent {
	noticeMsg := NoticeMsg{}
	notice := make(map[string]interface{}, 0)
	err := json.Unmarshal([]byte(msg["msg"].(string)), &notice)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	switch notice["real_roomid"].(type) {
	case float64:
		notice["real_roomid"] = strconv.FormatFloat(notice["real_roomid"].(float64), 'f', -1, 64)
	case int:
		notice["real_roomid"] = strconv.Itoa(notice["real_roomid"].(int))
	}
	dataJson, err := json.Marshal(notice)
	if err != nil {
		log.Printf("Marshal cmd json failed: %v", err)
	}
	err = json.Unmarshal(dataJson, &noticeMsg)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}

	return MsgEvent{Cmd: CmdNoticeMsg, NoticeMsg: &noticeMsg, RoomId: msg["RoomId"].(int)}
}

func (_ *Handler) SetSuperChatMessage(msg map[string]interface{}) MsgEvent {
	superChatMsg := SuperChatMessage{}
	superChatMsg.Cmd = CmdSuperChatMessage
	sc := make(map[string]interface{}, 0)
	err := json.Unmarshal([]byte(msg["msg"].(string)), &sc)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
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
	dataJson, err := json.Marshal(sc["data"])
	if err != nil {
		log.Printf("Marshal cmd json failed: %v", err)
	}
	err = json.Unmarshal(dataJson, &superChatMsg.Data)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}

	return MsgEvent{Cmd: CmdSuperChatMessage, SuperChatMessage: &superChatMsg, RoomId: msg["RoomId"].(int)}
}

func (_ *Handler) SetSendGift(msg map[string]interface{}) MsgEvent {
	sendGift := SendGift{}
	err := json.Unmarshal([]byte(msg["msg"].(string)), &sendGift)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}

	return MsgEvent{Cmd: CmdSendGift, SendGift: &sendGift, RoomId: msg["RoomId"].(int)}
}

func (_ *Handler) SetOnlineRankV2(msg map[string]interface{}) MsgEvent {
	onlineRankV2 := OnlineRankV2{}
	err := json.Unmarshal([]byte(msg["msg"].(string)), &onlineRankV2)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}

	return MsgEvent{Cmd: CmdOnlineRankV2, OnlineRankV2: &onlineRankV2, RoomId: msg["RoomId"].(int)}
}

func (_ *Handler) SetOnlineRankTop3(msg map[string]interface{}) MsgEvent {
	onlineRankTop3 := OnlineRankTop3{}
	err := json.Unmarshal([]byte(msg["msg"].(string)), &onlineRankTop3)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}

	}

	return MsgEvent{Cmd: CmdOnlineRankTop3, OnlineRankTop3: &onlineRankTop3, RoomId: msg["RoomId"].(int)}
}

func (_ *Handler) SetLikeInfoV3Click(msg map[string]interface{}) MsgEvent {
	likeInfoV3Click := LikeInfoV3Click{}
	err := json.Unmarshal([]byte(msg["msg"].(string)), &likeInfoV3Click)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}

	}

	return MsgEvent{Cmd: CmdLikeInfoV3Click, LikeInfoV3Click: &likeInfoV3Click, RoomId: msg["RoomId"].(int)}
}

func (_ *Handler) SetStopLiveRoomList(msg map[string]interface{}) MsgEvent {
	stopLiveRoomList := StopLiveRoomList{}
	err := json.Unmarshal([]byte(msg["msg"].(string)), &stopLiveRoomList)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}

	}
	return MsgEvent{Cmd: CmdStopLiveRoomList, StopLiveRoomList: &stopLiveRoomList, RoomId: msg["RoomId"].(int)}
}

func (_ *Handler) SetLikeInfoV3Update(msg map[string]interface{}) MsgEvent {
	likeInfoV3Update := LikeInfoV3Update{}
	err := json.Unmarshal([]byte(msg["msg"].(string)), &likeInfoV3Update)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}

	}
	return MsgEvent{Cmd: CmdLikeInfoV3Update, LikeInfoV3Update: &likeInfoV3Update, RoomId: msg["RoomId"].(int)}
}

func (_ *Handler) SetHotRankChange(msg map[string]interface{}) MsgEvent {
	hotRankChange := HotRankChange{}
	err := json.Unmarshal([]byte(msg["msg"].(string)), &hotRankChange)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}

	return MsgEvent{Cmd: CmdHotRankChange, HotRankChange: &hotRankChange, RoomId: msg["RoomId"].(int)}
}

func (_ *Handler) SetRoomRealTimeMessageUpdate(msg map[string]interface{}) MsgEvent {
	roomRealTimeMessageUpdate := RoomRealTimeMessageUpdate{}
	err := json.Unmarshal([]byte(msg["msg"].(string)), &roomRealTimeMessageUpdate)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}

	}
	return MsgEvent{Cmd: CmdRoomRealTimeMessageUpdate, RoomRealTimeMessageUpdate: &roomRealTimeMessageUpdate, RoomId: msg["RoomId"].(int)}
}

func (_ *Handler) SetWidgetBanner(msg map[string]interface{}) MsgEvent {
	widgetBanner := WidgetBanner{}
	err := json.Unmarshal([]byte(msg["msg"].(string)), &widgetBanner)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}

	return MsgEvent{Cmd: CmdWidgetBanner, WidgetBanner: &widgetBanner, RoomId: msg["RoomId"].(int)}
}

func (_ *Handler) SetHotRankChangedV2(msg map[string]interface{}) MsgEvent {
	hotRankChangedV2 := HotRankChangedV2{}
	err := json.Unmarshal([]byte(msg["msg"].(string)), &hotRankChangedV2)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}

	return MsgEvent{Cmd: CmdHotRankChangedV2, HotRankChangedV2: &hotRankChangedV2, RoomId: msg["RoomId"].(int)}
}

func (_ *Handler) SetGuardHonorThousand(msg map[string]interface{}) MsgEvent {
	guardHonorThousand := GuardHonorThousand{}
	err := json.Unmarshal([]byte(msg["msg"].(string)), &guardHonorThousand)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}

	return MsgEvent{Cmd: CmdGuardHonorThousand, GuardHonorThousand: &guardHonorThousand, RoomId: msg["RoomId"].(int)}
}

func (_ *Handler) SetLive(msg map[string]interface{}) MsgEvent {
	live := Live{}
	err := json.Unmarshal([]byte(msg["msg"].(string)), &live)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}

	return MsgEvent{Cmd: CmdLive, Live: &live, RoomId: msg["RoomId"].(int)}
}

func (_ *Handler) SetRoomChange(msg map[string]interface{}) MsgEvent {
	roomChange := RoomChange{}
	err := json.Unmarshal([]byte(msg["msg"].(string)), &roomChange)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}

	return MsgEvent{Cmd: CmdRoomChange, RoomChange: &roomChange, RoomId: msg["RoomId"].(int)}
}

func (_ *Handler) SetRoomBlockMsg(msg map[string]interface{}) MsgEvent {
	roomBlockMsg := RoomBlockMsg{}
	err := json.Unmarshal([]byte(msg["msg"].(string)), &roomBlockMsg)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}

	return MsgEvent{Cmd: CmdRoomBlockMsg, RoomBlockMsg: &roomBlockMsg, RoomId: msg["RoomId"].(int)}
}

func (_ *Handler) SetFullScreenSpecialEffect(msg map[string]interface{}) MsgEvent {
	fullScreenSpecialEffect := FullScreenSpecialEffect{}
	err := json.Unmarshal([]byte(msg["msg"].(string)), &fullScreenSpecialEffect)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}

	return MsgEvent{Cmd: CmdFullScreenSpecialEffect, FullScreenSpecialEffect: &fullScreenSpecialEffect, RoomId: msg["RoomId"].(int)}
}

func (_ *Handler) SetCommonNoticeDanmaku(msg map[string]interface{}) MsgEvent {
	commonNoticeDanmaku := CommonNoticeDanmaku{}
	err := json.Unmarshal([]byte(msg["msg"].(string)), &commonNoticeDanmaku)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}

	return MsgEvent{Cmd: CmdCommonNoticeDanmaku, CommonNoticeDanmaku: &commonNoticeDanmaku, RoomId: msg["RoomId"].(int)}
}

func (_ *Handler) SetTradingScore(msg map[string]interface{}) MsgEvent {
	tradingScore := TradingScore{}
	err := json.Unmarshal([]byte(msg["msg"].(string)), &tradingScore)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}

	return MsgEvent{Cmd: CmdTradingScore, TradingScore: &tradingScore, RoomId: msg["RoomId"].(int)}
}

func (_ *Handler) SetPreparing(msg map[string]interface{}) MsgEvent {
	preparing := Preparing{}
	preparing.Cmd = CmdPreparing
	tmp := make(map[string]interface{})
	err := json.Unmarshal([]byte(msg["msg"].(string)), &tmp)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}

	preparing.RoomId = msg["RoomId"].(string)
	return MsgEvent{Cmd: CmdPreparing, Preparing: &preparing, RoomId: msg["RoomId"].(int)}
}

func (_ *Handler) SetGuardBuy(msg map[string]interface{}) MsgEvent {
	guardBuy := GuardBuy{}
	err := json.Unmarshal([]byte(msg["msg"].(string)), &guardBuy)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}

	return MsgEvent{Cmd: CmdGuardBuy, GuardBuy: &guardBuy, RoomId: msg["RoomId"].(int)}
}

func (_ *Handler) SetGiftStarProcess(msg map[string]interface{}) MsgEvent {
	giftStarProcess := GiftStarProcess{}
	err := json.Unmarshal([]byte(msg["msg"].(string)), &giftStarProcess)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}

	return MsgEvent{Cmd: CmdGiftStarProcess, GiftStarProcess: &giftStarProcess, RoomId: msg["RoomId"].(int)}
}

func (_ *Handler) SetRoomSkinMsg(msg map[string]interface{}) MsgEvent {
	roomSkinMsg := RoomSkinMsg{}
	err := json.Unmarshal([]byte(msg["msg"].(string)), &roomSkinMsg)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}

	return MsgEvent{Cmd: CmdRoomSkinMsg, RoomSkinMsg: &roomSkinMsg, RoomId: msg["RoomId"].(int)}
}

func (_ *Handler) SetEntryEffect(msg map[string]interface{}) MsgEvent {
	enterEffect := EntryEffect{}
	err := json.Unmarshal([]byte(msg["msg"].(string)), &enterEffect)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}

	return MsgEvent{Cmd: CmdEntryEffect, EntryEffect: &enterEffect, RoomId: msg["RoomId"].(int)}
}
