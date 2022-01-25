package server

import (
	"github.com/go-netty/go-netty"
	"github.com/yuwnloyblog/gxgchat/services/connectmanager/server/codec"
)

type IMMessageHandler struct {
	listener ImListener
}

func (handler IMMessageHandler) HandleActive(ctx netty.ActiveContext) {
	if handler.listener != nil {
		handler.listener.Create(ctx)
	}
	ctx.HandleActive()
}

func (handler IMMessageHandler) HandleRead(ctx netty.InboundContext, message netty.Message) {
	if handler.listener != nil {
		switch msg := message.(type) {
		case *codec.ConnectMessage:
			handler.listener.Connected(msg, ctx)
		case *codec.DisconnectMessage:
			handler.listener.Diconnected(msg, ctx)
		case *codec.PingMessage:
			handler.listener.PingArrived(msg, ctx)
		case *codec.UserPublishMessage:
			handler.listener.PublishArrived(msg, ctx)
		case *codec.ServerPublishAckMessage:
			handler.listener.PubAckArrived(msg, ctx)
		case *codec.QueryMessage:
			handler.listener.QueryArrived(msg, ctx)
		case *codec.QueryConfirmMessage:
			handler.listener.QueryConfirmArrived(msg, ctx)
		default:
			break
		}
	}
	ctx.HandleRead(message)
}

func (handler IMMessageHandler) HandleInactive(ctx netty.InactiveContext, ex netty.Exception) {
	if handler.listener != nil {
		handler.listener.Close(ctx)
	}
	ctx.Close(ex)
	ctx.HandleInactive(ex)
}

func (handler IMMessageHandler) HandleException(ctx netty.ExceptionContext, ex netty.Exception) {
	if handler.listener != nil {
		handler.listener.ExceptionCaught(ctx, ex)
	}
	ctx.Close(ex)
	ctx.HandleException(ex)
}
