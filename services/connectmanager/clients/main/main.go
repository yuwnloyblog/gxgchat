package main

import (
	"fmt"

	"github.com/yuwnloyblog/gxgchat/services/connectmanager/server/codec"
)

func main() {
	code := [8]byte{1, 2, 3, 3, 2, 1, 5, 75}
	data := []byte("aabbccdd")

	codec.DoObfuscation(code, data)
	fmt.Println(string(data))

	codec.DoObfuscation(code, data)
	fmt.Println(string(data))
}
