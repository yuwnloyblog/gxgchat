package codec

import "github.com/yuwnloyblog/gxgchat/commons/tools"

type ConnectAckMessage struct {
	MsgHeader
	MsgBody *ConnectAckMsgBody
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
