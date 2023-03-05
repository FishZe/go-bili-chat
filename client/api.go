package client

import (
	"github.com/go-ping/ping"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"time"
)

const (
	UserAgent      = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36"
	BiliLiveApiUrl = "api.live.bilibili.com"
	MainWsUrl      = "broadcastlv.chat.bilibili.com"
)

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
	req.Header = http.Header{
		"User-Agent": {UserAgent},
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	if resp.StatusCode != 200 {
		return nil, ErrespCodeNot
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
	s, err := getReq(data, getUrl.String())
	if err != nil {
		return ApiLiveAuth{}, err
	}
	var jBA ApiLiveAuth
	err = JsonCoder.Unmarshal(s, &jBA)
	if err != nil {
		return ApiLiveAuth{}, err
	}
	return liveRoomSort(jBA), nil
}

func liveRoomSort(j ApiLiveAuth) ApiLiveAuth {
	if SequenceMode == DefaultSequence {
		return j
	} else if SequenceMode == DelaySequence {
		pings := make(map[string]float64, 10)
		for _, v := range j.Data.HostList {
			pings[v.Host] = getPing(v.Host)
		}
		sort.Slice(j.Data.HostList, func(i, k int) bool {
			return pings[j.Data.HostList[i].Host] < pings[j.Data.HostList[k].Host]
		})
		return j
	}
	return j
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
	getUrl := url.URL{Scheme: "https", Host: BiliLiveApiUrl, Path: "/xlive/web-room/v1/index/getRoomPlayInfo"}
	data := url.Values{}
	data.Set("room_id", strconv.Itoa(roomId))
	s, err := getReq(data, getUrl.String())
	if err != nil {
		return 0, err
	}
	var jBA ApiLiveRoomId
	err = JsonCoder.Unmarshal(s, &jBA)
	if err != nil {
		return 0, err
	}
	return jBA.Data.RoomID, nil
}
