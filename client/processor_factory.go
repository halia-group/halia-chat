package client

import "halia-chat/protocol"

func NewProcessors(cc *ChatClient) map[uint16]protocol.Processor {
	processors := make(map[uint16]protocol.Processor)
	processors[protocol.OpPong] = &pongProcessor{cc: cc}
	processors[protocol.OpRegisterResp] = &registerProcessor{cc: cc}
	processors[protocol.OpLoginResp] = &loginProcessor{cc: cc}
	processors[protocol.OpPublicChatResp] = &chatRespProcessor{cc: cc}
	processors[protocol.OpPublicMessage] = &chatMessageProcessor{cc: cc}
	return processors
}
