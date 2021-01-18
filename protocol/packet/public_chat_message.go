package packet

// 接收到聊天消息
import (
	"encoding/binary"
	"fmt"
	"halia-chat/protocol"
	"io"
	"time"
)

type PublicChatMessage struct {
	protocol.BasePacket
	Time      time.Time // 发布时间
	Publisher string    // 发布者
	MsgType   uint8     // 消息类型
	Message   string    // 消息内容
}

func (p PublicChatMessage) String() string {
	return fmt.Sprintf("PublicChatMessage{Time=%s,Publisher=%s,MsgType=%d,Message=%s}", p.Time, p.Publisher, p.MsgType, p.Message)
}

func (PublicChatMessage) Opcode() uint16 {
	return protocol.OpPublicMessage
}

func (p *PublicChatMessage) Write(w io.Writer) (err error) {
	var (
		timestamp = p.Time.Unix()
	)
	if err = binary.Write(w, binary.BigEndian, &timestamp); err != nil {
		return
	}
	if err = p.WriteString(w, p.Publisher); err != nil {
		return
	}
	if err = binary.Write(w, binary.BigEndian, &p.MsgType); err != nil {
		return
	}
	if err = p.WriteShortString(w, p.Message); err != nil {
		return
	}
	return
}

func (p *PublicChatMessage) Read(r io.Reader) (err error) {
	var (
		timestamp int64
	)
	if err = binary.Read(r, binary.BigEndian, &timestamp); err != nil {
		return
	}
	p.Time = time.Unix(timestamp, 0)
	if p.Publisher, err = p.ReadString(r); err != nil {
		return
	}
	if err = binary.Read(r, binary.BigEndian, &p.MsgType); err != nil {
		return
	}
	if p.Message, err = p.ReadShortString(r); err != nil {
		return
	}
	return
}

func NewChatMessage(time time.Time, publisher string, msgType uint8, message string) *PublicChatMessage {
	return &PublicChatMessage{
		BasePacket: protocol.BasePacket{
			MagicNumber: protocol.MagicNumber,
			Opcode:      protocol.OpPublicMessage,
		},
		Time:      time,
		Publisher: publisher,
		MsgType:   msgType,
		Message:   message,
	}
}
