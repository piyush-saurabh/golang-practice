package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/roguesecurity/golang-practice/interceptor/pb"
	"google.golang.org/grpc"
)

const (
	serverIP   = "0.0.0.0"
	serverPort = "50053"
)

type server struct{}

// Before calling this function, server interceptor will be invoked
func (s *server) GetStatus(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {

	fmt.Println("[roguesecurity] GetStatus() method invoked")

	// name := req.GetDetails().GetName()
	// departmant := req.GetDetails().GetDepartment()
	role := req.GetDetails().GetRole()

	var result string

	// Check if metadata is set
	if role == "admin" {
		result = "allowed"
	} else {
		result = "not allowed"
	}

	res := &pb.HelloResponse{
		Status: result,
	}

	return res, nil
}

// Start the server
func main() {

	serverAddress := serverIP + ":" + serverPort

	// port binding
	lis, err := net.Listen("tcp", serverAddress)

	fmt.Println("[roguesecurity] gRPC server on " + serverAddress)

	if err != nil {
		log.Fatalf("[roguesecurity] Failed to listen: %v\n", err)
	}

	// Start a normal gRPC server. (without any interceptor)
	// s := grpc.NewServer()

	// Start the server with UnaryServerInterceptor. location: server/interceptor.go
	// This interceptor will add the metadata 'x-interceptor: client' before making request to the server
	s := grpc.NewServer(grpc.UnaryInterceptor(serverInterceptor))

	pb.RegisterHelloServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("[roguesecurity] Failed to serve: %v", err)
	}
}
