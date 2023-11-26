package grpc_server

import (
	"context"
	"fmt"
	"github.com/brianvoe/gofakeit"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	desc "github.com/Yorherim/m_practice_auth/pkg/user_v1"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedUserV1Server
}

func InitGRPCServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("User id: %d", req.GetId())

	return &desc.GetResponse{
		Id:        req.GetId(),
		Name:      gofakeit.BeerName(),
		Email:     gofakeit.Email(),
		Role:      desc.Role(1),
		CreatedAt: timestamppb.New(gofakeit.Date()),
		UpdatedAt: timestamppb.New(gofakeit.Date()),
	}, nil
}

func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("User Name: %s", req.GetName())
	log.Printf("User Email: %s", req.GetEmail())
	log.Printf("User Role: %d", req.GetRole())
	log.Printf("User Password: %s", req.GetPassword())
	log.Printf("User Password: %s", req.GetPasswordConfirm())

	return &desc.CreateResponse{
		Id: 1,
	}, nil
}
