package packet

// 服务器响应
import (
	"halia-chat/protocol"
	"io"
)

type Pong struct{}

func (p Pong) Opcode() uint16 {
	return protocol.OpPong
}

func (p Pong) Write(w io.Writer) error {
	return nil
}

func (p Pong) Read(r io.Reader) error {
	return nil
}
