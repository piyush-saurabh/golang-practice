package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
)

// Unary Interceptor
// It authorizes the request if 'authorization' header/metadata is present
// Note: function name can be anything, but its signature should match to UnaryServerInterceptor
func serverInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	fmt.Println("UnaryServerInterceptor interceptor invoked")

	h, err := handler(ctx, req)
	return h, err
}
