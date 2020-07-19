package xrpc_pkg

import (
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 调试命令：
// curl localhost:12347/jsonrpc -X POST --data '{"method":"HelloService.Hello","params":["hello"],"id":0}'
func HTTPRPC_Server() {
	rpc.RegisterName("HelloService", new(HelloService))

	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: r.Body,
			Writer:     w,
		}

		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})

	http.ListenAndServe(":12347", nil)
}