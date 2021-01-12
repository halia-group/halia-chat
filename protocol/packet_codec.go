package protocol

// 数据包编解码器
import (
	"bytes"
	"encoding/binary"
	"github.com/halia-group/halia/channel"
	"github.com/halia-group/halia/handler/codec"
)

type PacketDecoder struct {
	codec.Decoder
	factory map[uint16]func() Packet
}

func NewPacketDecoder(factory map[uint16]func() Packet) *PacketDecoder {
	return &PacketDecoder{factory: factory}
}

func (p PacketDecoder) ChannelRead(c channel.HandlerContext, msg interface{}) {
	var (
		buf = msg.([]byte)
		bp  = BasePacket{}
	)
	if err := bp.readCommonField(bytes.NewReader(buf)); err != nil {
		c.FireOnError(err)
		return
	}
	if p.factory[bp.Opcode] == nil {
		c.FireOnError(ErrUnknownOpcode)
		return
	}
	packet := p.factory[bp.Opcode]()
	if err := packet.Read(bytes.NewReader(bp.Data)); err != nil {
		c.FireOnError(err)
		return
	}
	c.FireChannelRead(packet)
}

type PacketEncoder struct{}

func NewPacketEncoder() *PacketEncoder {
	return &PacketEncoder{}
}

func (p PacketEncoder) OnError(c channel.HandlerContext, err error) {
	c.FireOnError(err)
}

func (p PacketEncoder) Write(c channel.HandlerContext, msg interface{}) error {
	var (
		packet = msg.(Packet)
		buf    = bytes.Buffer{}
	)
	if err := packet.Write(&buf); err != nil {
		return err
	}
	var (
		body   = buf.Bytes()
		opcode = packet.Opcode()
		length = uint16(len(body))
	)
	buf = bytes.Buffer{}
	if err := binary.Write(&buf, binary.BigEndian, &MagicNumber); err != nil {
		return err
	}
	if err := binary.Write(&buf, binary.BigEndian, &opcode); err != nil {
		return err
	}
	if err := binary.Write(&buf, binary.BigEndian, &length); err != nil {
		return err
	}
	if err := binary.Write(&buf, binary.BigEndian, &body); err != nil {
		return err
	}
	return c.Write(buf.Bytes())
}

func (p PacketEncoder) Flush(c channel.HandlerContext) error {
	return c.Channel().Flush()
}
