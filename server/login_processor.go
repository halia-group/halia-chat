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

type loginProcessor struct {
	dao    Dao
	server *ChatServer
	log    *log.Entry
}

func NewLoginProcessor(dao Dao, server *ChatServer) *loginProcessor {
	return &loginProcessor{dao: dao, server: server, log: log.WithField("component", "loginProcessor")}
}

func (p *loginProcessor) Process(ctx context.Context, c channel.HandlerContext, msg protocol.Packet) error {
	req := msg.(*packet.LoginReq)
	// 登录失败
	if err := p.dao.Login(ctx, req.Username, req.Password); err != nil {
		return c.WriteAndFlush(packet.NewLoginResp(1, err.Error()))
	}
	// 登录成功
	if err := c.WriteAndFlush(packet.NewLoginResp(0, "登录成功")); err != nil {
		p.log.WithField("peer", c.Channel().RemoteAddr()).WithError(err).Warnln("write error")
		return err
	}
	// 设置channel状态
	c.Channel().SetAttribute(AttrLogged, true)
	c.Channel().SetAttribute(AttrUsername, req.Username)
	// 添加channel到server
	p.server.addChannel(c.Channel())
	// 广播登录消息
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	loginMsg := packet.NewChatMessage(time.Now(), SystemUsername, protocol.MsgText, fmt.Sprintf("<%s>已登录", req.Username))
	return p.server.broadcast(ctx, loginMsg)
}
