package Handler

const (
	CmdDanmuMsg                  = "DANMU_MSG"
	CmdSuperChatMessage          = "SUPER_CHAT_MESSAGE"
	CmdSuperChatMessageJpn       = "SUPER_CHAT_MESSAGE_JPN"
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
)

type User struct {
	Uid    int64
	Name   string
	RoomId int64
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
		Sender                   User
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
		Magnification     int         `json:"magnification"`
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
	Cmd  string `json:"cmd"`
	Data struct {
		Ts                    int       `json:"ts"`
		ID                    string    `json:"id"`
		UID                   string    `json:"uid"`
		Price                 int       `json:"price"`
		Rate                  int       `json:"rate"`
		Time                  int       `json:"time"`
		StartTime             int       `json:"start_time"`
		EndTime               int       `json:"end_time"`
		Token                 string    `json:"token"`
		Message               string    `json:"message"`
		MessageJpn            string    `json:"message_jpn"`
		IsRanked              int       `json:"is_ranked"`
		BackgroundImage       string    `json:"background_image"`
		BackgroundColor       string    `json:"background_color"`
		BackgroundIcon        string    `json:"background_icon"`
		BackgroundPriceColor  string    `json:"background_price_color"`
		BackgroundBottomColor string    `json:"background_bottom_color"`
		FansMedal             FansMedal `json:"medal_info"`
		UserInfo              struct {
			Uname      string `json:"uname"`
			Face       string `json:"face"`
			FaceFrame  string `json:"face_frame"`
			GuardLevel int    `json:"guard_level"`
			UserLevel  int    `json:"user_level"`
			LevelColor string `json:"level_color"`
			IsVip      int    `json:"is_vip"`
			IsSvip     int    `json:"is_svip"`
			IsMainVip  int    `json:"is_main_vip"`
			Title      string `json:"title"`
			Manager    int    `json:"manager"`
		} `json:"user_info"`
		Gift struct {
			Num      int    `json:"num"`
			GiftID   int    `json:"gift_id"`
			GiftName string `json:"gift_name"`
		} `json:"gift"`
	} `json:"data"`
	Roomid string `json:"roomid"`
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
	} `json:"side"`
	Roomid     int    `json:"roomid"`
	RealRoomid int    `json:"real_roomid"`
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
