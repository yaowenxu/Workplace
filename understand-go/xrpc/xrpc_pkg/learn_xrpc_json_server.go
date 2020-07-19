package xrpc_pkg

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"
)

const XHelloServiceName = "XHelloService"

type XHelloServiceInterface = interface {
	XHello(request string, reply *string) error
}

type XHelloService struct {

}

// 进行回答
func (p *XHelloService) XHello(request string, reply *string) error {
	*reply = "hello: " + request
	fmt.Println(time.Now().String() + " 一次Hello服务结束")
	return nil
}

func RegisterXHelloService(svc XHelloServiceInterface) error {
	return rpc.RegisterName(XHelloServiceName, svc)
}

// 基于json编码 rpc 服务,
// 使用 echo -e '{"method":"HelloService.Hello","params":["hello"],"id":1}' | nc localhost 1234 调试
func XrpcServer() {
	fmt.Println("启动xrpcserver线程")
	RegisterXHelloService(new(XHelloService)) // 进行注册RPC服务
	listener, err := net.Listen("tcp", ":12346")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	for true { // 通过循环支持多个链接，然后每个循环支持特定的服务;
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
