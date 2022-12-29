package client

import (
	"encoding/json"
	"errors"
	"io"
	"log"
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
		log.Printf("Error orrured when parsing the url: %v", err)
		return nil, "", err
	}
	u.RawQuery = data.Encode()
	client := http.Client{}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		log.Printf("Error orrured when creating the request: %v", err)
		return nil, "", err
	}
	req.Header = http.Header{
		"User-Agent": {UserAgent},
		"cookie":     {cookies},
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error orrured when sending the request: %v", err)
		return nil, "", err
	}
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			log.Printf("Error orrured when closing the session: %v", err)
		}
	}(resp.Body)
	if resp.StatusCode != 200 {
		log.Printf("Status code is not 200: %v", resp.StatusCode)
		return nil, "", errors.New("status code error: " + strconv.Itoa(resp.StatusCode) + " " + resp.Status)
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
		log.Printf("Error orrured when reading the response: %v", err)
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
	err = json.Unmarshal(s, &jBA)
	if err != nil {
		log.Println(err)
		return ApiLiveAuth{}, err
	}
	return jBA, nil
}
