package client

import (
	"bytes"
	"compress/zlib"
	"encoding/json"
	"github.com/andybalholm/brotli"
	"io"
	"log"
)

type MsgHandler struct {
	RoomId  int
	CmdChan chan map[string]interface{}
}

func (msgHandler *MsgHandler) CmdHandler(wsHeader *WsHeader, msg []byte) {
	cmdJson := map[string]interface{}{}
	err := json.Unmarshal(msg[wsHeader.HeaderLen:wsHeader.PackageLen], &cmdJson)
	if err != nil {
		log.Printf("Unmarshal cmd json failed: %v", err)
		return
	}
	msgHandler.CmdChan <- cmdJson
	//fmt.Println(string(msg[wsHeader.HeaderLen:wsHeader.PackageLen]))
	//fmt.Println(cmdJson["cmd"])
}

func (msgHandler *MsgHandler) CmdBrotliProtoDecoder(wsHeader *WsHeader, msg []byte) []byte {
	reader := brotli.NewReader(bytes.NewReader(msg[wsHeader.HeaderLen:wsHeader.PackageLen]))
	resp, err := io.ReadAll(reader)
	if err != nil {
		log.Println("Brotli decode failed: ", err)
		return []byte{}
	}
	return resp
}

func (msgHandler *MsgHandler) CmdZlibProtoDecoder(wsHeader *WsHeader, msg []byte) []byte {
	var resp bytes.Buffer
	w := zlib.NewWriter(&resp)
	_, err := w.Write(msg[wsHeader.HeaderLen:wsHeader.PackageLen])
	if err != nil {
		log.Println("Zlib decode failed: ", err)
		return []byte{}
	}
	err = w.Close()
	if err != nil {
		log.Println("Zlib decode failed: ", err)
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
		//fmt.Printf("Heartbeat: %v\n", wsHeartBeatReply.Hot)
	case OpCmd:
		msgBody := msg
		cmdHeader := wsHeader
		switch wsHeader.ProtoVer {
		case CmdZlibProto:
			msgBody = msgHandler.CmdZlibProtoDecoder(&wsHeader, msg)
			cmdHeader = WsHeaderDecoder(msgBody)
			fallthrough
		case CmdBrotliProto:
			msgBody = msgHandler.CmdBrotliProtoDecoder(&wsHeader, msg)
			cmdHeader = WsHeaderDecoder(msgBody)
			fallthrough
		default:
			for {
				msgHandler.CmdHandler(&cmdHeader, msgBody[:int(cmdHeader.PackageLen)])
				msgBody = msgBody[cmdHeader.PackageLen:]
				if len(msgBody) == 0 {
					break
				}
				cmdHeader = WsHeaderDecoder(msgBody)
			}
		}
	case OpAuthReply:
		wsAuthReplyMessage := WsAuthReplyMessage{}
		wsAuthReplyMessage.SetPackage(wsHeader, msg)
		//fmt.Printf("Auth: %v\b", wsAuthReplyMessage.Body.Code)
	}
}
