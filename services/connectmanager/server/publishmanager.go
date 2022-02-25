package server

import (
	"time"

	"github.com/yuwnloyblog/gxgchat/commons/logs"
	"github.com/yuwnloyblog/gxgchat/services/connectmanager/server/codec"
	"github.com/yuwnloyblog/gxgchat/services/connectmanager/server/utils"
)

func PublishMessage(appkey, userid, session string, serverPubMsg *codec.PublishMsgBody, publishType int, callback func()) {
	userCtxMap := GetConnectCtxByUser(appkey, userid)
	if len(userCtxMap) > 0 {
		for kSess, vCtx := range userCtxMap {
			if publishType == 1 && kSess != session { //publishType:1, 只给指定的session发送
				continue
			}
			if publishType == 2 && kSess == session { //publishType:2, 除了指定session以外，给该用户其他登录端发送
				continue
			}
			if vCtx.Channel().IsActive() {
				index := utils.GetServerIndexAfterIncrease(vCtx)
				tmpPubMsg := &codec.PublishMsgBody{
					Index:     int32(index),
					Topic:     serverPubMsg.Topic,
					TargetId:  serverPubMsg.TargetId,
					Timestamp: time.Now().UnixMilli(),
					Data:      serverPubMsg.Data,
				}
				vCtx.Write(tmpPubMsg)
				logs.Info(utils.GetConnSession(vCtx), utils.Action_ServerPub, tmpPubMsg.Index, tmpPubMsg.Topic, len(tmpPubMsg.Data))

			}
		}
	}
}

func PublishQryAckMessage() {

}

func PublishUserPubAckMessage() {

}
