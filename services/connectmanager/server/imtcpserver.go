package server

import (
	"fmt"

	"github.com/go-netty/go-netty"
	"github.com/yuwnloyblog/gxgchat/services/connectmanager/server/codec"
)

const (
	StateKey_ConnectSession    string = "state.connect_session"
	StateKey_ConnectCreateTime string = "state.connect_timestamp"
	StateKey_ServerMsgIndex    string = "state.server_msg_index"
	StateKey_ClientMsgIndex    string = "state.client_msg_index"
	StateKey_AppId             string = "state.appid"
	StateKey_UserID            string = "state.userid"
	StateKey_Platform          string = "state.platform"
	StateKey_Version           string = "state.version"
)

type ImTcpServer struct {
	MessageListener ImListener
	bootstrap       netty.Bootstrap
}

func (server *ImTcpServer) SyncStart(port int) {
	childInitializer := func(channel netty.Channel) {
		channel.Pipeline().
			AddLast(codec.IMCodecHandler{}).
			AddLast(IMMessageHandler{server.MessageListener})
	}

	// new bootstrap
	server.bootstrap = netty.NewBootstrap(netty.WithChildInitializer(childInitializer))

	// setup bootstrap & startup server.
	server.bootstrap.Listen(fmt.Sprintf("0.0.0.0:%d", port)).Sync()
}

func (server *ImTcpServer) Stop() {
	if server.bootstrap != nil {
		server.bootstrap.Shutdown()
	}
}
