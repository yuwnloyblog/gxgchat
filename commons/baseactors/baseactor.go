package baseactors

import (
	"errors"

	"github.com/yuwnloyblog/gmicro/actorsystem"
	"github.com/yuwnloyblog/gxgchat/commons/baseactors/ssrequests"
	"github.com/yuwnloyblog/gxgchat/commons/tools"
	"google.golang.org/protobuf/proto"
)

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

func (actor *BaseActor) CreateInputObj() proto.Message {
	return &ssrequests.SSRequest{}
}

func (actor *BaseActor) HandleInput(input, msg proto.Message) error {
	var err error
	if input != nil {
		ssRequest, ok := input.(*ssrequests.SSRequest)
		if ok {
			actor.Context = BaseContext{
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
			msgBytes := ssRequest.AppMessage
			err = tools.PbUnMarshal(msgBytes, msg)
		} else {
			err = errors.New("Not SSRequest!!!")
		}
	}
	return err
}
