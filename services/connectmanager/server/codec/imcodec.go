package codec

import (
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
	// bs := make([]byte, 4)
	// reader.Read(bs)
	// for _, b := range bs {
	// 	fmt.Println("decode:", b)
	// }

	imMsg := &BaseMessage{}
	imMsg.Decode(reader)

	//fmt.Println("aaa=decode\tname:", imMsg.Name, "\tage:", imMsg.Age)

	// post text
	ctx.HandleRead(imMsg)
}
func (*IMCodecHandler) HandleWrite(ctx netty.OutboundContext, message netty.Message) {
	switch s := message.(type) {
	case *BaseMessage:
		// buf := bytes.NewBuffer(tools.Int2Bytes(s.MsgLength()))
		// buf.Write(s.Encode().Bytes())
		//fmt.Println(s.Name, ":", s.Age)
		buf := s.Encode()
		ctx.HandleWrite(buf)
	default:
		fmt.Println(s)
		ctx.HandleWrite(message)
	}
}
