package clusters

import (
	"github.com/yuwnloyblog/gmicro/actorsystem"
	"github.com/yuwnloyblog/gxgchat/commons/clusters/ssrequests"
	"github.com/yuwnloyblog/gxgchat/commons/tools"
	"google.golang.org/protobuf/proto"
)

func BaseProcessActor(actor actorsystem.IUntypedActor) actorsystem.IUntypedActor {
	return &baseProcessActor{exeActor: actor}
}

type IContextHandler interface {
	SetContext(ctx BaseContext)
}

type baseProcessActor struct {
	exeActor actorsystem.IUntypedActor
}

type BaseActor struct {
	actorsystem.UntypedActor
	Context BaseContext
}

type BaseContext struct {
	SequenceId    int
	AppKey        string
	ClientOs      string
	DeviceId      string
	ClientAddr    string
	SdkVersion    string
	Qos           int
	Package       string
	SessionId     string
	Method        string
	RequestMethod string
	RequestId     string
	TargetId      string
	TerminalCount int
	PublishType   int
}

func (actor *BaseActor) SetContext(ctx BaseContext) {
	actor.Context = ctx
}

func (actor *baseProcessActor) CreateInputObj() proto.Message {
	return &ssrequests.SSRequest{}
}

func (actor *baseProcessActor) OnReceive(input proto.Message) {
	var err error
	if input != nil {
		ssRequest, ok := input.(*ssrequests.SSRequest)
		if ok {
			ctx := BaseContext{
				SequenceId:    int(ssRequest.SequenceId),
				AppKey:        ssRequest.AppKey,
				ClientOs:      ssRequest.ClientOs,
				DeviceId:      ssRequest.DeviceId,
				ClientAddr:    ssRequest.ClientAddr,
				SdkVersion:    ssRequest.SdkVersion,
				Qos:           int(ssRequest.Qos),
				Package:       ssRequest.Package,
				SessionId:     ssRequest.SessionId,
				Method:        ssRequest.Method,
				RequestMethod: ssRequest.RequestMethod,
				RequestId:     ssRequest.RequestId,
				TargetId:      ssRequest.TargetId,
				TerminalCount: int(ssRequest.TerminalCount),
				PublishType:   int(ssRequest.PublishType),
			}

			ctxHandler, ok := actor.exeActor.(IContextHandler)
			if ok {
				ctxHandler.SetContext(ctx)
			}

			msgBytes := ssRequest.AppMessage
			msg := actor.CreateInputObj()
			err = tools.PbUnMarshal(msgBytes, msg)
			if err == nil {
				receiveHandler, ok := actor.exeActor.(actorsystem.IReceiveHandler)
				if ok {
					receiveHandler.OnReceive(msg)
				}
			}
		}
	}
}

func (actor *baseProcessActor) SetSender(sender actorsystem.ActorRef) {
	senderHandler, ok := actor.exeActor.(actorsystem.ISenderHandler)
	if ok {
		senderHandler.SetSender(sender)
	}
}
func (actor *baseProcessActor) SetSelf(self actorsystem.ActorRef) {
	selfHandler, ok := actor.exeActor.(actorsystem.ISelfHandler)
	if ok {
		selfHandler.SetSelf(self)
	}
}

func (actor *baseProcessActor) OnTimeout() {
	timeoutHandler, ok := actor.exeActor.(actorsystem.ITimeoutHandler)
	if ok {
		timeoutHandler.OnTimeout()
	}
}
