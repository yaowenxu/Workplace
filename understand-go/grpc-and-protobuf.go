package main

import (
	"./xrpc"
	"fmt"
	"time"
)

func main() {
	fmt.Println("RPC")
	//xrpc.Run() // 进行运行xrpc包中的内容
	go xrpc.XrpcServer()
	//xrpc.HTTPRPC_Server()
	for true {
		time.Sleep(1000 * time.Millisecond) // 进行睡眠 1 秒钟，并起一个客户端;
		xrpc.XrpcClient("Hello:")
	}
}