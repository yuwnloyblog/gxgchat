package clients

import (
	"fmt"

	"github.com/go-netty/go-netty"
	"github.com/yuwnloyblog/gxgchat/services/connectmanager/server/codec"
)

type ImClientMessageHandler struct {
}

func (handler ImClientMessageHandler) HandleActive(ctx netty.ActiveContext) {

	ctx.HandleActive()
}

func (handler ImClientMessageHandler) HandleRead(ctx netty.InboundContext, message netty.Message) {
	switch msg := message.(type) {
	case *codec.ConnectAckMessage:
		fmt.Println(msg.MsgBody)
	case *codec.DisconnectMessage:

	case *codec.PongMessage:

	case *codec.UserPublishMessage:

	case *codec.ServerPublishAckMessage:

	case *codec.QueryMessage:

	case *codec.QueryConfirmMessage:

	default:
		break
	}

	ctx.HandleRead(message)
}

func (handler ImClientMessageHandler) HandleInactive(ctx netty.InactiveContext, ex netty.Exception) {

	ctx.Close(ex)
	ctx.HandleInactive(ex)
}

func (handler ImClientMessageHandler) HandleException(ctx netty.ExceptionContext, ex netty.Exception) {

	ctx.Close(ex)
	ctx.HandleException(ex)
}
