package packet

// 注册响应
import (
	"encoding/binary"
	"fmt"
	"halia-chat/protocol"
	"io"
)

type RegisterResp struct {
	protocol.BasePacket
	Code    uint8
	Message string
}

func (p RegisterResp) String() string {
	return fmt.Sprintf("RegisterResp{Code=%d,Message=%s}", p.Code, p.Message)
}

func NewRegisterResp(code uint8, message string) *RegisterResp {
	return &RegisterResp{
		BasePacket: protocol.BasePacket{
			MagicNumber: protocol.MagicNumber,
			Opcode:      protocol.OpRegisterResp,
		},
		Code:    code,
		Message: message,
	}
}

func (RegisterResp) Opcode() uint16 {
	return protocol.OpRegisterResp
}

func (p *RegisterResp) Write(w io.Writer) (err error) {
	if err = binary.Write(w, binary.BigEndian, &p.Code); err != nil {
		return
	}
	if err = p.WriteString(w, p.Message); err != nil {
		return
	}
	return
}

func (p *RegisterResp) Read(r io.Reader) (err error) {
	if err = binary.Read(r, binary.BigEndian, &p.Code); err != nil {
		return
	}
	if p.Message, err = p.ReadString(r); err != nil {
		return
	}
	return
}
