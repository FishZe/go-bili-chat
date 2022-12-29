package handler

const (
	CmdDanmuMsg                  = "DANMU_MSG"
	CmdSuperChatMessage          = "SUPER_CHAT_MESSAGE"
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
	CmdEntryEffect               = "ENTRY_EFFECT"
)

var CmdName = map[string]string{
	"DANMU_MSG":                     "DanMuMsg",
	"SUPER_CHAT_MESSAGE":            "SuperChatMessage",
	"SUPER_CHAT_MESSAGE_JPN":        "SuperChatMessage",
	"WATCHED_CHANGE":                "WatchedChange",
	"SEND_GIFT":                     "SendGift",
	"ONLINE_RANK_COUNT":             "OnlineRankCount",
	"ONLINE_RANK_V2":                "OnlineRankV2",
	"ONLINE_RANK_TOP3":              "OnlineRankTop3",
	"LIKE_INFO_V3_CLICK":            "LikeInfoV3Click",
	"INTERACT_WORD":                 "InteractWord",
	"STOP_LIVE_ROOM_LIST":           "StopLiveRoomList",
	"LIKE_INFO_V3_UPDATE":           "LikeInfoV3Update",
	"HOT_RANK_CHANGED":              "HotRankChanged",
	"NOTICE_MSG":                    "NoticeMsg",
	"ROOM_REAL_TIME_MESSAGE_UPDATE": "RoomRealTimeMessageUpdate",
	"WIDGET_BANNER":                 "WidgetBanner",
	"HOT_RANK_CHANGED_V2":           "HotRankChangedV2",
	"GUARD_HONOR_THOUSAND":          "GuardHonorThousand",
	"LIVE":                          "Live",
	"ROOM_CHANGE":                   "RoomChange",
	"ROOM_BLOCK_MSG":                "RoomBlockMsg",
	"FULL_SCREEN_SPECIAL_EFFECT":    "FullScreenSpecialEffect",
	"COMMON_NOTICE_DANMAKU":         "CommonNoticeDanmaku",
	"TRADING_SCORE":                 "TradingScore",
	"PREPARING":                     "Preparing",
	"GUARD_BUY":                     "GuardBuy",
	"GIFT_STAR_PROCESS":             "GiftStarProcess",
	"ROOM_SKIN_MSG":                 "RoomSkinMsg",
	"ENTRY_EFFECT":                  "EntryEffect",
}

type MsgEvent struct {
	Cmd                       string
	RoomId                    int
	DanMuMsg                  *DanMuMsg
	SuperChatMessage          *SuperChatMessage
	WatchedChange             *WatchedChange
	SendGift                  *SendGift
	OnlineRankCount           *OnlineRankCount
	OnlineRankV2              *OnlineRankV2
	OnlineRankTop3            *OnlineRankTop3
	LikeInfoV3Click           *LikeInfoV3Click
	InteractWord              *InteractWord
	StopLiveRoomList          *StopLiveRoomList
	LikeInfoV3Update          *LikeInfoV3Update
	HotRankChange             *HotRankChange
	NoticeMsg                 *NoticeMsg
	RoomRealTimeMessageUpdate *RoomRealTimeMessageUpdate
	WidgetBanner              *WidgetBanner
	HotRankChangedV2          *HotRankChangedV2
	GuardHonorThousand        *GuardHonorThousand
	Live                      *Live
	RoomChange                *RoomChange
	RoomBlockMsg              *RoomBlockMsg
	FullScreenSpecialEffect   *FullScreenSpecialEffect
	CommonNoticeDanmaku       *CommonNoticeDanmaku
	TradingScore              *TradingScore
	Preparing                 *Preparing
	GuardBuy                  *GuardBuy
	GiftStarProcess           *GiftStarProcess
	RoomSkinMsg               *RoomSkinMsg
	EntryEffect               *EntryEffect
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
			UID        int    `json:"uid,omitempty"`
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
		UID              int       `json:"uid"`
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
		UID           int       `json:"uid"`
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

type SendGift struct {
	Cmd  string `json:"cmd"`
	Data struct {
		UID               int         `json:"uid"`
		Name              string      `json:"uname"`
		NameColor         string      `json:"name_color"`
		Action            string      `json:"action"`
		BatchComboID      string      `json:"batch_combo_id"`
		BatchComboSend    interface{} `json:"batch_combo_send"`
		BeatID            string      `json:"beatId"`
		BizSource         string      `json:"biz_source"`
		BlindGift         interface{} `json:"blind_gift"`
		BroadcastID       int         `json:"broadcast_id"`
		CoinType          string      `json:"coin_type"`
		ComboResourcesID  int         `json:"combo_resources_id"`
		ComboSend         interface{} `json:"combo_send"`
		ComboStayTime     int         `json:"combo_stay_time"`
		ComboTotalCoin    int         `json:"combo_total_coin"`
		CritProb          int         `json:"crit_prob"`
		Demarcation       int         `json:"demarcation"`
		DiscountPrice     int         `json:"discount_price"`
		Dmscore           int         `json:"dmscore"`
		Draw              int         `json:"draw"`
		Effect            int         `json:"effect"`
		EffectBlock       int         `json:"effect_block"`
		Face              string      `json:"face"`
		FaceEffectID      int         `json:"face_effect_id"`
		FaceEffectType    int         `json:"face_effect_type"`
		FloatScResourceID int         `json:"float_sc_resource_id"`
		GiftID            int         `json:"giftId"`
		GiftName          string      `json:"giftName"`
		GiftType          int         `json:"giftType"`
		Gold              int         `json:"gold"`
		GuardLevel        int         `json:"guard_level"`
		IsFirst           bool        `json:"is_first"`
		IsNaming          bool        `json:"is_naming"`
		IsSpecialBatch    int         `json:"is_special_batch"`
		Magnification     float64     `json:"magnification"`
		Num               int         `json:"num"`
		OriginalGiftName  string      `json:"original_gift_name"`
		Price             int         `json:"price"`
		Rcost             int         `json:"rcost"`
		Remain            int         `json:"remain"`
		Rnd               string      `json:"rnd"`
		SendMaster        interface{} `json:"send_master"`
		Silver            int         `json:"silver"`
		Super             int         `json:"super"`
		SuperBatchGiftNum int         `json:"super_batch_gift_num"`
		SuperGiftNum      int         `json:"super_gift_num"`
		SvgaBlock         int         `json:"svga_block"`
		Switch            bool        `json:"switch"`
		TagImage          string      `json:"tag_image"`
		Tid               string      `json:"tid"`
		Timestamp         int         `json:"timestamp"`
		TopList           interface{} `json:"top_list"`
		TotalCoin         int         `json:"total_coin"`
		FansMedal         FansMedal   `json:"medal_info"`
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
	ShieldUID  int    `json:"shield_uid"`
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
		BubbleShowTime int `json:"bubble_show_time"`
		Num            int `json:"num"`
		ScoreID        int `json:"score_id"`
		UID            int `json:"uid"`
		UpdateTime     int `json:"update_time"`
		UpdateType     int `json:"update_type"`
	} `json:"data"`
}

type Preparing struct {
	Cmd    string `json:"cmd"`
	RoomId string `json:"roomid"`
}

type GuardBuy struct {
	Cmd  string `json:"cmd"`
	Data struct {
		UID        int    `json:"uid"`
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
		UID                  int           `json:"uid"`
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
