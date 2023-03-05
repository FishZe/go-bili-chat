package client

import (
	"io"
	"net/http"
	"net/url"
	"strconv"
)

const (
	UserAgent      = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36"
	BiliLiveApiUrl = "api.live.bilibili.com"
	MainWsUrl      = "broadcastlv.chat.bilibili.com"
)

var cookies = ""

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
		return nil, RespCodeNotError
	}
	s, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func GetLiveRoomAuth(roomId int) (ApiLiveAuth, error) {
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
	return jBA, nil
}
