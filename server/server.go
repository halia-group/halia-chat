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
	"halia-chat/server/dao"
	"net"
	"sync"
)

const (
	SystemUsername = "system"
)

const (
	AttrLogged   = "logged"
	AttrUserId   = "uid"
	AttrNickname = "nickname"
)

type ChatServer struct {
	channels   map[int]channel.Channel
	lock       sync.RWMutex
	log        *log.Entry
	server     *bootstrap.Server
	processors map[uint16]protocol.Processor
}

func NewServer() (*ChatServer, error) {
	cs := &ChatServer{
		channels: make(map[int]channel.Channel),
		log:      log.WithField("component", "ChatServer"),
	}
	d, err := dao.New()
	if err != nil {
		return nil, err
	}
	cs.processors = NewProcessors(d, cs)
	return cs, nil
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

func (p *ChatServer) addChannel(uid int, c channel.Channel) {
	p.lock.Lock()
	p.channels[uid] = c
	p.lock.Unlock()
}

func (p *ChatServer) removeChannel(uid int) {
	p.lock.Lock()
	delete(p.channels, uid)
	p.lock.Unlock()
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
