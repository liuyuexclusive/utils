package main

import (
	"github.com/yuexclusive/utils/config"
	"github.com/yuexclusive/utils/logger"
	"github.com/yuexclusive/utils/rpc"

	"github.com/yuexclusive/utils/srv/basic/handler/role_handler"
	"github.com/yuexclusive/utils/srv/basic/handler/user_handler"
	"github.com/yuexclusive/utils/srv/basic/proto/role"
	"github.com/yuexclusive/utils/srv/basic/proto/user"
)

func main() {
	cfg := config.MustGet()
	s, err := rpc.NewServer(
		rpc.Registry(cfg.Name, config.MustGet().ETCDAddress, cfg.Address, 5),
		rpc.TLS(cfg.TLS.CertFile, cfg.TLS.KeyFile),
		rpc.Auth(),
	)

	if err != nil {
		logger.Sugar.Fatal(err)
	}

	role.RegisterRoleServer(s.Server, new(role_handler.Handler))
	user.RegisterUserServer(s.Server, new(user_handler.Handler))

	logger.Sugar.Fatal(s.Serve())
}
