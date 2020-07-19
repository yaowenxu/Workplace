package xrpc

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 使用 nc -l 1234 调试
func XrpcClient(str string) {
	conn, err := net.Dial("tcp", "localhost:12346")
	if err != nil {
		log.Fatal("net.Dial:", err)
	}

	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	var reply string
	err = client.Call("XHelloService.XHello", "hello", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}

//func XrpcClient(str string) {
//	client, err := XDialHelloService("tcp", "127.0.0.1:12346")
//	if err != nil {
//		log.Fatal("dialing: ", err)
//	}
//	var reply string
//	err = client.Hello(str+time.Now().String(), &reply)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(reply) //进行输出返回值
//	client.Close()     // 需要进行关闭连接;
//}
//
//type XHelloServiceClient struct {
//	*rpc.Client
//}
//
//func (p *XHelloServiceClient) Hello(request string, reply *string) error {
//	return p.Client.Call(XHelloServiceName+".XHello", request, reply)
//}
//
//func XDialHelloService(network, address string) (*XHelloServiceClient, error) {
//	c, err := rpc.Dial(network, address)
//	if err != nil {
//		return nil, err
//	}
//	return &XHelloServiceClient{Client: c}, nil
//}