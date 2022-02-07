package clusters

import (
	"strings"

	"github.com/yuwnloyblog/gmicro"
	"github.com/yuwnloyblog/gmicro/actorsystem"
	"github.com/yuwnloyblog/gxgchat/commons/configures"
)

const (
	Mod_Signal  = "signal"
	Mod_Cluster = "cluster"
)

var Cluster *ImCluster

func InitCluster() error {
	clusterMod := configures.Config.ClusterMod
	if clusterMod == Mod_Cluster {
		zkAddress := strings.Split(configures.Config.Zookeeper.Address, ",")
		Cluster = &ImCluster{
			Cluster: gmicro.NewCluster(configures.Config.ClusterName, configures.Config.NodeName, configures.Config.RpcHost, configures.Config.RpcPort, zkAddress),
		}
	} else {
		Cluster = &ImCluster{
			Cluster: gmicro.NewSingleCluster(configures.Config.ClusterName, configures.Config.NodeName),
		}
	}
	return nil
}

type ImCluster struct {
	Cluster *gmicro.Cluster
}

func (cluster *ImCluster) RegisterActor(method string, actorCreator func() actorsystem.IUntypedActor, concurrentCount int) {
	cluster.Cluster.RegisterActor(method, func() actorsystem.IUntypedActor {
		return BaseProcessActor(actorCreator())
	}, concurrentCount)
}

func (cluster *ImCluster) StartUp() {
	if cluster.Cluster != nil {
		cluster.Cluster.StartUp()
	}
}
