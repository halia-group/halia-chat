package client

import (
	"context"
	"github.com/halia-group/halia/channel"
	"github.com/halia-group/halia/examples/chat/common"
	"halia-chat/protocol"
)

// todo UI句柄
type pongProcessor struct {
	cc *ChatClient
}

func (p pongProcessor) Process(ctx context.Context, c channel.HandlerContext, msg protocol.Packet) error {
	return c.WriteAndFlush(common.NewRegisterReq("xialei", "111111"))
}

func newPongProcessor(cc *ChatClient) *pongProcessor {
	return &pongProcessor{cc: cc}
}
