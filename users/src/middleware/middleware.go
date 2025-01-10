package middleware

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func UnaryAuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	ctx, err := AuthInterceptor(ctx)
	if err != nil {
		return nil, err
	}
	return handler(ctx, req)
}

func AuthInterceptor(ctx context.Context) (context.Context, error) {
	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		log.Println("No metadata found")
		return nil, status.Errorf(codes.Unauthenticated, "Unauthorized")
	}

	log.Printf("Metadata received: %v", md)

	token := md["authorization"]
	if len(token) == 0 || token[0] != "Bearer valid-token" {
		log.Println("Invalid or Missing token")
		return nil, status.Errorf(codes.Unauthenticated, "Unauthorized")
	}

	log.Println("Token validated successfully")

	return ctx, nil
}
