package client

import (
	"context"
	"errors"
	"log"
	"net/url"
	"strconv"
	"time"
)
import "github.com/gorilla/websocket"

type Client struct {
	RoomId    int
	Connected bool
	ctx       context.Context
	cancel    context.CancelFunc
	connect   *websocket.Conn
	revMsg    chan []byte
}

func (c *Client) biliChatConnect(url string) error {
	err := errors.New("")
	c.connect, _, err = websocket.DefaultDialer.Dial(url, nil)
	if nil != err {
		return err
	}
	return nil
}

func (c *Client) sendAuthMsg(wsAuthMsg WsAuthMessage) error {
	wsPackage := wsAuthMsg.GetPackage()
	err := c.connect.WriteMessage(websocket.BinaryMessage, wsPackage)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) receiveWsMsg() {
	for {
		select {
		case <-c.ctx.Done():
			_ = c.connect.Close()
			return
		default:
			if c.connect != nil && c.Connected {
				_, message, err := c.connect.ReadMessage()
				if err != nil {
					log.Println("read:", err)
					c.Connected = false
					c.connectLoop()
				}
				c.revMsg <- message
			}
		}
	}
}

func (c *Client) heartBeat() {
	for {
		select {
		case <-c.ctx.Done():
			_ = c.connect.Close()
			return
		default:
			if c.Connected && c.connect != nil {
				heartBeatPackage := WsHeartBeatMessage{Body: []byte{}}
				_ = c.connect.WriteMessage(websocket.TextMessage, heartBeatPackage.GetPackage())
				time.Sleep(30 * time.Second)
			}
		}
	}
}

func (c *Client) revHandler(handler MsgHandler) {
	for {
		select {
		case <-c.ctx.Done():
			c.revMsg = nil
			return
		case msg, ok := <-c.revMsg:
			if ok {
				handler.MsgHandler(msg)
			}
		default:
		}
	}
}

func (c *Client) sendConnect() error {
	wsAuthMsg := WsAuthMessage{Body: WsAuthBody{UID: 0, Roomid: c.RoomId, Protover: 3, Platform: "web", Type: 2}}
	apiLiveAuth, err := GetLiveRoomAuth(c.RoomId)
	if err != nil {
		return err
	} else if apiLiveAuth.Code != 0 {
		return RespCodeNotError
	}
	wsAuthMsg.Body.Key = apiLiveAuth.Data.Token
	for nowSum, i := range apiLiveAuth.Data.HostList {
		u := url.URL{Scheme: "wss", Host: i.Host + ":" + strconv.Itoa(i.WssPort), Path: "/sub"}
		err = c.biliChatConnect(u.String())
		if err != nil {
			if nowSum == 2 {
				return err
			}
		} else {
			break
		}
		time.Sleep(5 * time.Second)
	}
	err = c.sendAuthMsg(wsAuthMsg)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) connectLoop() {
	for {
		c.Connected = false
		err := c.sendConnect()
		if err != nil {
			time.Sleep(5 * time.Second)
		} else {
			c.Connected = true
			break
		}
	}
}

func (c *Client) Close() {
	c.cancel()
}

func (c *Client) BiliChat(CmdChan chan map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("run time panic: %v", err)
		}
	}()
	c.connectLoop()
	c.revMsg = make(chan []byte, 1000)
	handler := MsgHandler{RoomId: c.RoomId, CmdChan: CmdChan}
	c.ctx, c.cancel = context.WithCancel(context.Background())
	go c.revHandler(handler)
	go c.receiveWsMsg()
	go c.heartBeat()
}
