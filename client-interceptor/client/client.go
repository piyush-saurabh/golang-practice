package main

import (
	"context"
	"fmt"
	"log"

	"github.com/roguesecurity/golang-practice/client-interceptor/pb"

	"google.golang.org/grpc"
)

const (
	serverIP   = "localhost"
	serverPort = "50053"
)

func main() {

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
			Role:       "admin",
		},
	}

	// Make gRPC call to the server
	res, err := c.GetStatus(context.Background(), req)

	if err != nil {
		log.Fatalf("[roguesecurity] Error while calling GetStatus() RPC: %v", err)
	}

	log.Printf("[roguesecurity] Response from GetStatus() rpc :%v", res.Status)
}
