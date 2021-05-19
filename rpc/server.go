package rpc

import (
	"context"
	"crypto/tls"
	"errors"
	"net"
	"strings"

	"github.com/yuexclusive/utils/config"
	"github.com/yuexclusive/utils/registry"
	"github.com/yuexclusive/utils/srv/auth/proto/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	errMissingMetadata = status.Errorf(codes.InvalidArgument, "missing metadata")
	errInvalidToken    = status.Errorf(codes.Unauthenticated, "invalid token")
)

type Server struct {
	option serverOption
	*grpc.Server
}

type serverOption struct {
	name      string
	endpoints []string
	address   string
	lease     int64
	certFile  string
	keyFile   string
	tls       bool
	auth      bool
}

type ServerOption interface {
	apply(*serverOption)
}

type registryServerOption struct {
	name      string
	endpoints []string
	address   string
	lease     int64
}

func (r *registryServerOption) apply(s *serverOption) {
	s.name = r.name
	s.endpoints = r.endpoints
	s.address = r.address
	s.lease = r.lease
}

func Registry(name string, endpoints []string, address string, lease int64) ServerOption {
	return &registryServerOption{
		name:      name,
		endpoints: endpoints,
		address:   address,
		lease:     lease,
	}
}

type tlsServerOption struct {
	certFile string
	keyFile  string
}

func (r *tlsServerOption) apply(s *serverOption) {
	s.certFile = r.certFile
	s.keyFile = r.keyFile
	s.tls = true
}

func TLS(certFile, keyFile string) ServerOption {
	return &tlsServerOption{certFile: certFile, keyFile: keyFile}
}

type authServerOption struct {
}

func (r *authServerOption) apply(s *serverOption) {
	s.auth = true
}

func Auth() ServerOption {
	return &authServerOption{}
}

func NewServer(serverOptions ...ServerOption) (*Server, error) {
	var option serverOption

	for _, s := range serverOptions {
		s.apply(&option)
	}

	var grpcServerOptions []grpc.ServerOption

	if option.tls {
		cert, err := tls.LoadX509KeyPair(option.certFile, option.keyFile)
		if err != nil {
			return nil, err
		}
		grpcServerOptions = append(grpcServerOptions, grpc.Creds(credentials.NewServerTLSFromCert(&cert)))
	}

	if option.auth {
		grpcServerOptions = append(grpcServerOptions, grpc.UnaryInterceptor(ensureValidToken))
	}

	server := grpc.NewServer(
		grpcServerOptions...,
	)

	res := &Server{Server: server, option: option}

	return res, nil
}

func (s *Server) Serve() error {
	option := s.option

	var address string
	if s.option.address == "" {
		address = "127.0.0.1:0"
	} else {
		address = s.option.address
	}
	listener, err := net.Listen("tcp", address)
	option.address = listener.Addr().String()
	if err != nil {
		return err
	}

	//registry
	_, err = registry.NewService(option.endpoints, option.name, option.address, option.lease)
	if err != nil {
		return err
	}

	return s.Server.Serve(listener)
}

// ensureValidToken ensures a valid token exists within a request's metadata. If
// the token is missing or invalid, the interceptor blocks execution of the
// handler and returns an error. Otherwise, the interceptor invokes the unary
// handler.
func ensureValidToken(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errMissingMetadata
	}

	// The keys within metadata.MD are normalized to lowercase.
	// See: https://godoc.org/google.golang.org/grpc/metadata#New
	if err := validToken(md["authorization"]); err != nil {
		return nil, err
	}
	// Continue execution of handler after ensuring a valid token.
	return handler(ctx, req)
}

// validToken validates the authorization.
func validToken(authorization []string) error {
	if len(authorization) < 1 {
		return errors.New("please pass a authorization")
	}
	token := strings.TrimPrefix(authorization[0], "Bearer ")

	//validate through the auth server

	cfg := config.MustGet()

	client, err := Dial(
		Discovery(cfg.AuthServiceName, []string{cfg.AuthServiceName}, cfg.ETCDAddress),
		TLSClient(cfg.TLS.CACertFile, cfg.TLS.ServerNameOverride),
	)

	if err != nil {
		return err
	}

	_, err = auth.NewAuthClient(client).Validate(context.Background(), &auth.ValidateRequest{Token: token})

	if err != nil {
		return err
	}

	return nil
}
