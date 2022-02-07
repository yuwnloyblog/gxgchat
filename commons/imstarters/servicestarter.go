package imstarters

import (
	"github.com/yuwnloyblog/gmicro"
	"github.com/yuwnloyblog/gxgchat/commons/clusters"
)

type IServiceStarter interface{}

type IRegisterActorsHandler interface {
	RegisterActors(register gmicro.IActorRegister)
}

type IStartupHandler interface {
	Startup(args map[string]interface{})
}
type IShundownHandler interface {
	Shundown()
}

var serverList []IServiceStarter

func Loaded(server IServiceStarter) {
	if server != nil {
		//register actors
		if regActorsHandler, ok := server.(IRegisterActorsHandler); ok {
			regActorsHandler.RegisterActors(clusters.Cluster)
		}
		serverList = append(serverList, server)
	}
}

func Startup() {
	for _, server := range serverList {
		//execute startup
		if startHandler, ok := server.(IStartupHandler); ok {
			startHandler.Startup(map[string]interface{}{})
		}
	}
	clusters.Cluster.StartUp()
}

func Shundown() {
	for _, server := range serverList {
		//execute startup
		if shundownHandler, ok := server.(IShundownHandler); ok {
			shundownHandler.Shundown()
		}
	}
}
