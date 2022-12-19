package client

import (
	"errors"
	"log"
	"net/url"
	"strconv"
	"time"
)
import "github.com/gorilla/websocket"

type Client struct {
	RoomId  int
	connect *websocket.Conn
	revMsg  chan []byte
}

func (c *Client) biliChatConnect(url string) error {
	//fmt.Println(url)
	err := errors.New("")
	c.connect, _, err = websocket.DefaultDialer.Dial(url, nil)
	if nil != err {
		log.Println("")
		return err
	}
	return nil
}

func (c *Client) sendAuthMsg(wsAuthMsg WsAuthMessage) error {
	wsPackage := wsAuthMsg.GetPackage()
	err := c.connect.WriteMessage(websocket.BinaryMessage, wsPackage)
	if err != nil {
		log.Printf("Send auth msg failed: %v\n", err)
		return err
	}
	return nil
}

func (c *Client) receiveWsMsg() {
	for {
		_, message, err := c.connect.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		c.revMsg <- message
	}
}

func (c *Client) heartBeat() {
	for {
		heartBeatPackage := WsHeartBeatMessage{Body: []byte{}}
		err := c.connect.WriteMessage(websocket.TextMessage, heartBeatPackage.GetPackage())
		if err != nil {
			log.Printf("Send heart beat failed %v", err)
		}
		time.Sleep(30 * time.Second)
	}
}

func (c *Client) revHandler(handler MsgHandler) {
	for {
		select {
		case msg, ok := <-c.revMsg:
			if ok {
				handler.MsgHandler(msg)
			}
		default:
			time.Sleep(10 * time.Microsecond)
			continue
		}

	}
}

func (c *Client) BiliChat() (MsgHandler, error) {
	wsAuthMsg := WsAuthMessage{Body: WsAuthBody{UID: 0, Roomid: c.RoomId, Protover: 3, Platform: "web", Type: 2}}
	apiLiveAuth, err := GetLiveRoomAuth(c.RoomId)
	if err != nil {
		log.Printf("Get live room auth failed, not use key")
		return MsgHandler{}, err
	}
	wsAuthMsg.Body.Key = apiLiveAuth.Data.Token
	for nowSum, i := range apiLiveAuth.Data.HostList {
		u := url.URL{Scheme: "wss", Host: i.Host + ":" + strconv.Itoa(i.WssPort), Path: "/sub"}
		err = c.biliChatConnect(u.String())
		if err != nil {
			if nowSum == 2 {
				log.Println("Connect to all server failed")
				return MsgHandler{}, err
			}
			log.Println("Connect to bili chat failed, i'll try again")
		} else {
			break
		}
	}
	c.revMsg = make(chan []byte, 1)
	err = c.sendAuthMsg(wsAuthMsg)
	if err != nil {
		log.Println("Send auth msg failed")
		return MsgHandler{}, err
	}
	cmdChan := make(chan map[string]interface{}, 1)
	handler := MsgHandler{RoomId: c.RoomId, CmdChan: cmdChan}
	go c.revHandler(handler)
	go c.receiveWsMsg()
	go c.heartBeat()
	return handler, nil
}
