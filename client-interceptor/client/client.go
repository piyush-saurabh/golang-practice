package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc/metadata"

	"github.com/roguesecurity/golang-practice/client-interceptor/pb"

	"google.golang.org/grpc"
)

const (
	serverPort = "50053"
	role       = "admin"
)

var (
	serverIP string
)

func main() {

	// Read the gRPC server IP from the client
	fmt.Printf("Enter the gRPC server IP: ")
	_, err := fmt.Scanf("%s\n", &serverIP)
	if err != nil {
		log.Fatalf("[roguesecurity] There was an error while reading gRPC server IP: %v", err)
	}

	serverAddress := serverIP + ":" + serverPort
	fmt.Println("[roguesecurity] gRPC client is trying to connect to: " + serverAddress)

	// create a normal connection to server. (without any interceptor)
	// con, err := grpc.Dial(serverAddress, grpc.WithInsecure())

	// Create a connection to the server with UnaryClientInterceptor. location: client/interceptor.go
	// This interceptor will add the metadata 'x-interceptor: client' before making request to the server
	i := &clientInterceptor{}
	con, err := grpc.Dial(serverAddress, grpc.WithUnaryInterceptor(i.UnaryClientInterceptor), grpc.WithInsecure())

	// Close the connection at the end
	defer con.Close()

	if err != nil {
		log.Fatalf("[roguesecurity] There was an error in the connection: %v", err)
	}

	client := pb.NewHelloServiceClient(con)

	//fmt.Printf("Created client: %f", c)

	checkResult(client)
}

// Make request to gRPC server
func checkResult(c pb.HelloServiceClient) {
	req := &pb.HelloRequest{
		Details: &pb.Hello{
			Name:       "Piyush Saurabh",
			Department: "Security",
			Role:       role,
		},
	}

	// Get the root context
	ctx := context.Background()

	// Create a metadata
	md := metadata.Pairs("authorization", "Bearer some.jwt.token")

	// Attach this metadata to the current context
	ctx = metadata.NewOutgoingContext(ctx, md)

	// Make gRPC call to the server
	fmt.Println("[roguesecurity] Before calling GetStatus() RPC on the server")
	res, err := c.GetStatus(ctx, req)
	fmt.Println("[roguesecurity] After calling GetStatus() RPC on the server")

	if err != nil {
		log.Fatalf("[roguesecurity] Error while calling GetStatus() RPC: %v", err)
	}

	log.Printf("[roguesecurity] Response from GetStatus() rpc :%v", res.Status)
}
