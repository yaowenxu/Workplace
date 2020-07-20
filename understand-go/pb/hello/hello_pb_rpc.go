package hello

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"time"
)

func Run() int {
	go rpcServer()
	for true {
		time.Sleep(1000 * time.Millisecond) // 进行睡眠 1 秒钟，并起一个客户端;
		rpcClient("Hello:")
	}
	return 0
}

func rpcClient(str string) {
	client, err := DialHelloService("tcp", "127.0.0.1:12345")
	if err != nil {
		log.Fatal("dialing: ", err)
	}
	var tmpreq String
	tmpreq.Value = str+time.Now().String()
	var tmprep String
	err = client.Hello(&tmpreq, &tmprep)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tmprep.GetValue()) //进行输出返回值
	client.Close()     // 需要进行关闭连接;
}

type HelloServiceClient struct {
	*rpc.Client
}

func (p *HelloServiceClient) Hello(request *String, reply *String) error {
	return p.Client.Call(HelloServiceName+".Hello", &request, reply)
}

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: c}, nil
}

const HelloServiceName = "HelloService"

type HelloServiceInterface = interface {
	Hello(request *String, reply *String) error
}

type HelloService struct {
}

// 进行回答
func (p *HelloService) Hello(request *String, reply *String) error {
	reply.Value = "hello: " + request.GetValue()
	fmt.Println(time.Now().String() + " 一次Hello服务结束")
	return nil
}

func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}

func rpcServer() {
	fmt.Println("启动rpcserver线程")
	RegisterHelloService(new(HelloService)) // 进行注册RPC服务
	listener, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	for true { // 通过循环支持多个链接，然后每个循环支持特定的服务;
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		go rpc.ServeConn(conn) // rpc进行服务rpc请求
	}
}
