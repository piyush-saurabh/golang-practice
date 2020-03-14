package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc"
)

// UnaryClientInterceptor is the unary client interceptor
// This interceptor will add a metadata 'x-interceptor: client'
func clientInterceptor(ctx context.Context, method string, req interface{}, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {

	fmt.Println("clientInterceptor interceptor invoked")

	// Add the metadata to the current context
	// 'x-interceptor: client'
	ctx = metadata.AppendToOutgoingContext(ctx, "x-interceptor", "client")

	return invoker(ctx, method, req, reply, cc, opts...)

}
