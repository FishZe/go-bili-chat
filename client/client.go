package client

import (
	"crypto/tls"
	"github.com/lxzan/gws"
	"net/http"
	"net/url"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

type jsonCoder interface {
	Unmarshal(data []byte, v interface{}) error
	Marshal(v interface{}) ([]byte, error)
}

var JsonCoder jsonCoder
var UID = int64(1)
var Buvid = ""
var Header http.Header

type Client struct {
	// 一个直播间的状态: 关闭 / 已连接 / 连接中
	RoomInfo WsAuthBody
	connect  *gws.Conn
	handler  MsgHandler
	done     chan struct{}
}

func (c *Client) OnOpen(socket *gws.Conn) {
	log.Info("connected to blive success: ", c.RoomInfo.Roomid)
	wsAuthMsg := WsAuthMessage{Body: c.RoomInfo}
	// 连接成功
	if err := c.sendAuthMsg(wsAuthMsg); err != nil {
		log.Warn("send auth msg to websocket error: ", err)
		return
	}
	go c.heartBeat()
}

func (c *Client) OnClose(socket *gws.Conn, err error) {
	log.Info("disconnected from blive: ", c.RoomInfo.Roomid)
	select {
	case <-c.done:
		return
	default:
		c.cancel()
		time.Sleep(2 * time.Second)
		c.connectLoop()
	}
}

func (c *Client) OnPing(socket *gws.Conn, payload []byte) {
}

func (c *Client) OnPong(socket *gws.Conn, payload []byte) {
}

func (c *Client) OnMessage(socket *gws.Conn, message *gws.Message) {
	defer message.Close()
	c.handler.MsgHandler(message.Bytes())
}

func init() {
	Header = make(http.Header)
	Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/116.0.0.0 Safari/537.36 Edg/116.0.1938.62")
	Header.Set("Origin", "https://live.bilibili.com")
}

func (c *Client) biliChatConnect(url string) error {
	var err error
	c.connect, _, err = gws.NewClient(c, &gws.ClientOption{
		Addr:      url,
		TlsConfig: &tls.Config{InsecureSkipVerify: true},
		PermessageDeflate: gws.PermessageDeflate{
			Enabled:               true,
			ServerContextTakeover: true,
			ClientContextTakeover: true,
		},
		RequestHeader: Header,
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) sendAuthMsg(wsAuthMsg WsAuthMessage) error {
	wsPackage := wsAuthMsg.GetPackage()
	err := c.connect.WriteMessage(gws.OpcodeBinary, wsPackage)
	if err != nil {
		return err
	}
	log.Debug("send auth msg to blive success")
	return nil
}

func (c *Client) heartBeat() {
	t := time.NewTicker(time.Second * 30)
	for {
		select {
		case <-c.done:
			log.Debug("heartBeat exit...")
			//_ = c.connect.Close()
			return
		case <-t.C:
			heartBeatPackage := WsHeartBeatMessage{Body: []byte{}}
			log.Debug("send heart beat to blive...")
			err := c.connect.WriteMessage(gws.OpcodeBinary, heartBeatPackage.GetPackage())
			if err != nil {
				log.Error("failed to send heart beat to blive: ", err)
				return
			}
		}
	}
}

func (c *Client) sendConnect() error {
	switch PriorityMode {
	// No CDN Mode
	case NoCDNPriority:
		{
			u := url.URL{Scheme: "wss", Host: MainWsUrl, Path: "/sub"}
			log.Debug("connect to blive websocket: ", u.String())
			if err := c.biliChatConnect(u.String()); err != nil {
				return err
			}
		}
	// others MODE or No CDN Mode failed
	default:
		{
			apiLiveAuth, err := getLiveRoomAuth(c.RoomInfo.Roomid)
			if err != nil {
				return err
			} else if apiLiveAuth.Code != 0 {
				log.Warnf("get live room info error: %v", apiLiveAuth.Message)
				return RespCodeNotError
			}
			if c.RoomInfo.Key == "" {
				c.RoomInfo.Key = apiLiveAuth.Data.Token
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
		}
	}
	return nil
}

func (c *Client) connectLoop() {
	c.done = make(chan struct{})
	for {
		err := c.sendConnect()
		if err != nil {
			log.Warn("connect to blive error: ", err)
			time.Sleep(5 * time.Second)
		} else {
			// Start websocket message loop
			go c.connect.ReadLoop()
			break
		}
	}
}

func (c *Client) cancel() {
	close(c.done)
}

func (c *Client) Close() {
	c.cancel()
	c.connect.WriteClose(1000, nil)
}

func (c *Client) BiliChat(CmdChan chan map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Warnf("start blive panic: %v", err)
		}
	}()
	c.handler = MsgHandler{RoomId: c.RoomInfo.Roomid, CmdChan: CmdChan}
	// Try to start websocket
	c.connectLoop()
	log.Debug("start blive success: ", c.RoomInfo.Roomid)
}
