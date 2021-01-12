package client

import (
	"encoding/binary"
	"github.com/halia-group/halia/bootstrap"
	"github.com/halia-group/halia/channel"
	"github.com/halia-group/halia/handler/codec"
	log "github.com/sirupsen/logrus"
	"halia-chat/protocol"
	"halia-chat/protocol/packet"
	"net"
)

type ChatClient struct {
	client     *bootstrap.Client
	processors map[uint16]protocol.Processor
	log        *log.Entry
}

func NewChatClient() *ChatClient {
	c := &ChatClient{
		log: log.WithField("component", "ChatClient"),
	}
	c.processors = NewProcessors(c)
	return c
}

func (p *ChatClient) Dial(network, addr string) error {
	p.client = bootstrap.NewClient(&bootstrap.ClientOptions{ChannelFactory: func(conn net.Conn) channel.Channel {
		c := channel.NewDefaultChannel(conn)
		c.Pipeline().AddInbound("frameDecoder", codec.NewLengthFieldBasedFrameDecoder(2, 4, binary.BigEndian))
		c.Pipeline().AddInbound("packetDecoder", protocol.NewPacketDecoder(packet.Factory))
		// todo: 传入UI句柄

		c.Pipeline().AddInbound("handler", newHandler(p))
		c.Pipeline().AddOutbound("packetEncoder", protocol.NewPacketEncoder())
		return c
	}})

	return p.client.Dial(network, addr)
}
