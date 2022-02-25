package connectmanager

import (
	"fmt"

	"github.com/yuwnloyblog/gmicro"
	"github.com/yuwnloyblog/gmicro/actorsystem"
	"github.com/yuwnloyblog/gxgchat/commons/clusters"
	"github.com/yuwnloyblog/gxgchat/commons/configures"
	"github.com/yuwnloyblog/gxgchat/services/connectmanager/server"
	"github.com/yuwnloyblog/gxgchat/services/connectmanager/server/actors"
)

type ConnectManager struct {
	tcpServer *server.ImServer
	wsServer  *server.ImWebsocketServer
}

func (ser *ConnectManager) RegisterActors(register gmicro.IActorRegister) {
	register.RegisterActor("connect", func() actorsystem.IUntypedActor {
		return clusters.BaseProcessActor(&actors.ConnectActor{})
	}, 64)
}
func (ser *ConnectManager) Startup(args map[string]interface{}) {
	tcpPort := configures.Config.ConnectManager.TcpPort
	wsPort := configures.Config.ConnectManager.WsPort
	ser.tcpServer = &server.ImServer{
		MessageListener: &server.ImListenerImpl{},
	}
	go ser.tcpServer.SyncStart(tcpPort)
	fmt.Println("start with tcp port :", tcpPort)
	ser.wsServer = &server.ImWebsocketServer{
		MessageListener: &server.ImListenerImpl{},
	}
	go ser.wsServer.SyncStart(wsPort)
	fmt.Println("start with ws port :", wsPort)
}

func (ser *ConnectManager) Shutdown() {
	if ser.tcpServer != nil {
		ser.tcpServer.Stop()
	}
	if ser.wsServer != nil {
		ser.wsServer.Stop()
	}
}
