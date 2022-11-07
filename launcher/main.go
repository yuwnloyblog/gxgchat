package main

import (
	"sync"

	"github.com/yuwnloyblog/gxgchat/commons/clusters"
	"github.com/yuwnloyblog/gxgchat/commons/configures"
	"github.com/yuwnloyblog/gxgchat/commons/dbs"
	"github.com/yuwnloyblog/gxgchat/commons/imstarters"
	"github.com/yuwnloyblog/gxgchat/commons/logs"
	"github.com/yuwnloyblog/gxgchat/services/chatroom"
	"github.com/yuwnloyblog/gxgchat/services/connectmanager"
	"github.com/yuwnloyblog/gxgchat/services/message"
)

func main() {
	var waitgroup sync.WaitGroup
	waitgroup.Add(1)
	//init configures
	if err := configures.InitConfigures(); err != nil {
		logs.Error("Init Configures failed.", err)
		return
	}
	//init logs
	logs.InitLogs()
	//init mysql
	if err := dbs.InitMysql(); err != nil {
		logs.Error("Init Mysql failed.", err)
		return
	}
	//init cluster
	if err := clusters.InitCluster(); err != nil {
		logs.Error("Init Cluster failed.", err)
		return
	}

	imstarters.Loaded(&connectmanager.ConnectManager{})
	imstarters.Loaded(&message.MessageManager{})
	imstarters.Loaded(&chatroom.ChatRoomManager{})

	imstarters.Startup()

	waitgroup.Wait()
}
