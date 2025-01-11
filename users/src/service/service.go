// service/user_service.go
package service

import (
	"context"
	"fmt"
	"log"
	"users/entity"
	"users/src/repository"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"

	pb "github.com/dharmasatrya/proto-repo/user"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	userRepo repository.UserRepository
}

var jwtSecret = []byte("secret")

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

func (s *UserService) LoginUser(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {

	userInput := entity.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	}

	res, err := s.userRepo.LoginUser(userInput)
	if err != nil {
		fmt.Println("error", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(req.Password)); err != nil {
		fmt.Printf("Password comparison error: %v\n", err)
		return nil, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": res.ID.Hex(),
	})

	tokenString, err2 := token.SignedString(jwtSecret)
	if err2 != nil {
		fmt.Println("error stringify")
		return nil, err2
	}

	return &pb.LoginResponse{
		Token:        tokenString,
		Success:      true,
		ErrorMessage: "",
	}, nil
}
