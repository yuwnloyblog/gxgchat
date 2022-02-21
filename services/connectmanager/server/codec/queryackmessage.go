package codec

import "github.com/yuwnloyblog/gxgchat/commons/tools"

type QueryAckMessage struct {
	MsgHeader
	MsgBody *QueryAckMsgBody
}

func NewQueryAckMessageWithHeader(header *MsgHeader) *QueryAckMessage {
	msg := &QueryAckMessage{
		MsgHeader: MsgHeader{
			Version:     Version_0,
			HeaderCode:  header.HeaderCode,
			Checksum:    header.Checksum,
			MsgBodySize: header.MsgBodySize,
		},
	}
	return msg
}

func (msg *QueryAckMessage) EncodeBody() ([]byte, error) {
	if msg.MsgBody != nil {
		return tools.PbMarshal(msg.MsgBody)
	}
	return nil, &CodecError{"MsgBody's length is 0."}
}

func (msg *QueryAckMessage) DecodeBody(msgBodyBytes []byte) error {
	msg.MsgBody = &QueryAckMsgBody{}
	return tools.PbUnMarshal(msgBodyBytes, msg.MsgBody)
}
