package clients

import (
	"github.com/go-netty/go-netty"
	"github.com/yuwnloyblog/gxgchat/services/connectmanager/server/codec"
)

type ImClient struct {
	Address string
	Token   string
	Appkey  string

	UserId string

	channel     netty.Channel
	isConnected bool
}

func (client *ImClient) Connect() {
	if !client.isConnected {
		initializer := func(channel netty.Channel) {
			channel.Pipeline().
				AddLast(ImClientCodecHandler{}).
				AddLast(ImClientMessageHandler{})
		}
		bootstrap := netty.NewBootstrap(netty.WithChildInitializer(initializer))
		var err error
		client.channel, err = bootstrap.Connect(client.Address, nil)
		if err == nil {
			connectMsg := codec.NewConnectMessage(&codec.ConnectMsgBody{
				ProtoId:    codec.ProtoId,
				SdkVersion: "1.0.1",
				AppId:      client.Appkey,
				Token:      client.Token,
			})
			client.channel.Write(connectMsg)
		}
	}
}

func (client *ImClient) Reconnect() {

}
