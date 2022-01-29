package server

import "sync"

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

var OnlineUserConnectMap sync.Map    //map[userid]map[session]netty.HandlerContext
var OnlineSessionConnectMap sync.Map // map[session]netty.HandlerContext
