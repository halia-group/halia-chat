package packet

// 登录响应
import (
	"encoding/binary"
	"halia-chat/protocol"
	"io"
)

type LoginResp struct {
	protocol.BasePacket
	Code    uint8
	Message string
}

func NewLoginResp(code uint8, message string) *LoginResp {
	return &LoginResp{
		BasePacket: protocol.BasePacket{
			MagicNumber: protocol.MagicNumber,
			Opcode:      protocol.OpLoginResp,
		},
		Code:    code,
		Message: message,
	}
}

func (LoginResp) Opcode() uint16 {
	return protocol.OpLoginResp
}

func (p *LoginResp) Write(w io.Writer) (err error) {
	if err = binary.Write(w, binary.BigEndian, &p.Code); err != nil {
		return
	}
	if err = p.WriteString(w, p.Message); err != nil {
		return
	}
	return
}

func (p *LoginResp) Read(r io.Reader) (err error) {
	if err = binary.Read(r, binary.BigEndian, &p.Code); err != nil {
		return
	}
	if p.Message, err = p.ReadString(r); err != nil {
		return
	}
	return
}
