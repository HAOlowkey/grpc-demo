package auth

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	ClientIdKey     = "client-id"
	ClientSecretKey = "client-secret"
)

func GrpcAuthUnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return newAuth().auth
}

func newAuth() *auth {
	return &auth{}
}

type auth struct {
}

func (a *auth) auth(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (
	resp interface{}, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		fmt.Println(md)
	}

	clientId, clientSecret := a.getCredentialsFromMedata(md)
	err = a.validateServiceCredential(clientId, clientSecret)
	if err != nil {
		return
	}
	return handler(ctx, req)
}

func (a *auth) validateServiceCredential(clientId, clientSecret string) error {
	if clientId == "" && clientSecret == "" {
		return status.Errorf(codes.Unauthenticated, "%s or %s not provided", ClientIdKey, ClientSecretKey)
	}

	if !(clientId == "admin" && clientSecret == "123456") {
		return status.Errorf(codes.Unauthenticated, "client-Id or client-Secret is not correct")
	}

	return nil
}

func (a *auth) getCredentialsFromMedata(md metadata.MD) (clientId, clientSecret string) {
	cids := md.Get(ClientIdKey)
	sids := md.Get(ClientSecretKey)
	if len(cids) > 0 {
		clientId = cids[0]
	}

	if len(sids) > 0 {
		clientSecret = sids[0]
	}

	return
}
