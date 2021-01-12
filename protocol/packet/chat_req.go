package packet

// 聊天请求

import (
	"encoding/binary"
	"fmt"
	"halia-chat/protocol"
	"io"
)

type ChatReq struct {
	protocol.BasePacket
	MsgType uint8  // 消息类型
	Message string // 消息内容
}

func (p ChatReq) String() string {
	return fmt.Sprintf("ChatReq{MsgType=%d,Message=%s}", p.MsgType, p.Message)
}

func NewChatReq(msgType uint8, message string) *ChatReq {
	return &ChatReq{
		BasePacket: protocol.BasePacket{
			MagicNumber: protocol.MagicNumber,
			Opcode:      protocol.OpChatReq,
		},
		MsgType: msgType,
		Message: message,
	}
}

func (p *ChatReq) Opcode() uint16 {
	return protocol.OpChatReq
}

func (p *ChatReq) Write(w io.Writer) (err error) {
	if err = binary.Write(w, binary.BigEndian, &p.MsgType); err != nil {
		return
	}
	if err = p.WriteShortString(w, p.Message); err != nil {
		return
	}
	return
}

func (p *ChatReq) Read(r io.Reader) (err error) {
	if err = binary.Read(r, binary.BigEndian, &p.MsgType); err != nil {
		return
	}
	if p.Message, err = p.ReadShortString(r); err != nil {
		return
	}
	return
}
