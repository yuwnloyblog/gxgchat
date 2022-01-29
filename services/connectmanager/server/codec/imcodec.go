package codec

import (
	"bytes"
	"fmt"

	"github.com/go-netty/go-netty"
	"github.com/go-netty/go-netty/utils"
)

type IMCodecHandler struct{}

func (*IMCodecHandler) CodecName() string {
	return "im-codec"
}

func (IMCodecHandler) HandleRead(ctx netty.InboundContext, message netty.Message) {
	reader := utils.MustToReader(message)
	var imMsg IMessage
	tmpBs := make([]byte, 1)
	reader.Read(tmpBs)
	version := tmpBs[0]
	if version == Version_0 {
		msgHeader := &MsgHeader{Version: Version_0}
		msgHeader.DecodeHeader(reader)
		var msgBodyBytes []byte
		if msgHeader.MsgBodySize > 0 {
			msgBodyBytes = make([]byte, msgHeader.MsgBodySize)
			reader.Read(msgBodyBytes)
		} else {
			msgBodyBytes = []byte{}
		}

		//validate checksum  TODO
		switch msgHeader.GetCmd() {
		case Cmd_Connect:
			imMsg = NewConnectMessage(msgHeader)
		case Cmd_Disconnect:
			imMsg = NewDisconnectMessage(msgHeader)
		case Cmd_Ping:
			imMsg = NewPingMessage(msgHeader)
		case Cmd_Publish:
			imMsg = NewUserPublishMessage(msgHeader)
		case Cmd_PublishAck:
			imMsg = NewServerPublishAckMessage(msgHeader)
		case Cmd_Query:
			imMsg = NewQueryMessage(msgHeader)
		case Cmd_QueryConfirm:
			imMsg = NewQueryConfirmMessage(msgHeader)
		default:
			return
		}
		imMsg.DecodeBody(msgBodyBytes)
	}
	ctx.HandleRead(imMsg)
}
func (IMCodecHandler) HandleWrite(ctx netty.OutboundContext, message netty.Message) {
	switch s := message.(type) {
	case IMessage:
		msgBody, err := s.EncodeBody()
		if err == nil {
			buf := bytes.NewBuffer([]byte{})
			s.EncodeHeader(buf, msgBody) //encode header
			if len(msgBody) > 0 {
				buf.Write(msgBody) //write msg body
			}
			ctx.HandleWrite(buf)
		} else {
			ctx.HandleWrite(message)
		}
	default:
		fmt.Println(s)
		ctx.HandleWrite(message)
	}
}
