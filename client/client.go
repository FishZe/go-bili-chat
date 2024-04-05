package client

import (
	"crypto/tls"
	"fmt"
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
	RoomId  int
	connect *gws.Conn
	handler MsgHandler
}

func (c *Client) OnOpen(socket *gws.Conn) {
	wsAuthMsg := WsAuthMessage{Body: WsAuthBody{
		Roomid:   c.RoomId,
		Protover: 3,
		UID:      UID,
		Buvid:    Buvid,
		Type:     2,
		Platform: "web",
		Key:      "hJLzABYiPe0mLR9mxK9wTcM1Q6BvNj-MFhvbcsnxd8E_MJ_TH434rhC1RvlZYhb-_sfgyq5_Jm5if-O4RRNexX9YY8lHgSITQ2qN1e8Qrdqw4XIGnVpOvk97YlZNt7IDRmj7sStUMgEeBQWVvwJ0Se9l6r_MRVsXHs92U9n5cH4ez5lPv4BcUO7PX-h_xt3H",
	}}
	// 连接成功
	log.Debug("connect to blive websocket success", "roomId", c.RoomId)
	if err := c.sendAuthMsg(wsAuthMsg); err != nil {
		log.Warn("send auth msg to websocket error: ", err)
		return
	}
	go c.heartBeat()
}

func (c *Client) OnClose(socket *gws.Conn, err error) {
	fmt.Println("disconnected to ", c.RoomId)
	if err != nil {
		time.Sleep(2 * time.Second)
		c.connectLoop()
	}
}

func (c *Client) OnPing(socket *gws.Conn, payload []byte) {
}

func (c *Client) OnPong(socket *gws.Conn, payload []byte) {
}

func (c *Client) OnMessage(socket *gws.Conn, message *gws.Message) {
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
	})
	if nil != err {
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
		<-t.C
		heartBeatPackage := WsHeartBeatMessage{Body: []byte{}}
		log.Debug("send heart beat to blive...")
		err := c.connect.WriteMessage(gws.OpcodeBinary, heartBeatPackage.GetPackage())
		if err != nil {
			break
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
	for {
		err := c.sendConnect()
		if err != nil {
			log.Warn("connect to blive error: ", err)
			time.Sleep(5 * time.Second)
		} else {
			log.Info("connected to blive success: ", c.RoomId)
			break
		}
	}
}

func (c *Client) Close() {
	c.connect.WriteClose(1000, nil)
}

func (c *Client) BiliChat(CmdChan chan map[string]interface{}) {
	defer func() {
		if err := recover(); err != nil {
			log.Warnf("start blive panic: %v", err)
		}
	}()
	c.handler = MsgHandler{RoomId: c.RoomId, CmdChan: CmdChan}
	// Try to start websocket
	c.connectLoop()
	// Start websocket message loop
	go c.connect.ReadLoop()
	log.Debug("start blive success: ", c.RoomId)
}
