package client

import (
	"context"
	"github.com/halia-group/halia/channel"
	log "github.com/sirupsen/logrus"
	"halia-chat/protocol"
	"halia-chat/protocol/packet"
)

type handler struct {
	log *log.Entry
	cc  *ChatClient
}

func newHandler(cc *ChatClient) *handler {
	return &handler{log: log.WithField("component", "handler"), cc: cc}
}

func (p handler) OnError(c channel.HandlerContext, err error) {
	p.log.WithError(err).Warnln("an error was caused")
}

func (p handler) ChannelActive(c channel.HandlerContext) {
	p.log.Debugln("connected")
	c.WriteAndFlush(packet.NewPing())
}

func (p handler) ChannelInActive(c channel.HandlerContext) {
	p.log.Debugln("disconnected")
}

func (p handler) ChannelRead(c channel.HandlerContext, msg interface{}) {
	resp := msg.(protocol.Packet)
	p.log.Debugln("packet", resp)
	if p.cc.processors[resp.Opcode()] == nil {
		p.log.WithField("packet", resp).Warnln("unknown packet")
		return
	}
	if err := p.cc.processors[resp.Opcode()].Process(context.Background(), c, resp); err != nil {
		p.OnError(c, err)
	}
}
