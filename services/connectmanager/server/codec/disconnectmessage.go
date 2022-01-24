package codec

import "github.com/yuwnloyblog/gxgchat/commons/tools"

type DisconnectMessage struct {
	MsgHeader
	MsgBody *DisconnectMsgBody
}

func (msg *DisconnectMessage) EncodeBody() ([]byte, error) {
	if msg.MsgBody != nil {
		return tools.PbMarshal(msg.MsgBody)
	}
	return nil, &CodecError{"MsgBody's length is 0."}
}

func (msg *DisconnectMessage) DecodeBody(msgBodyBytes []byte) error {
	msg.MsgBody = &DisconnectMsgBody{}
	return tools.PbUnMarshal(msgBodyBytes, msg.MsgBody)
}
