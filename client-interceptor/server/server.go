package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/roguesecurity/golang-practice/client-interceptor/pb"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) GetStatus(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {

	fmt.Println("[roguesecurity] Inside GetStatus() method")

	name := req.GetName()
	departmant := req.GetDepartment()
	role := req.GetRole()

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
	fmt.Println("[roguesecurity] gRPC server started. Now serving client interceptor microservice ...")

	// port binding
	lis, err := net.Listen("tcp", "0.0.0.0:50053")
	if err != nil {
		log.Fatalf("[roguesecurity] Failed to listen: %v\n", err)
	}

	// Create gRPC server
	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("[roguesecurity] Failed to serve: %v", err)
	}
}
