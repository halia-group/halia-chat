package client

import (
	"context"
	"github.com/halia-group/halia/channel"
	log "github.com/sirupsen/logrus"
	"halia-chat/protocol"
	"halia-chat/protocol/packet"
)

type registerProcessor struct {
	cc *ChatClient
}

func (p registerProcessor) Process(ctx context.Context, c channel.HandlerContext, msg protocol.Packet) error {
	resp := msg.(*packet.RegisterResp)
	if resp.Code != 0 {
		log.WithField("component", "registerProcessor").Warnln(resp.Message)
		return nil
	}

	return c.WriteAndFlush(packet.NewLoginReq("xialei", "111111"))
}
