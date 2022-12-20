package Handler

import (
	"encoding/json"
	"fmt"
	"log"
)

func SetDanMuMsg(msg map[string]interface{}) {
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
	fmt.Println(danMu)
}

func SetInteractWord(msg map[string]interface{}) {
	interactMsg := InteractWord{}
	interactMsg.Cmd = CmdInteractWord
	dataJson, err := json.Marshal(msg["data"])
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(dataJson, &interactMsg.Data)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return
	}
	fmt.Println(interactMsg)
}

func SetOnlineRankCount(msg map[string]interface{}) {
	onlineRankCount := OnlineRankCount{}
	onlineRankCount.Cmd = CmdOnlineRankCount
	dataJson, err := json.Marshal(msg["data"])
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(dataJson, &onlineRankCount.Data)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return
	}
	fmt.Println(onlineRankCount)
}

func SetWatchedChange(msg map[string]interface{}) {
	watchedChange := WatchedChange{}
	watchedChange.Cmd = CmdWatchedChange
	dataJson, err := json.Marshal(msg["data"])
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(dataJson, &watchedChange.Data)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return
	}
	fmt.Println(watchedChange)
}

func SetNoticeMsg(msg map[string]interface{}) {
	noticeMsg := NoticeMsg{}
	dataJson, err := json.Marshal(msg)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(dataJson, &noticeMsg)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return
	}
	fmt.Println(noticeMsg)
}

func SetSuperChatMessage(msg map[string]interface{}) {
	superChatMsg := SuperChatMessage{}
	superChatMsg.Cmd = CmdSuperChatMessage
	dataJson, err := json.Marshal(msg["data"])
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(dataJson, &superChatMsg.Data)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return
	}
	fmt.Println(superChatMsg)
}

func SetSendGift(msg map[string]interface{}) {
	sendGift := SendGift{}
	sendGift.Cmd = CmdSendGift
	dataJson, err := json.Marshal(msg["data"])
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(dataJson, &sendGift.Data)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return
	}
	fmt.Println(sendGift)
}

func SetOnlineRankV2(msg map[string]interface{}) {
	onlineRankV2 := OnlineRankV2{}
	onlineRankV2.Cmd = CmdOnlineRankV2
	dataJson, err := json.Marshal(msg["data"])
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(dataJson, &onlineRankV2.Data)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return
	}
	fmt.Println(onlineRankV2)
}

func SetOnlineRankTop3(msg map[string]interface{}) {
	onlineRankTop3 := OnlineRankTop3{}
	onlineRankTop3.Cmd = CmdOnlineRankTop3
	dataJson, err := json.Marshal(msg["data"])
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(dataJson, &onlineRankTop3.Data)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return
	}
	fmt.Println(onlineRankTop3)
}

func SetLikeInfoV3Click(msg map[string]interface{}) {
	likeInfoV3Click := LikeInfoV3Click{}
	likeInfoV3Click.Cmd = CmdLikeInfoV3Click
	dataJson, err := json.Marshal(msg["data"])
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(dataJson, &likeInfoV3Click.Data)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return
	}
	fmt.Println(likeInfoV3Click)
}

func SetStopLiveRoomList(msg map[string]interface{}) {
	stopLiveRoomList := StopLiveRoomList{}
	stopLiveRoomList.Cmd = CmdStopLiveRoomList
	dataJson, err := json.Marshal(msg["data"])
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(dataJson, &stopLiveRoomList.Data)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return
	}
	fmt.Println(stopLiveRoomList)
}

func SetLikeInfoV3Update(msg map[string]interface{}) {
	likeInfoV3Update := LikeInfoV3Update{}
	likeInfoV3Update.Cmd = CmdLikeInfoV3Update
	dataJson, err := json.Marshal(msg["data"])
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(dataJson, &likeInfoV3Update.Data)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return
	}
	fmt.Println(likeInfoV3Update)
}

func SetHotRankChange(msg map[string]interface{}) {
	hotRankChange := HotRankChange{}
	hotRankChange.Cmd = CmdHotRankChange
	dataJson, err := json.Marshal(msg["data"])
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(dataJson, &hotRankChange.Data)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return
	}
	fmt.Println(hotRankChange)
}

func SetRoomRealTimeMessageUpdate(msg map[string]interface{}) {
	roomRealTimeMessageUpdate := RoomRealTimeMessageUpdate{}
	roomRealTimeMessageUpdate.Cmd = CmdRoomRealTimeMessageUpdate
	dataJson, err := json.Marshal(msg["data"])
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(dataJson, &roomRealTimeMessageUpdate.Data)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return
	}
	fmt.Println(roomRealTimeMessageUpdate)
}

func SetWidgetBanner(msg map[string]interface{}) {
	widgetBanner := WidgetBanner{}
	widgetBanner.Cmd = CmdWidgetBanner
	dataJson, err := json.Marshal(msg["data"])
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(dataJson, &widgetBanner.Data)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return
	}
	fmt.Println(widgetBanner)
}

func SetHotRankChangedV2(msg map[string]interface{}) {
	hotRankChangedV2 := HotRankChangedV2{}
	hotRankChangedV2.Cmd = CmdHotRankChangedV2
	dataJson, err := json.Marshal(msg["data"])
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(dataJson, &hotRankChangedV2.Data)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return
	}
	fmt.Println(hotRankChangedV2)
}
