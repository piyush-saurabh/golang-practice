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
