package server

import "sync"

var OnlineUserConnectMap sync.Map    //map[userid]map[session]netty.HandlerContext
var OnlineSessionConnectMap sync.Map // map[session]netty.HandlerContext
