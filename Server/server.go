package main

import (
	"context"
	"log"
	"net"

	pb "grpc-project/server" // Replace with your actual Go package

	"google.golang.org/grpc"
)

// Server struct implements the Calculator service
type Server struct {
	pb.UnimplementedCalculatorServer
}

// Square calculates the square of a number
func (s *Server) Square(ctx context.Context, req *pb.SquareRequest) (*pb.SquareResponse, error) {
	result := req.GetNumber() * req.GetNumber()
	return &pb.SquareResponse{Result: result}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCalculatorServer(grpcServer, &Server{})

	log.Println("Server is running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
