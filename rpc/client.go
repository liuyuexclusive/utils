package rpc

import (
	"fmt"

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
	names              []string
	name               string
	endpoints          []string
	certFile           string
	serverNameOverride string
}

type discoveryDialOption struct {
	names     []string
	name      string
	endpoints []string
}

func (r *discoveryDialOption) apply(s *dialOption) {
	s.endpoints = r.endpoints
	s.name = r.name
	s.names = r.names
}

func Discovery(name string, names []string, endpoints []string) DialOption {
	return &discoveryDialOption{
		name:      name,
		names:     names,
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
}

func (r *authClientDialOption) apply(s *dialOption) {
	s.auth = true
}

func AuthClient() DialOption {
	return &authClientDialOption{}
}

func Dial(dialOptions ...DialOption) (*grpc.ClientConn, error) {
	var option dialOption

	for _, v := range dialOptions {
		v.apply(&option)
	}

	var opts []grpc.DialOption

	if option.auth {
		perRPC := oauth.NewOauthAccess(fetchToken())
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

	dis := registry.NewDiscovery(option.endpoints, option.names)

	address, err := dis.Get(option.name)
	if err != nil {
		return nil, err
	}

	return grpc.Dial(address, opts...)
}

func fetchToken() *oauth2.Token {
	return &oauth2.Token{
		AccessToken: "some-secret-token",
	}
}
