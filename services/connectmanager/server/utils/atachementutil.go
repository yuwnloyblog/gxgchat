package utils

import "github.com/go-netty/go-netty"

func SetContextAttr(ctx netty.HandlerContext, key string, value interface{}) {
	if ctx.Attachment() == nil {
		attMap := make(map[string]interface{})
		ctx.SetAttachment(attMap)
	}
	attMap := ctx.Attachment().(map[string]interface{})
	attMap[key] = value
	ctx.SetAttachment(attMap)
}
func GetContextAttr(ctx netty.HandlerContext, key string) interface{} {
	if ctx.Attachment() != nil {
		attMap := ctx.Attachment().(map[string]interface{})
		return attMap[key]
	}
	return nil
}
func GetContextAttrString(ctx netty.HandlerContext, key string) string {
	ret := GetContextAttr(ctx, key)
	if ret != nil {
		str, ok := ret.(string)
		if ok {
			return str
		}
	}
	return ""
}
func GetConnSession(ctx netty.HandlerContext) string {
	return GetContextAttrString(ctx, StateKey_ConnectSession)
}
func GetServerIndexAfterIncrease(ctx netty.HandlerContext) uint16 {
	var index uint16 = 0
	indexObj := GetContextAttr(ctx, StateKey_ServerMsgIndex)
	if indexObj != nil {
		index = indexObj.(uint16)
	}
	index = index + 1
	SetContextAttr(ctx, StateKey_ServerMsgIndex, index)
	return index
}
