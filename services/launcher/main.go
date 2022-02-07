package main

import (
	"github.com/yuwnloyblog/gxgchat/commons/clusters"
	"github.com/yuwnloyblog/gxgchat/commons/imstarters"
	"github.com/yuwnloyblog/gxgchat/services/connectmanager"
)

func main() {
	imstarters.Loaded(connectmanager.ConnectManager{})

	clusters.ImCluster.StartUp()
}
