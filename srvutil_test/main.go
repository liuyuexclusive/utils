package main

import (
	"log"

	"github.com/liuyuexclusive/utils/srvutil"
	"github.com/micro/go-micro/v2"
)

type start struct {
}

func (s *start) Start(service micro.Service) {
	service.Options()
}

func main() {
	err := srvutil.Startup("go.micro.srv.test", new(start))
	if err != nil {
		log.Fatal(err)
	}
}
