package packet

import (
	"halia-chat/protocol"
	"io"
)

type UnAuthorization struct {
	protocol.BasePacket
}

func (p UnAuthorization) String() string {
	return "UnAuthorization"
}

func NewUnAuthorization() *UnAuthorization {
	return &UnAuthorization{
		BasePacket: protocol.BasePacket{
			MagicNumber: protocol.MagicNumber,
			Opcode:      protocol.OpUnAuthorization,
		},
	}
}

func (p UnAuthorization) Opcode() uint16 {
	return protocol.OpUnAuthorization
}

func (p UnAuthorization) Write(w io.Writer) error {
	return nil
}

func (p UnAuthorization) Read(r io.Reader) error {
	return nil
}
