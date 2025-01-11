// main.go
package main

import (
	"books/src/middleware"
	"context"
	"log"
	"net"

	"books/config"
	"books/src/repository"
	"books/src/service"

	pb "github.com/dharmasatrya/proto-repo/book"
	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", ":50052")
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

	bookRepository := repository.NewBookRepository(db)

	bookService := service.NewBookService(bookRepository)
	pb.RegisterBookServiceServer(grpcServer, bookService)

	log.Println("Server is running on port 50052...")
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
