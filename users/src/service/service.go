// service/user_service.go
package service

import (
	"context"
	"log"

	pb "github.com/dharmasatrya/proto-repo/user"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) RegisterUser(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	log.Printf("Received : %s", req.Username)
	return &pb.RegisterResponse{
		Id:       "123",
		Username: req.Username,
	}, nil
}
