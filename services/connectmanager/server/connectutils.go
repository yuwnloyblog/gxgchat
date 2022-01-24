package server

import (
	"github.com/go-netty/go-netty"
	"github.com/yuwnloyblog/gxgchat/services/connectmanager/server/codec"
)

type ImListener interface {
	Create(ctx netty.InboundContext)
	Close(ctx netty.InboundContext)

	Connected(msg *codec.ConnectMessage, ctx netty.InboundContext)
	Diconnected(msg *codec.DisconnectMessage, ctx netty.InboundContext)
	PublishArrived(msg codec.UserPublishMessage, ctx netty.InboundContext)
	PubAckArrived(msg codec.ServerPublishAckMessage, ctx netty.InboundContext)
	QueryArrived(msg codec.QueryMessage, ctx netty.InboundContext)
	QueryConfirmArrived(msg codec.QueryConfirmMessage, ctx netty.InboundContext)
	PingArrived(msg codec.PingMessage, ctx netty.InboundContext)
}
