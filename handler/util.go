package handler

import (
	"strconv"
)

// SetDanMuMsg 设置弹幕消息
// 该消息为list结构, 部分字段含义未知, 因此目前只有部分内容
func (*Handler) SetDanMuMsg(msg map[string]interface{}) (m MsgEvent) {
	danMu := DanMuMsg{}
	danMu.Cmd = CmdDanmuMsg
	danMuMsg := make(map[string]interface{}, 0)
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &danMuMsg); err == nil {
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
func (*Handler) SetInteractWord(msg map[string]interface{}) (m MsgEvent) {
	interactMsg := InteractWord{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &interactMsg); err == nil {
		m = MsgEvent{Cmd: CmdInteractWord, InteractWord: &interactMsg, RoomId: msg["RoomId"].(int)}
	}
	return
}

// SetOnlineRankCount 暂时未知
func (*Handler) SetOnlineRankCount(msg map[string]interface{}) (m MsgEvent) {
	onlineRankCount := OnlineRankCount{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &onlineRankCount); err == nil {
		m = MsgEvent{Cmd: CmdOnlineRankCount, OnlineRankCount: &onlineRankCount, RoomId: msg["RoomId"].(int)}
	}
	return
}

// SetWatchedChange 暂时未知
func (*Handler) SetWatchedChange(msg map[string]interface{}) (m MsgEvent) {
	watchedChange := WatchedChange{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &watchedChange); err == nil {
		m = MsgEvent{Cmd: CmdWatchedChange, WatchedChange: &watchedChange, RoomId: msg["RoomId"].(int)}
	}
	return
}

// SetNoticeMsg 可能为系统消息
// TODO: 尝试优化
func (*Handler) SetNoticeMsg(msg map[string]interface{}) (m MsgEvent) {
	noticeMsg := NoticeMsg{}
	notice := make(map[string]interface{}, 0)
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &notice); err == nil {
		// 这个字段很奇怪, 类型不确定, 需要特判
		switch notice["real_roomid"].(type) {
		case float64:
			notice["real_roomid"] = strconv.FormatFloat(notice["real_roomid"].(float64), 'f', -1, 64)
		case int:
			notice["real_roomid"] = strconv.Itoa(notice["real_roomid"].(int))
		}
		if dataJson, err := JsonCoder.Marshal(notice); err == nil {
			if err = JsonCoder.Unmarshal(dataJson, &noticeMsg); err == nil {
				m = MsgEvent{Cmd: CmdNoticeMsg, NoticeMsg: &noticeMsg, RoomId: msg["RoomId"].(int)}
			}
		}
	}
	return
}

// SetSuperChatMessage 超级留言
// TODO: 尝试优化
func (*Handler) SetSuperChatMessage(msg map[string]interface{}) (m MsgEvent) {
	superChatMsg := SuperChatMessage{}
	superChatMsg.Cmd = CmdSuperChatMessage
	sc := make(map[string]interface{}, 0)
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &sc); err == nil {
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
		if dataJson, err := JsonCoder.Marshal(sc["data"]); err == nil {
			if err = JsonCoder.Unmarshal(dataJson, &superChatMsg.Data); err == nil {
				m = MsgEvent{Cmd: CmdSuperChatMessage, SuperChatMessage: &superChatMsg, RoomId: msg["RoomId"].(int)}
			}
		}
	}
	return
}

// SetSendGift 赠送礼物
func (*Handler) SetSendGift(msg map[string]interface{}) (m MsgEvent) {
	sendGift := SendGift{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &sendGift); err == nil {
		m = MsgEvent{Cmd: CmdSendGift, SendGift: &sendGift, RoomId: msg["RoomId"].(int)}
	}
	return
}

// SetOnlineRankV2 未知
func (*Handler) SetOnlineRankV2(msg map[string]interface{}) (m MsgEvent) {
	onlineRankV2 := OnlineRankV2{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &onlineRankV2); err == nil {
		m = MsgEvent{Cmd: CmdOnlineRankV2, OnlineRankV2: &onlineRankV2, RoomId: msg["RoomId"].(int)}
	}
	return
}

// SetOnlineRankTop3 未知
func (*Handler) SetOnlineRankTop3(msg map[string]interface{}) (m MsgEvent) {
	onlineRankTop3 := OnlineRankTop3{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &onlineRankTop3); err == nil {
		m = MsgEvent{Cmd: CmdOnlineRankTop3, OnlineRankTop3: &onlineRankTop3, RoomId: msg["RoomId"].(int)}
	}
	return
}

// SetLikeInfoV3Click 可能为点赞
func (*Handler) SetLikeInfoV3Click(msg map[string]interface{}) (m MsgEvent) {
	likeInfoV3Click := LikeInfoV3Click{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &likeInfoV3Click); err == nil {
		m = MsgEvent{Cmd: CmdLikeInfoV3Click, LikeInfoV3Click: &likeInfoV3Click, RoomId: msg["RoomId"].(int)}
	}
	return
}

// SetStopLiveRoomList 未知
func (*Handler) SetStopLiveRoomList(msg map[string]interface{}) (m MsgEvent) {
	stopLiveRoomList := StopLiveRoomList{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &stopLiveRoomList); err == nil {
		m = MsgEvent{Cmd: CmdStopLiveRoomList, StopLiveRoomList: &stopLiveRoomList, RoomId: msg["RoomId"].(int)}
	}
	return
}

// SetLikeInfoV3Update 未知
func (*Handler) SetLikeInfoV3Update(msg map[string]interface{}) (m MsgEvent) {
	likeInfoV3Update := LikeInfoV3Update{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &likeInfoV3Update); err == nil {
		m = MsgEvent{Cmd: CmdLikeInfoV3Update, LikeInfoV3Update: &likeInfoV3Update, RoomId: msg["RoomId"].(int)}
	}
	return
}

// SetHotRankChange 未知
func (*Handler) SetHotRankChange(msg map[string]interface{}) (m MsgEvent) {
	hotRankChange := HotRankChange{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &hotRankChange); err == nil {
		m = MsgEvent{Cmd: CmdHotRankChange, HotRankChange: &hotRankChange, RoomId: msg["RoomId"].(int)}
	}
	return
}

// SetRoomRealTimeMessageUpdate 未知
func (*Handler) SetRoomRealTimeMessageUpdate(msg map[string]interface{}) (m MsgEvent) {
	roomRealTimeMessageUpdate := RoomRealTimeMessageUpdate{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &roomRealTimeMessageUpdate); err == nil {
		m = MsgEvent{Cmd: CmdRoomRealTimeMessageUpdate, RoomRealTimeMessageUpdate: &roomRealTimeMessageUpdate, RoomId: msg["RoomId"].(int)}
	}
	return
}

// SetWidgetBanner 未知
func (*Handler) SetWidgetBanner(msg map[string]interface{}) (m MsgEvent) {
	widgetBanner := WidgetBanner{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &widgetBanner); err == nil {
		m = MsgEvent{Cmd: CmdWidgetBanner, WidgetBanner: &widgetBanner, RoomId: msg["RoomId"].(int)}
	}
	return
}

// SetHotRankChangedV2 未知
func (*Handler) SetHotRankChangedV2(msg map[string]interface{}) (m MsgEvent) {
	hotRankChangedV2 := HotRankChangedV2{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &hotRankChangedV2); err == nil {
		m = MsgEvent{Cmd: CmdHotRankChangedV2, HotRankChangedV2: &hotRankChangedV2, RoomId: msg["RoomId"].(int)}
	}
	return
}

// SetGuardHonorThousand 未知
func (*Handler) SetGuardHonorThousand(msg map[string]interface{}) (m MsgEvent) {
	guardHonorThousand := GuardHonorThousand{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &guardHonorThousand); err == nil {
		m = MsgEvent{Cmd: CmdGuardHonorThousand, GuardHonorThousand: &guardHonorThousand, RoomId: msg["RoomId"].(int)}
	}
	return
}

// SetLive 开始直播
func (*Handler) SetLive(msg map[string]interface{}) (m MsgEvent) {
	live := Live{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &live); err == nil {
		m = MsgEvent{Cmd: CmdLive, Live: &live, RoomId: msg["RoomId"].(int)}
	}
	return
}

// SetRoomChange 未知
func (*Handler) SetRoomChange(msg map[string]interface{}) (m MsgEvent) {
	roomChange := RoomChange{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &roomChange); err == nil {
		m = MsgEvent{Cmd: CmdRoomChange, RoomChange: &roomChange, RoomId: msg["RoomId"].(int)}
	}
	return
}

// SetRoomBlockMsg 未知
func (*Handler) SetRoomBlockMsg(msg map[string]interface{}) (m MsgEvent) {
	roomBlockMsg := RoomBlockMsg{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &roomBlockMsg); err == nil {
		m = MsgEvent{Cmd: CmdRoomBlockMsg, RoomBlockMsg: &roomBlockMsg, RoomId: msg["RoomId"].(int)}
	}
	return
}

// SetFullScreenSpecialEffect 可能为礼物特效
func (*Handler) SetFullScreenSpecialEffect(msg map[string]interface{}) (m MsgEvent) {
	fullScreenSpecialEffect := FullScreenSpecialEffect{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &fullScreenSpecialEffect); err == nil {
		m = MsgEvent{Cmd: CmdFullScreenSpecialEffect, FullScreenSpecialEffect: &fullScreenSpecialEffect, RoomId: msg["RoomId"].(int)}
	}
	return
}

// SetCommonNoticeDanmaku 未知
func (*Handler) SetCommonNoticeDanmaku(msg map[string]interface{}) (m MsgEvent) {
	commonNoticeDanmaku := CommonNoticeDanmaku{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &commonNoticeDanmaku); err == nil {
		m = MsgEvent{Cmd: CmdCommonNoticeDanmaku, CommonNoticeDanmaku: &commonNoticeDanmaku, RoomId: msg["RoomId"].(int)}
	}
	return
}

// SetTradingScore 未知
func (*Handler) SetTradingScore(msg map[string]interface{}) (m MsgEvent) {
	tradingScore := TradingScore{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &tradingScore); err == nil {
		m = MsgEvent{Cmd: CmdTradingScore, TradingScore: &tradingScore, RoomId: msg["RoomId"].(int)}
	}
	return
}

// SetPreparing 开始准备
func (*Handler) SetPreparing(msg map[string]interface{}) (m MsgEvent) {
	preparing := Preparing{}
	preparing.Cmd = CmdPreparing
	tmp := make(map[string]interface{})
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &tmp); err == nil {
		preparing.RoomId = msg["RoomId"].(int)
		m = MsgEvent{Cmd: CmdPreparing, Preparing: &preparing, RoomId: msg["RoomId"].(int)}
	}
	return
}

// SetGuardBuy 大航海购买
func (*Handler) SetGuardBuy(msg map[string]interface{}) (m MsgEvent) {
	guardBuy := GuardBuy{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &guardBuy); err == nil {
		m = MsgEvent{Cmd: CmdGuardBuy, GuardBuy: &guardBuy, RoomId: msg["RoomId"].(int)}
	}
	return
}

// SetGiftStarProcess 未知
func (*Handler) SetGiftStarProcess(msg map[string]interface{}) (m MsgEvent) {
	giftStarProcess := GiftStarProcess{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &giftStarProcess); err == nil {
		m = MsgEvent{Cmd: CmdGiftStarProcess, GiftStarProcess: &giftStarProcess, RoomId: msg["RoomId"].(int)}
	}
	return
}

// SetRoomSkinMsg 未知
func (*Handler) SetRoomSkinMsg(msg map[string]interface{}) (m MsgEvent) {
	roomSkinMsg := RoomSkinMsg{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &roomSkinMsg); err == nil {
		m = MsgEvent{Cmd: CmdRoomSkinMsg, RoomSkinMsg: &roomSkinMsg, RoomId: msg["RoomId"].(int)}
	}
	return
}

// SetEntryEffect 未知
func (*Handler) SetEntryEffect(msg map[string]interface{}) (m MsgEvent) {
	enterEffect := EntryEffect{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &enterEffect); err == nil {
		m = MsgEvent{Cmd: CmdEntryEffect, EntryEffect: &enterEffect, RoomId: msg["RoomId"].(int)}
	}
	return
}

// SetUserToastMsg 上舰长
func (*Handler) SetUserToastMsg(msg map[string]interface{}) (m MsgEvent) {
	userToastMsg := UserToastMsg{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &userToastMsg); err == nil {
		m = MsgEvent{Cmd: CmdUserToastMsg, UserToastMsg: &userToastMsg, RoomId: msg["RoomId"].(int)}
	}
	return
}

func (*Handler) SetHeartBeatReply(msg map[string]interface{}) (m MsgEvent) {
	heartBeatReply := HeartBeatReply{Sum: msg["msg"].(int)}
	m = MsgEvent{Cmd: CmdHeartBeatReply, RoomId: msg["RoomId"].(int), HeartBeatReply: &heartBeatReply}
	return
}

func (*Handler) SetPopularityRedPocketNew(msg map[string]interface{}) (m MsgEvent) {
	popularityRedPocketNew := PopularityRedPocketNew{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &popularityRedPocketNew); err == nil {
		m = MsgEvent{Cmd: CmdPopularityRedPocketNew, RoomId: msg["RoomId"].(int), PopularityRedPocketNew: &popularityRedPocketNew}
	}
	return
}

func (*Handler) SetAreaRankChanged(msg map[string]interface{}) (m MsgEvent) {
	areaRankChanged := AreaRankChanged{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &areaRankChanged); err == nil {
		m = MsgEvent{Cmd: CmdAreaRankChanged, RoomId: msg["RoomId"].(int), AreaRankChanged: &areaRankChanged}
	}
	return
}

func (*Handler) SetSuperChatEntrance(msg map[string]interface{}) (m MsgEvent) {
	superChatEntrance := SuperChatEntrance{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &superChatEntrance); err == nil {
		m = MsgEvent{Cmd: CmdSuperChatEntrance, RoomId: msg["RoomId"].(int), SuperChatEntrance: &superChatEntrance}
	}
	return
}

func (*Handler) SetPlayTogether(msg map[string]interface{}) (m MsgEvent) {
	playTogether := PlayTogether{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &playTogether); err == nil {
		m = MsgEvent{Cmd: CmdPlayTogether, RoomId: msg["RoomId"].(int), PlayTogether: &playTogether}
	}
	return
}

func (*Handler) SetComboSend(msg map[string]interface{}) (m MsgEvent) {
	comboSend := ComboSend{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &comboSend); err == nil {
		m = MsgEvent{Cmd: CmdComboSend, RoomId: msg["RoomId"].(int), ComboSend: &comboSend}
	}
	return
}

func (*Handler) SetPopularityRedPocketStart(msg map[string]interface{}) (m MsgEvent) {
	popularityRedPocketStart := PopularityRedPocketStart{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &popularityRedPocketStart); err == nil {
		m = MsgEvent{Cmd: CmdPopularityRedPocketStart, RoomId: msg["RoomId"].(int), PopularityRedPocketStart: &popularityRedPocketStart}
	}
	return
}

func (*Handler) SetPkBattleProcess(msg map[string]interface{}) (m MsgEvent) {
	pkBattleProcess := PkBattleProcess{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &pkBattleProcess); err == nil {
		m = MsgEvent{Cmd: CmdPkBattleProcess, RoomId: msg["RoomId"].(int), PkBattleProcess: &pkBattleProcess}
	}
	return
}

func (*Handler) SetPopularRankChanged(msg map[string]interface{}) (m MsgEvent) {
	popularRankChanged := PopularRankChanged{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &popularRankChanged); err == nil {
		m = MsgEvent{Cmd: CmdPopularRankChanged, RoomId: msg["RoomId"].(int), PopularRankChanged: &popularRankChanged}
	}
	return
}

func (*Handler) SetPkBattleStartNew(msg map[string]interface{}) (m MsgEvent) {
	pkBattleStartNew := PkBattleStartNew{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &pkBattleStartNew); err == nil {
		m = MsgEvent{Cmd: CmdPkBattleStartNew, RoomId: msg["RoomId"].(int), PkBattleStartNew: &pkBattleStartNew}
	}
	return
}

func (*Handler) SetDanMuAggregation(msg map[string]interface{}) (m MsgEvent) {
	danMuAggregation := DanMuAggregation{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &danMuAggregation); err == nil {
		m = MsgEvent{Cmd: CmdDanMuAggregation, RoomId: msg["RoomId"].(int), DanMuAggregation: &danMuAggregation}
	}
	return
}

func (*Handler) SetLiveInteractiveGame(msg map[string]interface{}) (m MsgEvent) {
	liveInteractiveGame := LiveInteractiveGame{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &liveInteractiveGame); err == nil {
		m = MsgEvent{Cmd: CmdLiveInteractiveGame, RoomId: msg["RoomId"].(int), LiveInteractiveGame: &liveInteractiveGame}
	}
	return
}

func (*Handler) SetRecommendCar(msg map[string]interface{}) (m MsgEvent) {
	recommendCard := RecommendCard{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &recommendCard); err == nil {
		m = MsgEvent{Cmd: CmdRecommendCard, RoomId: msg["RoomId"].(int), RecommendCard: &recommendCard}
	}
	return
}

func (*Handler) SetPkBattleProcessNew(msg map[string]interface{}) (m MsgEvent) {
	pkBattleProcessNew := PkBattleProcessNew{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &pkBattleProcessNew); err == nil {
		m = MsgEvent{Cmd: CmdPkBattleProcessNew, RoomId: msg["RoomId"].(int), PkBattleProcessNew: &pkBattleProcessNew}
	}
	return
}

func (*Handler) SetPkBattlePreNew(msg map[string]interface{}) (m MsgEvent) {
	pkBattlePreNew := PkBattlePreNew{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &pkBattlePreNew); err == nil {
		m = MsgEvent{Cmd: CmdPkBattlePreNew, RoomId: msg["RoomId"].(int), PkBattlePreNew: &pkBattlePreNew}
	}
	return
}

func (*Handler) SetPkBattlePre(msg map[string]interface{}) (m MsgEvent) {
	pkBattlePre := PkBattlePre{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &pkBattlePre); err == nil {
		m = MsgEvent{Cmd: CmdPkBattlePre, RoomId: msg["RoomId"].(int), PkBattlePre: &pkBattlePre}
	}
	return
}

func (*Handler) SetPkBattleFinalProcess(msg map[string]interface{}) (m MsgEvent) {
	pkBattleFinalProcess := PkBattleFinalProcess{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &pkBattleFinalProcess); err == nil {
		m = MsgEvent{Cmd: CmdPkBattleFinalProcess, RoomId: msg["RoomId"].(int), PkBattleFinalProcess: &pkBattleFinalProcess}
	}
	return
}

func (*Handler) SetPkBattleStart(msg map[string]interface{}) (m MsgEvent) {
	pkBattleStart := PkBattleStart{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &pkBattleStart); err == nil {
		m = MsgEvent{Cmd: CmdPkBattleStart, RoomId: msg["RoomId"].(int), PkBattleStart: &pkBattleStart}
	}
	return
}

func (*Handler) SetWidgetGiftStarProcess(msg map[string]interface{}) (m MsgEvent) {
	widgetGiftStarProcess := WidgetGiftStarProcess{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &widgetGiftStarProcess); err == nil {
		m = MsgEvent{Cmd: CmdWidgetGiftStarProcess, RoomId: msg["RoomId"].(int), WidgetGiftStarProcess: &widgetGiftStarProcess}
	}
	return
}

func (*Handler) SetPopularityRedPocketWinnerList(msg map[string]interface{}) (m MsgEvent) {
	popularityRedPocketWinnerList := PopularityRedPocketWinnerList{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &popularityRedPocketWinnerList); err == nil {
		m = MsgEvent{Cmd: CmdPopularityRedPocketWinnerList, RoomId: msg["RoomId"].(int), PopularityRedPocketWinnerList: &popularityRedPocketWinnerList}
	}
	return
}

func (*Handler) SetGotoBuyFlow(msg map[string]interface{}) (m MsgEvent) {
	gotoBuyFlow := GotoBuyFlow{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &gotoBuyFlow); err == nil {
		m = MsgEvent{Cmd: CmdGotoBuyFlow, RoomId: msg["RoomId"].(int), GotoBuyFlow: &gotoBuyFlow}
	}
	return
}

func (*Handler) SetPkBattleEnd(msg map[string]interface{}) (m MsgEvent) {
	pkBattleEndNew := PkBattleEnd{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &pkBattleEndNew); err == nil {
		m = MsgEvent{Cmd: CmdPkBattleEnd, RoomId: msg["RoomId"].(int), PkBattleEnd: &pkBattleEndNew}
	}
	return
}

func (*Handler) SetPkBattleSettleUser(msg map[string]interface{}) (m MsgEvent) {
	pkBattleSettleUser := PkBattleSettleUser{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &pkBattleSettleUser); err == nil {
		m = MsgEvent{Cmd: CmdPkBattleSettleUser, RoomId: msg["RoomId"].(int), PkBattleSettleUser: &pkBattleSettleUser}
	}
	return
}

func (*Handler) SetAnchorLotStart(msg map[string]interface{}) (m MsgEvent) {
	anchorLotStart := AnchorLotStart{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &anchorLotStart); err == nil {
		m = MsgEvent{Cmd: CmdAnchorLotStart, RoomId: msg["RoomId"].(int), AnchorLotStart: &anchorLotStart}
	}
	return
}

func (*Handler) SetPkBattleSettleV2(msg map[string]interface{}) (m MsgEvent) {
	pkBattleSettleV2 := PkBattleSettleV2{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &pkBattleSettleV2); err == nil {
		m = MsgEvent{Cmd: CmdPkBattleSettleV2, RoomId: msg["RoomId"].(int), PkBattleSettleV2: &pkBattleSettleV2}
	}
	return
}

func (*Handler) SetPkBattleSettle(msg map[string]interface{}) (m MsgEvent) {
	pkBattleSettle := PkBattleSettle{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &pkBattleSettle); err == nil {
		m = MsgEvent{Cmd: CmdPkBattleSettle, RoomId: msg["RoomId"].(int), PkBattleSettle: &pkBattleSettle}
	}
	return
}

func (*Handler) SetHotRoomNotify(msg map[string]interface{}) (m MsgEvent) {
	hotRoomNotify := HotRoomNotify{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &hotRoomNotify); err == nil {
		m = MsgEvent{Cmd: CmdHotRoomNotify, RoomId: msg["RoomId"].(int), HotRoomNotify: &hotRoomNotify}
	}
	return
}

func (*Handler) SetLiveOpenPlatformGame(msg map[string]interface{}) (m MsgEvent) {
	liveOpenPlatformGame := LiveOpenPlatformGame{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &liveOpenPlatformGame); err == nil {
		m = MsgEvent{Cmd: CmdLiveOpenPlatformGame, RoomId: msg["RoomId"].(int), LiveOpenPlatformGame: &liveOpenPlatformGame}
	}
	return
}

func (*Handler) SetLivePanelChangeContent(msg map[string]interface{}) (m MsgEvent) {
	livePanelChangeContent := LivePanelChangeContent{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &livePanelChangeContent); err == nil {
		m = MsgEvent{Cmd: CmdLivePanelChangeContent, RoomId: msg["RoomId"].(int), LivePanelChangeContent: &livePanelChangeContent}
	}
	return
}

func (*Handler) SetGiftPanelPlan(msg map[string]interface{}) (m MsgEvent) {
	giftPanelPlan := GiftPanelPlan{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &giftPanelPlan); err == nil {
		m = MsgEvent{Cmd: CmdGiftPanelPlan, RoomId: msg["RoomId"].(int), GiftPanelPlan: &giftPanelPlan}
	}
	return
}

func (*Handler) SetShoppingExplainCard(msg map[string]interface{}) (m MsgEvent) {
	shoppingExplainCard := ShoppingExplainCard{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &shoppingExplainCard); err == nil {
		m = MsgEvent{Cmd: CmdShoppingExplainCard, RoomId: msg["RoomId"].(int), ShoppingExplainCard: &shoppingExplainCard}
	}
	return
}

func (*Handler) SetAnchorLotCheckStatus(msg map[string]interface{}) (m MsgEvent) {
	anchorLotCheckStatus := AnchorLotCheckStatus{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &anchorLotCheckStatus); err == nil {
		m = MsgEvent{Cmd: CmdAnchorLotCheckStatus, RoomId: msg["RoomId"].(int), AnchorLotCheckStatus: &anchorLotCheckStatus}
	}
	return
}

func (*Handler) SetPkBattlePunishEnd(msg map[string]interface{}) (m MsgEvent) {
	pkBattlePunishEnd := PkBattlePunishEnd{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &pkBattlePunishEnd); err == nil {
		m = MsgEvent{Cmd: CmdPkBattlePunishEnd, RoomId: msg["RoomId"].(int), PkBattlePunishEnd: &pkBattlePunishEnd}
	}
	return
}

func (*Handler) SetAnchorLotEnd(msg map[string]interface{}) (m MsgEvent) {
	anchorLotEnd := AnchorLotEnd{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &anchorLotEnd); err == nil {
		m = MsgEvent{Cmd: CmdAnchorLotEnd, RoomId: msg["RoomId"].(int), AnchorLotEnd: &anchorLotEnd}
	}
	return
}

func (*Handler) SetAnchorLotAward(msg map[string]interface{}) (m MsgEvent) {
	anchorLotAward := AnchorLotAward{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &anchorLotAward); err == nil {
		m = MsgEvent{Cmd: CmdAnchorLotAward, RoomId: msg["RoomId"].(int), AnchorLotAward: &anchorLotAward}
	}
	return
}

func (*Handler) SetSpecialGift(msg map[string]interface{}) (m MsgEvent) {
	specialGift := SpecialGift{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &specialGift); err == nil {
		m = MsgEvent{Cmd: CmdSpecialGift, RoomId: msg["RoomId"].(int), SpecialGift: &specialGift}
	}
	return
}

func (*Handler) SetSuperChatMessageDelete(msg map[string]interface{}) (m MsgEvent) {
	superChatMessageDelete := SuperChatMessageDelete{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &superChatMessageDelete); err == nil {
		m = MsgEvent{Cmd: CmdSuperChatMessageDelete, RoomId: msg["RoomId"].(int), SuperChatMessageDelete: &superChatMessageDelete}
	}
	return
}

func (*Handler) SetVoiceJoinRoomCountInfo(msg map[string]interface{}) (m MsgEvent) {
	voiceJoinRoomCountInfo := VoiceJoinRoomCountInfo{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &voiceJoinRoomCountInfo); err == nil {
		m = MsgEvent{Cmd: CmdVoiceJoinRoomCountInfo, RoomId: msg["RoomId"].(int), VoiceJoinRoomCountInfo: &voiceJoinRoomCountInfo}
	}
	return
}

func (*Handler) SetVoiceJoinList(msg map[string]interface{}) (m MsgEvent) {
	voiceJoinList := VoiceJoinList{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &voiceJoinList); err == nil {
		m = MsgEvent{Cmd: CmdVoiceJoinList, RoomId: msg["RoomId"].(int), VoiceJoinList: &voiceJoinList}
	}
	return
}

func (*Handler) SetVoiceJoinStatus(msg map[string]interface{}) (m MsgEvent) {
	voiceJoinStatus := VoiceJoinStatus{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &voiceJoinStatus); err == nil {
		m = MsgEvent{Cmd: CmdVoiceJoinStatus, RoomId: msg["RoomId"].(int), VoiceJoinStatus: &voiceJoinStatus}
	}
	return
}

func (*Handler) SetWarning(msg map[string]interface{}) (m MsgEvent) {
	warning := Warning{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &warning); err == nil {
		m = MsgEvent{Cmd: CmdWarning, RoomId: msg["RoomId"].(int), Warning: &warning}
	}
	return
}

func (*Handler) SetPkBattleRankChange(msg map[string]interface{}) (m MsgEvent) {
	pkBattleRankChange := PkBattleRankChange{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &pkBattleRankChange); err == nil {
		m = MsgEvent{Cmd: CmdPkBattleRankChange, RoomId: msg["RoomId"].(int), PkBattleRankChange: &pkBattleRankChange}
	}
	return
}

func (*Handler) SetPkBattleSettleNew(msg map[string]interface{}) (m MsgEvent) {
	pkBattleSettleNew := PkBattleSettleNew{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &pkBattleSettleNew); err == nil {
		m = MsgEvent{Cmd: CmdPkBattleSettleNew, RoomId: msg["RoomId"].(int), PkBattleSettleNew: &pkBattleSettleNew}
	}
	return
}

func (*Handler) SetHotBuyNum(msg map[string]interface{}) (m MsgEvent) {
	hotBuyNum := HotBuyNum{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &hotBuyNum); err == nil {
		m = MsgEvent{Cmd: CmdHotBuyNum, RoomId: msg["RoomId"].(int), HotBuyNum: &hotBuyNum}
	}
	return
}

func (*Handler) SetShoppingCartShow(msg map[string]interface{}) (m MsgEvent) {
	shoppingCartShow := ShoppingCartShow{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &shoppingCartShow); err == nil {
		m = MsgEvent{Cmd: CmdShoppingCartShow, RoomId: msg["RoomId"].(int), ShoppingCartShow: &shoppingCartShow}
	}
	return
}

func (*Handler) SetVoiceJoinSwitch(msg map[string]interface{}) (m MsgEvent) {
	voiceJoinSwitch := VoiceJoinSwitch{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &voiceJoinSwitch); err == nil {
		m = MsgEvent{Cmd: CmdVoiceJoinSwitch, RoomId: msg["RoomId"].(int), VoiceJoinSwitch: &voiceJoinSwitch}
	}
	return
}

func (*Handler) SetCutOff(msg map[string]interface{}) (m MsgEvent) {
	cutOff := CutOff{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &cutOff); err == nil {
		m = MsgEvent{Cmd: CmdCutOff, RoomId: msg["RoomId"].(int), CutOff: &cutOff}
	}
	return
}

func (*Handler) SetRoomAdminRevoke(msg map[string]interface{}) (m MsgEvent) {
	roomAdminRevoke := RoomAdminRevoke{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &roomAdminRevoke); err == nil {
		m = MsgEvent{Cmd: CmdRoomAdminRevoke, RoomId: msg["RoomId"].(int), RoomAdminRevoke: &roomAdminRevoke}
	}
	return
}

func (*Handler) SetRoomSilentOf(msg map[string]interface{}) (m MsgEvent) {
	roomSilentOff := RoomSilentOff{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &roomSilentOff); err == nil {
		m = MsgEvent{Cmd: CmdRoomSilentOff, RoomId: msg["RoomId"].(int), RoomSilentOff: &roomSilentOff}
	}
	return
}

func (*Handler) SetRoomSilentOn(msg map[string]interface{}) (m MsgEvent) {
	roomSilentOn := RoomSilentOn{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &roomSilentOn); err == nil {
		m = MsgEvent{Cmd: CmdRoomSilentOn, RoomId: msg["RoomId"].(int), RoomSilentOn: &roomSilentOn}
	}
	return
}

func (*Handler) SetRoomAdminEntrance(msg map[string]interface{}) (m MsgEvent) {
	roomAdminEntrance := RoomAdminEntrance{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &roomAdminEntrance); err == nil {
		m = MsgEvent{Cmd: CmdRoomAdminEntrance, RoomId: msg["RoomId"].(int), RoomAdminEntrance: &roomAdminEntrance}
	}
	return
}

func (*Handler) SetRoomAdmins(msg map[string]interface{}) (m MsgEvent) {
	roomAdmins := RoomAdmins{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &roomAdmins); err == nil {
		m = MsgEvent{Cmd: CmdRoomAdmins, RoomId: msg["RoomId"].(int), RoomAdmins: &roomAdmins}
	}
	return
}

func (*Handler) SetVideoConnectionJoinStart(msg map[string]interface{}) (m MsgEvent) {
	videoConnectionJoinStart := VideoConnectionJoinStart{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &videoConnectionJoinStart); err == nil {
		m = MsgEvent{Cmd: CmdVideoConnectionJoinStart, RoomId: msg["RoomId"].(int), VideoConnectionJoinStart: &videoConnectionJoinStart}
	}
	return
}

func (*Handler) SetVideoConnectionMsg(msg map[string]interface{}) (m MsgEvent) {
	videoConnectionMsg := VideoConnectionMsg{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &videoConnectionMsg); err == nil {
		m = MsgEvent{Cmd: CmdVideoConnectionMsg, RoomId: msg["RoomId"].(int), VideoConnectionMsg: &videoConnectionMsg}
	}
	return
}

func (*Handler) SetVideoConnectionJoinEnd(msg map[string]interface{}) (m MsgEvent) {
	videoConnectionJoinEnd := VideoConnectionJoinEnd{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &videoConnectionJoinEnd); err == nil {
		m = MsgEvent{Cmd: CmdVideoConnectionJoinEnd, RoomId: msg["RoomId"].(int), VideoConnectionJoinEnd: &videoConnectionJoinEnd}
	}
	return
}

func (*Handler) SetRingStatusChange(msg map[string]interface{}) (m MsgEvent) {
	ringStatusChange := RingStatusChange{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &ringStatusChange); err == nil {
		m = MsgEvent{Cmd: CmdRingStatusChange, RoomId: msg["RoomId"].(int), RingStatusChange: &ringStatusChange}
	}
	return
}

func (*Handler) SetRingStatusChangeV2(msg map[string]interface{}) (m MsgEvent) {
	ringStatusChangeV2 := RingStatusChangeV2{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &ringStatusChangeV2); err == nil {
		m = MsgEvent{Cmd: CmdRingStatusChangeV2, RoomId: msg["RoomId"].(int), RingStatusChangeV2: &ringStatusChangeV2}
	}
	return
}

func (*Handler) SetRoomLock(msg map[string]interface{}) (m MsgEvent) {
	roomLock := RoomLock{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &roomLock); err == nil {
		m = MsgEvent{Cmd: CmdRoomLock, RoomId: msg["RoomId"].(int), RoomLock: &roomLock}
	}
	return
}

func (*Handler) SetShoppingBubblesStyle(msg map[string]interface{}) (m MsgEvent) {
	shoppingBubblesStyle := ShoppingBubblesStyle{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &shoppingBubblesStyle); err == nil {
		m = MsgEvent{Cmd: CmdShoppingBubblesStyle, RoomId: msg["RoomId"].(int), ShoppingBubblesStyle: &shoppingBubblesStyle}
	}
	return
}

func (*Handler) SetMultiVoiceOperating(msg map[string]interface{}) (m MsgEvent) {
	multiVoiceOperating := MultiVoiceOperating{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &multiVoiceOperating); err == nil {
		m = MsgEvent{Cmd: CmdMultiVoiceOperating, RoomId: msg["RoomId"].(int), MultiVoiceOperating: &multiVoiceOperating}
	}
	return
}

func (*Handler) SetMultiVoiceApplicationUser(msg map[string]interface{}) (m MsgEvent) {
	multiVoiceApplicationUser := MultiVoiceApplicationUser{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &multiVoiceApplicationUser); err == nil {
		m = MsgEvent{Cmd: CmdMultiVoiceApplicationUser, RoomId: msg["RoomId"].(int), MultiVoiceApplicationUser: &multiVoiceApplicationUser}
	}
	return
}

func (*Handler) SetPkBattleMatchTimeout(msg map[string]interface{}) (m MsgEvent) {
	pkBattleMatchTimeout := PkBattleMatchTimeout{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &pkBattleMatchTimeout); err == nil {
		m = MsgEvent{Cmd: CmdPkBattleMatchTimeout, RoomId: msg["RoomId"].(int), PkBattleMatchTimeout: &pkBattleMatchTimeout}
	}
	return
}

func (*Handler) SetChangeRoomInfo(msg map[string]interface{}) (m MsgEvent) {
	changeRoomInfo := ChangeRoomInfo{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &changeRoomInfo); err == nil {
		m = MsgEvent{Cmd: CmdChangeRoomInfo, RoomId: msg["RoomId"].(int), ChangeRoomInfo: &changeRoomInfo}
	}
	return
}

func (*Handler) SetLiveMultiViewChange(msg map[string]interface{}) (m MsgEvent) {
	liveMultiViewChange := LiveMultiViewChange{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &liveMultiViewChange); err == nil {
		m = MsgEvent{Cmd: CmdLiveMultiViewChange, RoomId: msg["RoomId"].(int), LiveMultiViewChange: &liveMultiViewChange}
	}
	return
}

func (*Handler) SetGuardAchievementRoom(msg map[string]interface{}) (m MsgEvent) {
	guardAchievementRoom := GuardAchievementRoom{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &guardAchievementRoom); err == nil {
		m = MsgEvent{Cmd: CmdGuardAchievementRoom, RoomId: msg["RoomId"].(int), GuardAchievementRoom: &guardAchievementRoom}
	}
	return
}

func (*Handler) SetSysMsg(msg map[string]interface{}) (m MsgEvent) {
	sysMsg := SysMsg{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &sysMsg); err == nil {
		m = MsgEvent{Cmd: CmdSysMsg, RoomId: msg["RoomId"].(int), SysMsg: &sysMsg}
	}
	return
}

func (*Handler) SetMvRoleChange(msg map[string]interface{}) (m MsgEvent) {
	mvRoleChange := MvRoleChange{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &mvRoleChange); err == nil {
		m = MsgEvent{Cmd: CmdMvRoleChange, RoomId: msg["RoomId"].(int), MvRoleChange: &mvRoleChange}
	}
	return
}

func (*Handler) SetSelectedGoodsInfo(msg map[string]interface{}) (m MsgEvent) {
	selectedGoodsInfo := SelectedGoodsInfo{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &selectedGoodsInfo); err == nil {
		m = MsgEvent{Cmd: CmdSelectedGoodsInfo, RoomId: msg["RoomId"].(int), SelectedGoodsInfo: &selectedGoodsInfo}
	}
	return
}

func (*Handler) SetMultiVoiceOperatin(msg map[string]interface{}) (m MsgEvent) {
	multiVoiceOperatin := MultiVoiceOperatin{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &multiVoiceOperatin); err == nil {
		m = MsgEvent{Cmd: CmdMultiVoiceOperatin, RoomId: msg["RoomId"].(int), MultiVoiceOperatin: &multiVoiceOperatin}
	}
	return
}

func (*Handler) SetPanelInteractiveNotifyChange(msg map[string]interface{}) (m MsgEvent) {
	panelInteractiveNotifyChange := PanelInteractiveNotifyChange{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &panelInteractiveNotifyChange); err == nil {
		m = MsgEvent{Cmd: CmdPanelInteractiveNotifyChange, RoomId: msg["RoomId"].(int), PanelInteractiveNotifyChange: &panelInteractiveNotifyChange}
	}
	return
}

func (*Handler) SetInteractiveUser(msg map[string]interface{}) (m MsgEvent) {
	interactiveUser := InteractiveUser{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &interactiveUser); err == nil {
		m = MsgEvent{Cmd: CmdInteractiveUser, RoomId: msg["RoomId"].(int), InteractiveUser: &interactiveUser}
	}
	return
}

func (*Handler) SetUserVirtualMvp(msg map[string]interface{}) (m MsgEvent) {
	userVirtualMvp := UserVirtualMvp{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &userVirtualMvp); err == nil {
		m = MsgEvent{Cmd: CmdUserVirtualMvp, RoomId: msg["RoomId"].(int), UserVirtualMvp: &userVirtualMvp}
	}
	return
}

func (*Handler) SetWidgetWishList(msg map[string]interface{}) (m MsgEvent) {
	widgetWishList := WidgetWishList{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &widgetWishList); err == nil {
		m = MsgEvent{Cmd: CmdWidgetWishList, RoomId: msg["RoomId"].(int), WidgetWishList: &widgetWishList}
	}
	return
}

func (*Handler) SetCheckSingStatus(msg map[string]interface{}) (m MsgEvent) {
	checkSingStatus := CheckSingStatus{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &checkSingStatus); err == nil {
		m = MsgEvent{Cmd: CmdCheckSingStatus, RoomId: msg["RoomId"].(int), CheckSingStatus: &checkSingStatus}
	}
	return
}

func (*Handler) SetRoomModuleDisplay(msg map[string]interface{}) (m MsgEvent) {
	roomModuleDisplay := RoomModuleDisplay{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &roomModuleDisplay); err == nil {
		m = MsgEvent{Cmd: CmdRoomModuleDisplay, RoomId: msg["RoomId"].(int), RoomModuleDisplay: &roomModuleDisplay}
	}
	return
}

func (*Handler) SetVoiceChatUpdate(msg map[string]interface{}) (m MsgEvent) {
	voiceChatUpdate := VoiceChatUpdate{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &voiceChatUpdate); err == nil {
		m = MsgEvent{Cmd: CmdVoiceChatUpdate, RoomId: msg["RoomId"].(int), VoiceChatUpdate: &voiceChatUpdate}
	}
	return
}

func (*Handler) SetReenterLiveRoom(msg map[string]interface{}) (m MsgEvent) {
	reenterLiveRoom := ReenterLiveRoom{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &reenterLiveRoom); err == nil {
		m = MsgEvent{Cmd: CmdReenterLiveRoom, RoomId: msg["RoomId"].(int), ReenterLiveRoom: &reenterLiveRoom}
	}
	return
}

func (*Handler) SetOfficialRoomEvent(msg map[string]interface{}) (m MsgEvent) {
	officialRoomEvent := OfficialRoomEvent{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &officialRoomEvent); err == nil {
		m = MsgEvent{Cmd: CmdOfficialRoomEvent, RoomId: msg["RoomId"].(int), OfficialRoomEvent: &officialRoomEvent}
	}
	return
}

func (*Handler) SetActivityBannerChangeV2(msg map[string]interface{}) (m MsgEvent) {
	activityBannerChangeV2 := ActivityBannerChangeV2{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &activityBannerChangeV2); err == nil {
		m = MsgEvent{Cmd: CmdActivityBannerChangeV2, RoomId: msg["RoomId"].(int), ActivityBannerChangeV2: &activityBannerChangeV2}
	}
	return
}

func (*Handler) SetActivityBannerChange(msg map[string]interface{}) (m MsgEvent) {
	activityBannerChange := ActivityBannerChange{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &activityBannerChange); err == nil {
		m = MsgEvent{Cmd: CmdActivityBannerChange, RoomId: msg["RoomId"].(int), ActivityBannerChange: &activityBannerChange}
	}
	return
}

func (*Handler) SetVideoConnectionStart(msg map[string]interface{}) (m MsgEvent) {
	videoConnectionStart := VideoConnectionStart{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &videoConnectionStart); err == nil {
		m = MsgEvent{Cmd: CmdVideoConnectionStart, RoomId: msg["RoomId"].(int), VideoConnectionStart: &videoConnectionStart}
	}
	return
}

func (*Handler) SetGuideInfoStatus(msg map[string]interface{}) (m MsgEvent) {
	guideInfoStatus := GuideInfoStatus{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &guideInfoStatus); err == nil {
		m = MsgEvent{Cmd: CmdGuideInfoStatus, RoomId: msg["RoomId"].(int), GuideInfoStatus: &guideInfoStatus}
	}
	return
}

func (*Handler) SetObsShieldStatusUpdate(msg map[string]interface{}) (m MsgEvent) {
	obsShieldStatusUpdate := ObsShieldStatusUpdate{}
	if err := JsonCoder.Unmarshal([]byte(msg["msg"].(string)), &obsShieldStatusUpdate); err == nil {
		m = MsgEvent{Cmd: CmdObsShieldStatusUpdate, RoomId: msg["RoomId"].(int), ObsShieldStatusUpdate: &obsShieldStatusUpdate}
	}
	return
}
