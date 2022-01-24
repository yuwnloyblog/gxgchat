package codec

type PongMessage struct {
	MsgHeader
}

func NewPongMessage(header *MsgHeader) *PongMessage {
	msg := &PongMessage{
		MsgHeader: MsgHeader{
			Version:     Version_0,
			HeaderCode:  header.HeaderCode,
			Checksum:    header.Checksum,
			MsgBodySize: header.MsgBodySize,
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