// main.go
package main

import (
	"borrow/src/middleware"
	"context"
	"log"
	"net"

	"borrow/config"
	"borrow/src/repository"
	"borrow/src/service"

	pb "github.com/dharmasatrya/proto-repo/borrow"
	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.UnaryAuthInterceptor),
	)

	// grpcServer := grpc.NewServer()

	db, err := config.ConnectionDB(context.Background())
	if err != nil {
		log.Fatalf("Error connecting to db")
	}

	borrowRepository := repository.NewBorrowRepository(db)

	borrowService := service.NewBorrowService(borrowRepository)
	pb.RegisterBorrowServiceServer(grpcServer, borrowService)

	log.Println("Server is running on port 50053...")
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
