package codec

import (
	"bytes"
	"fmt"

	"github.com/go-netty/go-netty"
	"github.com/go-netty/go-netty/utils"
	"github.com/yuwnloyblog/gxgchat/commons/tools"
)

type IMCodecHandler struct{}

func (*IMCodecHandler) CodecName() string {
	return "im-codec"
}

func (IMCodecHandler) HandleRead(ctx netty.InboundContext, message netty.Message) {
	reader := utils.MustToReader(message)
	var imMsg IMessage
	tmpBs := make([]byte, 1)
	reader.Read(tmpBs)
	version := tmpBs[0]
	if version == Version_0 {
		msgHeader := &MsgHeader{Version: Version_0}
		msgHeader.DecodeHeader(reader)
		var msgBodyBytes []byte
		if msgHeader.MsgBodySize > 0 {
			msgBodyBytes = make([]byte, msgHeader.MsgBodySize)
			reader.Read(msgBodyBytes)
		} else {
			msgBodyBytes = []byte{}
		}

		//validate checksum  TODO
		switch msgHeader.GetCmd() {
		case Cmd_Connect:
			imMsg = NewConnectMessage(msgHeader)
		case Cmd_Disconnect:
			imMsg = NewDisconnectMessage(msgHeader)
		case Cmd_Ping:
			imMsg = NewPingMessage(msgHeader)
		case Cmd_Publish:
			imMsg = NewUserPublishMessage(msgHeader)
		case Cmd_PublishAck:
			imMsg = NewServerPublishAckMessage(msgHeader)
		case Cmd_Query:
			imMsg = NewQueryMessage(msgHeader)
		case Cmd_QueryConfirm:
			imMsg = NewQueryConfirmMessage(msgHeader)
		default:
			return
		}
		var obfuscationCode []byte
		if msgHeader.GetCmd() == Cmd_Connect {
			obfuscationCode = calObfuscationCode(msgBodyBytes)
			SetContextAttr(ctx, StateKey_ObfuscationCode, obfuscationCode)
		} else {
			obfuscationCode = getObfuscationCodeFromCtx(ctx)
		}
		doObfuscation(obfuscationCode, msgBodyBytes)
		imMsg.DecodeBody(msgBodyBytes)
	}
	ctx.HandleRead(imMsg)
}
func (IMCodecHandler) HandleWrite(ctx netty.OutboundContext, message netty.Message) {
	switch s := message.(type) {
	case IMessage:
		msgBody, err := s.EncodeBody()
		if err == nil {
			obfuscationCode := getObfuscationCodeFromCtx(ctx)
			doObfuscation(obfuscationCode, msgBody)
			buf := bytes.NewBuffer([]byte{})
			s.EncodeHeader(buf, msgBody) //encode header
			if len(msgBody) > 0 {
				buf.Write(msgBody) //write msg body
			}
			ctx.HandleWrite(buf)
		} else {
			ctx.HandleWrite(message)
		}
	default:
		fmt.Println(s)
		ctx.HandleWrite(message)
	}
}

func getObfuscationCodeFromCtx(ctx netty.HandlerContext) []byte {
	obfuscationCodeObj := GetContextAttr(ctx, StateKey_ObfuscationCode)
	if obfuscationCodeObj != nil {
		obfuscationCode := obfuscationCodeObj.([]byte)
		return obfuscationCode
	}
	return []byte{}
}

var StateKey_ObfuscationCode string = "state.connect_session"

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

var fixedConnMsgBytes []byte

func getFixedConnMsgBytes() []byte {
	if len(fixedConnMsgBytes) != 8 {
		connMsg := &ConnectMsgBody{
			ProtoId: "IamGxg",
		}
		bs, err := tools.PbMarshal(connMsg)
		if err == nil {
			fixedConnMsgBytes = bs[:8]
		}
	}
	return fixedConnMsgBytes
}
func calObfuscationCode(connectData []byte) []byte {
	fixedConnBytes := getFixedConnMsgBytes()
	if len(connectData) > 8 && len(fixedConnBytes) == 8 {
		code := make([]byte, 8)
		for i := 0; i < 8; i++ {
			code[i] = connectData[i] ^ fixedConnBytes[i]
		}
		return code
	}
	return []byte{}
}

func doObfuscation(code, data []byte) {
	dataLen := len(data)
	if dataLen > 0 && len(code) == 8 {
		for i := 0; i < dataLen; i += 8 {
			for j := 0; (j < 8) && (i+j < dataLen); j++ {
				data[i+j] = data[i+j] ^ code[j]
			}
		}
	}
}
