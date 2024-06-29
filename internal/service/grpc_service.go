package service

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/tanush-128/openzo_backend/store/config"
	"github.com/tanush-128/openzo_backend/store/internal/pb"
	"github.com/tanush-128/openzo_backend/store/internal/repository"

	// "github.com/tanush-128/openzo_backend/store/internal/repository"

	"google.golang.org/grpc"
)

type Server struct {
	pb.StoreServiceServer
	StoreRepository repository.StoreRepository
}

func GrpcServer(
	cfg *config.Config,
	server *Server,

) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.GRPCPort))

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Printf("Server listening at %v", lis.Addr())
	// Initialize gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterStoreServiceServer(grpcServer, server)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}

func (s *Server) GetFCMToken(ctx context.Context, req *pb.StoreId) (*pb.FCMToken, error) {
	// Implement your business logic here
	token, err := s.StoreRepository.GetFCMTokenByStoreID(req.GetId())
	if err != nil {
		return nil, err
	}
	return &pb.FCMToken{
		Token: token,
	}, nil

}
