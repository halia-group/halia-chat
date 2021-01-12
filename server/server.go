package server

import (
	"context"
	"encoding/binary"
	"github.com/halia-group/halia/bootstrap"
	"github.com/halia-group/halia/channel"
	"github.com/halia-group/halia/handler/codec"
	log "github.com/sirupsen/logrus"
	"halia-chat/protocol"
	"halia-chat/protocol/packet"
	"net"
	"sync"
)

const (
	SystemUsername = "system"
)

const (
	AttrLogged   = "logged"
	AttrUsername = "username"
)

type ChatServer struct {
	channels   []channel.Channel
	lock       sync.RWMutex
	log        *log.Entry
	server     *bootstrap.Server
	processors map[uint16]protocol.Processor
}

func NewServer() *ChatServer {
	cs := &ChatServer{
		channels: make([]channel.Channel, 0),
		log:      log.WithField("component", "ChatServer"),
	}
	cs.processors = NewProcessors(newDao(), cs)
	return cs
}
func (p *ChatServer) Run(network, addr string) (err error) {
	options := &bootstrap.ServerOptions{ChannelFactory: func(conn net.Conn) channel.Channel {
		c := channel.NewDefaultChannel(conn)
		c.Pipeline().AddInbound("frameDecoder", codec.NewLengthFieldBasedFrameDecoder(2, 4, binary.BigEndian))
		c.Pipeline().AddInbound("packetDecoder", protocol.NewPacketDecoder(packet.Factory))
		c.Pipeline().AddInbound("handler", newHandler(p))

		c.Pipeline().AddOutbound("packetEncoder", protocol.NewPacketEncoder())
		return c
	}}
	p.server = bootstrap.NewServer(options)
	return p.server.Listen(network, addr)
}

func (p *ChatServer) addChannel(c channel.Channel) {
	index := p.indexOfChannel(c)
	if index != -1 {
		return
	}
	p.lock.Lock()
	p.channels = append(p.channels, c)
	p.lock.Unlock()
}

func (p *ChatServer) removeChannel(c channel.Channel) {
	index := p.indexOfChannel(c)
	if index == -1 {
		return
	}
	p.lock.Lock()
	p.channels = append(p.channels[:index], p.channels[index+1:]...)
	p.lock.Unlock()
}

func (p *ChatServer) indexOfChannel(c channel.Channel) int {
	var index = -1
	p.lock.RLock()
	for i := range p.channels {
		if p.channels[i] == c {
			index = i
		}
	}
	p.lock.RUnlock()
	return index
}

// 广播消息
func (p *ChatServer) broadcast(ctx context.Context, msg protocol.Packet) error {
	var (
		ch     = make(chan struct{})
		logger = p.log.WithField("component", "processor")
	)
	go func() {
		for _, c := range p.channels {
			if err := c.Pipeline().WriteAndFlush(msg); err != nil {
				logger.WithField("peer", c.RemoteAddr()).WithError(err).Warnln("write error")
			}
		}
		ch <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-ch:
		return nil
	}
}
