package client

import (
	"context"
	"github.com/halia-group/halia/channel"
	"halia-chat/protocol"
	"halia-chat/protocol/packet"
)

type pongProcessor struct {
	cc *ChatClient
}

func (p pongProcessor) Process(ctx context.Context, c channel.HandlerContext, msg protocol.Packet) error {
	return c.WriteAndFlush(packet.NewRegisterReq("xialei", "111111", "夏磊"))
}

func newPongProcessor(cc *ChatClient) *pongProcessor {
	return &pongProcessor{cc: cc}
}
