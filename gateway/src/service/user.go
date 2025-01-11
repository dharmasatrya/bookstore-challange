package service

import (
	"context"
	"gateway/entity"
	"log"
	"net/http"

	pb "github.com/dharmasatrya/proto-repo/user"
	userConn "github.com/dharmasatrya/proto-repo/user"
	"google.golang.org/grpc"
)

type UserService interface {
	RegisterUser(order entity.RegisterRequest) (int, *entity.User)
	LoginUser(conn entity.LoginRequest) (int, *entity.LoginResponse)
}

type userService struct {
	conn *grpc.ClientConn
}

func NewUserService(conn *grpc.ClientConn) *userService {
	return &userService{conn}
}

func (u *userService) RegisterUser(input entity.RegisterRequest) (int, *entity.User) {
	client := userConn.NewUserServiceClient(u.conn)

	// token := "Bearer valid-token"

	// md := metadata.Pairs("authorization", token)
	// ctx := metadata.NewOutgoingContext(context.Background(), md)

	res, err := client.RegisterUser(context.Background(), &pb.RegisterRequest{Username: input.Username, Password: input.Password})
	if err != nil {
		log.Fatalf("error while create request %v", err)
	}

	response := entity.User{
		ID:       res.Id,
		Username: res.Username,
	}

	return http.StatusOK, &response
}

func (u *userService) LoginUser(input entity.LoginRequest) (int, *entity.LoginResponse) {
	client := userConn.NewUserServiceClient(u.conn)

	res, err := client.LoginUser(context.Background(), &pb.LoginRequest{Username: input.Username, Password: input.Password})
	if err != nil {
		log.Fatalf("error while create request %v", err)
	}

	response := entity.LoginResponse{
		Token:        res.Token,
		Success:      true,
		ErrorMessage: "",
	}

	return http.StatusOK, &response
}
