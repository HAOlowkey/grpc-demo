package auth

import (
	"context"
)

func NewClientAuth(clientId, clientSecret string) *clientAuth {
	return &clientAuth{
		clientId:     clientId,
		clientSecret: clientSecret,
	}
}

type clientAuth struct {
	clientId     string
	clientSecret string
}

func (c *clientAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		ClientIdKey:     c.clientId,
		ClientSecretKey: c.clientSecret,
	}, nil
}

func (c *clientAuth) RequireTransportSecurity() bool {
	return false
}
