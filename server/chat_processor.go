package server

import (
	"context"
	"github.com/halia-group/halia/channel"
	"halia-chat/protocol"
	"halia-chat/protocol/packet"
	"time"
)

type chatProcessor struct {
	dao    Dao
	server *ChatServer
}

func (p chatProcessor) Process(ctx context.Context, c channel.HandlerContext, msg protocol.Packet) error {
	if !c.Channel().GetBoolAttribute(AttrLogged) {
		return c.WriteAndFlush(packet.NewUnAuthorization())
	}

	req := msg.(*packet.ChatReq)
	username := c.Channel().GetStringAttribute(AttrUsername)

	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	loginMsg := packet.NewChatMessage(time.Now(), username, req.MsgType, req.Message)
	return p.server.broadcast(ctx, loginMsg)
}
