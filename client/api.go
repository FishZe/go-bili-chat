package client

import (
	"io"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"time"

	"github.com/go-ping/ping"
	log "github.com/sirupsen/logrus"
)

const (
	UserAgent      = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36"
	BiliLiveApiUrl = "api.live.bilibili.com"
	MainWsUrl      = "broadcastlv.chat.bilibili.com"
)

const DefaultPriority = 1 << 0
const DelayPriority = 1 << 1
const NoCDNPriority = 1 << 2

var PriorityMode = DefaultPriority
var shortRoomId2Long = make(map[int]int)

func ChangeSequenceMode(mode int) {
	log.Debug("change sequence mode: ", mode)
	PriorityMode = mode
}

func getReq(data url.Values, getUrl string) ([]byte, error) {
	u, err := url.ParseRequestURI(getUrl)
	if err != nil {
		return nil, err
	}
	u.RawQuery = data.Encode()
	client := http.Client{}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header = Header
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	if resp.StatusCode != 200 {
		return nil, RespCodeNotError
	}
	s, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func getLiveRoomAuth(roomId int) (ApiLiveAuth, error) {
	getUrl := url.URL{Scheme: "https", Host: BiliLiveApiUrl, Path: "/xlive/web-room/v1/index/getDanmuInfo"}
	data := url.Values{}
	data.Set("id", strconv.Itoa(roomId))
	data.Set("type", "0")
	log.Debug("try to get live room info: ", roomId)
	s, err := getReq(data, getUrl.String())
	if err != nil {
		return ApiLiveAuth{}, err
	}
	var j ApiLiveAuth
	err = JsonCoder.Unmarshal(s, &j)
	if err != nil {
		return ApiLiveAuth{}, err
	}
	return liveRoomSort(j), nil
}

func liveRoomSort(j ApiLiveAuth) ApiLiveAuth {
	switch PriorityMode {
	case DefaultPriority:
		return j
	case DelayPriority:
		pings := make(map[string]float64, 10)
		for _, v := range j.Data.HostList {
			pings[v.Host] = getPing(v.Host)
		}
		sort.Slice(j.Data.HostList, func(i, k int) bool {
			return pings[j.Data.HostList[i].Host] < pings[j.Data.HostList[k].Host]
		})
		return j
	case NoCDNPriority:
		sort.Slice(j.Data.HostList, func(i, k int) bool {
			return j.Data.HostList[i].Host == MainWsUrl
		})
		return j
	default:
		return j
	}
}

func getPing(pingUrl string) float64 {
	p, err := ping.NewPinger(pingUrl)
	p.SetPrivileged(true)
	if err != nil {
		return float64(INF)
	}
	p.Count = 3
	p.Interval = 100 * time.Millisecond
	err = p.Run()
	if err != nil {
		return float64(INF)
	}
	stats := p.Statistics()
	return stats.AvgRtt.Seconds()
}

func GetRealRoomId(roomId int) (int, error) {
	if _, ok := shortRoomId2Long[roomId]; ok {
		return shortRoomId2Long[roomId], nil
	}
	getUrl := url.URL{Scheme: "https", Host: BiliLiveApiUrl, Path: "/xlive/web-room/v1/index/getRoomPlayInfo"}
	data := url.Values{}
	data.Set("room_id", strconv.Itoa(roomId))
	log.Debug("try to get real room id: ", roomId)
	s, err := getReq(data, getUrl.String())
	if err != nil {
		return 0, err
	}
	var j ApiLiveRoomId
	err = JsonCoder.Unmarshal(s, &j)
	if err != nil {
		return 0, err
	}
	shortRoomId2Long[roomId] = j.Data.RoomID
	return j.Data.RoomID, nil
}
