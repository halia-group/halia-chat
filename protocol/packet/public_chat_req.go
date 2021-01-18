package packet

// 聊天请求

import (
	"encoding/binary"
	"fmt"
	"halia-chat/protocol"
	"io"
)

type PublicChatReq struct {
	protocol.BasePacket
	MsgType uint8  // 消息类型
	Message string // 消息内容
}

func (p PublicChatReq) String() string {
	return fmt.Sprintf("PublicChatReq{MsgType=%d,Message=%s}", p.MsgType, p.Message)
}

func NewChatReq(msgType uint8, message string) *PublicChatReq {
	return &PublicChatReq{
		BasePacket: protocol.BasePacket{
			MagicNumber: protocol.MagicNumber,
			Opcode:      protocol.OpPublicChatReq,
		},
		MsgType: msgType,
		Message: message,
	}
}

func (p *PublicChatReq) Opcode() uint16 {
	return protocol.OpPublicChatReq
}

func (p *PublicChatReq) Write(w io.Writer) (err error) {
	if err = binary.Write(w, binary.BigEndian, &p.MsgType); err != nil {
		return
	}
	if err = p.WriteShortString(w, p.Message); err != nil {
		return
	}
	return
}

func (p *PublicChatReq) Read(r io.Reader) (err error) {
	if err = binary.Read(r, binary.BigEndian, &p.MsgType); err != nil {
		return
	}
	if p.Message, err = p.ReadShortString(r); err != nil {
		return
	}
	return
}
