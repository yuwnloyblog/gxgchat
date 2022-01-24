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

func (*IMCodecHandler) HandleRead(ctx netty.InboundContext, message netty.Message) {
	reader := utils.MustToReader(message)
	ctx.SetAttachment("zafdsfsfa")

	var imMsg IMessage

	tmpBs := make([]byte, 1)
	reader.Read(tmpBs)
	version := tmpBs[0]
	if version == Version_0 {
		msgHeader := &MsgHeader{Version: Version_0}
		msgHeader.DecodeHeader(reader)

		msgBodyBytes := make([]byte, msgHeader.MsgBodySize)
		reader.Read(msgBodyBytes)

		//validate checksum  TODO
		switch msgHeader.GetCmd() {
		case Cmd_Connect:
			imMsg = NewConnectMessage(msgHeader)
			imMsg.DecodeBody(msgBodyBytes)
		}
	}
	ctx.HandleRead(imMsg)
}
func (*IMCodecHandler) HandleWrite(ctx netty.OutboundContext, message netty.Message) {
	switch s := message.(type) {
	case IMessage:
		msgBody, err := s.EncodeBody()
		if err == nil {
			buf := bytes.NewBuffer([]byte{})

			s.EncodeHeader(buf, msgBody) //encode header
			buf.Write(msgBody)           //write msg body
			ctx.HandleWrite(buf)
		}
	default:
		fmt.Println(s)
		ctx.HandleWrite(message)
	}
}
