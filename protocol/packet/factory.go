package packet

// 数据包Opcode映射共产
import "halia-chat/protocol"

var Factory = map[uint16]func() protocol.Packet{
	protocol.OpPing: func() protocol.Packet {
		return new(Ping)
	},
	protocol.OpPong: func() protocol.Packet {
		return new(Pong)
	},
	protocol.OpRegisterReq: func() protocol.Packet {
		return new(RegisterReq)
	},
	protocol.OpRegisterResp: func() protocol.Packet {
		return new(RegisterResp)
	},
}
