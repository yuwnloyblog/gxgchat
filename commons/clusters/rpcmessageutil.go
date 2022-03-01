package clusters

import (
	"github.com/yuwnloyblog/gxgchat/commons/pbdefines/pbobjs"
)

func CreateUserPubAckWraper(code int, msgId string, msgSendTime int64, ctx BaseContext) *pbobjs.RpcMessageWraper {
	userPubAck := &pbobjs.RpcMessageWraper{
		RpcMsgType:  pbobjs.RpcMsgType_UserPubAck,
		ResultCode:  int32(code),
		MsgId:       msgId,
		MsgSendTime: msgSendTime,
	}
	handleBaseContext(userPubAck, ctx)
	return userPubAck
}

func handleBaseContext(rpcMsg *pbobjs.RpcMessageWraper, ctx BaseContext) {
	rpcMsg.ReqIndex = int32(ctx.SeqIndex)
	rpcMsg.AppKey = ctx.AppKey
	rpcMsg.ClientOs = ctx.ClientOs
	rpcMsg.DeviceId = ctx.DeviceId
	rpcMsg.ClientAddress = ctx.ClientAddr
	rpcMsg.SdkVersion = ctx.SdkVersion
	rpcMsg.Qos = int32(ctx.Qos)
	rpcMsg.PackageName = ctx.Package
	rpcMsg.Session = ctx.Session
	rpcMsg.Method = ctx.Method
	rpcMsg.SourceMethod = ctx.SourceMethod
	rpcMsg.RequesterId = ctx.RequesterId
	rpcMsg.TargetId = ctx.TargetId
	rpcMsg.TerminalCount = int32(ctx.TerminalCount)
	rpcMsg.PublishType = int32(ctx.PublishType)
}
