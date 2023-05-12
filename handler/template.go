package handler

const (
	CmdDanmuMsg                      = "DANMU_MSG"
	CmdSuperChatMessage              = "SUPER_CHAT_MESSAGE"
	CmdWatchedChange                 = "WATCHED_CHANGE"
	CmdSendGift                      = "SEND_GIFT"
	CmdOnlineRankCount               = "ONLINE_RANK_COUNT"
	CmdOnlineRankV2                  = "ONLINE_RANK_V2"
	CmdOnlineRankTop3                = "ONLINE_RANK_TOP3"
	CmdLikeInfoV3Click               = "LIKE_INFO_V3_CLICK"
	CmdInteractWord                  = "INTERACT_WORD"
	CmdStopLiveRoomList              = "STOP_LIVE_ROOM_LIST"
	CmdLikeInfoV3Update              = "LIKE_INFO_V3_UPDATE"
	CmdHotRankChange                 = "HOT_RANK_CHANGED"
	CmdNoticeMsg                     = "NOTICE_MSG"
	CmdRoomRealTimeMessageUpdate     = "ROOM_REAL_TIME_MESSAGE_UPDATE"
	CmdWidgetBanner                  = "WIDGET_BANNER"
	CmdHotRankChangedV2              = "HOT_RANK_CHANGED_V2"
	CmdGuardHonorThousand            = "GUARD_HONOR_THOUSAND"
	CmdLive                          = "LIVE"
	CmdRoomChange                    = "ROOM_CHANGE"
	CmdRoomBlockMsg                  = "ROOM_BLOCK_MSG"
	CmdFullScreenSpecialEffect       = "FULL_SCREEN_SPECIAL_EFFECT"
	CmdCommonNoticeDanmaku           = "COMMON_NOTICE_DANMAKU"
	CmdTradingScore                  = "TRADING_SCORE"
	CmdPreparing                     = "PREPARING"
	CmdGuardBuy                      = "GUARD_BUY"
	CmdGiftStarProcess               = "GIFT_STAR_PROCESS"
	CmdRoomSkinMsg                   = "ROOM_SKIN_MSG"
	CmdEntryEffect                   = "ENTRY_EFFECT"
	CmdUserToastMsg                  = "USER_TOAST_MSG"
	CmdHeartBeatReply                = "HEARTBEAT_REPLY"
	CmdPopularityRedPocketNew        = "POPULARITY_RED_POCKET_NEW"
	CmdAreaRankChanged               = "AREA_RANK_CHANGED"
	CmdSuperChatEntrance             = "SUPER_CHAT_ENTRANCE"
	CmdPlayTogether                  = "PLAY_TOGETHER"
	CmdComboSend                     = "COMBO_SEND"
	CmdPopularityRedPocketStart      = "POPULARITY_RED_POCKET_START"
	CmdPkBattleProcess               = "PK_BATTLE_PROCESS"
	CmdPopularRankChanged            = "POPULAR_RANK_CHANGED"
	CmdPkBattleStartNew              = "PK_BATTLE_START_NEW"
	CmdDanMuAggregation              = "DANMU_AGGREGATION"
	CmdLiveInteractiveGame           = "LIVE_INTERACTIVE_GAME"
	CmdRecommendCard                 = "RECOMMEND_CARD"
	CmdPkBattleProcessNew            = "PK_BATTLE_PROCESS_NEW"
	CmdPkBattlePreNew                = "PK_BATTLE_PRE_NEW"
	CmdPkBattlePre                   = "PK_BATTLE_PRE"
	CmdPkBattleFinalProcess          = "PK_BATTLE_FINAL_PROCESS"
	CmdPkBattleStart                 = "PK_BATTLE_START"
	CmdWidgetGiftStarProcess         = "WIDGET_GIFT_STAR_PROCESS"
	CmdPopularityRedPocketWinnerList = "POPULARITY_RED_POCKET_WINNER_LIST"
	CmdGotoBuyFlow                   = "GOTO_BUY_FLOW"
	CmdPkBattleEnd                   = "PK_BATTLE_END"
	CmdPkBattleSettleUser            = "PK_BATTLE_SETTLE_USER"
	CmdAnchorLotStart                = "ANCHOR_LOT_START"
	CmdPkBattleSettleV2              = "PK_BATTLE_SETTLE_V2"
	CmdPkBattleSettle                = "PK_BATTLE_SETTLE"
	CmdHotRoomNotify                 = "HOT_ROOM_NOTIFY"
	CmdLiveOpenPlatformGame          = "LIVE_OPEN_PLATFORM_GAME"
	CmdLivePanelChangeContent        = "LIVE_PANEL_CHANGE_CONTENT"
	CmdGiftPanelPlan                 = "GIFT_PANEL_PLAN"
	CmdShoppingExplainCard           = "SHOPPING_EXPLAIN_CARD"
	CmdAnchorLotCheckStatus          = "ANCHOR_LOT_CHECK_STATUS"
	CmdPkBattlePunishEnd             = "PK_BATTLE_PUNISH_END"
	CmdAnchorLotEnd                  = "ANCHOR_LOT_END"
	CmdAnchorLotAward                = "ANCHOR_LOT_AWARD"
	CmdSpecialGift                   = "SPECIAL_GIFT"
	CmdSuperChatMessageDelete        = "SUPER_CHAT_MESSAGE_DELETE"
	CmdVoiceJoinRoomCountInfo        = "VOICE_JOIN_ROOM_COUNT_INFO"
	CmdVoiceJoinList                 = "VOICE_JOIN_LIST"
	CmdVoiceJoinStatus               = "VOICE_JOIN_STATUS"
	CmdWarning                       = "WARNING"
	CmdPkBattleRankChange            = "PK_BATTLE_RANK_CHANGE"
	CmdPkBattleSettleNew             = "PK_BATTLE_SETTLE_NEW"
	CmdHotBuyNum                     = "HOT_BUY_NUM"
	CmdShoppingCartShow              = "SHOPPING_CART_SHOW"
	CmdVoiceJoinSwitch               = "VOICE_JOIN_SWITCH"
	CmdCutOff                        = "CUT_OFF"
	CmdRoomAdminRevoke               = "ROOM_ADMIN_REVOKE"
	CmdRoomSilentOff                 = "ROOM_SILENT_OFF"
	CmdRoomSilentOn                  = "ROOM_SILENT_ON"
	CmdRoomAdminEntrance             = "room_admin_entrance"
	CmdRoomAdmins                    = "ROOM_ADMINS"
	CmdVideoConnectionJoinStart      = "VIDEO_CONNECTION_JOIN_START"
	CmdVideoConnectionMsg            = "VIDEO_CONNECTION_MSG"
	CmdVideoConnectionJoinEnd        = "VIDEO_CONNECTION_JOIN_END"
	CmdRingStatusChange              = "RING_STATUS_CHANGE"
	CmdRingStatusChangeV2            = "RING_STATUS_CHANGE_V2"
	CmdRoomLock                      = "ROOM_LOCK"
	CmdShoppingBubblesStyle          = "SHOPPING_BUBBLES_STYLE"
	CmdMultiVoiceOperating           = "MULTI_VOICE_OPERATING"
	CmdMultiVoiceApplicationUser     = "MULTI_VOICE_APPLICATION_USER"
	CmdPkBattleMatchTimeout          = "PK_BATTLE_MATCH_TIMEOUT"
	CmdChangeRoomInfo                = "CHANGE_ROOM_INFO"
	CmdLiveMultiViewChange           = "LIVE_MULTI_VIEW_CHANGE"
	CmdGuardAchievementRoom          = "GUARD_ACHIEVEMENT_ROOM"
	CmdSysMsg                        = "SYS_MSG"
	CmdMvRoleChange                  = "MV_ROLE_CHANGE"
	CmdSelectedGoodsInfo             = "SELECTED_GOODS_INFO"
	CmdMultiVoiceOperatin            = "MULTI_VOICE_OPERATING"
	CmdPanelInteractiveNotifyChange  = "PANEL_INTERACTIVE_NOTIFY_CHANGE"
	CmdInteractiveUser               = "INTERACTIVE_USER"
	CmdUserVirtualMvp                = "USER_VIRTUAL_MVP"
	CmdWidgetWishList                = "WIDGET_WISH_LIST"
	CmdCheckSingStatus               = "CHECK_SING_STATUS"
	CmdRoomModuleDisplay             = "ROOM_MODULE_DISPLAY"
	CmdVoiceChatUpdate               = "VOICE_CHAT_UPDATE"
	CmdReenterLiveRoom               = "REENTER_LIVE_ROOM"
)

var CmdName = map[string]string{
	"DANMU_MSG":                         "DanMuMsg",
	"SUPER_CHAT_MESSAGE":                "SuperChatMessage",
	"SUPER_CHAT_MESSAGE_JPN":            "SuperChatMessage",
	"WATCHED_CHANGE":                    "WatchedChange",
	"SEND_GIFT":                         "SendGift",
	"ONLINE_RANK_COUNT":                 "OnlineRankCount",
	"ONLINE_RANK_V2":                    "OnlineRankV2",
	"ONLINE_RANK_TOP3":                  "OnlineRankTop3",
	"LIKE_INFO_V3_CLICK":                "LikeInfoV3Click",
	"INTERACT_WORD":                     "InteractWord",
	"STOP_LIVE_ROOM_LIST":               "StopLiveRoomList",
	"LIKE_INFO_V3_UPDATE":               "LikeInfoV3Update",
	"HOT_RANK_CHANGED":                  "HotRankChanged",
	"NOTICE_MSG":                        "NoticeMsg",
	"ROOM_REAL_TIME_MESSAGE_UPDATE":     "RoomRealTimeMessageUpdate",
	"WIDGET_BANNER":                     "WidgetBanner",
	"HOT_RANK_CHANGED_V2":               "HotRankChangedV2",
	"GUARD_HONOR_THOUSAND":              "GuardHonorThousand",
	"LIVE":                              "Live",
	"ROOM_CHANGE":                       "RoomChange",
	"ROOM_BLOCK_MSG":                    "RoomBlockMsg",
	"FULL_SCREEN_SPECIAL_EFFECT":        "FullScreenSpecialEffect",
	"COMMON_NOTICE_DANMAKU":             "CommonNoticeDanmaku",
	"TRADING_SCORE":                     "TradingScore",
	"PREPARING":                         "Preparing",
	"GUARD_BUY":                         "GuardBuy",
	"GIFT_STAR_PROCESS":                 "GiftStarProcess",
	"ROOM_SKIN_MSG":                     "RoomSkinMsg",
	"ENTRY_EFFECT":                      "EntryEffect",
	"USER_TOAST_MSG":                    "UserToastMsg",
	"HEARTBEAT_REPLY":                   "HeartBeatReply",
	"POPULARITY_RED_POCKET_NEW":         "PopularityRedPocketNew",
	"POPULARITY_RED_POCKET_START":       "PopularityRedPocketStart",
	"AREA_RANK_CHANGED":                 "AreaRankChanged",
	"SUPER_CHAT_ENTRANCE":               "SuperChatEntrance",
	"PLAY_TOGETHER":                     "PlayTogether",
	"COMBO_SEND":                        "ComboSend",
	"PK_BATTLE_PROCESS":                 "PkBattleProcess",
	"POPULAR_RANK_CHANGED":              "PopularRankChanged",
	"PK_BATTLE_START_NEW":               "PkBattleStartNew",
	"DANMU_AGGREGATION":                 "DanMuAggregation",
	"LIVE_INTERACTIVE_GAME":             "LiveInteractiveGame",
	"RECOMMEND_CARD":                    "RecommendCard",
	"PK_BATTLE_PROCESS_NEW":             "PkBattleProcessNew",
	"PK_BATTLE_PRE_NEW":                 "PkBattlePreNew",
	"PK_BATTLE_PRE":                     "PkBattlePre",
	"PK_BATTLE_FINAL_PROCESS":           "PkBattleFinalProcess",
	"PK_BATTLE_START":                   "PkBattleStart",
	"WIDGET_GIFT_STAR_PROCESS":          "WidgetGiftStarProcess",
	"POPULARITY_RED_POCKET_WINNER_LIST": "PopularityRedPocketWinnerList",
	"GOTO_BUY_FLOW":                     "GotoBuyFlow",
	"PK_BATTLE_END":                     "PkBattleEnd",
	"PK_BATTLE_SETTLE_USER":             "PkBattleSettleUser",
	"ANCHOR_LOT_START":                  "AnchorLotStart",
	"PK_BATTLE_SETTLE_V2":               "PkBattleSettleV2",
	"PK_BATTLE_SETTLE":                  "PkBattleSettle",
	"HOT_ROOM_NOTIFY":                   "HotRoomNotify",
	"LIVE_OPEN_PLATFORM_GAME":           "LiveOpenPlatformGame",
	"LIVE_PANEL_CHANGE_CONTENT":         "LivePanelChangeContent",
	"GIFT_PANEL_PLAN":                   "GiftPanelPlan",
	"SHOPPING_EXPLAIN_CARD":             "ShoppingExplainCard",
	"ANCHOR_LOT_CHECKSTATUS":            "AnchorLotCheckStatus",
	"PK_BATTLE_PUNISH_END":              "PkBattlePunishEnd",
	"ANCHOR_LOT_END":                    "AnchorLotEnd",
	"ANCHOR_LOT_AWARD":                  "AnchorLotAward",
	"SPECIAL_GIFT":                      "SpecialGift",
	"SUPER_CHAT_MESSAGE_DELETE":         "SuperChatMessageDelete",
	"VOICE_JOIN_ROOM_COUNT_INFO":        "VoiceJoinRoomCountInfo",
	"VOICE_JOIN_LIST":                   "VoiceJoinList",
	"VOICE_JOIN_STATUS":                 "VoiceJoinStatus",
	"WARNING":                           "Warning",
	"PK_BATTLE_RANK_CHANGE":             "PkBattleRankChange",
	"PK_BATTLE_SETTLE_NEW":              "PkBattleSettleNew",
	"HOT_BUY_NUM":                       "HotBuyNum",
	"SHOPPING_CART_SHOW":                "ShoppingCartShow",
	"VOICE_JOIN_SWITCH":                 "VoiceJoinSwitch",
	"CUT_OFF":                           "CutOff",
	"ROOM_ADMIN_REVOKE":                 "RoomAdminRevoke",
	"ROOM_SILENT_OFF":                   "RoomSilentOff",
	"ROOM_SILENT_ON":                    "RoomSilentOn",
	"room_admin_entrance":               "RoomAdminEntrance",
	"ROOM_ADMINS":                       "RoomAdmins",
	"VIDEO_CONNECTION_JOIN_START":       "VideoConnectionJoinStart",
	"VIDEO_CONNECTION_MSG":              "VideoConnectionMsg",
	"VIDEO_CONNECTION_JOIN_END":         "VideoConnectionJoinEnd",
	"RING_STATUS_CHANGE":                "RingStatusChange",
	"RING_STATUS_CHANGE_V2":             "RingStatusChangeV2",
	"ROOM_LOCK":                         "RoomLock",
	"SHOPPING_BUBBLES_STYLE":            "ShoppingBubblesStyle",
	"MULTI_VOICE_OPERATING":             "MultiVoiceOperating",
	"MULTI_VOICE_APPLICATION_USER":      "MultiVoiceApplicationUser",
	"PK_BATTLE_MATCH_TIMEOUT":           "PkBattleMatchTimeout",
	"CHANGE_ROOM_INFO":                  "ChangeRoomInfo",
	"LIVE_MULTI_VIEW_CHANGE":            "LiveMultiViewChange",
	"GUARD_ACHIEVEMENT_ROOM":            "GuardAchievementRoom",
	"SYS_MSG":                           "SysMsg",
	"MVROLECHANGE":                      "MvRoleChange",
	"SELECTED_GOODS_INFO":               "SelectedGoodsInfo",
	"MULTI_VOICE_OPERATIN":              "MultiVoiceOperatin",
	"PANEL_INTERACTIVE_NOTIFY_CHANGE":   "PanelInteractiveNotifyChange",
	"INTERACTIVE_USER":                  "InteractiveUser",
	"USER_VIRTUAL_MVP":                  "UserVirtualMvp",
	"WIDGET_WISH_LIST":                  "WidgetWishList",
	"CHECK_SING_STATUS":                 "CheckSingStatus",
	"ROOM_MODULE_DISPLAY":               "RoomModuleDisplay",
	"VOICE_CHAT_UPDATE":                 "VoiceChatUpdate",
	"REENTER_LIVE_ROOM":                 "ReenterLiveRoom",
}

type MsgEvent struct {
	Cmd                           string
	RoomId                        int
	DanMuMsg                      *DanMuMsg
	SuperChatMessage              *SuperChatMessage
	WatchedChange                 *WatchedChange
	SendGift                      *SendGift
	OnlineRankCount               *OnlineRankCount
	OnlineRankV2                  *OnlineRankV2
	OnlineRankTop3                *OnlineRankTop3
	LikeInfoV3Click               *LikeInfoV3Click
	InteractWord                  *InteractWord
	StopLiveRoomList              *StopLiveRoomList
	LikeInfoV3Update              *LikeInfoV3Update
	HotRankChange                 *HotRankChange
	NoticeMsg                     *NoticeMsg
	RoomRealTimeMessageUpdate     *RoomRealTimeMessageUpdate
	WidgetBanner                  *WidgetBanner
	HotRankChangedV2              *HotRankChangedV2
	GuardHonorThousand            *GuardHonorThousand
	Live                          *Live
	RoomChange                    *RoomChange
	RoomBlockMsg                  *RoomBlockMsg
	FullScreenSpecialEffect       *FullScreenSpecialEffect
	CommonNoticeDanmaku           *CommonNoticeDanmaku
	TradingScore                  *TradingScore
	Preparing                     *Preparing
	GuardBuy                      *GuardBuy
	GiftStarProcess               *GiftStarProcess
	RoomSkinMsg                   *RoomSkinMsg
	EntryEffect                   *EntryEffect
	UserToastMsg                  *UserToastMsg
	HeartBeatReply                *HeartBeatReply
	PopularityRedPocketNew        *PopularityRedPocketNew
	AreaRankChanged               *AreaRankChanged
	SuperChatEntrance             *SuperChatEntrance
	PlayTogether                  *PlayTogether
	ComboSend                     *ComboSend
	PopularityRedPocketStart      *PopularityRedPocketStart
	PkBattleProcess               *PkBattleProcess
	PopularRankChanged            *PopularRankChanged
	PkBattleStartNew              *PkBattleStartNew
	DanMuAggregation              *DanMuAggregation
	LiveInteractiveGame           *LiveInteractiveGame
	RecommendCard                 *RecommendCard
	PkBattleProcessNew            *PkBattleProcessNew
	PkBattlePreNew                *PkBattlePreNew
	PkBattlePre                   *PkBattlePre
	PkBattleFinalProcess          *PkBattleFinalProcess
	PkBattleStart                 *PkBattleStart
	WidgetGiftStarProcess         *WidgetGiftStarProcess
	PopularityRedPocketWinnerList *PopularityRedPocketWinnerList
	GotoBuyFlow                   *GotoBuyFlow
	PkBattleEnd                   *PkBattleEnd
	PkBattleSettleUser            *PkBattleSettleUser
	AnchorLotStart                *AnchorLotStart
	PkBattleSettleV2              *PkBattleSettleV2
	PkBattleSettle                *PkBattleSettle
	HotRoomNotify                 *HotRoomNotify
	LiveOpenPlatformGame          *LiveOpenPlatformGame
	LivePanelChangeContent        *LivePanelChangeContent
	GiftPanelPlan                 *GiftPanelPlan
	ShoppingExplainCard           *ShoppingExplainCard
	AnchorLotCheckStatus          *AnchorLotCheckStatus
	PkBattlePunishEnd             *PkBattlePunishEnd
	AnchorLotEnd                  *AnchorLotEnd
	AnchorLotAward                *AnchorLotAward
	SpecialGift                   *SpecialGift
	SuperChatMessageDelete        *SuperChatMessageDelete
	VoiceJoinRoomCountInfo        *VoiceJoinRoomCountInfo
	VoiceJoinList                 *VoiceJoinList
	VoiceJoinStatus               *VoiceJoinStatus
	Warning                       *Warning
	PkBattleRankChange            *PkBattleRankChange
	PkBattleSettleNew             *PkBattleSettleNew
	HotBuyNum                     *HotBuyNum
	ShoppingCartShow              *ShoppingCartShow
	VoiceJoinSwitch               *VoiceJoinSwitch
	CutOff                        *CutOff
	RoomAdminRevoke               *RoomAdminRevoke
	RoomSilentOff                 *RoomSilentOff
	RoomSilentOn                  *RoomSilentOn
	RoomAdminEntrance             *RoomAdminEntrance
	RoomAdmins                    *RoomAdmins
	VideoConnectionJoinStart      *VideoConnectionJoinStart
	VideoConnectionMsg            *VideoConnectionMsg
	VideoConnectionJoinEnd        *VideoConnectionJoinEnd
	RingStatusChange              *RingStatusChange
	RingStatusChangeV2            *RingStatusChangeV2
	RoomLock                      *RoomLock
	ShoppingBubblesStyle          *ShoppingBubblesStyle
	MultiVoiceOperating           *MultiVoiceOperating
	MultiVoiceApplicationUser     *MultiVoiceApplicationUser
	PkBattleMatchTimeout          *PkBattleMatchTimeout
	ChangeRoomInfo                *ChangeRoomInfo
	LiveMultiViewChange           *LiveMultiViewChange
	GuardAchievementRoom          *GuardAchievementRoom
	SysMsg                        *SysMsg
	MvRoleChange                  *MvRoleChange
	SelectedGoodsInfo             *SelectedGoodsInfo
	MultiVoiceOperatin            *MultiVoiceOperatin
	PanelInteractiveNotifyChange  *PanelInteractiveNotifyChange
	InteractiveUser               *InteractiveUser
	UserVirtualMvp                *UserVirtualMvp
	WidgetWishList                *WidgetWishList
	CheckSingStatus               *CheckSingStatus
	RoomModuleDisplay             *RoomModuleDisplay
	VoiceChatUpdate               *VoiceChatUpdate
	ReenterLiveRoom               *ReenterLiveRoom
}

type FansMedal struct {
	AnchorRoomId     int    `json:"anchor_roomid"`
	GuardLevel       int    `json:"guard_level"`
	IconID           int    `json:"icon_id"`
	IsLighted        int    `json:"is_lighted"`
	MedalColor       int    `json:"medal_color"`
	MedalColorBorder int    `json:"medal_color_border"`
	MedalColorEnd    int    `json:"medal_color_end"`
	MedalColorStart  int    `json:"medal_color_start"`
	MedalLevel       int    `json:"medal_level"`
	MedalName        string `json:"medal_name"`
	Score            int    `json:"score"`
	Special          string `json:"special"`
	TargetID         int    `json:"target_id"`
}

type DanMuMsg struct {
	Cmd  string `json:"cmd"`
	Data struct {
		Sender struct {
			Uid    int64
			Name   string
			RoomId int64
		}
		Medal                    FansMedal
		Content                  string
		SendTimeStamp            int
		SendMillionTimeStamp     int64
		SenderEnterRoomTimeStamp int
	}
}

type OnlineRankV2 struct {
	Cmd  string `json:"cmd"`
	Data struct {
		List []struct {
			UID        int64  `json:"uid,omitempty"`
			Face       string `json:"face,omitempty"`
			Score      string `json:"score"`
			Name       string `json:"uname"`
			Rank       int    `json:"rank,omitempty"`
			GuardLevel int    `json:"guard_level,omitempty"`
		} `json:"list"`
		RankType string `json:"rank_type"`
	} `json:"data"`
}

type OnlineRankTop3 struct {
	Cmd  string `json:"cmd"`
	Data struct {
		Dmscore int `json:"dmscore"`
		List    []struct {
			Msg  string `json:"msg"`
			Rank int    `json:"rank"`
		} `json:"list"`
	} `json:"data"`
}

type StopLiveRoomList struct {
	Cmd  string `json:"cmd"`
	Data struct {
		RoomIDList []int `json:"room_id_list"`
	} `json:"data"`
}

type WatchedChange struct {
	Cmd  string `json:"cmd"`
	Data struct {
		Num       int    `json:"num"`
		TextSmall string `json:"text_small"`
		TextLarge string `json:"text_large"`
	} `json:"data"`
}

type OnlineRankCount struct {
	Cmd  string `json:"cmd"`
	Data struct {
		Count int `json:"count"`
	} `json:"data"`
}

type LikeInfoV3Click struct {
	Cmd  string `json:"cmd"`
	Data struct {
		UID              int64     `json:"uid"`
		Name             string    `json:"uname"`
		NameColor        string    `json:"uname_color"`
		ShowArea         int       `json:"show_area"`
		MsgType          int       `json:"msg_type"`
		LikeIcon         string    `json:"like_icon"`
		LikeText         string    `json:"like_text"`
		Identities       []int     `json:"identities"`
		Dmscore          int       `json:"dmscore"`
		FansMedal        FansMedal `json:"fans_medal"`
		ContributionInfo struct {
			Grade int `json:"grade"`
		} `json:"contribution_info"`
	} `json:"data"`
}

type InteractWord struct {
	Cmd  string `json:"cmd"`
	Data struct {
		UID           int64     `json:"uid"`
		Name          string    `json:"uname"`
		NameColor     string    `json:"uname_color"`
		Dmscore       int       `json:"dmscore"`
		Identities    []int     `json:"identities"`
		IsSpread      int       `json:"is_spread"`
		MsgType       int       `json:"msg_type"`
		PrivilegeType int       `json:"privilege_type"`
		RoomId        int       `json:"roomid"`
		Score         int64     `json:"score"`
		SpreadDesc    string    `json:"spread_desc"`
		SpreadInfo    string    `json:"spread_info"`
		TailIcon      int       `json:"tail_icon"`
		Timestamp     int       `json:"timestamp"`
		TriggerTime   int64     `json:"trigger_time"`
		FansMedal     FansMedal `json:"fans_medal"`
		Contribution  struct {
			Grade int `json:"grade"`
		} `json:"contribution"`
	} `json:"data"`
}

type RoomRealTimeMessageUpdate struct {
	Cmd  string `json:"cmd"`
	Data struct {
		Roomid    int `json:"roomid"`
		Fans      int `json:"fans"`
		RedNotice int `json:"red_notice"`
		FansClub  int `json:"fans_club"`
	} `json:"data"`
}

type BlindGift struct {
	BlindGiftConfigId int    `json:"blind_gift_config_id"`
	From              int    `json:"from"`
	GiftAction        string `json:"gift_action"`
	GiftTipPrice      int    `json:"gift_tip_price"`
	OriginalGiftId    int    `json:"original_gift_id"`
	OriginalGiftName  string `json:"original_gift_name"`
	OriginalGiftPrice int    `json:"original_gift_price"`
}

type ReceiveUserInfo struct {
	Uid   int    `json:"uid"`
	Uname string `json:"uname"`
}

type SendGift struct {
	Cmd  string `json:"cmd"`
	Data struct {
		UID            int64  `json:"uid"`
		Name           string `json:"uname"`
		NameColor      string `json:"name_color"`
		Action         string `json:"action"`
		BatchComboID   string `json:"batch_combo_id"`
		BatchComboSend struct {
			Action        string      `json:"action"`
			BatchComboId  string      `json:"batch_combo_id"`
			BatchComboNum int         `json:"batch_combo_num"`
			BlindGift     *BlindGift  `json:"blind_gift"`
			GiftId        int         `json:"gift_id"`
			GiftName      string      `json:"gift_name"`
			GiftNum       int         `json:"gift_num"`
			SendMaster    interface{} `json:"send_master"`
			Uid           int         `json:"uid"`
			Uname         string      `json:"uname"`
		} `json:"batch_combo_send"`
		BeatID            string          `json:"beatId"`
		BizSource         string          `json:"biz_source"`
		BlindGift         *BlindGift      `json:"blind_gift"`
		BroadcastID       int             `json:"broadcast_id"`
		CoinType          string          `json:"coin_type"`
		ComboResourcesID  int             `json:"combo_resources_id"`
		ComboSend         interface{}     `json:"combo_send"`
		ComboStayTime     int             `json:"combo_stay_time"`
		ComboTotalCoin    int             `json:"combo_total_coin"`
		CritProb          int             `json:"crit_prob"`
		Demarcation       int             `json:"demarcation"`
		DiscountPrice     int             `json:"discount_price"`
		Dmscore           int             `json:"dmscore"`
		Draw              int             `json:"draw"`
		Effect            int             `json:"effect"`
		EffectBlock       int             `json:"effect_block"`
		Face              string          `json:"face"`
		FaceEffectID      int             `json:"face_effect_id"`
		FaceEffectType    int             `json:"face_effect_type"`
		FloatScResourceID int             `json:"float_sc_resource_id"`
		GiftID            int             `json:"giftId"`
		GiftName          string          `json:"giftName"`
		GiftType          int             `json:"giftType"`
		Gold              int             `json:"gold"`
		GuardLevel        int             `json:"guard_level"`
		IsFirst           bool            `json:"is_first"`
		IsNaming          bool            `json:"is_naming"`
		IsSpecialBatch    int             `json:"is_special_batch"`
		Magnification     float64         `json:"magnification"`
		Num               int             `json:"num"`
		OriginalGiftName  string          `json:"original_gift_name"`
		Price             int             `json:"price"`
		Rcost             int             `json:"rcost"`
		ReceiveUserInfo   ReceiveUserInfo `json:"receive_user_info"`
		Remain            int             `json:"remain"`
		Rnd               string          `json:"rnd"`
		SendMaster        interface{}     `json:"send_master"`
		Silver            int             `json:"silver"`
		Super             int             `json:"super"`
		SuperBatchGiftNum int             `json:"super_batch_gift_num"`
		SuperGiftNum      int             `json:"super_gift_num"`
		SvgaBlock         int             `json:"svga_block"`
		Switch            bool            `json:"switch"`
		TagImage          string          `json:"tag_image"`
		Tid               string          `json:"tid"`
		Timestamp         int             `json:"timestamp"`
		TopList           interface{}     `json:"top_list"`
		TotalCoin         int             `json:"total_coin"`
		FansMedal         FansMedal       `json:"medal_info"`
	} `json:"data"`
}

type SuperChatMessage struct {
	RoomID int    `json:"RoomId"`
	Cmd    string `json:"cmd"`
	Data   struct {
		BackgroundBottomColor string `json:"background_bottom_color"`
		BackgroundColor       string `json:"background_color"`
		BackgroundIcon        string `json:"background_icon"`
		BackgroundImage       string `json:"background_image"`
		BackgroundPriceColor  string `json:"background_price_color"`
		EndTime               int    `json:"end_time"`
		Gift                  struct {
			GiftID   int    `json:"gift_id"`
			GiftName string `json:"gift_name"`
			Num      int    `json:"num"`
		} `json:"gift"`
		ID        string `json:"id"`
		IsRanked  int    `json:"is_ranked"`
		MedalInfo struct {
			AnchorRoomid int    `json:"anchor_roomid"`
			AnchorUname  string `json:"anchor_uname"`
			IconID       int    `json:"icon_id"`
			MedalColor   string `json:"medal_color"`
			MedalLevel   int    `json:"medal_level"`
			MedalName    string `json:"medal_name"`
			Special      string `json:"special"`
			TargetID     int    `json:"target_id"`
		} `json:"medal_info"`
		Message    string `json:"message"`
		MessageJpn string `json:"message_jpn"`
		Price      int    `json:"price"`
		Rate       int    `json:"rate"`
		StartTime  int    `json:"start_time"`
		Time       int    `json:"time"`
		Token      string `json:"token"`
		Ts         int    `json:"ts"`
		UID        string `json:"uid"`
		UserInfo   struct {
			Face       string `json:"face"`
			FaceFrame  string `json:"face_frame"`
			GuardLevel int    `json:"guard_level"`
			IsMainVip  int    `json:"is_main_vip"`
			IsSvip     int    `json:"is_svip"`
			IsVip      int    `json:"is_vip"`
			LevelColor string `json:"level_color"`
			Manager    int    `json:"manager"`
			Title      string `json:"title"`
			Uname      string `json:"uname"`
			UserLevel  int    `json:"user_level"`
		} `json:"user_info"`
	} `json:"data"`
}

type LikeInfoV3Update struct {
	Cmd  string `json:"cmd"`
	Data struct {
		ClickCount int `json:"click_count"`
	} `json:"data"`
}

type HotRankChange struct {
	Cmd  string `json:"cmd"`
	Data struct {
		Rank        int    `json:"rank"`
		Trend       int    `json:"trend"`
		Countdown   int    `json:"countdown"`
		Timestamp   int    `json:"timestamp"`
		WebURL      string `json:"web_url"`
		LiveURL     string `json:"live_url"`
		BlinkURL    string `json:"blink_url"`
		LiveLinkURL string `json:"live_link_url"`
		PcLinkURL   string `json:"pc_link_url"`
		Icon        string `json:"icon"`
		AreaName    string `json:"area_name"`
		RankDesc    string `json:"rank_desc"`
	} `json:"data"`
}
type NoticeMsg struct {
	Cmd  string `json:"cmd"`
	ID   int    `json:"id"`
	Name string `json:"name"`
	Full struct {
		HeadIcon    string `json:"head_icon"`
		TailIcon    string `json:"tail_icon"`
		HeadIconFa  string `json:"head_icon_fa"`
		TailIconFa  string `json:"tail_icon_fa"`
		HeadIconFan int    `json:"head_icon_fan"`
		TailIconFan int    `json:"tail_icon_fan"`
		Background  string `json:"background"`
		Color       string `json:"color"`
		Highlight   string `json:"highlight"`
		Time        int    `json:"time"`
	} `json:"full"`
	Half struct {
		HeadIcon   string `json:"head_icon"`
		TailIcon   string `json:"tail_icon"`
		Background string `json:"background"`
		Color      string `json:"color"`
		Highlight  string `json:"highlight"`
		Time       int    `json:"time"`
	} `json:"half"`
	Side struct {
		HeadIcon   string `json:"head_icon"`
		Background string `json:"background"`
		Color      string `json:"color"`
		Highlight  string `json:"highlight"`
		Border     string `json:"border"`
		Time       int    `json:"time"`
	} `json:"side"`
	Roomid     int    `json:"roomid"`
	RealRoomid string `json:"real_roomid"`
	MsgCommon  string `json:"msg_common"`
	MsgSelf    string `json:"msg_self"`
	LinkURL    string `json:"link_url"`
	MsgType    int    `json:"msg_type"`
	ShieldUID  int64  `json:"shield_uid"`
	BusinessID string `json:"business_id"`
	Scatter    struct {
		Min int `json:"min"`
		Max int `json:"max"`
	} `json:"scatter"`
	MarqueeID  string `json:"marquee_id"`
	NoticeType int    `json:"notice_type"`
}

type WidgetBanner struct {
	Cmd  string `json:"cmd"`
	Data struct {
		Timestamp  int `json:"timestamp"`
		WidgetList struct {
			Num293 struct {
				ID             int      `json:"id"`
				Title          string   `json:"title"`
				Cover          string   `json:"cover"`
				WebCover       string   `json:"web_cover"`
				TipText        string   `json:"tip_text"`
				TipTextColor   string   `json:"tip_text_color"`
				TipBottomColor string   `json:"tip_bottom_color"`
				JumpURL        string   `json:"jump_url"`
				URL            string   `json:"url"`
				StayTime       int      `json:"stay_time"`
				Site           int      `json:"site"`
				PlatformIn     []string `json:"platform_in"`
				Type           int      `json:"type"`
				BandID         int      `json:"band_id"`
				SubKey         string   `json:"sub_key"`
				SubData        string   `json:"sub_data"`
				IsAdd          bool     `json:"is_add"`
			} `json:"293"`
		} `json:"widget_list"`
	} `json:"data"`
}

type HotRankChangedV2 struct {
	Cmd  string `json:"cmd"`
	Data struct {
		Rank        int    `json:"rank"`
		Trend       int    `json:"trend"`
		Countdown   int    `json:"countdown"`
		Timestamp   int    `json:"timestamp"`
		WebURL      string `json:"web_url"`
		LiveURL     string `json:"live_url"`
		BlinkURL    string `json:"blink_url"`
		LiveLinkURL string `json:"live_link_url"`
		PcLinkURL   string `json:"pc_link_url"`
		Icon        string `json:"icon"`
		AreaName    string `json:"area_name"`
		RankDesc    string `json:"rank_desc"`
	} `json:"data"`
}

type GuardHonorThousand struct {
	Cmd  string `json:"cmd"`
	Data struct {
		Add []interface{} `json:"add"`
		Del []int         `json:"del"`
	} `json:"data"`
}

type Live struct {
	Cmd             string `json:"cmd"`
	LiveKey         string `json:"live_key"`
	VoiceBackground string `json:"voice_background"`
	SubSessionKey   string `json:"sub_session_key"`
	LivePlatform    string `json:"live_platform"`
	LiveModel       int    `json:"live_model"`
	LiveTime        int    `json:"live_time"`
	Roomid          int    `json:"roomid"`
}

type RoomChange struct {
	Cmd  string `json:"cmd"`
	Data struct {
		Title          string `json:"title"`
		AreaID         int    `json:"area_id"`
		ParentAreaID   int    `json:"parent_area_id"`
		AreaName       string `json:"area_name"`
		ParentAreaName string `json:"parent_area_name"`
		LiveKey        string `json:"live_key"`
		SubSessionKey  string `json:"sub_session_key"`
	} `json:"data"`
}

type RoomBlockMsg struct {
	Cmd  string `json:"cmd"`
	Data struct {
		Dmscore  int    `json:"dmscore"`
		Operator int    `json:"operator"`
		UID      int64  `json:"uid"`
		Uname    string `json:"uname"`
	} `json:"data"`
	UID  string `json:"uid"`
	Name string `json:"uname"`
}

type FullScreenSpecialEffect struct {
	Cmd  string `json:"cmd"`
	Data struct {
		Type       int   `json:"type"`
		Ids        []int `json:"ids"`
		Queue      int   `json:"queue"`
		PlatformIn []int `json:"platform_in"`
	} `json:"data"`
}

type CommonNoticeDanmaku struct {
	Cmd  string `json:"cmd"`
	Data struct {
		ContentSegments []struct {
			FontColor              string `json:"font_color"`
			FontColorDark          string `json:"font_color_dark"`
			HighlightFontColor     string `json:"highlight_font_color"`
			HighlightFontColorDark string `json:"highlight_font_color_dark"`
			Text                   string `json:"text"`
			Type                   int    `json:"type"`
		} `json:"content_segments"`
		Dmscore   int   `json:"dmscore"`
		Terminals []int `json:"terminals"`
	} `json:"data"`
}

type TradingScore struct {
	Cmd  string `json:"cmd"`
	Data struct {
		BubbleShowTime int   `json:"bubble_show_time"`
		Num            int   `json:"num"`
		ScoreID        int   `json:"score_id"`
		UID            int64 `json:"uid"`
		UpdateTime     int   `json:"update_time"`
		UpdateType     int   `json:"update_type"`
	} `json:"data"`
}

type Preparing struct {
	Cmd    string `json:"cmd"`
	RoomId int    `json:"roomid"`
}

type GuardBuy struct {
	Cmd  string `json:"cmd"`
	Data struct {
		UID        int64  `json:"uid"`
		Username   string `json:"username"`
		GuardLevel int    `json:"guard_level"`
		Num        int    `json:"num"`
		Price      int    `json:"price"`
		GiftID     int    `json:"gift_id"`
		GiftName   string `json:"gift_name"`
		StartTime  int    `json:"start_time"`
		EndTime    int    `json:"end_time"`
	} `json:"data"`
}

type GiftStarProcess struct {
	Cmd  string `json:"cmd"`
	Data struct {
		Status int    `json:"status"`
		Tip    string `json:"tip"`
	} `json:"data"`
}

type RoomSkinMsg struct {
	Cmd         string `json:"cmd"`
	SkinID      int    `json:"skin_id"`
	Status      int    `json:"status"`
	EndTime     int    `json:"end_time"`
	CurrentTime int    `json:"current_time"`
	OnlyLocal   bool   `json:"only_local"`
	Scatter     struct {
		Min int `json:"min"`
		Max int `json:"max"`
	} `json:"scatter"`
	SkinConfig struct {
		Android struct {
			Num1 struct {
				Zip string `json:"zip"`
				Md5 string `json:"md5"`
			} `json:"1"`
		} `json:"android"`
		Ios struct {
			Num1 struct {
				Zip string `json:"zip"`
				Md5 string `json:"md5"`
			} `json:"1"`
		} `json:"ios"`
		Ipad struct {
			Num1 struct {
				Zip string `json:"zip"`
				Md5 string `json:"md5"`
			} `json:"1"`
		} `json:"ipad"`
		Web struct {
			Num1 struct {
				Zip              string `json:"zip"`
				Md5              string `json:"md5"`
				Platform         string `json:"platform"`
				Version          string `json:"version"`
				HeadInfoBgPic    string `json:"headInfoBgPic"`
				GiftControlBgPic string `json:"giftControlBgPic"`
				RankListBgPic    string `json:"rankListBgPic"`
				MainText         string `json:"mainText"`
				NormalText       string `json:"normalText"`
				HighlightContent string `json:"highlightContent"`
				Border           string `json:"border"`
				ButtonText       string `json:"buttonText"`
			} `json:"1"`
		} `json:"web"`
	} `json:"skin_config"`
}

type EntryEffect struct {
	Cmd  string `json:"cmd"`
	Data struct {
		ID                   int           `json:"id"`
		UID                  int64         `json:"uid"`
		TargetID             int           `json:"target_id"`
		MockEffect           int           `json:"mock_effect"`
		Face                 string        `json:"face"`
		PrivilegeType        int           `json:"privilege_type"`
		CopyWriting          string        `json:"copy_writing"`
		CopyColor            string        `json:"copy_color"`
		HighlightColor       string        `json:"highlight_color"`
		Priority             int           `json:"priority"`
		BasemapURL           string        `json:"basemap_url"`
		ShowAvatar           int           `json:"show_avatar"`
		EffectiveTime        int           `json:"effective_time"`
		WebBasemapURL        string        `json:"web_basemap_url"`
		WebEffectiveTime     int           `json:"web_effective_time"`
		WebEffectClose       int           `json:"web_effect_close"`
		WebCloseTime         int           `json:"web_close_time"`
		Business             int           `json:"business"`
		CopyWritingV2        string        `json:"copy_writing_v2"`
		IconList             []interface{} `json:"icon_list"`
		MaxDelayTime         int           `json:"max_delay_time"`
		TriggerTime          int64         `json:"trigger_time"`
		Identities           int           `json:"identities"`
		EffectSilentTime     int           `json:"effect_silent_time"`
		EffectiveTimeNew     float64       `json:"effective_time_new"`
		WebDynamicURLWebp    string        `json:"web_dynamic_url_webp"`
		WebDynamicURLApng    string        `json:"web_dynamic_url_apng"`
		MobileDynamicURLWebp string        `json:"mobile_dynamic_url_webp"`
	} `json:"data"`
}

type UserToastMsg struct {
	Cmd  string `json:"cmd"`
	Data struct {
		AnchorShow       bool   `json:"anchor_show"`
		Color            string `json:"color"`
		Dmscore          int    `json:"dmscore"`
		EffectID         int    `json:"effect_id"`
		EndTime          int    `json:"end_time"`
		FaceEffectID     int    `json:"face_effect_id"`
		GiftID           int    `json:"gift_id"`
		GuardLevel       int    `json:"guard_level"`
		IsShow           int    `json:"is_show"`
		Num              int    `json:"num"`
		OpType           int    `json:"op_type"`
		PayflowID        string `json:"payflow_id"`
		Price            int    `json:"price"`
		RoleName         string `json:"role_name"`
		RoomEffectID     int    `json:"room_effect_id"`
		StartTime        int    `json:"start_time"`
		SvgaBlock        int    `json:"svga_block"`
		TargetGuardCount int    `json:"target_guard_count"`
		ToastMsg         string `json:"toast_msg"`
		UID              int64  `json:"uid"`
		Unit             string `json:"unit"`
		UserShow         bool   `json:"user_show"`
		Username         string `json:"username"`
	} `json:"data"`
}

type HeartBeatReply struct {
	Sum int `json:"sum"`
}

type PopularityRedPocketNew struct {
	Cmd  string `json:"cmd"`
	Data struct {
		LotId       int    `json:"lot_id"`
		StartTime   int    `json:"start_time"`
		CurrentTime int    `json:"current_time"`
		WaitNum     int    `json:"wait_num"`
		Uname       string `json:"uname"`
		Uid         int64  `json:"uid"`
		Action      string `json:"action"`
		Num         int    `json:"num"`
		GiftName    string `json:"gift_name"`
		GiftId      int    `json:"gift_id"`
		Price       int    `json:"price"`
		NameColor   string `json:"name_color"`
		MedalInfo   struct {
			TargetId         int    `json:"target_id"`
			Special          string `json:"special"`
			IconId           int    `json:"icon_id"`
			AnchorUname      string `json:"anchor_uname"`
			AnchorRoomid     int    `json:"anchor_roomid"`
			MedalLevel       int    `json:"medal_level"`
			MedalName        string `json:"medal_name"`
			MedalColor       int    `json:"medal_color"`
			MedalColorStart  int    `json:"medal_color_start"`
			MedalColorEnd    int    `json:"medal_color_end"`
			MedalColorBorder int    `json:"medal_color_border"`
			IsLighted        int    `json:"is_lighted"`
			GuardLevel       int    `json:"guard_level"`
		} `json:"medal_info"`
	} `json:"data"`
}

type AreaRankChanged struct {
	Cmd  string `json:"cmd"`
	Data struct {
		ConfId      int    `json:"conf_id"`
		RankName    string `json:"rank_name"`
		Uid         int64  `json:"uid"`
		Rank        int    `json:"rank"`
		IconUrlBlue string `json:"icon_url_blue"`
		IconUrlPink string `json:"icon_url_pink"`
		IconUrlGrey string `json:"icon_url_grey"`
		ActionType  int    `json:"action_type"`
		Timestamp   int    `json:"timestamp"`
		MsgId       string `json:"msg_id"`
		JumpUrlLink string `json:"jump_url_link"`
		JumpUrlPc   string `json:"jump_url_pc"`
		JumpUrlPink string `json:"jump_url_pink"`
		JumpUrlWeb  string `json:"jump_url_web"`
	} `json:"data"`
}

type SuperChatEntrance struct {
	Cmd  string `json:"cmd"`
	Data struct {
		Status        int    `json:"status"`
		JumpUrl       string `json:"jump_url"`
		Icon          string `json:"icon"`
		BroadcastType int    `json:"broadcast_type"`
	} `json:"data"`
	Roomid string `json:"roomid"`
}

type PlayTogether struct {
	Cmd  string `json:"cmd"`
	Data struct {
		Ruid        int    `json:"ruid"`
		Roomid      int    `json:"roomid"`
		Action      string `json:"action"`
		Uid         int64  `json:"uid"`
		Timestamp   int    `json:"timestamp"`
		Message     string `json:"message"`
		MessageType int    `json:"message_type"`
		JumpUrl     string `json:"jump_url"`
		WebUrl      string `json:"web_url"`
		ApplyNumber int    `json:"apply_number"`
		RefreshTool bool   `json:"refresh_tool"`
		CurFleetNum int    `json:"cur_fleet_num"`
		MaxFleetNum int    `json:"max_fleet_num"`
	} `json:"data"`
}

type ComboSend struct {
	Cmd  string `json:"cmd"`
	Data struct {
		Action         string `json:"action"`
		BatchComboId   string `json:"batch_combo_id"`
		BatchComboNum  int    `json:"batch_combo_num"`
		ComboId        string `json:"combo_id"`
		ComboNum       int    `json:"combo_num"`
		ComboTotalCoin int    `json:"combo_total_coin"`
		Dmscore        int    `json:"dmscore"`
		GiftId         int    `json:"gift_id"`
		GiftName       string `json:"gift_name"`
		GiftNum        int    `json:"gift_num"`
		IsJoinReceiver bool   `json:"is_join_receiver"`
		IsNaming       bool   `json:"is_naming"`
		IsShow         int    `json:"is_show"`
		MedalInfo      struct {
			AnchorRoomid     int    `json:"anchor_roomid"`
			AnchorUname      string `json:"anchor_uname"`
			GuardLevel       int    `json:"guard_level"`
			IconId           int    `json:"icon_id"`
			IsLighted        int    `json:"is_lighted"`
			MedalColor       int    `json:"medal_color"`
			MedalColorBorder int    `json:"medal_color_border"`
			MedalColorEnd    int    `json:"medal_color_end"`
			MedalColorStart  int    `json:"medal_color_start"`
			MedalLevel       int    `json:"medal_level"`
			MedalName        string `json:"medal_name"`
			Special          string `json:"special"`
			TargetId         int64  `json:"target_id"`
		} `json:"medal_info"`
		NameColor       string `json:"name_color"`
		RUname          string `json:"r_uname"`
		ReceiveUserInfo struct {
			Uid   int64  `json:"uid"`
			Uname string `json:"uname"`
		} `json:"receive_user_info"`
		Ruid       int         `json:"ruid"`
		SendMaster interface{} `json:"send_master"`
		TotalNum   int         `json:"total_num"`
		Uid        int64       `json:"uid"`
		Uname      string      `json:"uname"`
	} `json:"data"`
}

type PopularityRedPocketStart struct {
	Cmd  string `json:"cmd"`
	Data struct {
		LotId           int    `json:"lot_id"`
		SenderUid       int64  `json:"sender_uid"`
		SenderName      string `json:"sender_name"`
		SenderFace      string `json:"sender_face"`
		JoinRequirement int    `json:"join_requirement"`
		Danmu           string `json:"danmu"`
		CurrentTime     int    `json:"current_time"`
		StartTime       int    `json:"start_time"`
		EndTime         int    `json:"end_time"`
		LastTime        int    `json:"last_time"`
		RemoveTime      int    `json:"remove_time"`
		ReplaceTime     int    `json:"replace_time"`
		LotStatus       int    `json:"lot_status"`
		H5Url           string `json:"h5_url"`
		UserStatus      int    `json:"user_status"`
		Awards          []struct {
			GiftId   int    `json:"gift_id"`
			GiftName string `json:"gift_name"`
			GiftPic  string `json:"gift_pic"`
			Num      int    `json:"num"`
		} `json:"awards"`
		LotConfigId int `json:"lot_config_id"`
		TotalPrice  int `json:"total_price"`
		WaitNum     int `json:"wait_num"`
	} `json:"data"`
}

type PkBattleProcess struct {
	Cmd  string `json:"cmd"`
	Data struct {
		BattleType int `json:"battle_type"`
		InitInfo   struct {
			RoomId     int    `json:"room_id"`
			Votes      int    `json:"votes"`
			BestUname  string `json:"best_uname"`
			VisionDesc int    `json:"vision_desc"`
		} `json:"init_info"`
		MatchInfo struct {
			RoomId     int    `json:"room_id"`
			Votes      int    `json:"votes"`
			BestUname  string `json:"best_uname"`
			VisionDesc int    `json:"vision_desc"`
		} `json:"match_info"`
	} `json:"data"`
	PkId      int `json:"pk_id"`
	PkStatus  int `json:"pk_status"`
	Timestamp int `json:"timestamp"`
}

type PopularRankChanged struct {
	Cmd  string `json:"cmd"`
	Data struct {
		Uid       int64  `json:"uid"`
		Rank      int    `json:"rank"`
		Countdown int    `json:"countdown"`
		Timestamp int    `json:"timestamp"`
		CacheKey  string `json:"cache_key"`
	} `json:"data"`
}

type PkBattleStartNew struct {
	Cmd       string `json:"cmd"`
	PkId      int    `json:"pk_id"`
	PkStatus  int    `json:"pk_status"`
	Timestamp int    `json:"timestamp"`
	Data      struct {
		BattleType    int    `json:"battle_type"`
		FinalHitVotes int    `json:"final_hit_votes"`
		PkStartTime   int    `json:"pk_start_time"`
		PkFrozenTime  int    `json:"pk_frozen_time"`
		PkEndTime     int    `json:"pk_end_time"`
		PkVotesType   int    `json:"pk_votes_type"`
		PkVotesAdd    int    `json:"pk_votes_add"`
		PkVotesName   string `json:"pk_votes_name"`
		StarLightMsg  string `json:"star_light_msg"`
		PkCountdown   int    `json:"pk_countdown"`
		FinalConf     struct {
			Switch    int `json:"switch"`
			StartTime int `json:"start_time"`
			EndTime   int `json:"end_time"`
		} `json:"final_conf"`
		InitInfo struct {
			RoomId     int `json:"room_id"`
			DateStreak int `json:"date_streak"`
		} `json:"init_info"`
		MatchInfo struct {
			RoomId     int `json:"room_id"`
			DateStreak int `json:"date_streak"`
		} `json:"match_info"`
	} `json:"data"`
	Roomid string `json:"roomid"`
}

type DanMuAggregation struct {
	Cmd  string `json:"cmd"`
	Data struct {
		ActivityIdentity string `json:"activity_identity"`
		ActivitySource   int    `json:"activity_source"`
		AggregationCycle int    `json:"aggregation_cycle"`
		AggregationIcon  string `json:"aggregation_icon"`
		AggregationNum   int    `json:"aggregation_num"`
		BroadcastMsgType int    `json:"broadcast_msg_type"`
		Dmscore          int    `json:"dmscore"`
		Msg              string `json:"msg"`
		ShowRows         int    `json:"show_rows"`
		ShowTime         int    `json:"show_time"`
		Timestamp        int    `json:"timestamp"`
	} `json:"data"`
}

type LiveInteractiveGame struct {
	Cmd  string `json:"cmd"`
	Data struct {
		Type           int         `json:"type"`
		Uid            int64       `json:"uid"`
		Uname          string      `json:"uname"`
		Uface          string      `json:"uface"`
		GiftId         int         `json:"gift_id"`
		GiftName       string      `json:"gift_name"`
		GiftNum        int         `json:"gift_num"`
		Price          int         `json:"price"`
		Paid           bool        `json:"paid"`
		Msg            string      `json:"msg"`
		FansMedalLevel int         `json:"fans_medal_level"`
		GuardLevel     int         `json:"guard_level"`
		Timestamp      int         `json:"timestamp"`
		AnchorLottery  interface{} `json:"anchor_lottery"`
		PkInfo         interface{} `json:"pk_info"`
		AnchorInfo     interface{} `json:"anchor_info"`
		ComboInfo      interface{} `json:"combo_info"`
	} `json:"data"`
}

type RecommendCard struct {
	Cmd  string `json:"cmd"`
	Data struct {
		TitleIcon     string `json:"title_icon"`
		RecommendList []struct {
			ShoppingCardDetail struct {
				GoodsId             string      `json:"goods_id"`
				GoodsName           string      `json:"goods_name"`
				GoodsPrice          string      `json:"goods_price"`
				GoodsMaxPrice       string      `json:"goods_max_price"`
				SaleStatus          int         `json:"sale_status"`
				CouponName          string      `json:"coupon_name"`
				GoodsIcon           string      `json:"goods_icon"`
				GoodsStatus         int         `json:"goods_status"`
				Source              int         `json:"source"`
				H5Url               string      `json:"h5_url"`
				JumpLink            string      `json:"jump_link"`
				SchemaUrl           string      `json:"schema_url"`
				IsPreSale           int         `json:"is_pre_sale"`
				ActivityInfo        interface{} `json:"activity_info"`
				PreSaleInfo         interface{} `json:"pre_sale_info"`
				EarlyBirdInfo       interface{} `json:"early_bird_info"`
				Timestamp           int         `json:"timestamp"`
				CouponDiscountPrice string      `json:"coupon_discount_price"`
				SellingPoint        string      `json:"selling_point"`
				HotBuyNum           int         `json:"hot_buy_num"`
				GiftBuyInfo         interface{} `json:"gift_buy_info"`
				IsExclusive         bool        `json:"is_exclusive"`
				CouponId            string      `json:"coupon_id"`
				RewardInfo          interface{} `json:"reward_info"`
				GoodsTagList        interface{} `json:"goods_tag_list"`
				VirtualExtraInfo    struct {
					GoodsType        int `json:"goods_type"`
					WebContainerType int `json:"web_container_type"`
				} `json:"virtual_extra_info"`
				PriceInfo struct {
					Normal struct {
						PrefixPrice   string `json:"prefix_price"`
						SalePrice     string `json:"sale_price"`
						SuffixPrice   string `json:"suffix_price"`
						StrockPrice   string `json:"strock_price"`
						SaleStartTime int    `json:"sale_start_time"`
						SaleEndTime   int    `json:"sale_end_time"`
					} `json:"normal"`
					Activity interface{} `json:"activity"`
				} `json:"price_info"`
				BtnInfo struct {
					CardBtnStatus int    `json:"card_btn_status"`
					CardBtnTitle  string `json:"card_btn_title"`
					CardBtnStyle  int    `json:"card_btn_style"`
				} `json:"btn_info"`
				GoodsSortId int `json:"goods_sort_id"`
			} `json:"shopping_card_detail"`
			RecommendCardExtra interface{} `json:"recommend_card_extra"`
		} `json:"recommend_list"`
		Timestamp int `json:"timestamp"`
	} `json:"data"`
}

type PkBattleProcessNew struct {
	Cmd  string `json:"cmd"`
	Data struct {
		BattleType int `json:"battle_type"`
		InitInfo   struct {
			RoomId     int    `json:"room_id"`
			Votes      int    `json:"votes"`
			BestUname  string `json:"best_uname"`
			AssistInfo []struct {
				Rank  int    `json:"rank"`
				Uid   int64  `json:"uid"`
				Face  string `json:"face"`
				Uname string `json:"uname"`
			} `json:"assist_info"`
		} `json:"init_info"`
		MatchInfo struct {
			RoomId     int         `json:"room_id"`
			Votes      int         `json:"votes"`
			BestUname  string      `json:"best_uname"`
			AssistInfo interface{} `json:"assist_info"`
		} `json:"match_info"`
	} `json:"data"`
	PkId      int `json:"pk_id"`
	PkStatus  int `json:"pk_status"`
	Timestamp int `json:"timestamp"`
}

type PkBattlePreNew struct {
	Cmd       string `json:"cmd"`
	PkStatus  int    `json:"pk_status"`
	PkId      int    `json:"pk_id"`
	Timestamp int    `json:"timestamp"`
	Data      struct {
		BattleType  int         `json:"battle_type"`
		MatchType   int         `json:"match_type"`
		Uname       string      `json:"uname"`
		Face        string      `json:"face"`
		Uid         int64       `json:"uid"`
		RoomId      int         `json:"room_id"`
		SeasonId    int         `json:"season_id"`
		PreTimer    int         `json:"pre_timer"`
		PkVotesName string      `json:"pk_votes_name"`
		EndWinTask  interface{} `json:"end_win_task"`
	} `json:"data"`
	Roomid int `json:"roomid"`
}

type PkBattlePre struct {
	Cmd       string `json:"cmd"`
	PkStatus  int    `json:"pk_status"`
	PkId      int    `json:"pk_id"`
	Timestamp int    `json:"timestamp"`
	Data      struct {
		BattleType  int         `json:"battle_type"`
		MatchType   int         `json:"match_type"`
		Uname       string      `json:"uname"`
		Face        string      `json:"face"`
		Uid         int64       `json:"uid"`
		RoomId      int         `json:"room_id"`
		SeasonId    int         `json:"season_id"`
		PreTimer    int         `json:"pre_timer"`
		PkVotesName string      `json:"pk_votes_name"`
		EndWinTask  interface{} `json:"end_win_task"`
	} `json:"data"`
	Roomid int `json:"roomid"`
}

type PkBattleFinalProcess struct {
	Cmd  string `json:"cmd"`
	Data struct {
		BattleType   int `json:"battle_type"`
		PkFrozenTime int `json:"pk_frozen_time"`
	} `json:"data"`
	PkId      int `json:"pk_id"`
	PkStatus  int `json:"pk_status"`
	Timestamp int `json:"timestamp"`
}

type PkBattleStart struct {
	Cmd       string `json:"cmd"`
	PkId      int    `json:"pk_id"`
	PkStatus  int    `json:"pk_status"`
	Timestamp int    `json:"timestamp"`
	Data      struct {
		BattleType    int    `json:"battle_type"`
		FinalHitVotes int    `json:"final_hit_votes"`
		PkStartTime   int    `json:"pk_start_time"`
		PkFrozenTime  int    `json:"pk_frozen_time"`
		PkEndTime     int    `json:"pk_end_time"`
		PkVotesType   int    `json:"pk_votes_type"`
		PkVotesAdd    int    `json:"pk_votes_add"`
		PkVotesName   string `json:"pk_votes_name"`
		StarLightMsg  string `json:"star_light_msg"`
		PkCountdown   int    `json:"pk_countdown"`
		FinalConf     struct {
			Switch    int `json:"switch"`
			StartTime int `json:"start_time"`
			EndTime   int `json:"end_time"`
		} `json:"final_conf"`
		InitInfo struct {
			RoomId     int `json:"room_id"`
			DateStreak int `json:"date_streak"`
		} `json:"init_info"`
		MatchInfo struct {
			RoomId     int `json:"room_id"`
			DateStreak int `json:"date_streak"`
		} `json:"match_info"`
	} `json:"data"`
	Roomid string `json:"roomid"`
}

type WidgetGiftStarProcess struct {
	Cmd  string `json:"cmd"`
	Data struct {
		StartDate   int `json:"start_date"`
		ProcessList []struct {
			GiftId       int    `json:"gift_id"`
			GiftImg      string `json:"gift_img"`
			GiftName     string `json:"gift_name"`
			CompletedNum int    `json:"completed_num"`
			TargetNum    int    `json:"target_num"`
		} `json:"process_list"`
		Finished       bool   `json:"finished"`
		DdlTimestamp   int    `json:"ddl_timestamp"`
		Version        int64  `json:"version"`
		RewardGift     int    `json:"reward_gift"`
		RewardGiftImg  string `json:"reward_gift_img"`
		RewardGiftName string `json:"reward_gift_name"`
	} `json:"data"`
}

type PopularityRedPocketWinnerList struct {
	Cmd  string `json:"cmd"`
	Data struct {
		LotId      int             `json:"lot_id"`
		TotalNum   int             `json:"total_num"`
		WinnerInfo [][]interface{} `json:"winner_info"`
		Awards     struct {
			Field1 struct {
				AwardType   int    `json:"award_type"`
				AwardName   string `json:"award_name"`
				AwardPic    string `json:"award_pic"`
				AwardBigPic string `json:"award_big_pic"`
				AwardPrice  int    `json:"award_price"`
			} `json:"31225"`
			Field2 struct {
				AwardType   int    `json:"award_type"`
				AwardName   string `json:"award_name"`
				AwardPic    string `json:"award_pic"`
				AwardBigPic string `json:"award_big_pic"`
				AwardPrice  int    `json:"award_price"`
			} `json:"31251"`
			Field3 struct {
				AwardType   int    `json:"award_type"`
				AwardName   string `json:"award_name"`
				AwardPic    string `json:"award_pic"`
				AwardBigPic string `json:"award_big_pic"`
				AwardPrice  int    `json:"award_price"`
			} `json:"31278"`
		} `json:"awards"`
		Version int `json:"version"`
	} `json:"data"`
}

type GotoBuyFlow struct {
	Cmd  string `json:"cmd"`
	Data struct {
		Text string `json:"text"`
	} `json:"data"`
}

type PkBattleEnd struct {
	Cmd       string `json:"cmd"`
	PkId      string `json:"pk_id"`
	PkStatus  int    `json:"pk_status"`
	Timestamp int    `json:"timestamp"`
	Data      struct {
		BattleType int `json:"battle_type"`
		Timer      int `json:"timer"`
		InitInfo   struct {
			RoomId     int    `json:"room_id"`
			Votes      int    `json:"votes"`
			WinnerType int    `json:"winner_type"`
			BestUname  string `json:"best_uname"`
		} `json:"init_info"`
		MatchInfo struct {
			RoomId     int    `json:"room_id"`
			Votes      int    `json:"votes"`
			WinnerType int    `json:"winner_type"`
			BestUname  string `json:"best_uname"`
		} `json:"match_info"`
	} `json:"data"`
}

type PkBattleSettleUser struct {
	Cmd          string `json:"cmd"`
	PkId         int    `json:"pk_id"`
	PkStatus     int    `json:"pk_status"`
	SettleStatus int    `json:"settle_status"`
	Timestamp    int    `json:"timestamp"`
	Data         struct {
		PkId         string `json:"pk_id"`
		SeasonId     int    `json:"season_id"`
		SettleStatus int    `json:"settle_status"`
		ResultType   int    `json:"result_type"`
		BattleType   int    `json:"battle_type"`
		ResultInfo   struct {
			TotalScore        int           `json:"total_score"`
			ResultTypeScore   int           `json:"result_type_score"`
			PkVotes           int           `json:"pk_votes"`
			PkVotesName       string        `json:"pk_votes_name"`
			PkCritScore       int           `json:"pk_crit_score"`
			PkResistCritScore int           `json:"pk_resist_crit_score"`
			PkExtraScoreSlot  string        `json:"pk_extra_score_slot"`
			PkExtraValue      int           `json:"pk_extra_value"`
			PkExtraScore      int           `json:"pk_extra_score"`
			PkTaskScore       int           `json:"pk_task_score"`
			PkTimesScore      int           `json:"pk_times_score"`
			PkDoneTimes       int           `json:"pk_done_times"`
			PkTotalTimes      int           `json:"pk_total_times"`
			WinCount          int           `json:"win_count"`
			WinFinalHit       int           `json:"win_final_hit"`
			WinnerCountScore  int           `json:"winner_count_score"`
			TaskScoreList     []interface{} `json:"task_score_list"`
		} `json:"result_info"`
		Winner struct {
			RoomId    int    `json:"room_id"`
			Uid       int64  `json:"uid"`
			Uname     string `json:"uname"`
			Face      string `json:"face"`
			FaceFrame string `json:"face_frame"`
			Exp       struct {
				Color       int `json:"color"`
				UserLevel   int `json:"user_level"`
				MasterLevel struct {
					Color int `json:"color"`
					Level int `json:"level"`
				} `json:"master_level"`
			} `json:"exp"`
			BestUser struct {
				Uid         int64  `json:"uid"`
				Uname       string `json:"uname"`
				Face        string `json:"face"`
				PkVotes     int    `json:"pk_votes"`
				PkVotesName string `json:"pk_votes_name"`
				Exp         struct {
					Color int `json:"color"`
					Level int `json:"level"`
				} `json:"exp"`
				FaceFrame string `json:"face_frame"`
				Badge     struct {
					Url      string `json:"url"`
					Desc     string `json:"desc"`
					Position int    `json:"position"`
				} `json:"badge"`
				AwardInfo           interface{}   `json:"award_info"`
				AwardInfoList       []interface{} `json:"award_info_list"`
				EndWinAwardInfoList struct {
					List []interface{} `json:"list"`
				} `json:"end_win_award_info_list"`
			} `json:"best_user"`
		} `json:"winner"`
		MyInfo struct {
			RoomId    int    `json:"room_id"`
			Uid       int64  `json:"uid"`
			Uname     string `json:"uname"`
			Face      string `json:"face"`
			FaceFrame string `json:"face_frame"`
			Exp       struct {
				Color       int `json:"color"`
				UserLevel   int `json:"user_level"`
				MasterLevel struct {
					Color int `json:"color"`
					Level int `json:"level"`
				} `json:"master_level"`
			} `json:"exp"`
			BestUser struct {
				Uid         int64  `json:"uid"`
				Uname       string `json:"uname"`
				Face        string `json:"face"`
				PkVotes     int    `json:"pk_votes"`
				PkVotesName string `json:"pk_votes_name"`
				Exp         struct {
					Color int `json:"color"`
					Level int `json:"level"`
				} `json:"exp"`
				FaceFrame string `json:"face_frame"`
				Badge     struct {
					Url      string `json:"url"`
					Desc     string `json:"desc"`
					Position int    `json:"position"`
				} `json:"badge"`
				AwardInfo           interface{}   `json:"award_info"`
				AwardInfoList       []interface{} `json:"award_info_list"`
				EndWinAwardInfoList struct {
					List []interface{} `json:"list"`
				} `json:"end_win_award_info_list"`
			} `json:"best_user"`
		} `json:"my_info"`
		LevelInfo struct {
			FirstRankName  string `json:"first_rank_name"`
			SecondRankNum  int    `json:"second_rank_num"`
			FirstRankImg   string `json:"first_rank_img"`
			SecondRankIcon string `json:"second_rank_icon"`
		} `json:"level_info"`
	} `json:"data"`
}

type AnchorLotStart struct {
	Cmd  string `json:"cmd"`
	Data struct {
		AssetIcon     string `json:"asset_icon"`
		AssetIconWebp string `json:"asset_icon_webp"`
		AwardImage    string `json:"award_image"`
		AwardName     string `json:"award_name"`
		AwardNum      int    `json:"award_num"`
		AwardType     int    `json:"award_type"`
		CurGiftNum    int    `json:"cur_gift_num"`
		CurrentTime   int    `json:"current_time"`
		Danmu         string `json:"danmu"`
		DanmuNew      []struct {
			Danmu     string `json:"danmu"`
			DanmuView string `json:"danmu_view"`
			Reject    bool   `json:"reject"`
		} `json:"danmu_new"`
		DanmuType      int    `json:"danmu_type"`
		GiftId         int    `json:"gift_id"`
		GiftName       string `json:"gift_name"`
		GiftNum        int    `json:"gift_num"`
		GiftPrice      int    `json:"gift_price"`
		GoawayTime     int    `json:"goaway_time"`
		GoodsId        int    `json:"goods_id"`
		Id             int    `json:"id"`
		IsBroadcast    int    `json:"is_broadcast"`
		JoinType       int    `json:"join_type"`
		LotStatus      int    `json:"lot_status"`
		MaxTime        int    `json:"max_time"`
		RequireText    string `json:"require_text"`
		RequireType    int    `json:"require_type"`
		RequireValue   int    `json:"require_value"`
		RoomId         int    `json:"room_id"`
		SendGiftEnsure int    `json:"send_gift_ensure"`
		ShowPanel      int    `json:"show_panel"`
		StartDontPopup int    `json:"start_dont_popup"`
		Status         int    `json:"status"`
		Time           int    `json:"time"`
		Url            string `json:"url"`
		WebUrl         string `json:"web_url"`
	} `json:"data"`
}

type PkBattleSettleV2 struct {
	Cmd          string `json:"cmd"`
	PkId         int    `json:"pk_id"`
	PkStatus     int    `json:"pk_status"`
	SettleStatus int    `json:"settle_status"`
	Timestamp    int    `json:"timestamp"`
	Data         struct {
		PkId       string `json:"pk_id"`
		SeasonId   int    `json:"season_id"`
		PkType     int    `json:"pk_type"`
		ResultType int    `json:"result_type"`
		ResultInfo struct {
			TotalScore   int    `json:"total_score"`
			PkVotes      int    `json:"pk_votes"`
			PkVotesName  string `json:"pk_votes_name"`
			PkExtraValue int    `json:"pk_extra_value"`
		} `json:"result_info"`
		LevelInfo struct {
			Uid            string `json:"uid"`
			FirstRankName  string `json:"first_rank_name"`
			SecondRankNum  int    `json:"second_rank_num"`
			FirstRankImg   string `json:"first_rank_img"`
			SecondRankIcon string `json:"second_rank_icon"`
		} `json:"level_info"`
		AssistList []struct {
			Id    int    `json:"id"`
			Uname string `json:"uname"`
			Face  string `json:"face"`
			Score int    `json:"score"`
		} `json:"assist_list"`
		StarLightMsg string `json:"star_light_msg"`
	} `json:"data"`
}

type PkBattleSettle struct {
	Cmd          string `json:"cmd"`
	PkId         int    `json:"pk_id"`
	PkStatus     int    `json:"pk_status"`
	SettleStatus int    `json:"settle_status"`
	Timestamp    int    `json:"timestamp"`
	Data         struct {
		PkId         string `json:"pk_id"`
		SeasonId     int    `json:"season_id"`
		SettleStatus int    `json:"settle_status"`
		ResultType   int    `json:"result_type"`
		BattleType   int    `json:"battle_type"`
		ResultInfo   struct {
			TotalScore        int           `json:"total_score"`
			ResultTypeScore   int           `json:"result_type_score"`
			PkVotes           int           `json:"pk_votes"`
			PkVotesName       string        `json:"pk_votes_name"`
			PkCritScore       int           `json:"pk_crit_score"`
			PkResistCritScore int           `json:"pk_resist_crit_score"`
			PkExtraScoreSlot  string        `json:"pk_extra_score_slot"`
			PkExtraValue      int           `json:"pk_extra_value"`
			PkExtraScore      int           `json:"pk_extra_score"`
			PkTaskScore       int           `json:"pk_task_score"`
			PkTimesScore      int           `json:"pk_times_score"`
			PkDoneTimes       int           `json:"pk_done_times"`
			PkTotalTimes      int           `json:"pk_total_times"`
			WinCount          int           `json:"win_count"`
			WinFinalHit       int           `json:"win_final_hit"`
			WinnerCountScore  int           `json:"winner_count_score"`
			TaskScoreList     []interface{} `json:"task_score_list"`
		} `json:"result_info"`
		Winner struct {
			RoomId    int    `json:"room_id"`
			Uid       int64  `json:"uid"`
			Uname     string `json:"uname"`
			Face      string `json:"face"`
			FaceFrame string `json:"face_frame"`
			Exp       struct {
				Color       int `json:"color"`
				UserLevel   int `json:"user_level"`
				MasterLevel struct {
					Color int `json:"color"`
					Level int `json:"level"`
				} `json:"master_level"`
			} `json:"exp"`
			BestUser struct {
				Uid         int64  `json:"uid"`
				Uname       string `json:"uname"`
				Face        string `json:"face"`
				PkVotes     int    `json:"pk_votes"`
				PkVotesName string `json:"pk_votes_name"`
				Exp         struct {
					Color int `json:"color"`
					Level int `json:"level"`
				} `json:"exp"`
				FaceFrame string `json:"face_frame"`
				Badge     struct {
					Url      string `json:"url"`
					Desc     string `json:"desc"`
					Position int    `json:"position"`
				} `json:"badge"`
				AwardInfo           interface{}   `json:"award_info"`
				AwardInfoList       []interface{} `json:"award_info_list"`
				EndWinAwardInfoList struct {
					List []interface{} `json:"list"`
				} `json:"end_win_award_info_list"`
			} `json:"best_user"`
		} `json:"winner"`
		MyInfo struct {
			RoomId    int    `json:"room_id"`
			Uid       int64  `json:"uid"`
			Uname     string `json:"uname"`
			Face      string `json:"face"`
			FaceFrame string `json:"face_frame"`
			Exp       struct {
				Color       int `json:"color"`
				UserLevel   int `json:"user_level"`
				MasterLevel struct {
					Color int `json:"color"`
					Level int `json:"level"`
				} `json:"master_level"`
			} `json:"exp"`
			BestUser struct {
				Uid         int64  `json:"uid"`
				Uname       string `json:"uname"`
				Face        string `json:"face"`
				PkVotes     int    `json:"pk_votes"`
				PkVotesName string `json:"pk_votes_name"`
				Exp         struct {
					Color int `json:"color"`
					Level int `json:"level"`
				} `json:"exp"`
				FaceFrame string `json:"face_frame"`
				Badge     struct {
					Url      string `json:"url"`
					Desc     string `json:"desc"`
					Position int    `json:"position"`
				} `json:"badge"`
				AwardInfo           interface{}   `json:"award_info"`
				AwardInfoList       []interface{} `json:"award_info_list"`
				EndWinAwardInfoList struct {
					List []interface{} `json:"list"`
				} `json:"end_win_award_info_list"`
			} `json:"best_user"`
		} `json:"my_info"`
		LevelInfo struct {
			FirstRankName  string `json:"first_rank_name"`
			SecondRankNum  int    `json:"second_rank_num"`
			FirstRankImg   string `json:"first_rank_img"`
			SecondRankIcon string `json:"second_rank_icon"`
		} `json:"level_info"`
	} `json:"data"`
}

type HotRoomNotify struct {
	Cmd  string `json:"cmd"`
	Data struct {
		Threshold        int `json:"threshold"`
		Ttl              int `json:"ttl"`
		ExitNoRefresh    int `json:"exit_no_refresh"`
		RandomDelayReqV2 []struct {
			Path  string `json:"path"`
			Delay int    `json:"delay"`
		} `json:"random_delay_req_v2"`
	} `json:"data"`
}

type LiveOpenPlatformGame struct {
	Cmd  string `json:"cmd"`
	Data struct {
		MsgType              string      `json:"msg_type"`
		MsgSubType           string      `json:"msg_sub_type"`
		GameName             string      `json:"game_name"`
		GameCode             string      `json:"game_code"`
		GameId               string      `json:"game_id"`
		GameStatus           string      `json:"game_status"`
		GameMsg              string      `json:"game_msg"`
		GameConf             string      `json:"game_conf"`
		InteractivePanelConf string      `json:"interactive_panel_conf"`
		Timestamp            int         `json:"timestamp"`
		BlockUids            interface{} `json:"block_uids"`
	} `json:"data"`
}

type LivePanelChangeContent struct {
	Cmd  string `json:"cmd"`
	Data struct {
		SettingList []struct {
			BizId         int         `json:"biz_id"`
			Icon          string      `json:"icon"`
			Title         string      `json:"title"`
			Note          string      `json:"note"`
			Weight        float64     `json:"weight"`
			StatusType    int         `json:"status_type"`
			Notification  interface{} `json:"notification"`
			Custom        interface{} `json:"custom"`
			JumpUrl       string      `json:"jump_url"`
			TypeId        int         `json:"type_id"`
			Tab           interface{} `json:"tab"`
			DynamicIcon   string      `json:"dynamic_icon"`
			SubIcon       string      `json:"sub_icon"`
			PanelIcon     string      `json:"panel_icon"`
			MatchEntrance int         `json:"match_entrance"`
			IconInfo      interface{} `json:"icon_info"`
		} `json:"setting_list"`
		InteractionList interface{} `json:"interaction_list"`
		OuterList       []struct {
			BizId        int         `json:"biz_id"`
			Icon         string      `json:"icon"`
			Title        string      `json:"title"`
			Note         string      `json:"note"`
			Weight       int         `json:"weight"`
			StatusType   int         `json:"status_type"`
			Notification interface{} `json:"notification"`
			Custom       []struct {
				Icon    string `json:"icon"`
				Title   string `json:"title"`
				Note    string `json:"note"`
				JumpUrl string `json:"jump_url"`
				Status  int    `json:"status"`
				SubIcon string `json:"sub_icon"`
			} `json:"custom"`
			JumpUrl string `json:"jump_url"`
			TypeId  int    `json:"type_id"`
			Tab     *struct {
				Type       string `json:"type"`
				BizType    string `json:"biz_type"`
				TabComment struct {
				} `json:"tab_comment"`
				TabTopic struct {
				} `json:"tab_topic"`
				Aggregation          int    `json:"aggregation"`
				Id                   int    `json:"id"`
				SubTitle             string `json:"sub_title"`
				SubIcon              string `json:"sub_icon"`
				ShowOuterAggregation int    `json:"show_outer_aggregation"`
				ShowGuideBubble      string `json:"show_guide_bubble"`
				GlobalId             string `json:"global_id"`
			} `json:"tab"`
			DynamicIcon   string      `json:"dynamic_icon"`
			SubIcon       string      `json:"sub_icon"`
			PanelIcon     string      `json:"panel_icon"`
			MatchEntrance int         `json:"match_entrance"`
			IconInfo      interface{} `json:"icon_info"`
		} `json:"outer_list"`
		PanelData     interface{} `json:"panel_data"`
		IsFixed       int         `json:"is_fixed"`
		IsMatch       int         `json:"is_match"`
		MatchCristina string      `json:"match_cristina"`
		MatchIcon     string      `json:"match_icon"`
		MatchBgImage  string      `json:"match_bg_image"`
	} `json:"data"`
}

type GiftPanelPlan struct {
	Cmd  string `json:"cmd"`
	Data struct {
		GiftList []struct {
			GiftId int `json:"gift_id"`
			Config struct {
				Id                int    `json:"id"`
				Name              string `json:"name"`
				Price             int    `json:"price"`
				Type              int    `json:"type"`
				CoinType          string `json:"coin_type"`
				BagGift           int    `json:"bag_gift"`
				Effect            int    `json:"effect"`
				CornerMark        string `json:"corner_mark"`
				CornerBackground  string `json:"corner_background"`
				Broadcast         int    `json:"broadcast"`
				Draw              int    `json:"draw"`
				StayTime          int    `json:"stay_time"`
				AnimationFrameNum int    `json:"animation_frame_num"`
				Desc              string `json:"desc"`
				Rule              string `json:"rule"`
				Rights            string `json:"rights"`
				PrivilegeRequired int    `json:"privilege_required"`
				CountMap          []struct {
					Num            int    `json:"num"`
					Text           string `json:"text"`
					Desc           string `json:"desc"`
					WebSvga        string `json:"web_svga"`
					VerticalSvga   string `json:"vertical_svga"`
					HorizontalSvga string `json:"horizontal_svga"`
					SpecialColor   string `json:"special_color"`
					EffectId       int    `json:"effect_id"`
				} `json:"count_map"`
				ImgBasic             string      `json:"img_basic"`
				ImgDynamic           string      `json:"img_dynamic"`
				FrameAnimation       string      `json:"frame_animation"`
				Gif                  string      `json:"gif"`
				Webp                 string      `json:"webp"`
				FullScWeb            string      `json:"full_sc_web"`
				FullScHorizontal     string      `json:"full_sc_horizontal"`
				FullScVertical       string      `json:"full_sc_vertical"`
				FullScHorizontalSvga string      `json:"full_sc_horizontal_svga"`
				FullScVerticalSvga   string      `json:"full_sc_vertical_svga"`
				BulletHead           string      `json:"bullet_head"`
				BulletTail           string      `json:"bullet_tail"`
				LimitInterval        int         `json:"limit_interval"`
				BindRuid             int         `json:"bind_ruid"`
				BindRoomid           int         `json:"bind_roomid"`
				GiftType             int         `json:"gift_type"`
				ComboResourcesId     int         `json:"combo_resources_id"`
				MaxSendLimit         int         `json:"max_send_limit"`
				Weight               int         `json:"weight"`
				GoodsId              int         `json:"goods_id"`
				HasImagedGift        int         `json:"has_imaged_gift"`
				LeftCornerText       string      `json:"left_corner_text"`
				LeftCornerBackground string      `json:"left_corner_background"`
				GiftBanner           interface{} `json:"gift_banner"`
				DiyCountMap          int         `json:"diy_count_map"`
				EffectId             int         `json:"effect_id"`
				FirstTips            string      `json:"first_tips"`
				GiftAttrs            []int       `json:"gift_attrs"`
			} `json:"config"`
			FullScEffect  interface{} `json:"full_sc_effect"`
			FloatScEffect interface{} `json:"float_sc_effect"`
			SpecialType   int         `json:"special_type"`
			Show          bool        `json:"show"`
		} `json:"gift_list"`
		SpecialTypeSort []int `json:"special_type_sort"`
		Action          int   `json:"action"`
	} `json:"data"`
}

type ShoppingExplainCard struct {
	Cmd  string `json:"cmd"`
	Data struct {
		GoodsId       string `json:"goods_id"`
		GoodsName     string `json:"goods_name"`
		GoodsPrice    string `json:"goods_price"`
		GoodsMaxPrice string `json:"goods_max_price"`
		SaleStatus    int    `json:"sale_status"`
		CouponName    string `json:"coupon_name"`
		GoodsIcon     string `json:"goods_icon"`
		Status        int    `json:"status"`
		H5Url         string `json:"h5_url"`
		Source        int    `json:"source"`
		Timestamp     int    `json:"timestamp"`
		IsPreSale     int    `json:"is_pre_sale"`
		ActivityInfo  struct {
			ActivityId         int    `json:"activity_id"`
			ActivityStatus     int    `json:"activity_status"`
			StartTime          int    `json:"start_time"`
			EndTime            int    `json:"end_time"`
			IsAllSku           int    `json:"is_all_sku"`
			Type               int    `json:"type"`
			LowerDiscountPrice string `json:"lower_discount_price"`
			UpperDiscountPrice string `json:"upper_discount_price"`
			WarmUpTime         int    `json:"warm_up_time"`
			ActivitySaleOut    bool   `json:"activity_sale_out"`
		} `json:"activity_info"`
		PreSaleInfo struct {
			Deposit                     string `json:"deposit"`
			DepositType                 int    `json:"deposit_type"`
			MaxDeposit                  string `json:"max_deposit"`
			PresaleStartOrderTime       int    `json:"presale_start_order_time"`
			PresaleEndOrderTime         int    `json:"presale_end_order_time"`
			PreSaleSupplyMoneyStartTime int    `json:"pre_sale_supply_money_start_time"`
			PreSaleSupplyMoneyEndTime   int    `json:"pre_sale_supply_money_end_time"`
		} `json:"pre_sale_info"`
		EarlyBirdInfo       interface{} `json:"early_bird_info"`
		UniqueId            string      `json:"unique_id"`
		Uid                 int64       `json:"uid"`
		SellingPoint        string      `json:"selling_point"`
		CouponDiscountPrice string      `json:"coupon_discount_price"`
		SeiStatus           int         `json:"sei_status"`
		GiftBuyInfo         interface{} `json:"gift_buy_info"`
		RewardInfo          interface{} `json:"reward_info"`
		IsExclusive         bool        `json:"is_exclusive"`
		CouponId            string      `json:"coupon_id"`
		GoodsTagList        interface{} `json:"goods_tag_list"`
		VirtualExtraInfo    interface{} `json:"virtual_extra_info"`
		PriceInfo           interface{} `json:"price_info"`
		BtnInfo             interface{} `json:"btn_info"`
		GoodsSortId         int         `json:"goods_sort_id"`
	} `json:"data"`
}

type AnchorLotCheckStatus struct {
	Cmd  string `json:"cmd"`
	Data struct {
		Id     int   `json:"id"`
		Status int   `json:"status"`
		Uid    int64 `json:"uid"`
	} `json:"data"`
}

type PkBattlePunishEnd struct {
	Cmd       string `json:"cmd"`
	PkId      string `json:"pk_id"`
	PkStatus  int    `json:"pk_status"`
	StatusMsg string `json:"status_msg"`
	Timestamp int    `json:"timestamp"`
	Data      struct {
		BattleType int `json:"battle_type"`
	} `json:"data"`
}

type AnchorLotEnd struct {
	Cmd  string `json:"cmd"`
	Data struct {
		Id int `json:"id"`
	} `json:"data"`
}

type AnchorLotAward struct {
	Cmd  string `json:"cmd"`
	Data struct {
		AwardDontPopup int    `json:"award_dont_popup"`
		AwardImage     string `json:"award_image"`
		AwardName      string `json:"award_name"`
		AwardNum       int    `json:"award_num"`
		AwardType      int    `json:"award_type"`
		AwardUsers     []struct {
			Uid   int64  `json:"uid"`
			Uname string `json:"uname"`
			Face  string `json:"face"`
			Level int    `json:"level"`
			Color int    `json:"color"`
			Num   int    `json:"num"`
		} `json:"award_users"`
		Id        int    `json:"id"`
		LotStatus int    `json:"lot_status"`
		Url       string `json:"url"`
		WebUrl    string `json:"web_url"`
	} `json:"data"`
}

type SpecialGift struct {
	Cmd  string `json:"cmd"`
	Data struct {
		Field1 struct {
			Action   string `json:"action"`
			Content  string `json:"content"`
			HadJoin  int    `json:"hadJoin"`
			Id       string `json:"id"`
			Num      int    `json:"num"`
			StormGif string `json:"storm_gif"`
			Time     int    `json:"time"`
		} `json:"39"`
	} `json:"data"`
}

type SuperChatMessageDelete struct {
	Cmd  string `json:"cmd"`
	Data struct {
		Ids []int `json:"ids"`
	} `json:"data"`
	Roomid int `json:"roomid"`
}

type VoiceJoinRoomCountInfo struct {
	Cmd  string `json:"cmd"`
	Data struct {
		Cmd         string `json:"cmd"`
		RoomId      int    `json:"room_id"`
		RootStatus  int    `json:"root_status"`
		RoomStatus  int    `json:"room_status"`
		ApplyCount  int    `json:"apply_count"`
		NotifyCount int    `json:"notify_count"`
		RedPoint    int    `json:"red_point"`
	} `json:"data"`
	RoomId int `json:"room_id"`
}

type VoiceJoinList struct {
	Cmd  string `json:"cmd"`
	Data struct {
		Cmd        string `json:"cmd"`
		RoomId     int    `json:"room_id"`
		Category   int    `json:"category"`
		ApplyCount int    `json:"apply_count"`
		RedPoint   int    `json:"red_point"`
		Refresh    int    `json:"refresh"`
	} `json:"data"`
	RoomId int `json:"room_id"`
}

type VoiceJoinStatus struct {
	Cmd  string `json:"cmd"`
	Data struct {
		RoomId       int    `json:"room_id"`
		Status       int    `json:"status"`
		Channel      string `json:"channel"`
		ChannelType  string `json:"channel_type"`
		Uid          int64  `json:"uid"`
		UserName     string `json:"user_name"`
		HeadPic      string `json:"head_pic"`
		Guard        int    `json:"guard"`
		StartAt      int    `json:"start_at"`
		CurrentTime  int    `json:"current_time"`
		WebShareLink string `json:"web_share_link"`
	} `json:"data"`
	RoomId int `json:"room_id"`
}

type Warning struct {
	Cmd    string `json:"cmd"`
	Roomid int    `json:"roomid"`
	Msg    string `json:"msg"`
}

type PkBattleRankChange struct {
	Cmd       string `json:"cmd"`
	Timestamp int    `json:"timestamp"`
	Data      struct {
		FirstRankImgUrl string `json:"first_rank_img_url"`
		RankName        string `json:"rank_name"`
	} `json:"data"`
}

type PkBattleSettleNew struct {
	Cmd  string `json:"cmd"`
	Data struct {
		BattleType int `json:"battle_type"`
		DmConf     struct {
			BgColor   string `json:"bg_color"`
			FontColor string `json:"font_color"`
		} `json:"dm_conf"`
		Dmscore  int `json:"dmscore"`
		InitInfo struct {
			AssistInfo []struct {
				Face  string `json:"face"`
				Rank  int    `json:"rank"`
				Uid   int64  `json:"uid"`
				Uname string `json:"uname"`
			} `json:"assist_info"`
			ResultType int `json:"result_type"`
			RoomId     int `json:"room_id"`
			Votes      int `json:"votes"`
		} `json:"init_info"`
		MatchInfo struct {
			AssistInfo []struct {
				Face  string `json:"face"`
				Rank  int    `json:"rank"`
				Uid   int64  `json:"uid"`
				Uname string `json:"uname"`
			} `json:"assist_info"`
			ResultType int `json:"result_type"`
			RoomId     int `json:"room_id"`
			Votes      int `json:"votes"`
		} `json:"match_info"`
		PkId          int `json:"pk_id"`
		PkStatus      int `json:"pk_status"`
		PunishEndTime int `json:"punish_end_time"`
		SettleStatus  int `json:"settle_status"`
		Timestamp     int `json:"timestamp"`
	} `json:"data"`
	PkId      int `json:"pk_id"`
	PkStatus  int `json:"pk_status"`
	Timestamp int `json:"timestamp"`
}

type HotBuyNum struct {
	Cmd  string `json:"cmd"`
	Data struct {
		GoodsId string `json:"goods_id"`
		Num     int    `json:"num"`
	} `json:"data"`
}

type ShoppingCartShow struct {
	Cmd  string `json:"cmd"`
	Data struct {
		Status int `json:"status"`
	} `json:"data"`
}

type VoiceJoinSwitch struct {
	Cmd  string `json:"cmd"`
	Data struct {
		RoomId     int `json:"room_id"`
		RoomStatus int `json:"room_status"`
		RootStatus int `json:"root_status"`
	} `json:"data"`
	Roomid int `json:"roomid"`
}

type CutOff struct {
	Cmd    string `json:"cmd"`
	Msg    string `json:"msg"`
	Roomid int    `json:"roomid"`
}

type RoomAdminRevoke struct {
	Cmd string `json:"cmd"`
	Msg string `json:"msg"`
	Uid int64  `json:"uid"`
}

type RoomSilentOff struct {
	Data struct {
		Type   string `json:"type"`
		Level  int    `json:"level"`
		Second int    `json:"second"`
	} `json:"data"`
	Cmd string `json:"cmd"`
}

type RoomSilentOn struct {
	Data struct {
		Type   string `json:"type"`
		Level  int    `json:"level"`
		Second int    `json:"second"`
	} `json:"data"`
	Cmd string `json:"cmd"`
}

type RoomAdminEntrance struct {
	Cmd     string `json:"cmd"`
	Dmscore int    `json:"dmscore"`
	Level   int    `json:"level"`
	Msg     string `json:"msg"`
	Uid     int64  `json:"uid"`
}

type RoomAdmins struct {
	Cmd  string  `json:"cmd"`
	Uids []int64 `json:"uids"`
}

type VideoConnectionJoinStart struct {
	Cmd  string `json:"cmd"`
	Data struct {
		Status       int    `json:"status"`
		InvitedUid   int64  `json:"invited_uid"`
		ChannelId    string `json:"channel_id"`
		InvitedUname string `json:"invited_uname"`
		InvitedFace  string `json:"invited_face"`
		StartAt      int    `json:"start_at"`
		CurrentTime  int    `json:"current_time"`
	} `json:"data"`
	Roomid int `json:"roomid"`
}

type VideoConnectionMsg struct {
	Cmd  string `json:"cmd"`
	Data struct {
		ChannelId   string `json:"channel_id"`
		CurrentTime int    `json:"current_time"`
		Dmscore     int    `json:"dmscore"`
		Toast       string `json:"toast"`
	} `json:"data"`
}

type VideoConnectionJoinEnd struct {
	Cmd  string `json:"cmd"`
	Data struct {
		ChannelId   string `json:"channel_id"`
		StartAt     int    `json:"start_at"`
		Toast       string `json:"toast"`
		CurrentTime int    `json:"current_time"`
	} `json:"data"`
	Roomid int `json:"roomid"`
}

type RingStatusChange struct {
	Cmd  string `json:"cmd"`
	Data struct {
		Status int `json:"status"`
	} `json:"data"`
}

type RingStatusChangeV2 struct {
	Cmd  string `json:"cmd"`
	Data struct {
		Status int `json:"status"`
	} `json:"data"`
}

type RoomLock struct {
	Cmd    string `json:"cmd"`
	Expire string `json:"expire"`
	Roomid int    `json:"roomid"`
}

type ShoppingBubblesStyle struct {
	Cmd  string `json:"cmd"`
	Data struct {
		IntervalBetweenBubbles int    `json:"interval_between_bubbles"`
		IntervalBetweenQueues  int    `json:"interval_between_queues"`
		CycleTime              int    `json:"cycle_time"`
		GoodsCount             int    `json:"goods_count"`
		Checksum               string `json:"checksum"`
		BubblesList            []struct {
			Tag        string        `json:"tag"`
			Name       string        `json:"name"`
			Priority   int           `json:"priority"`
			ShowBanner int           `json:"show_banner"`
			GoodsList  []interface{} `json:"goods_list"`
		} `json:"bubbles_list"`
	} `json:"data"`
}

type MultiVoiceOperating struct {
	Cmd  string `json:"cmd"`
	Data struct {
		Uid        int64 `json:"uid"`
		TotalPrice int   `json:"total_price"`
		Ts         int64 `json:"ts"`
	} `json:"data"`
}

type MultiVoiceApplicationUser struct {
	Cmd  string `json:"cmd"`
	Data struct {
		Count        int    `json:"count"`
		Uid          int64  `json:"uid"`
		AnchorUid    int64  `json:"anchor_uid"`
		OperateUid   int64  `json:"operate_uid"`
		WantPosition int    `json:"want_position"`
		Event        int    `json:"event"`
		Toast        string `json:"toast"`
		Channel      string `json:"channel"`
		RoomId       int    `json:"roomId"`
		Role         int    `json:"role"`
	} `json:"data"`
}

type PkBattleMatchTimeout struct {
	Cmd  string `json:"cmd"`
	Data struct {
		BattleType int `json:"battle_type"`
	} `json:"data"`
}

type ChangeRoomInfo struct {
	Cmd        string `json:"cmd"`
	Background string `json:"background"`
	Roomid     int    `json:"roomid"`
}

type LiveMultiViewChange struct {
	Cmd  string `json:"cmd"`
	Data struct {
		Scatter struct {
			Max int `json:"max"`
			Min int `json:"min"`
		} `json:"scatter"`
	} `json:"data"`
}

type GuardAchievementRoom struct {
	Cmd  string `json:"cmd"`
	Data struct {
		AnchorBasemapUrl        string `json:"anchor_basemap_url"`
		AnchorGuardAchieveLevel int    `json:"anchor_guard_achieve_level"`
		AnchorModal             struct {
			FirstLineContent  string `json:"first_line_content"`
			HighlightColor    string `json:"highlight_color"`
			SecondLineContent string `json:"second_line_content"`
			ShowTime          int    `json:"show_time"`
		} `json:"anchor_modal"`
		AppBasemapUrl            string `json:"app_basemap_url"`
		CurrentAchievementLevel  int    `json:"current_achievement_level"`
		Dmscore                  int    `json:"dmscore"`
		EventType                int    `json:"event_type"`
		Face                     string `json:"face"`
		FirstLineContent         string `json:"first_line_content"`
		FirstLineHighlightColor  string `json:"first_line_highlight_color"`
		FirstLineNormalColor     string `json:"first_line_normal_color"`
		HeadmapUrl               string `json:"headmap_url"`
		IsFirst                  bool   `json:"is_first"`
		IsFirstNew               bool   `json:"is_first_new"`
		RoomId                   int    `json:"room_id"`
		SecondLineContent        string `json:"second_line_content"`
		SecondLineHighlightColor string `json:"second_line_highlight_color"`
		SecondLineNormalColor    string `json:"second_line_normal_color"`
		ShowTime                 int    `json:"show_time"`
		WebBasemapUrl            string `json:"web_basemap_url"`
	} `json:"data"`
}

type SysMsg struct {
	Cmd string `json:"cmd"`
	Msg string `json:"msg"`
	Url string `json:"url"`
}

type MvRoleChange struct {
	Cmd  string `json:"cmd"`
	Data struct {
		ChangeUid int64 `json:"change_uid"`
		Role      int   `json:"role"`
		RoomId    int   `json:"room_id"`
		Ts        int   `json:"ts"`
	} `json:"data"`
}

type SelectedGoodsInfo struct {
	Cmd  string `json:"cmd"`
	Data struct {
		ChangeType int `json:"change_type"`
		Item       []struct {
			GoodsId             string      `json:"goods_id"`
			GoodsName           string      `json:"goods_name"`
			Source              int         `json:"source"`
			GoodsIcon           string      `json:"goods_icon"`
			IsPreSale           int         `json:"is_pre_sale"`
			ActivityInfo        interface{} `json:"activity_info"`
			PreSaleInfo         interface{} `json:"pre_sale_info"`
			EarlyBirdInfo       interface{} `json:"early_bird_info"`
			CouponDiscountPrice string      `json:"coupon_discount_price"`
			SelectedText        string      `json:"selected_text"`
			IsGiftBuy           int         `json:"is_gift_buy"`
			GoodsPrice          string      `json:"goods_price"`
			GoodsMaxPrice       string      `json:"goods_max_price"`
			RewardInfo          interface{} `json:"reward_info"`
			GoodsTagList        interface{} `json:"goods_tag_list"`
		} `json:"item"`
		Title string `json:"title"`
	} `json:"data"`
}

type MultiVoiceOperatin struct {
	Cmd  string `json:"cmd"`
	Data struct {
		Uid        int64 `json:"uid"`
		TotalPrice int   `json:"total_price"`
		Ts         int64 `json:"ts"`
	} `json:"data"`
}

type PanelInteractiveNotifyChange struct {
	Cmd  string `json:"cmd"`
	Data struct {
		BizId    int    `json:"biz_id"`
		EndTime  int    `json:"end_time"`
		Icon     string `json:"icon"`
		LastTime int    `json:"last_time"`
		Level    int    `json:"level"`
		Text     string `json:"text"`
	} `json:"data"`
}

type InteractiveUser struct {
	Cmd  string `json:"cmd"`
	Data struct {
		Type  int `json:"type"`
		Value struct {
			Delay         int    `json:"delay"`
			DmMsg         string `json:"dm_msg"`
			ProphetStatus int    `json:"prophet_status"`
			SendMsg       int    `json:"send_msg"`
		} `json:"value"`
	} `json:"data"`
}

type UserVirtualMvp struct {
	Cmd  string `json:"cmd"`
	Data struct {
		GoodsId        int    `json:"goods_id"`
		EffectId       int    `json:"effect_id"`
		EffectQueue    int    `json:"effect_queue"`
		Uid            int64  `json:"uid"`
		Uname          string `json:"uname"`
		UnameColor     string `json:"uname_color"`
		UserGuardLevel int    `json:"user_guard_level"`
		GoodsName      string `json:"goods_name"`
		GoodsNum       int    `json:"goods_num"`
		GoodsPrice     int    `json:"goods_price"`
		GoodsIcon      string `json:"goods_icon"`
		Action         string `json:"action"`
		OrderId        string `json:"order_id"`
		Timestamp      int    `json:"timestamp"`
		SuccessToast   string `json:"success_toast"`
		AnimationBlock int    `json:"animation_block"`
	} `json:"data"`
}

type WidgetWishList struct {
	Cmd  string `json:"cmd"`
	Data struct {
		Wish []struct {
			Type       int    `json:"type"`
			GiftId     int    `json:"gift_id"`
			GiftName   string `json:"gift_name"`
			GiftImg    string `json:"gift_img"`
			GiftPrice  int    `json:"gift_price"`
			TargetNum  int    `json:"target_num"`
			CurrentNum int    `json:"current_num"`
		} `json:"wish"`
		WishStatus     int `json:"wish_status"`
		Sid            int `json:"sid"`
		WishStatusInfo []struct {
			WishStatusMsg string `json:"wish_status_msg"`
			WishStatusImg string `json:"wish_status_img"`
			WishStatus    int    `json:"wish_status"`
		} `json:"wish_status_info"`
		WishName string `json:"wish_name"`
	} `json:"data"`
}

type CheckSingStatus struct {
	Cmd  string `json:"cmd"`
	Data struct {
		ShortTimeSize int   `json:"ShortTimeSize"`
		ShortTimeSing int   `json:"ShortTimeSing"`
		LongTimeSize  int   `json:"LongTimeSize"`
		LongTimeSing  int   `json:"LongTimeSing"`
		OpenArea      []int `json:"OpenArea"`
	} `json:"data"`
}

type RoomModuleDisplay struct {
	Cmd  string `json:"cmd"`
	Data struct {
		Timestamp int `json:"timestamp"`
		Modules   struct {
			BottomBanner int `json:"bottom_banner"`
			TopBanner    int `json:"top_banner"`
			WidgetBanner int `json:"widget_banner"`
		} `json:"modules"`
	} `json:"data"`
}

type VoiceChatUpdate struct {
	Data struct {
		Url string `json:"url"`
	} `json:"data"`
	Cmd string `json:"cmd"`
}

type ReenterLiveRoom struct {
	Cmd  string `json:"cmd"`
	Data struct {
		RoomId                int `json:"room_id"`
		RequestRandomSecRange int `json:"request_random_sec_range"`
		Reason                int `json:"reason"`
	} `json:"data"`
	Roomid int `json:"roomid"`
}
