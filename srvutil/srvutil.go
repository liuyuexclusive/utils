package srvutil

import (
	"time"

	"github.com/liuyuexclusive/utils/appconfigutil"
	"github.com/liuyuexclusive/utils/logutil"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/broker/nats"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/sirupsen/logrus"
)

type Options struct {
	// 是否记录日志到ES 默认为false
	IsLogToES bool
}

type Option func(ops *Options)

type Starter interface {
	Start(s micro.Service)
}

func Startup(name string, starter Starter, opts ...Option) {
	options := &Options{
		IsLogToES: false,
	}

	for _, opt := range opts {
		opt(options)
	}

	if options.IsLogToES {
		logutil.LogToES(name)
	}

	// New Service
	service := micro.NewService(
		micro.Name(name),
		micro.Version("latest"),
		micro.Registry(etcd.NewRegistry(registry.Addrs(appconfigutil.MustGet().ETCDAddress))),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
		micro.Broker(nats.NewBroker(broker.Addrs(appconfigutil.MustGet().NatsAddress))),
	)

	// Initialise service
	service.Init()

	starter.Start(service)

	// Run service
	if err := service.Run(); err != nil {
		logrus.Fatal(err)
	}
}
