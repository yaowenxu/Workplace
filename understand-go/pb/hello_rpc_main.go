package main

import (
	"main/hello"
	"fmt"
)

// 完成初步的调用；
func main() {
	fmt.Println("RPC")
	hello.Run() // 进行运行xrpc包中的内容
}