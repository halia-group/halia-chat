package server

import (
	"context"
	"github.com/halia-group/halia/channel"
	"halia-chat/protocol"
	"halia-chat/protocol/packet"
	"halia-chat/server/dao"
	"time"
)

type chatProcessor struct {
	dao    dao.Dao
	server *ChatServer
}

func (p chatProcessor) Process(ctx context.Context, c channel.HandlerContext, msg protocol.Packet) error {
	if !c.Channel().GetBoolAttribute(AttrLogged) {
		return c.WriteAndFlush(packet.NewUnAuthorization())
	}

	req := msg.(*packet.PublicChatReq)
	nickname := c.Channel().GetStringAttribute(AttrNickname)

	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	loginMsg := packet.NewChatMessage(time.Now(), nickname, req.MsgType, req.Message)
	return p.server.broadcast(ctx, loginMsg)
}
