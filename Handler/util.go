package Handler

import (
	"encoding/json"
	"log"
	"strconv"
)

func SetDanMuMsg(msg map[string]interface{}) MsgEvent {
	danMu := DanMuMsg{}
	danMu.Cmd = CmdDanmuMsg
	danMu.Data.Content = msg["info"].([]interface{})[1].(string)
	danMu.Data.SendTimeStamp = int(msg["info"].([]interface{})[9].(map[string]interface{})["ts"].(float64))
	danMu.Data.SenderEnterRoomTimeStamp = int(msg["info"].([]interface{})[0].([]interface{})[4].(float64))
	danMu.Data.SendMillionTimeStamp = int64(msg["info"].([]interface{})[0].([]interface{})[5].(float64))
	danMu.Data.Sender.Uid = int64(msg["info"].([]interface{})[2].([]interface{})[0].(float64))
	danMu.Data.Sender.Name = msg["info"].([]interface{})[2].([]interface{})[1].(string)
	if len(msg["info"].([]interface{})[3].([]interface{})) != 0 {
		danMu.Data.Medal.GuardLevel = int(msg["info"].([]interface{})[3].([]interface{})[0].(float64))
		danMu.Data.Medal.MedalName = msg["info"].([]interface{})[3].([]interface{})[1].(string)
		danMu.Data.Medal.TargetID = int(msg["info"].([]interface{})[3].([]interface{})[11].(float64))
		danMu.Data.Medal.AnchorRoomId = int(msg["info"].([]interface{})[3].([]interface{})[3].(float64))
	}
	return MsgEvent{Cmd: CmdDanmuMsg, DanMuMsg: danMu, RoomId: msg["RoomId"].(int)}
}

func SetInteractWord(msg map[string]interface{}) MsgEvent {
	interactMsg := InteractWord{}
	interactMsg.Cmd = CmdInteractWord
	dataJson, err := json.Marshal(msg["data"])
	if err != nil {
		log.Printf("Marshal cmd json failed: %v", err)
	}
	err = json.Unmarshal(dataJson, &interactMsg.Data)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	return MsgEvent{Cmd: CmdInteractWord, InteractWord: interactMsg, RoomId: msg["RoomId"].(int)}
}

func SetOnlineRankCount(msg map[string]interface{}) MsgEvent {
	onlineRankCount := OnlineRankCount{}
	onlineRankCount.Cmd = CmdOnlineRankCount
	dataJson, err := json.Marshal(msg["data"])
	if err != nil {
		log.Printf("Marshal cmd json failed: %v", err)
	}
	err = json.Unmarshal(dataJson, &onlineRankCount.Data)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	return MsgEvent{Cmd: CmdOnlineRankCount, OnlineRankCount: onlineRankCount, RoomId: msg["RoomId"].(int)}
}

func SetWatchedChange(msg map[string]interface{}) MsgEvent {
	watchedChange := WatchedChange{}
	watchedChange.Cmd = CmdWatchedChange
	dataJson, err := json.Marshal(msg["data"])
	if err != nil {
		log.Printf("Marshal cmd json failed: %v", err)
	}
	err = json.Unmarshal(dataJson, &watchedChange.Data)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	return MsgEvent{Cmd: CmdWatchedChange, WatchedChange: watchedChange, RoomId: msg["RoomId"].(int)}
}

func SetNoticeMsg(msg map[string]interface{}) MsgEvent {
	noticeMsg := NoticeMsg{}
	switch msg["real_roomid"].(type) {
	case float64:
		msg["real_roomid"] = strconv.FormatFloat(msg["real_roomid"].(float64), 'f', -1, 64)
	case int:
		msg["real_roomid"] = strconv.Itoa(msg["real_roomid"].(int))
	}
	dataJson, err := json.Marshal(msg)
	if err != nil {
		log.Printf("Marshal cmd json failed: %v", err)
	}
	err = json.Unmarshal(dataJson, &noticeMsg)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	return MsgEvent{Cmd: CmdNoticeMsg, NoticeMsg: noticeMsg, RoomId: msg["RoomId"].(int)}
}

func SetSuperChatMessage(msg map[string]interface{}) MsgEvent {
	superChatMsg := SuperChatMessage{}
	superChatMsg.Cmd = CmdSuperChatMessage
	switch msg["data"].(map[string]interface{})["id"].(type) {
	case float64:
		msg["data"].(map[string]interface{})["id"] = strconv.FormatFloat(msg["data"].(map[string]interface{})["id"].(float64), 'f', -1, 64)
	case int:
		msg["data"].(map[string]interface{})["id"] = strconv.Itoa(msg["data"].(map[string]interface{})["id"].(int))
	}
	switch msg["data"].(map[string]interface{})["uid"].(type) {
	case float64:
		msg["data"].(map[string]interface{})["uid"] = strconv.FormatFloat(msg["data"].(map[string]interface{})["uid"].(float64), 'f', -1, 64)
	case int:
		msg["data"].(map[string]interface{})["uid"] = strconv.Itoa(msg["data"].(map[string]interface{})["uid"].(int))
	}

	dataJson, err := json.Marshal(msg["data"])
	if err != nil {
		log.Printf("Marshal cmd json failed: %v", err)
	}
	err = json.Unmarshal(dataJson, &superChatMsg.Data)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	return MsgEvent{Cmd: CmdSuperChatMessage, SuperChatMessage: superChatMsg, RoomId: msg["RoomId"].(int)}
}

func SetSendGift(msg map[string]interface{}) MsgEvent {
	sendGift := SendGift{}
	sendGift.Cmd = CmdSendGift
	dataJson, err := json.Marshal(msg["data"])
	if err != nil {
		log.Printf("Marshal cmd json failed: %v", err)
	}
	err = json.Unmarshal(dataJson, &sendGift.Data)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	return MsgEvent{Cmd: CmdSendGift, SendGift: sendGift, RoomId: msg["RoomId"].(int)}
}

func SetOnlineRankV2(msg map[string]interface{}) MsgEvent {
	onlineRankV2 := OnlineRankV2{}
	onlineRankV2.Cmd = CmdOnlineRankV2
	dataJson, err := json.Marshal(msg["data"])
	if err != nil {
		log.Printf("Marshal cmd json failed: %v", err)
	}
	err = json.Unmarshal(dataJson, &onlineRankV2.Data)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	return MsgEvent{Cmd: CmdOnlineRankV2, OnlineRankV2: onlineRankV2, RoomId: msg["RoomId"].(int)}
}

func SetOnlineRankTop3(msg map[string]interface{}) MsgEvent {
	onlineRankTop3 := OnlineRankTop3{}
	onlineRankTop3.Cmd = CmdOnlineRankTop3
	dataJson, err := json.Marshal(msg["data"])
	if err != nil {
		log.Printf("Marshal cmd json failed: %v", err)
	}
	err = json.Unmarshal(dataJson, &onlineRankTop3.Data)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	return MsgEvent{Cmd: CmdOnlineRankTop3, OnlineRankTop3: onlineRankTop3, RoomId: msg["RoomId"].(int)}
}

func SetLikeInfoV3Click(msg map[string]interface{}) MsgEvent {
	likeInfoV3Click := LikeInfoV3Click{}
	likeInfoV3Click.Cmd = CmdLikeInfoV3Click
	dataJson, err := json.Marshal(msg["data"])
	if err != nil {
		log.Printf("Marshal cmd json failed: %v", err)
	}
	err = json.Unmarshal(dataJson, &likeInfoV3Click.Data)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	return MsgEvent{Cmd: CmdLikeInfoV3Click, LikeInfoV3Click: likeInfoV3Click, RoomId: msg["RoomId"].(int)}
}

func SetStopLiveRoomList(msg map[string]interface{}) MsgEvent {
	stopLiveRoomList := StopLiveRoomList{}
	stopLiveRoomList.Cmd = CmdStopLiveRoomList
	dataJson, err := json.Marshal(msg["data"])
	if err != nil {
		log.Printf("Marshal cmd json failed: %v", err)
	}
	err = json.Unmarshal(dataJson, &stopLiveRoomList.Data)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	return MsgEvent{Cmd: CmdStopLiveRoomList, StopLiveRoomList: stopLiveRoomList, RoomId: msg["RoomId"].(int)}
}

func SetLikeInfoV3Update(msg map[string]interface{}) MsgEvent {
	likeInfoV3Update := LikeInfoV3Update{}
	likeInfoV3Update.Cmd = CmdLikeInfoV3Update
	dataJson, err := json.Marshal(msg["data"])
	if err != nil {
		log.Printf("Marshal cmd json failed: %v", err)
	}
	err = json.Unmarshal(dataJson, &likeInfoV3Update.Data)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	return MsgEvent{Cmd: CmdLikeInfoV3Update, LikeInfoV3Update: likeInfoV3Update, RoomId: msg["RoomId"].(int)}
}

func SetHotRankChange(msg map[string]interface{}) MsgEvent {
	hotRankChange := HotRankChange{}
	hotRankChange.Cmd = CmdHotRankChange
	dataJson, err := json.Marshal(msg["data"])
	if err != nil {
		log.Printf("Marshal cmd json failed: %v", err)
	}
	err = json.Unmarshal(dataJson, &hotRankChange.Data)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	return MsgEvent{Cmd: CmdHotRankChange, HotRankChange: hotRankChange, RoomId: msg["RoomId"].(int)}
}

func SetRoomRealTimeMessageUpdate(msg map[string]interface{}) MsgEvent {
	roomRealTimeMessageUpdate := RoomRealTimeMessageUpdate{}
	roomRealTimeMessageUpdate.Cmd = CmdRoomRealTimeMessageUpdate
	dataJson, err := json.Marshal(msg["data"])
	if err != nil {
		log.Printf("Marshal cmd json failed: %v", err)
	}
	err = json.Unmarshal(dataJson, &roomRealTimeMessageUpdate.Data)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	return MsgEvent{Cmd: CmdRoomRealTimeMessageUpdate, RoomRealTimeMessageUpdate: roomRealTimeMessageUpdate, RoomId: msg["RoomId"].(int)}
}

func SetWidgetBanner(msg map[string]interface{}) MsgEvent {
	widgetBanner := WidgetBanner{}
	widgetBanner.Cmd = CmdWidgetBanner
	dataJson, err := json.Marshal(msg["data"])
	if err != nil {
		log.Printf("Marshal cmd json failed: %v", err)
	}
	err = json.Unmarshal(dataJson, &widgetBanner.Data)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	return MsgEvent{Cmd: CmdWidgetBanner, WidgetBanner: widgetBanner, RoomId: msg["RoomId"].(int)}
}

func SetHotRankChangedV2(msg map[string]interface{}) MsgEvent {
	hotRankChangedV2 := HotRankChangedV2{}
	hotRankChangedV2.Cmd = CmdHotRankChangedV2
	dataJson, err := json.Marshal(msg["data"])
	if err != nil {
		log.Printf("Marshal cmd json failed: %v", err)
	}
	err = json.Unmarshal(dataJson, &hotRankChangedV2.Data)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	return MsgEvent{Cmd: CmdHotRankChangedV2, HotRankChangedV2: hotRankChangedV2, RoomId: msg["RoomId"].(int)}
}

func SetGuardHonorThousand(msg map[string]interface{}) MsgEvent {
	guardHonorThousand := GuardHonorThousand{}
	guardHonorThousand.Cmd = CmdGuardHonorThousand
	dataJson, err := json.Marshal(msg["data"])
	if err != nil {
		log.Printf("Marshal cmd json failed: %v", err)
	}
	err = json.Unmarshal(dataJson, &guardHonorThousand.Data)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	return MsgEvent{Cmd: CmdGuardHonorThousand, GuardHonorThousand: guardHonorThousand, RoomId: msg["RoomId"].(int)}
}

func SetLive(msg map[string]interface{}) MsgEvent {
	live := Live{}
	dataJson, err := json.Marshal(msg)
	if err != nil {
		log.Printf("Marshal cmd json failed: %v", err)
	}
	err = json.Unmarshal(dataJson, &live)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	return MsgEvent{Cmd: CmdLive, Live: live, RoomId: msg["RoomId"].(int)}
}

func SetRoomChange(msg map[string]interface{}) MsgEvent {
	roomChange := RoomChange{}
	roomChange.Cmd = CmdRoomChange
	dataJson, err := json.Marshal(msg["data"])
	if err != nil {
		log.Printf("Marshal cmd json failed: %v", err)
	}
	err = json.Unmarshal(dataJson, &roomChange.Data)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	return MsgEvent{Cmd: CmdRoomChange, RoomChange: roomChange, RoomId: msg["RoomId"].(int)}
}

func SetRoomBlockMsg(msg map[string]interface{}) MsgEvent {
	roomBlockMsg := RoomBlockMsg{}
	roomBlockMsg.Cmd = CmdRoomBlockMsg
	dataJson, err := json.Marshal(msg["data"])
	if err != nil {
		log.Printf("Marshal cmd json failed: %v", err)
	}
	err = json.Unmarshal(dataJson, &roomBlockMsg.Data)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	if _, ok := msg["uid"]; ok {
		roomBlockMsg.UID = msg["uid"].(string)
	}
	if _, ok := msg["name"]; ok {
		roomBlockMsg.Name = msg["name"].(string)
	}
	return MsgEvent{Cmd: CmdRoomBlockMsg, RoomBlockMsg: roomBlockMsg, RoomId: msg["RoomId"].(int)}
}

func SetFullScreenSpecialEffect(msg map[string]interface{}) MsgEvent {
	fullScreenSpecialEffect := FullScreenSpecialEffect{}
	fullScreenSpecialEffect.Cmd = CmdFullScreenSpecialEffect
	dataJson, err := json.Marshal(msg["data"])
	if err != nil {
		log.Printf("Marshal cmd json failed: %v", err)
	}
	err = json.Unmarshal(dataJson, &fullScreenSpecialEffect.Data)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	return MsgEvent{Cmd: CmdFullScreenSpecialEffect, FullScreenSpecialEffect: fullScreenSpecialEffect, RoomId: msg["RoomId"].(int)}
}

func SetCommonNoticeDanmaku(msg map[string]interface{}) MsgEvent {
	commonNoticeDanmaku := CommonNoticeDanmaku{}
	commonNoticeDanmaku.Cmd = CmdCommonNoticeDanmaku
	dataJson, err := json.Marshal(msg["data"])
	if err != nil {
		log.Printf("Marshal cmd json failed: %v", err)
	}
	err = json.Unmarshal(dataJson, &commonNoticeDanmaku.Data)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	return MsgEvent{Cmd: CmdCommonNoticeDanmaku, CommonNoticeDanmaku: commonNoticeDanmaku, RoomId: msg["RoomId"].(int)}
}

func SetTradingScore(msg map[string]interface{}) MsgEvent {
	tradingScore := TradingScore{}
	tradingScore.Cmd = CmdTradingScore
	dataJson, err := json.Marshal(msg["data"])
	if err != nil {
		log.Printf("Marshal cmd json failed: %v", err)
	}
	err = json.Unmarshal(dataJson, &tradingScore.Data)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	return MsgEvent{Cmd: CmdTradingScore, TradingScore: tradingScore, RoomId: msg["RoomId"].(int)}
}

func SetPreparing(msg map[string]interface{}) MsgEvent {
	preparing := Preparing{}
	preparing.Cmd = CmdPreparing
	preparing.RoomId = strconv.Itoa(msg["RoomId"].(int))
	return MsgEvent{Cmd: CmdPreparing, Preparing: preparing, RoomId: msg["RoomId"].(int)}
}

func SetGuardBuy(msg map[string]interface{}) MsgEvent {
	guardBuy := GuardBuy{}
	guardBuy.Cmd = CmdGuardBuy
	dataJson, err := json.Marshal(msg["data"])
	if err != nil {
		log.Printf("Marshal cmd json failed: %v", err)
	}
	err = json.Unmarshal(dataJson, &guardBuy.Data)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	return MsgEvent{Cmd: CmdGuardBuy, GuardBuy: guardBuy, RoomId: msg["RoomId"].(int)}
}

func SetGiftStarProcess(msg map[string]interface{}) MsgEvent {
	giftStarProcess := GiftStarProcess{}
	giftStarProcess.Cmd = CmdGiftStarProcess
	dataJson, err := json.Marshal(msg["data"])
	if err != nil {
		log.Printf("Marshal cmd json failed: %v", err)
	}
	err = json.Unmarshal(dataJson, &giftStarProcess.Data)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	return MsgEvent{Cmd: CmdGiftStarProcess, GiftStarProcess: giftStarProcess, RoomId: msg["RoomId"].(int)}
}

func SetRoomSkinMsg(msg map[string]interface{}) MsgEvent {
	roomSkinMsg := RoomSkinMsg{}
	dataJson, err := json.Marshal(msg)
	if err != nil {
		log.Printf("Marshal cmd json failed: %v", err)
	}
	err = json.Unmarshal(dataJson, &roomSkinMsg)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	return MsgEvent{Cmd: CmdRoomSkinMsg, RoomSkinMsg: roomSkinMsg, RoomId: msg["RoomId"].(int)}
}

func SetEnterEffect(msg map[string]interface{}) MsgEvent {
	enterEffect := EnterEffect{}
	enterEffect.Cmd = CmdEnterEffect
	dataJson, err := json.Marshal(msg["data"])
	if err != nil {
		log.Printf("Marshal cmd json failed: %v", err)
	}
	err = json.Unmarshal(dataJson, &enterEffect.Data)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return MsgEvent{}
	}
	return MsgEvent{Cmd: CmdEnterEffect, EnterEffect: enterEffect, RoomId: msg["RoomId"].(int)}
}
