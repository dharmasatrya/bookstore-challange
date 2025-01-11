package repository

import (
	"context"
	"fmt"
	"time"
	"users/entity"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	CreateUser(user entity.CreateUserInput) (*entity.User, error)
	LoginUser(input entity.LoginRequest) (*entity.User, error)
}

type userRepository struct {
	db *mongo.Collection
}

func NewUserRepository(db *mongo.Collection) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) CreateUser(input entity.CreateUserInput) (*entity.User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user := entity.User{
		ID:       primitive.NewObjectID(),
		Username: input.Username,
		Password: input.Password,
	}

	_, err := r.db.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) LoginUser(input entity.LoginRequest) (*entity.User, error) {

	filter := bson.M{"username": input.Username}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user entity.User
	err := r.db.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		fmt.Println("error 59")
		return nil, err
	}

	fmt.Println("63")

	return &user, nil
}
