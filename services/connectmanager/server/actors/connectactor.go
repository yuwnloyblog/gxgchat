package actors

import (
	"fmt"

	"github.com/yuwnloyblog/gmicro/utils"
	"github.com/yuwnloyblog/gxgchat/commons/baseactors"
	"google.golang.org/protobuf/proto"
)

type ConnectActor struct {
	baseactors.BaseActor
}

func (actor *ConnectActor) OnReceive(input proto.Message) {
	stu := &utils.Student{}
	err := actor.HandleInput(input, stu)
	if err == nil {
		fmt.Println(stu.Name)
	}
}
