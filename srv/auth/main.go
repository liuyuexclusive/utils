package main

import (
	"github.com/yuexclusive/utils/config"
	"github.com/yuexclusive/utils/logger"
	"github.com/yuexclusive/utils/rpc"
	"github.com/yuexclusive/utils/srv/auth/handler"
	"github.com/yuexclusive/utils/srv/auth/proto/auth"
)

func main() {
	cfg := config.MustGet()
	s, err := rpc.NewServer(
		rpc.Registry(cfg.Name, cfg.ETCDAddress, cfg.Address, 5),
		rpc.TLS(cfg.TLS.CertFile, cfg.TLS.KeyFile),
	)

	if err != nil {
		logger.Sugar.Fatal(err)
	}

	auth.RegisterAuthServer(s.Server, new(handler.Handler))

	logger.Sugar.Fatal(s.Serve())
}
