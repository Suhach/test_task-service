package grpc

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/Suhach/test_protoc-cont/proto/task"
)

func RegisterServer(s *grpc.Server, h *Handler) {
	pb.RegisterTaskServiceServer(s, h)
}

func RunServer(h *Handler, port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	RegisterServer(s, h)
	log.Printf("Task service running on port %s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
