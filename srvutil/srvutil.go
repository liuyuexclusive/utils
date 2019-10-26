package srvutil

import (
	"time"

	"github.com/liuyuexclusive/utils/configutil"
	"github.com/liuyuexclusive/utils/logutil"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/broker/nats"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/sirupsen/logrus"
)

type Starter interface {
	Start(s micro.Service)
}

func Startup(name string, starter Starter) {
	logutil.LogToElastic(name)

	// New Service
	service := micro.NewService(
		micro.Name(name),
		micro.Version("latest"),
		micro.Registry(etcd.NewRegistry(registry.Addrs(configutil.MustGet().ETCDAddress))),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
		micro.Broker(nats.NewBroker(broker.Addrs(configutil.MustGet().NatsAddress))),
	)

	// Initialise service
	service.Init()

	starter.Start(service)

	// Run service
	if err := service.Run(); err != nil {
		logrus.Fatal(err)
	}
}
