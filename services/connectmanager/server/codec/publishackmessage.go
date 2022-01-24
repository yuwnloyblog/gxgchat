package codec

import "github.com/yuwnloyblog/gxgchat/commons/tools"

type ServerPublishAckMessage struct {
	MsgHeader
	MsgBody *PublishAckMsgBody
}

type UserPublishAckMessage struct {
	ServerPublishAckMessage
}

func NewServerPublishAckMessage(header *MsgHeader) *ServerPublishAckMessage {
	msg := &ServerPublishAckMessage{
		MsgHeader: MsgHeader{
			Version:     Version_0,
			HeaderCode:  header.HeaderCode,
			Checksum:    header.Checksum,
			MsgBodySize: header.MsgBodySize,
		},
	}
	return msg
}

func (msg *ServerPublishAckMessage) EncodeBody() ([]byte, error) {
	if msg.MsgBody != nil {
		return tools.PbMarshal(msg.MsgBody)
	}
	return nil, &CodecError{"MsgBody's length is 0."}
}

func (msg *ServerPublishAckMessage) DecodeBody(msgBodyBytes []byte) error {
	msg.MsgBody = &PublishAckMsgBody{}
	return tools.PbUnMarshal(msgBodyBytes, msg.MsgBody)
}