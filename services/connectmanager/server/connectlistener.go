package server

import (
	"time"

	"github.com/go-netty/go-netty"
	"github.com/yuwnloyblog/gxgchat/commons/tools"
	"github.com/yuwnloyblog/gxgchat/services/connectmanager/server/codec"
)

type ImListener interface {
	Create(ctx netty.ActiveContext)
	Close(ctx netty.InactiveContext)
	ExceptionCaught(ctx netty.ExceptionContext, ex netty.Exception)

	Connected(msg *codec.ConnectMessage, ctx netty.InboundContext)
	Diconnected(msg *codec.DisconnectMessage, ctx netty.InboundContext)
	PublishArrived(msg *codec.UserPublishMessage, ctx netty.InboundContext)
	PubAckArrived(msg *codec.ServerPublishAckMessage, ctx netty.InboundContext)
	QueryArrived(msg *codec.QueryMessage, ctx netty.InboundContext)
	QueryConfirmArrived(msg *codec.QueryConfirmMessage, ctx netty.InboundContext)
	PingArrived(msg *codec.PingMessage, ctx netty.InboundContext)
}

func SetContextAttr(ctx netty.HandlerContext, key string, value interface{}) {
	attMap := make(map[string]interface{})
	if ctx.Attachment() != nil {
		attMap = ctx.Attachment().(map[string]interface{})
	}
	attMap[key] = value
	ctx.SetAttachment(attMap)
}
func GetContextAttr(ctx netty.HandlerContext, key string) interface{} {
	if ctx.Attachment() != nil {
		attMap := ctx.Attachment().(map[string]interface{})
		return attMap[key]
	}
	return nil
}

type ConnectListener struct{}

func (ConnectListener) Create(ctx netty.ActiveContext) {
	SetContextAttr(ctx, StateKey_ConnectSession, tools.GenerateUUIDShortString())
	SetContextAttr(ctx, StateKey_ConnectCreateTime, time.Now().UnixMilli())
}
func (ConnectListener) Close(ctx netty.InactiveContext) {

}
func (ConnectListener) ExceptionCaught(ctx netty.ExceptionContext, ex netty.Exception)             {}
func (ConnectListener) Connected(msg *codec.ConnectMessage, ctx netty.InboundContext)              {}
func (ConnectListener) Diconnected(msg *codec.DisconnectMessage, ctx netty.InboundContext)         {}
func (ConnectListener) PublishArrived(msg *codec.UserPublishMessage, ctx netty.InboundContext)     {}
func (ConnectListener) PubAckArrived(msg *codec.ServerPublishAckMessage, ctx netty.InboundContext) {}
func (ConnectListener) QueryArrived(msg *codec.QueryMessage, ctx netty.InboundContext)             {}
func (ConnectListener) QueryConfirmArrived(msg *codec.QueryConfirmMessage, ctx netty.InboundContext) {
}
func (ConnectListener) PingArrived(msg *codec.PingMessage, ctx netty.InboundContext) {}
