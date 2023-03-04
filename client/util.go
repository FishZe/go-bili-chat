package client

import (
	"bytes"
	"encoding/json"
)

func uint32ToByte4(num uint32) []byte {
	if num > 4294967295 {
		return []byte{}
	}
	var bytesNum = make([]byte, 4)
	for i := 3; i >= 0; i-- {
		bytesNum[i] = byte(num & 255)
		num >>= 8
	}
	return bytesNum
}

func uint16ToByte2(num uint16) []byte {
	if num > 65535 {
		return []byte{}
	}
	var bytesNum = make([]byte, 2)
	for i := 1; i >= 0; i-- {
		bytesNum[i] = byte(num & 255)
		num >>= 8
	}
	return bytesNum
}

func byte2ToUint16(bytesNum []byte) uint16 {
	var num uint16
	for i := 0; i < 2; i++ {
		num <<= 8
		num += uint16(bytesNum[i])
	}
	return num
}

func byte4ToUint32(bytesNum []byte) uint32 {
	var num uint32
	for i := 0; i < 4; i++ {
		num <<= 8
		num += uint32(bytesNum[i])
	}
	return num
}

func WsHeaderDecoder(headerBytes []byte) WsHeader {
	if len(headerBytes) < 16 {
		return WsHeader{OpCode: OpError}
	}
	var wsHeader WsHeader
	wsHeader.PackageLen = byte4ToUint32(headerBytes[0:4])
	wsHeader.HeaderLen = byte2ToUint16(headerBytes[4:6])
	wsHeader.ProtoVer = byte2ToUint16(headerBytes[6:8])
	wsHeader.OpCode = byte4ToUint32(headerBytes[8:12])
	wsHeader.Sequence = byte4ToUint32(headerBytes[12:16])
	return wsHeader
}

func (wsHeader *WsHeader) HeaderEncoder(bodyLen uint32) []byte {
	var buffer bytes.Buffer
	buffer.Write(uint32ToByte4(bodyLen + 16))
	buffer.Write(uint16ToByte2(16))
	buffer.Write(uint16ToByte2(wsHeader.ProtoVer))
	buffer.Write(uint32ToByte4(wsHeader.OpCode))
	buffer.Write(uint32ToByte4(wsHeader.Sequence))
	return buffer.Bytes()
}

func (wsAuth *WsAuthMessage) GetPackage() []byte {
	wsAuth.WsHeader.OpCode = OpAuth
	wsAuth.WsHeader.Sequence = 1
	wsAuth.WsHeader.ProtoVer = AuthProto
	var buffer bytes.Buffer
	authBody := wsAuth.Body.getAuthBytes()
	buffer.Write(wsAuth.WsHeader.HeaderEncoder(uint32(len(authBody))))
	buffer.Write(authBody)
	return buffer.Bytes()
}

func (wsAuth *WsAuthBody) getAuthBytes() []byte {
	authBody, err := json.Marshal(wsAuth)
	if err != nil {
		return []byte{}
	}
	return authBody
}

func (wsHeartBeat *WsHeartBeatMessage) GetPackage() []byte {
	wsHeartBeat.WsHeader.OpCode = OpHeartBeat
	wsHeartBeat.WsHeader.Sequence = 1
	wsHeartBeat.WsHeader.ProtoVer = HeartBeatProto
	var buffer bytes.Buffer
	buffer.Write(wsHeartBeat.WsHeader.HeaderEncoder(0))
	buffer.Write(wsHeartBeat.Body)
	return buffer.Bytes()
}

func (wsAuthReplyMessage *WsAuthReplyMessage) SetPackage(header WsHeader, msg []byte) {
	wsAuthReplyMessage.WsHeader = header
	authBody := WsAuthReplyBody{}
	if err := json.Unmarshal(msg[header.HeaderLen:], &authBody); err == nil {
		wsAuthReplyMessage.Body = authBody
	}
}

func (wsHeartBeatReply *WsHeartBeatReply) SetPackage(header WsHeader, msg []byte) {
	wsHeartBeatReply.WsHeader = header
	wsHeartBeatReply.Hot = byte4ToUint32(msg[header.HeaderLen : header.HeaderLen+4])
	wsHeartBeatReply.Msg = msg[header.HeaderLen+4:]
}

func (wsCmdMessage *WsCmdMessage) SetPackage(header WsHeader, msg []byte) {
	wsCmdMessage.WsHeader = header
	wsCmdMessage.Body = msg[header.HeaderLen:]
}
