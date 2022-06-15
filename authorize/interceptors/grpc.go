package interceptors

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Neakxs/protoc-gen-authz/authorize"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

func NewGRPCUnaryInterceptor(authzHandler authorize.AuthzInterceptor) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		headers := http.Header{}
		if md, ok := metadata.FromIncomingContext(ctx); ok {
			for k, v := range md {
				for _, vv := range v {
					headers.Add(k, vv)
				}
			}
		}
		if p, ok := peer.FromContext(ctx); ok && p.Addr != nil {
			headers.Add("forwarded", fmt.Sprintf("for=%s", p.Addr.String()))
			headers.Add("x-forwarded-for", p.Addr.String())
		}
		if err := authzHandler.Authorize(ctx, info.FullMethod, headers, req); err != nil {
			return nil, status.Error(codes.PermissionDenied, err.Error())
		}
		return handler(ctx, req)
	}
}

func NewGRPCStreamInterceptor(authzHandler authorize.AuthzInterceptor) grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		ctx := ss.Context()
		headers := http.Header{}
		if md, ok := metadata.FromIncomingContext(ctx); ok {
			for k, v := range md {
				for _, vv := range v {
					headers.Add(k, vv)
				}
			}
		}
		if p, ok := peer.FromContext(ctx); ok && p.Addr != nil {
			headers.Add("forwarded", fmt.Sprintf("for=%s", p.Addr.String()))
			headers.Add("x-forwarded-for", p.Addr.String())
		}
		if err := authzHandler.Authorize(ctx, info.FullMethod, headers, nil); err != nil {
			return status.Error(codes.PermissionDenied, err.Error())
		}
		return handler(srv, ss)
	}
}
