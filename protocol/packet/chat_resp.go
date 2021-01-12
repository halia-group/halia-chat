package packet

// 聊天响应
import (
	"encoding/binary"
	"halia-chat/protocol"
	"io"
)

type ChatResp struct {
	protocol.BasePacket
	Code    uint8
	Message string
}

func NewChatResp(code uint8, message string) *ChatResp {
	return &ChatResp{
		BasePacket: protocol.BasePacket{
			MagicNumber: protocol.MagicNumber,
			Opcode:      protocol.OpChatResp,
		},
		Code:    code,
		Message: message,
	}
}

func (ChatResp) Opcode() uint16 {
	return protocol.OpChatResp
}

func (p *ChatResp) Write(w io.Writer) (err error) {
	if err = binary.Write(w, binary.BigEndian, &p.Code); err != nil {
		return
	}
	if err = p.WriteString(w, p.Message); err != nil {
		return
	}
	return
}

func (p *ChatResp) Read(r io.Reader) (err error) {
	if err = binary.Read(r, binary.BigEndian, &p.Code); err != nil {
		return
	}
	if p.Message, err = p.ReadString(r); err != nil {
		return
	}
	return
}
