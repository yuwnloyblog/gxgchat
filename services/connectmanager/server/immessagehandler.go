package server

import "github.com/go-netty/go-netty"

type IMMessageHandler struct{}

func (IMMessageHandler) HandleActive(ctx netty.ActiveContext) {

	ctx.HandleActive()
}

func (IMMessageHandler) HandleRead(ctx netty.InboundContext, message netty.Message) {

	ctx.HandleRead(message)
}

func (IMMessageHandler) HandleInactive(ctx netty.InactiveContext, ex netty.Exception) {

	ctx.HandleInactive(ex)
}

func (IMMessageHandler) HandleException(ctx netty.ExceptionContext, ex netty.Exception) {
	ctx.HandleException(ex)
}
