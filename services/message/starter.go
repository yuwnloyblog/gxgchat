package message

import (
	"fmt"

	"github.com/yuwnloyblog/gmicro"
	"github.com/yuwnloyblog/gmicro/actorsystem"
	"github.com/yuwnloyblog/gxgchat/commons/clusters"
	"github.com/yuwnloyblog/gxgchat/services/message/actors"
)

type MessageManager struct{}

func (manager MessageManager) RegisterActors(register gmicro.IActorRegister) {
	register.RegisterActor("pMsg", func() actorsystem.IUntypedActor {
		return clusters.BaseProcessActor(&actors.PrivateMsgActor{})
	}, 64)
}

func (manager MessageManager) Startup(args map[string]interface{}) {
	fmt.Println("Startup message.")
}
func (manager MessageManager) Shutdown() {
	fmt.Println("Shutdown message.")
}
