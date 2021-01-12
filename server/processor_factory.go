package server

import (
	"halia-chat/protocol"
)

func NewProcessors(dao Dao, cs *ChatServer) map[uint16]protocol.Processor {
	processors := make(map[uint16]protocol.Processor)
	processors[protocol.OpPing] = &pingProcessor{}
	processors[protocol.OpRegisterReq] = &registerProcessor{dao: dao}
	processors[protocol.OpLoginReq] = &loginProcessor{dao: dao, server: cs}
	processors[protocol.OpChatReq] = &chatProcessor{dao: dao, server: cs}
	return processors
}
