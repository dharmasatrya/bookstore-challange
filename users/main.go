// main.go
package main

import (
	"context"
	"log"
	"net"

	"users/config"
	"users/repository"
	"users/src/service"

	pb "github.com/dharmasatrya/proto-repo/user"
	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// grpcServer := grpc.NewServer(
	// 	grpc.UnaryInterceptor(middleware.UnaryAuthInterceptor),
	// )

	grpcServer := grpc.NewServer()

	db, err := config.ConnectionDB(context.Background())
	if err != nil {
		log.Fatalf("Error connecting to db")
	}

	userRepository := repository.NewUserRepository(db)

	userService := service.NewUserService(userRepository)
	pb.RegisterUserServiceServer(grpcServer, userService)

	log.Println("Server is running on port 50051...")
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
