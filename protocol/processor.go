package protocol

// 数据包处理器
import (
	"context"
	"github.com/halia-group/halia/channel"
)

type Processor interface {
	Process(ctx context.Context, c channel.HandlerContext, packet Packet) error
}
