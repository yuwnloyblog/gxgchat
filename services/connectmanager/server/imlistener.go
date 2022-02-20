package server

import (
	"time"

	"github.com/go-netty/go-netty"
	"github.com/yuwnloyblog/gxgchat/commons/tools"
	"github.com/yuwnloyblog/gxgchat/services/connectmanager/server/codec"
)

// type ImListener interface {
// 	Create(ctx netty.ActiveContext)
// 	Close(ctx netty.InactiveContext)
// 	ExceptionCaught(ctx netty.ExceptionContext, ex netty.Exception)

// 	Connected(msg *codec.ConnectMessage, ctx netty.InboundContext)
// 	Diconnected(msg *codec.DisconnectMessage, ctx netty.InboundContext)
// 	PublishArrived(msg *codec.UserPublishMessage, ctx netty.InboundContext)
// 	PubAckArrived(msg *codec.ServerPublishAckMessage, ctx netty.InboundContext)
// 	QueryArrived(msg *codec.QueryMessage, ctx netty.InboundContext)
// 	QueryConfirmArrived(msg *codec.QueryConfirmMessage, ctx netty.InboundContext)
// 	PingArrived(msg *codec.PingMessage, ctx netty.InboundContext)
// }

type ImListener interface {
	Create(ctx netty.ActiveContext)
	Close(ctx netty.InactiveContext)
	ExceptionCaught(ctx netty.ExceptionContext, ex netty.Exception)

	Connected(msg *codec.ConnectMsgBody, ctx netty.InboundContext)
	Diconnected(msg *codec.DisconnectMsgBody, ctx netty.InboundContext)
	PublishArrived(msg *codec.PublishMsgBody, ctx netty.InboundContext)
	PubAckArrived(msg *codec.PublishAckMsgBody, ctx netty.InboundContext)
	QueryArrived(msg *codec.QueryMsgBody, ctx netty.InboundContext)
	QueryConfirmArrived(msg *codec.QueryConfirmMsgBody, ctx netty.InboundContext)
	PingArrived(ctx netty.InboundContext)
}

type ImListenerImpl struct{}

func (*ImListenerImpl) Create(ctx netty.ActiveContext) {
	codec.SetContextAttr(ctx, StateKey_ConnectSession, tools.GenerateUUIDShortString())
	codec.SetContextAttr(ctx, StateKey_ConnectCreateTime, time.Now().UnixMilli())
}
func (*ImListenerImpl) Close(ctx netty.InactiveContext) {

}
func (*ImListenerImpl) ExceptionCaught(ctx netty.ExceptionContext, ex netty.Exception)       {}
func (*ImListenerImpl) Connected(msg *codec.ConnectMsgBody, ctx netty.InboundContext)        {}
func (*ImListenerImpl) Diconnected(msg *codec.DisconnectMsgBody, ctx netty.InboundContext)   {}
func (*ImListenerImpl) PublishArrived(msg *codec.PublishMsgBody, ctx netty.InboundContext)   {}
func (*ImListenerImpl) PubAckArrived(msg *codec.PublishAckMsgBody, ctx netty.InboundContext) {}
func (*ImListenerImpl) QueryArrived(msg *codec.QueryMsgBody, ctx netty.InboundContext)       {}
func (*ImListenerImpl) QueryConfirmArrived(msg *codec.QueryConfirmMsgBody, ctx netty.InboundContext) {
}
func (*ImListenerImpl) PingArrived(ctx netty.InboundContext) {}
