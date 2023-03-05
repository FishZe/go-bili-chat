package client

import (
	"bytes"
	"compress/zlib"
	"github.com/andybalholm/brotli"
	log "github.com/sirupsen/logrus"
	"io"
)

type MsgHandler struct {
	RoomId  int
	CmdChan chan map[string]interface{}
}

func getCmd(msg []byte) string {
	// 反正只需要一个cmd进行分发, 就不需要json解析整个数据了
	var layer = 0
	for i, v := range msg {
		if v == '{' || v == '[' {
			layer++
		} else if v == '}' || v == ']' {
			layer--
		} else if layer == 1 && v == '"' {
			if i+7 < len(msg) && msg[i+1] == 'c' && msg[i+2] == 'm' && msg[i+3] == 'd' && msg[i+4] == '"' {
				var from = i + 7
				var to int
				for to = from + 1; to < len(msg); to++ {
					if msg[to] == '"' {
						break
					}
				}
				return string(msg[from:to])
			}
		}
	}
	return ""
}

func (msgHandler *MsgHandler) CmdHandler(wsHeader *WsHeader, msg []byte) {
	cmd := getCmd(msg[wsHeader.HeaderLen:wsHeader.PackageLen])
	if cmd == "" {
		return
	}
	log.Debug("handle cmd: ", cmd)
	rev := make(map[string]interface{})
	rev["cmd"] = cmd
	rev["msg"] = string(msg[wsHeader.HeaderLen:wsHeader.PackageLen])
	rev["RoomId"] = msgHandler.RoomId
	msgHandler.CmdChan <- rev
}

func (msgHandler *MsgHandler) CmdBrotliProtoDecoder(wsHeader *WsHeader, msg []byte) []byte {
	reader := brotli.NewReader(bytes.NewReader(msg[wsHeader.HeaderLen:wsHeader.PackageLen]))
	resp, err := io.ReadAll(reader)
	if err != nil {
		return []byte{}
	}
	return resp
}

func (msgHandler *MsgHandler) CmdZlibProtoDecoder(wsHeader *WsHeader, msg []byte) []byte {
	// 这段还没测试过, 没遇到过Zlib的压缩方式
	var resp bytes.Buffer
	w := zlib.NewWriter(&resp)
	_, err := w.Write(msg[wsHeader.HeaderLen:wsHeader.PackageLen])
	if err != nil {
		return []byte{}
	}
	err = w.Close()
	if err != nil {
		return nil
	}
	return resp.Bytes()
}

func (msgHandler *MsgHandler) MsgHandler(msg []byte) {
	// TODO: 我好像还没写心跳包和认证包回复的处理
	wsHeader := WsHeaderDecoder(msg)
	switch wsHeader.OpCode {
	case OpHeartBeatReply:
		// wsHeartBeatReply := WsHeartBeatReply{}
		// wsHeartBeatReply.SetPackage(wsHeader, msg)
	case OpCmd:
		msgBody := msg
		cmdHeader := wsHeader
		switch wsHeader.ProtoVer {
		case CmdZlibProto:
			msgBody = msgHandler.CmdZlibProtoDecoder(&wsHeader, msg)
			cmdHeader = WsHeaderDecoder(msgBody)
			for {
				msgHandler.CmdHandler(&cmdHeader, msgBody[:int(cmdHeader.PackageLen)])
				msgBody = msgBody[cmdHeader.PackageLen:]
				if len(msgBody) == 0 {
					break
				}
				cmdHeader = WsHeaderDecoder(msgBody)
			}
		case CmdBrotliProto:
			msgBody = msgHandler.CmdBrotliProtoDecoder(&wsHeader, msg)
			cmdHeader = WsHeaderDecoder(msgBody)
			for {
				msgHandler.CmdHandler(&cmdHeader, msgBody[:int(cmdHeader.PackageLen)])
				msgBody = msgBody[cmdHeader.PackageLen:]
				if len(msgBody) == 0 {
					break
				}
				cmdHeader = WsHeaderDecoder(msgBody)
			}
		default:
			msgHandler.CmdHandler(&cmdHeader, msgBody[:int(cmdHeader.PackageLen)])
		}
	case OpAuthReply:
		// wsAuthReplyMessage := WsAuthReplyMessage{}
		// wsAuthReplyMessage.SetPackage(wsHeader, msg)
	case OpError:
		log.Warnf("recv error msg: %s", string(msg[wsHeader.HeaderLen:wsHeader.PackageLen]))
	}
}
