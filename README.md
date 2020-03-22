# golang-practice
This repo contains the source code for practicing golang concepts

## Context and Metadata
Metadata is sent as **HTTP/2 header**

### Add metadata information at client side

```go
ctx = metadata.AppendToOutgoingContext(ctx, "x-interceptor", "client")
```

### Read metadata information on server side
```go
md, err := metadata.FromIncomingContext(ctx)
```

## Interceptors

### Start the gRPC server with Unary Interceptor

```go
s := grpc.NewServer(grpc.UnaryInterceptor(serverInterceptor))
```

### Define the Server Unary Interceptor

From the **server interceptor**, invoke other methods for functionalities like **authentication, logging** etc
Make sure to pass the **current context** to those methods.

```go
// Server Unary Interceptor
// It authorizes the request if 'authorization' header/metadata is present
// Note: function name can be anything, but its signature should match to UnaryServerInterceptor
func serverInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	// Invoke function for authorization
  // Check if 'Authorization' header/metadata is set
	if err := authorize(ctx); err != nil {
		return nil, err
	}
  
  // Do not modify this
	h, err := handler(ctx, req)
	return h, err
}
```

### Using client with Client Unary Interceptor

```go
con, err := grpc.Dial(serverAddress, grpc.WithUnaryInterceptor(clientInterceptor), grpc.WithInsecure())
```

### Define the Client Unary Interceptor

Client interceptor is used to add **metadata information (headers)** before making request to the gRPC server

```go
// This interceptor will add a metadata 'x-interceptor: client'
func clientInterceptor(ctx context.Context, method string, req interface{}, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {

	fmt.Println("clientInterceptor interceptor invoked")

	// Add the metadata to the current context
	// 'x-interceptor: client'
	ctx = metadata.AppendToOutgoingContext(ctx, "x-interceptor", "client")

	return invoker(ctx, method, req, reply, cc, opts...)
}
```



