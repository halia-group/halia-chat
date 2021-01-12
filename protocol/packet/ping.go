package packet

// 客户端请求：检测服务器是否通畅
import (
	"halia-chat/protocol"
	"io"
)

type Ping struct{ protocol.BasePacket }

func (p Ping) String() string {
	return "Ping{}"
}

func NewPing() *Ping {
	return &Ping{
		BasePacket: protocol.BasePacket{
			MagicNumber: protocol.MagicNumber,
			Opcode:      protocol.OpPing,
		},
	}
}

func (p Ping) Opcode() uint16 {
	return protocol.OpPing
}

func (p Ping) Write(w io.Writer) error {
	return nil
}

func (p Ping) Read(r io.Reader) error {
	return nil
}
