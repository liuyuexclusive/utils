package rpc

import (
	"context"
	"crypto/tls"
	"net"
	"strings"

	"github.com/yuexclusive/utils/registry"
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
	if !validToken(md["authorization"]) {
		return nil, errInvalidToken
	}
	// Continue execution of handler after ensuring a valid token.
	return handler(ctx, req)
}

// validToken validates the authorization.
func validToken(authorization []string) bool {
	if len(authorization) < 1 {
		return false
	}
	token := strings.TrimPrefix(authorization[0], "Bearer ")
	// Perform the token validation here. For the sake of this example, the code
	// here forgoes any of the usual OAuth2 token validation and instead checks
	// for a token matching an arbitrary string.
	return token == "some-secret-token"
}
