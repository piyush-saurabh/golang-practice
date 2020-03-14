package main

import (
	"context"
	"fmt"
	"log"

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

	// create a connection to server
	con, err := grpc.Dial(serverAddress, grpc.WithInsecure())

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
	fmt.Println("[roguesecurity] Calling gRPC server from the method checkResult()")

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

	// Make gRPC call to the server
	res, err := c.GetStatus(ctx, req)

	if err != nil {
		log.Fatalf("[roguesecurity] Error while calling GetStatus() RPC: %v", err)
	}

	log.Printf("[roguesecurity] Response from GetStatus() rpc :%v", res.Status)
}
