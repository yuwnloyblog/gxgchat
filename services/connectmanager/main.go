package main

import (
	"bytes"
	"fmt"
	"time"

	"github.com/go-netty/go-netty"
	"github.com/go-netty/go-netty/utils"
	"github.com/yuwnloyblog/gxgchat/services/connectmanager/server/codec"
)

/*
int first = buf.readByte();
        int second = buf.readByte();
        int digit;
        int code = first;
        int msgLength = 0;
        int multiplier = 1;
        int lengthSize = 0;
        do {
            lengthSize++;
            digit = buf.readByte();
            code = code ^ digit;
            msgLength += (digit & 0x7f) * multiplier;
            multiplier *= 128;
            if ((digit & 0x80) > 0 && !buf.isReadable()) {
                resumeTimer(ctx);
                buf.resetReaderIndex();
                return null;
            }
        } while ((digit & 0x80) > 0);
        if (code != second) {
            close(ctx, buf);
            return null;
        }
*/

func main() {
	//	TestNetty()

	buf := bytes.NewBuffer([]byte{})
	buf.WriteByte(byte(2))

	cou, err := buf.WriteString("aabb")

	fmt.Println(cou, ":", err)

	fmt.Println(len(buf.Bytes()))

	//reader := bytes.NewReader(buf.Bytes())

	bx := byte(255)

	fmt.Println(bx)
	// ab := 4

	// xx := ab & 0x03

	// fmt.Println(xx)
}

func TestNetty() {
	// setup child pipeline initializer.
	childInitializer := func(channel netty.Channel) {
		channel.Pipeline().
			// AddLast(frame.LengthFieldCodec(binary.LittleEndian, 1024, 0, 2, 0, 2)).
			// AddLast(format.TextCodec()).
			AddLast(&codec.IMCodecHandler{}).
			AddLast(EchoHandler{
				role: "Server",
			})
	}

	// setup client pipeline initializer.
	clientInitializer := func(channel netty.Channel) {
		channel.Pipeline().
			// AddLast(frame.LengthFieldCodec(binary.LittleEndian, 1024, 0, 2, 0, 2)).
			// AddLast(format.TextCodec()).
			AddLast(&codec.IMCodecHandler{}).
			AddLast(EchoHandler{
				role: "Client",
				flag: true,
			})
	}

	// new bootstrap
	var bootstrap = netty.NewBootstrap(netty.WithChildInitializer(childInitializer), netty.WithClientInitializer(clientInitializer))

	// connect to the server after 1 second
	time.AfterFunc(time.Second, func() {
		_, err := bootstrap.Connect("127.0.0.1:6565", nil)
		utils.Assert(err)

	})

	// setup bootstrap & startup server.
	bootstrap.Listen("0.0.0.0:6565").Sync()
}

type EchoHandler struct {
	role string
	flag bool
}

func (l EchoHandler) HandleActive(ctx netty.ActiveContext) {
	if l.flag {
		fmt.Println(l.role, "->", "active:", ctx.Channel().RemoteAddr())
		msg := &codec.BaseMessage{
			Name: l.role,
			Age:  12,
		}
		ctx.Write(msg)
	}

	//ctx.Write("Hello I'm " + l.role)

	ctx.HandleActive()
}

func (l EchoHandler) HandleRead(ctx netty.InboundContext, message netty.Message) {
	if !l.flag {
		fmt.Println("atachemetn:", ctx.Attachment())
		fmt.Println(l.role, "->", "handle read:", message)
		m, ok := message.(*codec.BaseMessage)
		if ok {
			fmt.Println("name:", m.Name, "\tage:", m.Age)
		} else {
			fmt.Println("xxxxxxx")
		}
	}

	ctx.HandleRead(message)
}

func (l EchoHandler) HandleInactive(ctx netty.InactiveContext, ex netty.Exception) {
	if l.flag {
		fmt.Println(l.role, "->", "inactive:", ctx.Channel().RemoteAddr(), ex)
	}
	ctx.HandleInactive(ex)
}
