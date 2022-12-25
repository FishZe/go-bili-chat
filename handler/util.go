package handler

import (
	"encoding/json"
	"log"
	"strconv"
)

func (_ *Handler) SetDanMuMsg(msg map[string]string) MsgEvent {
	danMu := DanMuMsg{}
	danMu.Cmd = CmdDanmuMsg
	danMuMsg := make(map[string]interface{}, 0)
	err := json.Unmarshal([]byte(msg["msg"]), &danMuMsg)
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
	roomId, err := strconv.Atoi(msg["RoomId"])
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	return MsgEvent{Cmd: CmdDanmuMsg, DanMuMsg: &danMu, RoomId: roomId}
}

func (_ *Handler) SetInteractWord(msg map[string]string) MsgEvent {
	interactMsg := InteractWord{}
	err := json.Unmarshal([]byte(msg["msg"]), &interactMsg)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	roomId, err := strconv.Atoi(msg["RoomId"])
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	return MsgEvent{Cmd: CmdInteractWord, InteractWord: &interactMsg, RoomId: roomId}
}

func (_ *Handler) SetOnlineRankCount(msg map[string]string) MsgEvent {
	onlineRankCount := OnlineRankCount{}
	err := json.Unmarshal([]byte(msg["msg"]), &onlineRankCount)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	roomId, err := strconv.Atoi(msg["RoomId"])
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	return MsgEvent{Cmd: CmdOnlineRankCount, OnlineRankCount: &onlineRankCount, RoomId: roomId}
}

func (_ *Handler) SetWatchedChange(msg map[string]string) MsgEvent {
	watchedChange := WatchedChange{}
	err := json.Unmarshal([]byte(msg["msg"]), &watchedChange)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	roomId, err := strconv.Atoi(msg["RoomId"])
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	return MsgEvent{Cmd: CmdWatchedChange, WatchedChange: &watchedChange, RoomId: roomId}
}

func (_ *Handler) SetNoticeMsg(msg map[string]string) MsgEvent {
	noticeMsg := NoticeMsg{}
	notice := make(map[string]interface{}, 0)
	err := json.Unmarshal([]byte(msg["msg"]), &notice)
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
	roomId, err := strconv.Atoi(msg["RoomId"])
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	return MsgEvent{Cmd: CmdNoticeMsg, NoticeMsg: &noticeMsg, RoomId: roomId}
}

func (_ *Handler) SetSuperChatMessage(msg map[string]string) MsgEvent {
	superChatMsg := SuperChatMessage{}
	superChatMsg.Cmd = CmdSuperChatMessage
	sc := make(map[string]interface{}, 0)
	err := json.Unmarshal([]byte(msg["msg"]), &sc)
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
	roomId, err := strconv.Atoi(msg["RoomId"])
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	return MsgEvent{Cmd: CmdSuperChatMessage, SuperChatMessage: &superChatMsg, RoomId: roomId}
}

func (_ *Handler) SetSendGift(msg map[string]string) MsgEvent {
	sendGift := SendGift{}
	err := json.Unmarshal([]byte(msg["msg"]), &sendGift)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	roomId, err := strconv.Atoi(msg["RoomId"])
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	return MsgEvent{Cmd: CmdSendGift, SendGift: &sendGift, RoomId: roomId}
}

func (_ *Handler) SetOnlineRankV2(msg map[string]string) MsgEvent {
	onlineRankV2 := OnlineRankV2{}
	err := json.Unmarshal([]byte(msg["msg"]), &onlineRankV2)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	roomId, err := strconv.Atoi(msg["RoomId"])
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	return MsgEvent{Cmd: CmdOnlineRankV2, OnlineRankV2: &onlineRankV2, RoomId: roomId}
}

func (_ *Handler) SetOnlineRankTop3(msg map[string]string) MsgEvent {
	onlineRankTop3 := OnlineRankTop3{}
	err := json.Unmarshal([]byte(msg["msg"]), &onlineRankTop3)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}

	}
	roomId, err := strconv.Atoi(msg["RoomId"])
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	return MsgEvent{Cmd: CmdOnlineRankTop3, OnlineRankTop3: &onlineRankTop3, RoomId: roomId}
}

func (_ *Handler) SetLikeInfoV3Click(msg map[string]string) MsgEvent {
	likeInfoV3Click := LikeInfoV3Click{}
	err := json.Unmarshal([]byte(msg["msg"]), &likeInfoV3Click)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}

	}
	roomId, err := strconv.Atoi(msg["RoomId"])
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	return MsgEvent{Cmd: CmdLikeInfoV3Click, LikeInfoV3Click: &likeInfoV3Click, RoomId: roomId}
}

func (_ *Handler) SetStopLiveRoomList(msg map[string]string) MsgEvent {
	stopLiveRoomList := StopLiveRoomList{}
	err := json.Unmarshal([]byte(msg["msg"]), &stopLiveRoomList)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}

	}
	roomId, err := strconv.Atoi(msg["RoomId"])
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}

	}
	return MsgEvent{Cmd: CmdStopLiveRoomList, StopLiveRoomList: &stopLiveRoomList, RoomId: roomId}
}

func (_ *Handler) SetLikeInfoV3Update(msg map[string]string) MsgEvent {
	likeInfoV3Update := LikeInfoV3Update{}
	err := json.Unmarshal([]byte(msg["msg"]), &likeInfoV3Update)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}

	}
	roomId, err := strconv.Atoi(msg["RoomId"])
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}

	}
	return MsgEvent{Cmd: CmdLikeInfoV3Update, LikeInfoV3Update: &likeInfoV3Update, RoomId: roomId}
}

func (_ *Handler) SetHotRankChange(msg map[string]string) MsgEvent {
	hotRankChange := HotRankChange{}
	err := json.Unmarshal([]byte(msg["msg"]), &hotRankChange)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	roomId, err := strconv.Atoi(msg["RoomId"])
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	return MsgEvent{Cmd: CmdHotRankChange, HotRankChange: &hotRankChange, RoomId: roomId}
}

func (_ *Handler) SetRoomRealTimeMessageUpdate(msg map[string]string) MsgEvent {
	roomRealTimeMessageUpdate := RoomRealTimeMessageUpdate{}
	err := json.Unmarshal([]byte(msg["msg"]), &roomRealTimeMessageUpdate)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}

	}
	roomId, err := strconv.Atoi(msg["RoomId"])
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}

	}
	return MsgEvent{Cmd: CmdRoomRealTimeMessageUpdate, RoomRealTimeMessageUpdate: &roomRealTimeMessageUpdate, RoomId: roomId}
}

func (_ *Handler) SetWidgetBanner(msg map[string]string) MsgEvent {
	widgetBanner := WidgetBanner{}
	err := json.Unmarshal([]byte(msg["msg"]), &widgetBanner)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	roomId, err := strconv.Atoi(msg["RoomId"])
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	return MsgEvent{Cmd: CmdWidgetBanner, WidgetBanner: &widgetBanner, RoomId: roomId}
}

func (_ *Handler) SetHotRankChangedV2(msg map[string]string) MsgEvent {
	hotRankChangedV2 := HotRankChangedV2{}
	err := json.Unmarshal([]byte(msg["msg"]), &hotRankChangedV2)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	roomId, err := strconv.Atoi(msg["RoomId"])
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	return MsgEvent{Cmd: CmdHotRankChangedV2, HotRankChangedV2: &hotRankChangedV2, RoomId: roomId}
}

func (_ *Handler) SetGuardHonorThousand(msg map[string]string) MsgEvent {
	guardHonorThousand := GuardHonorThousand{}
	err := json.Unmarshal([]byte(msg["msg"]), &guardHonorThousand)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	roomId, err := strconv.Atoi(msg["RoomId"])
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	return MsgEvent{Cmd: CmdGuardHonorThousand, GuardHonorThousand: &guardHonorThousand, RoomId: roomId}
}

func (_ *Handler) SetLive(msg map[string]string) MsgEvent {
	live := Live{}
	err := json.Unmarshal([]byte(msg["msg"]), &live)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	roomId, err := strconv.Atoi(msg["RoomId"])
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	return MsgEvent{Cmd: CmdLive, Live: &live, RoomId: roomId}
}

func (_ *Handler) SetRoomChange(msg map[string]string) MsgEvent {
	roomChange := RoomChange{}
	err := json.Unmarshal([]byte(msg["msg"]), &roomChange)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	roomId, err := strconv.Atoi(msg["RoomId"])
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	return MsgEvent{Cmd: CmdRoomChange, RoomChange: &roomChange, RoomId: roomId}
}

func (_ *Handler) SetRoomBlockMsg(msg map[string]string) MsgEvent {
	roomBlockMsg := RoomBlockMsg{}
	err := json.Unmarshal([]byte(msg["msg"]), &roomBlockMsg)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	roomId, err := strconv.Atoi(msg["RoomId"])
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	return MsgEvent{Cmd: CmdRoomBlockMsg, RoomBlockMsg: &roomBlockMsg, RoomId: roomId}
}

func (_ *Handler) SetFullScreenSpecialEffect(msg map[string]string) MsgEvent {
	fullScreenSpecialEffect := FullScreenSpecialEffect{}
	err := json.Unmarshal([]byte(msg["msg"]), &fullScreenSpecialEffect)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	roomId, err := strconv.Atoi(msg["RoomId"])
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	return MsgEvent{Cmd: CmdFullScreenSpecialEffect, FullScreenSpecialEffect: &fullScreenSpecialEffect, RoomId: roomId}
}

func (_ *Handler) SetCommonNoticeDanmaku(msg map[string]string) MsgEvent {
	commonNoticeDanmaku := CommonNoticeDanmaku{}
	err := json.Unmarshal([]byte(msg["msg"]), &commonNoticeDanmaku)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	roomId, err := strconv.Atoi(msg["RoomId"])
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	return MsgEvent{Cmd: CmdCommonNoticeDanmaku, CommonNoticeDanmaku: &commonNoticeDanmaku, RoomId: roomId}
}

func (_ *Handler) SetTradingScore(msg map[string]string) MsgEvent {
	tradingScore := TradingScore{}
	err := json.Unmarshal([]byte(msg["msg"]), &tradingScore)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	roomId, err := strconv.Atoi(msg["RoomId"])
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	return MsgEvent{Cmd: CmdTradingScore, TradingScore: &tradingScore, RoomId: roomId}
}

func (_ *Handler) SetPreparing(msg map[string]string) MsgEvent {
	preparing := Preparing{}
	preparing.Cmd = CmdPreparing
	tmp := make(map[string]interface{})
	err := json.Unmarshal([]byte(msg["msg"]), &tmp)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	roomId, err := strconv.Atoi(msg["RoomId"])
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	preparing.RoomId = msg["RoomId"]
	return MsgEvent{Cmd: CmdPreparing, Preparing: &preparing, RoomId: roomId}
}

func (_ *Handler) SetGuardBuy(msg map[string]string) MsgEvent {
	guardBuy := GuardBuy{}
	err := json.Unmarshal([]byte(msg["msg"]), &guardBuy)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	roomId, err := strconv.Atoi(msg["RoomId"])
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	return MsgEvent{Cmd: CmdGuardBuy, GuardBuy: &guardBuy, RoomId: roomId}
}

func (_ *Handler) SetGiftStarProcess(msg map[string]string) MsgEvent {
	giftStarProcess := GiftStarProcess{}
	err := json.Unmarshal([]byte(msg["msg"]), &giftStarProcess)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	roomId, err := strconv.Atoi(msg["RoomId"])
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	return MsgEvent{Cmd: CmdGiftStarProcess, GiftStarProcess: &giftStarProcess, RoomId: roomId}
}

func (_ *Handler) SetRoomSkinMsg(msg map[string]string) MsgEvent {
	roomSkinMsg := RoomSkinMsg{}
	err := json.Unmarshal([]byte(msg["msg"]), &roomSkinMsg)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	roomId, err := strconv.Atoi(msg["RoomId"])
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	return MsgEvent{Cmd: CmdRoomSkinMsg, RoomSkinMsg: &roomSkinMsg, RoomId: roomId}
}

func (_ *Handler) SetEntryEffect(msg map[string]string) MsgEvent {
	enterEffect := EntryEffect{}
	err := json.Unmarshal([]byte(msg["msg"]), &enterEffect)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	roomId, err := strconv.Atoi(msg["RoomId"])
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	return MsgEvent{Cmd: CmdEntryEffect, EntryEffect: &enterEffect, RoomId: roomId}
}
