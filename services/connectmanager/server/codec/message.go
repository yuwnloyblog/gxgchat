package codec

import (
	"bytes"
	"io"

	"github.com/yuwnloyblog/gxgchat/commons/utils"
)

const (
	QoS_NoAck   int = 0
	QoS_NeedAck int = 1

	Cmd_Connect      int = 0
	Cmd_ConnectAck   int = 1
	Cmd_Disconnect   int = 2
	Cmd_Publish      int = 3
	Cmd_PublishAck   int = 4
	Cmd_Query        int = 5
	Cmd_QueryAck     int = 6
	Cmd_QueryConfirm int = 7
	Cmd_Ping         int = 8
	Cmd_Pong         int = 9
)

type IMessage interface {
	Encode() *bytes.Buffer
	Decode(io.Reader)
	MsgLength() int
}

type MsgHeader struct {
	version  byte
	cmd      int
	qos      int
	checksum byte
}

type BaseMessage struct {
	MsgHeader
	Name string
	Age  int
}

func (msg *BaseMessage) EncodeHeader(buf *bytes.Buffer) {
	buf.WriteByte(msg.version) // 1
	headerCode := msg.cmd & 0x0F
	headerCode = headerCode << 4
	qos := msg.qos & 0x03
	qos = qos << 6
	headerCode = headerCode | qos

	buf.WriteByte(byte(headerCode))
	buf.WriteByte(msg.checksum)
}

func (msg *BaseMessage) Encode() *bytes.Buffer {
	buf := bytes.NewBuffer([]byte{})
	msg.EncodeHeader(buf)
	buf.Write(utils.Int2Bytes(msg.MsgLength()))
	nameBytes := utils.String2Bytes(msg.Name)
	lenNameBytes := utils.Int2Bytes(len(nameBytes))
	buf.Write(lenNameBytes)
	buf.Write(nameBytes)

	ageBytes := utils.Int2Bytes(msg.Age)
	//fmt.Println("age:", msg.Age, "\t", utils.Bytes2Int(ageBytes))
	buf.Write(ageBytes)
	return buf
}

func (msg *BaseMessage) Decode(reader io.Reader) {
	b := make([]byte, 1)
	reader.Read(b)
	msg.version = b[0]
	reader.Read(b)
	msg.cmd = int(b[0])
	reader.Read(b)
	msg.checksum = b[0]

	lengthBytes := make([]byte, 4)
	reader.Read(lengthBytes)
	length := utils.Bytes2Int(lengthBytes)
	msgBytes := make([]byte, length)
	reader.Read(msgBytes)
	//encode or decode for msgBytes
	msgReader := bytes.NewReader(msgBytes)
	nameLenBytes := make([]byte, 4)
	msgReader.Read(nameLenBytes)
	nameLen := utils.Bytes2Int(nameLenBytes)

	nameBytes := make([]byte, nameLen)
	msgReader.Read(nameBytes)
	name := utils.Bytes2String(nameBytes)

	ageBytes := make([]byte, 4)
	msgReader.Read(ageBytes)
	age := utils.Bytes2Int(ageBytes)

	msg.Name = name
	msg.Age = age
}

func (msg *BaseMessage) MsgLength() int {
	length := 4 + 3
	nameBytes := utils.String2Bytes(msg.Name)
	length += len(nameBytes)
	length += 4
	return length
}
