package codec

type PongMessage struct {
	MsgHeader
}

func NewPongMessageWithHeader(header *MsgHeader) *PongMessage {
	msg := &PongMessage{
		MsgHeader: MsgHeader{
			Version:     Version_1,
			HeaderCode:  header.HeaderCode,
			Checksum:    header.Checksum,
			Sequence:    header.Sequence,
			MsgBodySize: header.MsgBodySize,
		},
	}
	msg.SetCmd(Cmd_Pong)
	msg.SetQoS(QoS_NoAck)
	return msg
}

func NewPongMessageWithSeq(seq [2]byte) *PongMessage {
	msg := &PongMessage{
		MsgHeader: MsgHeader{
			Version:  Version_1,
			Sequence: seq,
		},
	}
	msg.SetCmd(Cmd_Pong)
	msg.SetQoS(QoS_NoAck)
	return msg
}

func NewPongMessage() *PongMessage {
	msg := &PongMessage{
		MsgHeader: MsgHeader{
			Version: Version_1,
		},
	}
	msg.SetCmd(Cmd_Pong)
	msg.SetQoS(QoS_NoAck)
	return msg
}

func (msg *PongMessage) EncodeBody() ([]byte, error) {
	return []byte{}, nil
}

func (msg *PongMessage) DecodeBody(msgBodyBytes []byte) error {
	return nil
}
