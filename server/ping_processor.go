package server

import (
	"context"
	"github.com/halia-group/halia/channel"
	"halia-chat/protocol"
	"halia-chat/protocol/packet"
)

type pingProcessor struct{}

func (p pingProcessor) Process(ctx context.Context, c channel.HandlerContext, msg protocol.Packet) error {
	return c.WriteAndFlush(new(packet.Pong))
}
