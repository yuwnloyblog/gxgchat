package actors

import (
	"github.com/yuwnloyblog/gxgchat/commons/clusters"
	"google.golang.org/protobuf/proto"
)

type ConnectActor struct {
	clusters.BaseActor
}

func (actor *ConnectActor) OnReceive(input proto.Message) {

}
