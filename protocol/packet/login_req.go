package packet

// 登录请求
import (
	"fmt"
	"halia-chat/protocol"
	"io"
)

type LoginReq struct {
	protocol.BasePacket
	Username string
	Password string
}

func (p LoginReq) String() string {
	return fmt.Sprintf("LoginReq{Username=%s,Password=%s}", p.Username, p.Password)
}

func NewLoginReq(username string, password string) *LoginReq {
	return &LoginReq{
		BasePacket: protocol.BasePacket{
			MagicNumber: protocol.MagicNumber,
			Opcode:      protocol.OpLoginReq,
		},
		Username: username,
		Password: password,
	}
}

func (LoginReq) Opcode() uint16 {
	return protocol.OpLoginReq
}

func (p *LoginReq) Write(w io.Writer) (err error) {
	if err = p.WriteString(w, p.Username); err != nil {
		return
	}
	if err = p.WriteString(w, p.Password); err != nil {
		return
	}
	return
}

func (p *LoginReq) Read(r io.Reader) (err error) {
	if p.Username, err = p.ReadString(r); err != nil {
		return
	}
	if p.Password, err = p.ReadString(r); err != nil {
		return
	}
	return
}
