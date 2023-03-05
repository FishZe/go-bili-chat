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
	BiliUrl        = "www.bilibili.com"
	BiliLiveApiUrl = "api.live.bilibili.com"
)

const DefaultSequence = 1
const DelaySequence = 2

var SequenceMode = DefaultSequence

var cookies = ""

func ChangeSequenceMode(mode int) {
	SequenceMode = mode
}

func getReq(data url.Values, getUrl string, cookies string) ([]byte, string, error) {
	u, err := url.ParseRequestURI(getUrl)
	if err != nil {
		return nil, "", err
	}
	u.RawQuery = data.Encode()
	client := http.Client{}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, "", err
	}
	req.Header = http.Header{
		"User-Agent": {UserAgent},
		"cookie":     {cookies},
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, "", err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	if resp.StatusCode != 200 {
		return nil, "", RespCodeNotError
	}
	if resp.Header.Get("Set-Cookie") != "" {
		cookies = ""
		for _, v := range resp.Header["Set-Cookie"] {
			cookies += v + ";"
		}
	} else {
		cookies = ""
	}
	s, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "", err
	}
	return s, cookies, nil
}

func getCookies() string {
	getUrl := url.URL{Scheme: "https", Host: BiliUrl, Path: "/"}
	data := url.Values{}
	data.Set("spm_id_from", "333.999.0.0")
	_, cookies, err := getReq(data, getUrl.String(), "")
	if err != nil {
		return ""
	}
	return cookies
}

func getLiveRoomAuth(roomId int) (ApiLiveAuth, error) {
	if cookies == "" {
		cookies = getCookies()
	}
	getUrl := url.URL{Scheme: "https", Host: BiliLiveApiUrl, Path: "/xlive/web-room/v1/index/getDanmuInfo"}
	data := url.Values{}
	data.Set("id", strconv.Itoa(roomId))
	s, _, err := getReq(data, getUrl.String(), cookies)
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
