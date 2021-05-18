package main

import (
	"context"
	"fmt"
	"log"

	"github.com/yuexclusive/utils/rpc"

	"github.com/yuexclusive/utils/rpc/proto/hello"
)

func main() {

	conn, err := rpc.Dial(
		rpc.Discovery("test.srv.hello", []string{"test.srv.hello"}, []string{"localhost:2379"}),
		rpc.TLSClient("../../ssl/server_ca_cert.pem", "test.example.com"),
		rpc.AuthClient(),
	)
	if err != nil {
		log.Fatal(err)
	}

	client := hello.NewHelloClient(conn)
	res, err := client.Send(context.Background(), &hello.Request{Name: "somebody"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)

}
