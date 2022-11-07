package chatroom

import (
	"fmt"

	"github.com/yuwnloyblog/gmicro"
	"github.com/yuwnloyblog/gmicro/actorsystem"
	"github.com/yuwnloyblog/gxgchat/commons/clusters"
	"github.com/yuwnloyblog/gxgchat/services/message/actors"
)

type ChatRoomManager struct{}

func (manager *ChatRoomManager) RegisterActors(register gmicro.IActorRegister) {
	register.RegisterActor("cMsg", func() actorsystem.IUntypedActor {
		return clusters.BaseProcessActor(&actors.PrivateMsgActor{})
	}, 64)
}

func (manager *ChatRoomManager) Startup(args map[string]interface{}) {
	fmt.Println("Startup chatroom.")
}
func (manager *ChatRoomManager) Shutdown() {
	fmt.Println("Shutdown chatroom.")
}
