package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// Unary Interceptor
// It authorizes the request if 'authorization' header/metadata is present
// Note: function name can be anything, but its signature should match to UnaryServerInterceptor
func serverInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	fmt.Println("UnaryServerInterceptor interceptor invoked")

	// Authorization
	// Check if 'Authorization' header/metadata is set
	if err := authorize(ctx); err != nil {
		return nil, err
	}

	h, err := handler(ctx, req)
	return h, err
}

func authorize(ctx context.Context) error {

	var token string

	// Read the metadata from the context
	// Sample output:
	// md = map[:authority:[localhost:50053] authorization:[Bearer some.jwt.token] content-type:[application/grpc] user-agent:[grpc-go/1.29.0-dev]]
	if md, ok := metadata.FromIncomingContext(ctx); ok {

		token = md["authorization"][0]
		fmt.Printf("Authorization header value = %v \n", token)

	}
	if token != "some.jwt.token" {
		return status.Errorf(codes.Unauthenticated, "Authorization token missing")
	}
	return nil
}
