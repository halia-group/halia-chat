package server

import (
	"context"
	"github.com/halia-group/halia/channel"
	"halia-chat/protocol"
	"halia-chat/protocol/packet"
)

type registerProcessor struct {
	dao Dao
}

func NewRegisterProcessor(dao Dao) *registerProcessor {
	return &registerProcessor{dao: dao}
}

func (p registerProcessor) Process(ctx context.Context, c channel.HandlerContext, msg protocol.Packet) error {
	req := msg.(*packet.RegisterReq)
	// 注册失败
	if err := p.dao.Register(ctx, req.Username, req.Password); err != nil {
		return c.WriteAndFlush(packet.NewRegisterResp(1, err.Error()))
	}
	// 注册成功
	return c.WriteAndFlush(packet.NewRegisterResp(0, "注册成功"))
}
