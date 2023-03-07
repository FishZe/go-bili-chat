package client

import (
	"context"
	log "github.com/sirupsen/logrus"
	"net/url"
	"strconv"
	"time"
)
import "github.com/gorilla/websocket"

type jsonCoder interface {
	Unmarshal(data []byte, v interface{}) error
	Marshal(v interface{}) ([]byte, error)
}

var JsonCoder jsonCoder

type Client struct {
	RoomId    int
	Connected bool
	ctx       context.Context
	cancel    context.CancelFunc
	connect   *websocket.Conn
}

func (c *Client) biliChatConnect(url string) error {
	var err error
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
	log.Debug("send auth msg to blive success")
	return nil
}

func (c *Client) receiveWsMsg(handler MsgHandler) {
	for {
		select {
		case <-c.ctx.Done():
			log.Debug("receiveWsMsg exit...")
			_ = c.connect.Close()
			return
		default:
			if c.connect != nil && c.Connected {
				_, message, err := c.connect.ReadMessage()
				if err != nil {
					log.Warnf("read blive websocket msg error: %v", err)
					c.Connected = false
					c.connectLoop()
				}
				handler.MsgHandler(message)
			}
		}
	}
}

func (c *Client) heartBeat() {
	t := time.NewTicker(time.Second * 30)
	for {
		select {
		case <-c.ctx.Done():
			log.Debug("heartBeat exit...")
			//_ = c.connect.Close()
			return
		case <-t.C:
			if c.Connected && c.connect != nil {
				heartBeatPackage := WsHeartBeatMessage{Body: []byte{}}
				log.Debug("send heart beat to blive...")
				_ = c.connect.WriteMessage(websocket.TextMessage, heartBeatPackage.GetPackage())
			}
		}
	}
}

func (c *Client) sendConnect() error {
	wsAuthMsg := WsAuthMessage{Body: WsAuthBody{UID: 0, Roomid: c.RoomId, Protover: 3, Platform: "web", Type: 2}}
	// No CDN Mode
	if PriorityMode == NoCDNPriority {
		u := url.URL{Scheme: "wss", Host: MainWsUrl, Path: "/sub"}
		log.Debug("connect to blive websocket: ", u.String())
		if err := c.biliChatConnect(u.String()); err == nil {
			// 连接成功
			log.Debug("connect to blive websocket success")
			wsAuthMsg.Body.Key = ""
			if err = c.sendAuthMsg(wsAuthMsg); err != nil {
				log.Warn("send auth msg to websocket error: ", err)
				return err
			}
			return nil
		}
	}
	// others MODE or No CDN Mode failed
	apiLiveAuth, err := getLiveRoomAuth(c.RoomId)
	if err != nil {
		return err
	} else if apiLiveAuth.Code != 0 {
		log.Warnf("get live room info error: %v", apiLiveAuth.Message)
		return RespCodeNotError
	}
	for nowSum, i := range apiLiveAuth.Data.HostList {
		u := url.URL{Scheme: "wss", Host: i.Host + ":" + strconv.Itoa(i.WssPort), Path: "/sub"}
		log.Debug("connect to blive websocket: ", u.String())
		err = c.biliChatConnect(u.String())
		if err != nil {
			log.Warnf("connect to blive websocket error for %d time: %v\n", nowSum, err)
			if nowSum == 2 {
				return err
			}
		} else {
			log.Debug("connect to blive websocket success")
			if i.Host != MainWsUrl {
				wsAuthMsg.Body.Key = apiLiveAuth.Data.Token
				if PriorityMode == NoCDNPriority {
					log.Debug("use no cdn mode failed")
				}
			} else if PriorityMode == NoCDNPriority {
				log.Debug("use no cdn mode success")
			}
			break
		}
		time.Sleep(1 * time.Second)
	}
	if err = c.sendAuthMsg(wsAuthMsg); err != nil {
		log.Warn("send auth msg to websocket error: ", err)
		return err
	}
	return nil
}

func (c *Client) connectLoop() {
	for {
		c.Connected = false
		err := c.sendConnect()
		if err != nil {
			log.Warn("connect to blive error: ", err)
			time.Sleep(5 * time.Second)
		} else {
			log.Info("connected to blive success: ", c.RoomId)
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
			log.Warnf("start blive panic: %v", err)
		}
	}()
	c.connectLoop()
	handler := MsgHandler{RoomId: c.RoomId, CmdChan: CmdChan}
	c.ctx, c.cancel = context.WithCancel(context.Background())
	go c.receiveWsMsg(handler)
	go c.heartBeat()
	log.Debug("start blive success: ", c.RoomId)
}
