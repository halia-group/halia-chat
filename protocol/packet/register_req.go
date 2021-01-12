package packet

// 注册请求
import (
	"fmt"
	"halia-chat/protocol"
	"io"
)

type RegisterReq struct {
	protocol.BasePacket
	Username string
	Password string
}

func (p RegisterReq) String() string {
	return fmt.Sprintf("RegisterReq{Username=%s,Password=%s}", p.Username, p.Password)
}

func NewRegisterReq(username string, password string) *RegisterReq {
	return &RegisterReq{
		BasePacket: protocol.BasePacket{
			MagicNumber: protocol.MagicNumber,
			Opcode:      protocol.OpRegisterReq,
		},
		Username: username,
		Password: password,
	}
}

func (RegisterReq) Opcode() uint16 {
	return protocol.OpRegisterReq
}

func (p *RegisterReq) Write(w io.Writer) (err error) {
	if err = p.WriteString(w, p.Username); err != nil {
		return
	}
	if err = p.WriteString(w, p.Password); err != nil {
		return
	}
	return
}

func (p *RegisterReq) Read(r io.Reader) (err error) {
	if p.Username, err = p.ReadString(r); err != nil {
		return
	}
	if p.Password, err = p.ReadString(r); err != nil {
		return
	}
	return
}
