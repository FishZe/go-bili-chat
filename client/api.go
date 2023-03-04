package client

import (
	"github.com/bytedance/sonic"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

const (
	UserAgent      = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36"
	BiliUrl        = "www.bilibili.com"
	BiliLiveApiUrl = "api.live.bilibili.com"
)

var cookies = ""

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

func GetLiveRoomAuth(roomId int) (ApiLiveAuth, error) {
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
	err = sonic.Unmarshal(s, &jBA)
	if err != nil {
		return ApiLiveAuth{}, err
	}
	return jBA, nil
}
