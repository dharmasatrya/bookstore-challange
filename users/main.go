// main.go
package main

import (
	"log"
	"net"

	"users/src/middleware"
	"users/src/service"

	pb "github.com/dharmasatrya/proto-repo/user"
	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.UnaryAuthInterceptor),
	)

	userService := service.NewUserService()
	pb.RegisterUserServiceServer(grpcServer, userService)

	log.Println("Server is running on port 50051...")
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
