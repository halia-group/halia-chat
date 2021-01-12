package client

import (
	"context"
	"github.com/halia-group/halia/channel"
	log "github.com/sirupsen/logrus"
	"halia-chat/protocol"
	"halia-chat/protocol/packet"
)

type chatRespProcessor struct {
	cc *ChatClient
}

func (p chatRespProcessor) Process(ctx context.Context, c channel.HandlerContext, msg protocol.Packet) error {
	resp := msg.(*packet.ChatResp)
	if resp.Code != 0 {
		log.WithField("component", "chatProcessor").Warnln(resp.Message)
		return nil
	}
	return nil
}
