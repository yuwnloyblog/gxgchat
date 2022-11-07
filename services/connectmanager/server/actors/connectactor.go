package actors

import (
	"github.com/yuwnloyblog/gmicro/actorsystem"
	"github.com/yuwnloyblog/gxgchat/commons/pbdefines/pbobjs"
	"github.com/yuwnloyblog/gxgchat/services/connectmanager/server/codec"
	"github.com/yuwnloyblog/gxgchat/services/connectmanager/server/managers"
	"google.golang.org/protobuf/proto"
)

type ConnectActor struct {
	actorsystem.UntypedActor
}

func (actor *ConnectActor) OnReceive(input proto.Message) {
	if rpcMsg, ok := input.(*pbobjs.RpcMessageWraper); ok {
		if rpcMsg.RpcMsgType == pbobjs.RpcMsgType_UserPubAck {
			managers.PublishUserPubAckMessage(rpcMsg.AppKey, rpcMsg.RequesterId, rpcMsg.Session, &codec.PublishAckMsgBody{
				Index:     rpcMsg.ReqIndex,
				Code:      rpcMsg.ResultCode,
				MsgId:     rpcMsg.MsgId,
				Timestamp: rpcMsg.MsgSendTime,
			})
		} else if rpcMsg.RpcMsgType == pbobjs.RpcMsgType_QueryAck {
			var callback func()
			var timeoutCallback func()
			if int(rpcMsg.Qos) == codec.QoS_NeedAck || actor.Sender != actorsystem.NoSender {
				callback = func() {}
				timeoutCallback = func() {}
			}
			managers.PublishQryAckMessage(rpcMsg.Session, &codec.QueryAckMsgBody{
				Index:     rpcMsg.ReqIndex,
				Code:      rpcMsg.ResultCode,
				Timestamp: rpcMsg.MsgSendTime,
				Data:      rpcMsg.AppDataBytes,
			}, callback, timeoutCallback)
		} else if rpcMsg.RpcMsgType == pbobjs.RpcMsgType_ServerPub {
			var callback func()
			var timeoutCallback func()
			if int(rpcMsg.Qos) == codec.QoS_NeedAck || actor.Sender != actorsystem.NoSender {
				callback = func() {}
				timeoutCallback = func() {}
			}
			managers.PublishServerPubMessage(rpcMsg.AppKey, rpcMsg.TargetId, rpcMsg.Session, &codec.PublishMsgBody{
				Topic:     rpcMsg.Method,
				TargetId:  rpcMsg.TargetId,
				Timestamp: rpcMsg.MsgSendTime,
			}, int(rpcMsg.PublishType), callback, timeoutCallback)
		}
	}
}

func (actor *ConnectActor) CreateInputObj() proto.Message {
	return &pbobjs.RpcMessageWraper{}
}
