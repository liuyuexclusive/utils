package main

import (
	"context"
	"fmt"
	"log"

	"github.com/yuexclusive/utils/rpc/example/proto/hello"

	"github.com/yuexclusive/utils/rpc"
)

type handler struct{}

func (h *handler) Send(ctx context.Context, req *hello.Request) (*hello.Response, error) {
	// panic("not implemented") // TODO: Implement
	fmt.Println("call from client")

	return &hello.Response{Res: fmt.Sprintf("hello %s", req.Name)}, nil
}

func main() {
	server, err := rpc.NewServer(
		rpc.Registry("test.srv.hello", []string{"localhost:2379"}, "", 5),
		rpc.TLS("../../../ssl/server_cert.pem", "../../../ssl/server_key.pem"),
		rpc.Auth(),
	)

	if err != nil {
		log.Fatal(err)
	}

	hello.RegisterHelloServer(server.Server, new(handler))

	if err := server.Serve(); err != nil {
		log.Fatal(err)
	}
}
