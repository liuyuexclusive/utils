package main

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/yuexclusive/utils/logger"
	"github.com/yuexclusive/utils/rpc"
	"github.com/yuexclusive/utils/rpc/middleware/trace"
	"github.com/yuexclusive/utils/srv/auth/handler"
	"github.com/yuexclusive/utils/srv/auth/proto/auth"
	"google.golang.org/grpc"
)

func main() {
	tracer, closer, err := trace.Tracer()

	if err != nil {
		logger.Sugar.Fatal(err)
	}

	defer closer.Close()

	s, err := rpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_zap.UnaryServerInterceptor(logger.Logger),
			grpc_opentracing.UnaryServerInterceptor(grpc_opentracing.WithTracer(tracer)),
		)),
	)

	if err != nil {
		logger.Sugar.Fatal(err)
	}

	auth.RegisterAuthServer(s.Server, new(handler.Handler))

	logger.Sugar.Fatal(s.Serve())
}
