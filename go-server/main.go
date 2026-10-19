
package main

import (
	"context"
	"log"
	"net"

	pb "grpc-polyglot-lab/go-server/pb"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

func (s *server) GetUserProfile(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	log.Printf("Permintaan profil diterima untuk user_id: %d", req.GetUserId())

	return &pb.UserResponse{
		Name:  "Budi Santoso",
		Email: "budi@example.com",
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Gagal membuka port: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &server{})

	log.Println("Server gRPC berjalan pada port 50051")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Server gRPC gagal: %v", err)
	}
}
