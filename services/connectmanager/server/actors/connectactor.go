package actors

import (
	"github.com/yuwnloyblog/gxgchat/commons/baseactors"
	"google.golang.org/protobuf/proto"
)

type ConnectActor struct {
	baseactors.BaseActor
}

func (actor *ConnectActor) OnReceive(input proto.Message) {

}
