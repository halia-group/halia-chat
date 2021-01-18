package client

import (
	"context"
	"fmt"
	"github.com/halia-group/halia/channel"
	"halia-chat/protocol"
	"halia-chat/protocol/packet"
)

type chatMessageProcessor struct {
	cc *ChatClient
}

func (p chatMessageProcessor) Process(ctx context.Context, c channel.HandlerContext, msg protocol.Packet) error {
	resp := msg.(*packet.PublicChatMessage)
	fmt.Printf("%s <%s>: %s\n", resp.Time.String(), resp.Publisher, resp.Message)
	return nil
}
