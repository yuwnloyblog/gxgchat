package codec

import (
	"github.com/yuwnloyblog/gxgchat/commons/tools"
)

type ConnectMessage struct {
	MsgHeader
	MsgBody *ConnectMsgBody
}

func NewConnectMessage(header *MsgHeader) *ConnectMessage {
	msg := &ConnectMessage{
		MsgHeader: MsgHeader{
			Version:  Version_0,
			Checksum: header.Checksum,
		},
	}
	msg.SetCmd(Cmd_Connect)
	msg.SetQoS(QoS_NeedAck)
	return msg
}

func (msg *ConnectMessage) EncodeBody() ([]byte, error) {
	if msg.MsgBody != nil {
		return tools.PbMarshal(msg.MsgBody)
	}
	return nil, &CodecError{"MsgBody's length is 0."}
}

func (msg *ConnectMessage) DecodeBody(msgBodyBytes []byte) error {
	msg.MsgBody = &ConnectMsgBody{}
	return tools.PbUnMarshal(msgBodyBytes, msg.MsgBody)
}

func (msg *ConnectMessage) GetMsgBodySize() int {
	return msg.MsgBodySize
}
