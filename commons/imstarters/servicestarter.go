package imstarters

import (
	"github.com/yuwnloyblog/gmicro"
	"github.com/yuwnloyblog/gxgchat/commons/clusters"
)

type IServiceStarter interface {
	RegisterActors(register gmicro.IActorRegister)
	Startup(args map[string]interface{})
	Shundown()
}

func Loaded(server IServiceStarter) {
	if server != nil {
		//register actors
		server.RegisterActors(clusters.ImCluster)
		//execute startup
		server.Startup(map[string]interface{}{})
	}
}
