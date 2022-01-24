package codec

import "github.com/yuwnloyblog/gxgchat/commons/tools"

type QueryMessage struct {
	MsgHeader
	MsgBody *QueryMsgBody
}

func NewQueryMessage(header *MsgHeader) *QueryMessage {
	msg := &QueryMessage{
		MsgHeader: MsgHeader{
			Version:     Version_0,
			HeaderCode:  header.HeaderCode,
			Checksum:    header.Checksum,
			MsgBodySize: header.MsgBodySize,
		},
	}
	return msg
}

func (msg *QueryMessage) EncodeBody() ([]byte, error) {
	if msg.MsgBody != nil {
		return tools.PbMarshal(msg.MsgBody)
	}
	return nil, &CodecError{"MsgBody's length is 0."}
}

func (msg *QueryMessage) DecodeBody(msgBodyBytes []byte) error {
	msg.MsgBody = &QueryMsgBody{}
	return tools.PbUnMarshal(msgBodyBytes, msg.MsgBody)
}
