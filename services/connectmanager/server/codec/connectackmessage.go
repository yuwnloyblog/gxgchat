package codec

import "github.com/yuwnloyblog/gxgchat/commons/tools"

type ConnectAckMessage struct {
	MsgHeader
	MsgBody *ConnectAckMsgBody
}

func NewConnectAckMessageWithHeader(header *MsgHeader) *ConnectMessage {
	msg := &ConnectMessage{
		MsgHeader: MsgHeader{
			Version:     Version_0,
			HeaderCode:  header.HeaderCode,
			Checksum:    header.Checksum,
			MsgBodySize: header.MsgBodySize,
		},
	}
	msg.SetCmd(Cmd_Connect)
	msg.SetQoS(QoS_NoAck)
	return msg
}

func (msg *ConnectAckMessage) EncodeBody() ([]byte, error) {
	if msg.MsgBody != nil {
		return tools.PbMarshal(msg.MsgBody)
	}
	return nil, &CodecError{"MsgBody's length is 0."}
}

func (msg *ConnectAckMessage) DecodeBody(msgBodyBytes []byte) error {
	msg.MsgBody = &ConnectAckMsgBody{}
	return tools.PbUnMarshal(msgBodyBytes, msg.MsgBody)
}
