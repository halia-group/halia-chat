package server

import (
	"context"
	"fmt"
	"github.com/halia-group/halia/channel"
	log "github.com/sirupsen/logrus"
	"halia-chat/protocol"
	"halia-chat/protocol/packet"
	"time"
)

type handler struct {
	server *ChatServer
	log    *log.Entry
}

func newHandler(server *ChatServer) *handler {
	return &handler{server: server, log: log.WithField("component", "handler")}
}

func (p *handler) OnError(c channel.HandlerContext, err error) {
	p.log.WithField("peer", c.Channel().RemoteAddr()).WithError(err).Warnln("an error was caused")
}

func (p *handler) ChannelActive(c channel.HandlerContext) {
	p.log.WithField("peer", c.Channel().RemoteAddr()).Debugln("connected")
}

func (p *handler) ChannelInActive(c channel.HandlerContext) {
	p.log.WithField("peer", c.Channel().RemoteAddr()).Debugln("disconnected")
	// 广播离开消息
	if !c.Channel().HasAttribute(AttrUserId) {
		return
	}
	p.server.removeChannel(c.Channel().GetAttribute(AttrUserId).(int))

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer cancel()
	msg := fmt.Sprintf("<%s>已离开", c.Channel().GetStringAttribute(AttrNickname))
	loginMsg := packet.NewChatMessage(time.Now(), SystemUsername, protocol.MsgText, msg)
	p.server.broadcast(ctx, loginMsg)
}

func (p *handler) ChannelRead(c channel.HandlerContext, msg interface{}) {
	req := msg.(protocol.Packet)
	if p.server.processors[req.Opcode()] == nil {
		p.log.WithField("packet", req).Warnln("unknown packet")
		return
	}
	if err := p.server.processors[req.Opcode()].Process(context.Background(), c, req); err != nil {
		p.OnError(c, err)
	}
}
