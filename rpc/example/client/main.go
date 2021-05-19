package main

import (
	"context"
	"fmt"
	"log"

	"github.com/yuexclusive/utils/config"
	"github.com/yuexclusive/utils/rpc"
	"github.com/yuexclusive/utils/srv/auth/proto/auth"
	"github.com/yuexclusive/utils/srv/basic/proto/user"
)

func main() {
	cfg := config.MustGet()

	conn, err := rpc.DialByName(cfg.AuthServiceName)

	if err != nil {
		log.Fatal(err)
	}

	// client := hello.NewHelloClient(conn)
	// res, err := client.Send(context.Background(), &hello.Request{Name: "somebody"})
	ac := auth.NewAuthClient(conn)

	r1, err := ac.Auth(context.Background(), &auth.AuthRequest{Id: "super_admin", Key: "123"})
	if err != nil {
		log.Fatal(err)
	}

	conn2, err := rpc.DialByNameWithAuth("srv.basic", r1.Token)

	if err != nil {
		log.Fatal(err)
	}

	client := user.NewUserClient(conn2)
	res, err := client.Get(context.Background(), &user.GetRequest{Name: "super_admin"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)

}
