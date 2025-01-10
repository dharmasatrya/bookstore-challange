// service/user_service.go
package service

import (
	"context"
	"fmt"
	"log"
	"users/entity"
	"users/repository"

	pb "github.com/dharmasatrya/proto-repo/user"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	userRepo repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepository,
	}
}

func (s *UserService) RegisterUser(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	log.Printf("Received : %s", req.Username)

	userInput := entity.CreateUserInput{
		Username: req.Username,
		Password: req.Password,
	}

	res, err := s.userRepo.CreateUser(userInput)
	if err != nil {
		fmt.Println("error")
	}

	return &pb.RegisterResponse{
		Id:       res.ID.Hex(),
		Username: req.Username,
	}, nil
}
