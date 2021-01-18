package packet

// 聊天响应
import (
	"encoding/binary"
	"fmt"
	"halia-chat/protocol"
	"io"
)

type PublicChatResp struct {
	protocol.BasePacket
	Code    uint8
	Message string
}

func (p PublicChatResp) String() string {
	return fmt.Sprintf("PublicChatResp{Code=%d,Message=%s}", p.Code, p.Message)
}

func NewChatResp(code uint8, message string) *PublicChatResp {
	return &PublicChatResp{
		BasePacket: protocol.BasePacket{
			MagicNumber: protocol.MagicNumber,
			Opcode:      protocol.OpPublicChatResp,
		},
		Code:    code,
		Message: message,
	}
}

func (PublicChatResp) Opcode() uint16 {
	return protocol.OpPublicChatResp
}

func (p *PublicChatResp) Write(w io.Writer) (err error) {
	if err = binary.Write(w, binary.BigEndian, &p.Code); err != nil {
		return
	}
	if err = p.WriteString(w, p.Message); err != nil {
		return
	}
	return
}

func (p *PublicChatResp) Read(r io.Reader) (err error) {
	if err = binary.Read(r, binary.BigEndian, &p.Code); err != nil {
		return
	}
	if p.Message, err = p.ReadString(r); err != nil {
		return
	}
	return
}
