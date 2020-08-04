package main

import (
	"fmt"
	"time"
)
import "main/heartbeat-golang" // 注意，引入的是包的相对路径；

func main() {
	fmt.Println("heartbeat!")
	heartbeat.CommitHash = "yaoxu"
	go heartbeat.RunHeartbeatService(":10101")
	for true {
		time.Sleep(time.Second*1) // 客户端每间隔一秒钟进行轮询一次消息
		msg, err := heartbeat.Get("http://localhost:10101/heartbeat")
		if  err != nil{
			fmt.Println(err)
			continue
		}
		fmt.Println("客户端接收到的消息：" , msg)
	}
}