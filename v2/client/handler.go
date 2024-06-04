package client

import (
	"bytes"
	"compress/zlib"
	"github.com/FishZe/go-bili-chat/v2/events"
	"github.com/andybalholm/brotli"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	"io"
)

type MsgHandler struct {
	RoomId  int
	CmdChan chan *events.BLiveEvent
}

var HeartBeatReplyCmd = "HEARTBEAT_REPLY"

func (msgHandler *MsgHandler) CmdHandler(wsHeader *WsHeader, msg []byte) {
	body := msg[wsHeader.HeaderLen:wsHeader.PackageLen]
	result := gjson.GetBytes(body, "cmd")
	if !result.Exists() || result.Str == "" {
		return
	}
	log.Debug("handle cmd: ", result.Str)
	msgHandler.CmdChan <- &events.BLiveEvent{
		Cmd:        result.Str,
		RoomId:     msgHandler.RoomId,
		RawMessage: body,
	}
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
	wsHeader := WsHeaderDecoder(msg)
	switch wsHeader.OpCode {
	case OpHeartBeatReply:
		wsHeartBeatReply := WsHeartBeatReply{}
		wsHeartBeatReply.SetPackage(wsHeader, msg)
		msgHandler.CmdChan <- &events.BLiveEvent{
			Cmd:        events.CmdHeartBeatReply,
			RoomId:     msgHandler.RoomId,
			RawMessage: wsHeartBeatReply.Msg,
		}
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
