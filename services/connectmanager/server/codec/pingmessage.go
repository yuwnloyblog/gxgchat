package codec

type PingMessage struct {
	MsgHeader
}

func NewPingMessage(header *MsgHeader) *PingMessage {
	msg := &PingMessage{
		MsgHeader: MsgHeader{
			Version:     Version_0,
			HeaderCode:  header.HeaderCode,
			Checksum:    header.Checksum,
			MsgBodySize: header.MsgBodySize,
		},
	}
	msg.SetCmd(Cmd_Ping)
	msg.SetQoS(QoS_NeedAck)
	return msg
}
func (msg *PingMessage) EncodeBody() ([]byte, error) {
	return []byte{}, nil
}

func (msg *PingMessage) DecodeBody(msgBodyBytes []byte) error {
	return nil
}
