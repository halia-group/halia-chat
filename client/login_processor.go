package client

import (
	"context"
	"github.com/halia-group/halia/channel"
	log "github.com/sirupsen/logrus"
	"halia-chat/protocol"
	"halia-chat/protocol/packet"
)

type loginProcessor struct {
	cc *ChatClient
}

func (p loginProcessor) Process(ctx context.Context, c channel.HandlerContext, msg protocol.Packet) error {
	resp := msg.(*packet.LoginResp)
	if resp.Code != 0 {
		log.WithField("component", "loginProcessor").Warnln(resp.Message)
		return nil
	}

	return c.WriteAndFlush(packet.NewChatReq(protocol.MsgText, "大家好"))
}
