package rpc

import (
	"fmt"

	"github.com/yuexclusive/utils/config"
	"github.com/yuexclusive/utils/registry"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
)

type DialOption interface {
	apply(*dialOption)
}

type dialOption struct {
	auth               bool
	tls                bool
	name               string
	endpoints          []string
	certFile           string
	serverNameOverride string
	token              string
}

type discoveryDialOption struct {
	name      string
	endpoints []string
}

func (r *discoveryDialOption) apply(s *dialOption) {
	s.endpoints = r.endpoints
	s.name = r.name
}

func Discovery(name string, endpoints []string) DialOption {
	return &discoveryDialOption{
		name:      name,
		endpoints: endpoints,
	}
}

type tlsClientDialOption struct {
	certFile           string
	serverNameOverride string
}

func (r *tlsClientDialOption) apply(s *dialOption) {
	s.certFile = r.certFile
	s.serverNameOverride = r.serverNameOverride
	s.tls = true
}

func TLSClient(certFile, serverNameOverride string) DialOption {
	return &tlsClientDialOption{
		certFile:           certFile,
		serverNameOverride: serverNameOverride,
	}
}

type authClientDialOption struct {
	token string
}

func (r *authClientDialOption) apply(s *dialOption) {
	s.auth = true
	s.token = r.token
}

func AuthClient(token string) DialOption {
	return &authClientDialOption{token: token}
}

func Dial(dialOptions ...DialOption) (*grpc.ClientConn, error) {
	var option dialOption

	for _, v := range dialOptions {
		v.apply(&option)
	}

	var opts []grpc.DialOption

	if option.auth {
		perRPC := oauth.NewOauthAccess(fetchToken(option.token))
		opts = append(opts, grpc.WithPerRPCCredentials(perRPC))
	}

	if option.tls {
		creds, err := credentials.NewClientTLSFromFile(option.certFile, option.serverNameOverride)
		if err != nil {
			return nil, fmt.Errorf("failed to load credentials: %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	opts = append(opts, grpc.WithBlock())

	dis := registry.NewDiscovery(option.endpoints, option.name)

	address, err := dis.Get(option.name)
	if err != nil {
		return nil, err
	}

	return grpc.Dial(address, opts...)
}

func DialByName(name string) (*grpc.ClientConn, error) {
	cfg := config.MustGet()
	return Dial(
		Discovery(name, cfg.ETCDAddress),
		TLSClient(cfg.TLS.CACertFile, cfg.TLS.ServerNameOverride),
	)
}

func DialByNameWithAuth(name, token string) (*grpc.ClientConn, error) {
	cfg := config.MustGet()
	return Dial(
		Discovery(name, cfg.ETCDAddress),
		TLSClient(cfg.TLS.CACertFile, cfg.TLS.ServerNameOverride),
		AuthClient(token),
	)
}

func fetchToken(token string) *oauth2.Token {
	return &oauth2.Token{
		AccessToken: token,
	}
}
