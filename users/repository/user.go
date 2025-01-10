package repository

// import (
// 	"context"
// 	"fmt"
// 	"time"
// 	"users/entity"

// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// type UserRepository interface {
// 	CreateUser(user entity.CreateUserInput) (*entity.User, error)
// 	GetUserById(id string) (*entity.User, error)
// 	EditUserById(id string, input entity.UpdateUserInput) (*entity.User, error)
// 	DeleteUserById(id string) (*entity.User, error)
// }

// type userRepository struct {
// 	db *mongo.Collection
// }

// func NewUserRepository(db *mongo.Collection) *userRepository {
// 	return &userRepository{db}
// }

// func (r *userRepository) CreateUser(input entity.CreateUserInput) (*entity.User, error) {

// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	user := entity.User{
// 		ID:        primitive.NewObjectID(),
// 		Name:      input.Name,
// 		Email:     input.Email,
// 		CreatedAt: time.Now(),
// 	}

// 	_, err := r.db.InsertOne(ctx, user)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &user, nil
// }

// func (r *userRepository) GetUserById(id string) (*entity.User, error) {

// 	objectID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	filter := bson.M{"_id": objectID}

// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	result := r.db.FindOne(context.Background(), filter)

// 	fmt.Println(result)

// 	var dataUser entity.User
// 	err = r.db.FindOne(ctx, filter).Decode(&dataUser)
// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			return nil, nil
// 		}
// 		return nil, err
// 	}

// 	return &dataUser, nil
// }

// func (r *userRepository) EditUserById(id string, input entity.UpdateUserInput) (*entity.User, error) {
// 	objectID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	filter := bson.M{"_id": objectID}

// 	update := bson.M{
// 		"$set": input,
// 	}

// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

// 	var dataUser entity.User
// 	err = r.db.FindOneAndUpdate(ctx, filter, update, opts).Decode(&dataUser)
// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			return nil, nil
// 		}
// 		return nil, err
// 	}

// 	return &dataUser, nil
// }

// func (r *userRepository) DeleteUserById(id string) (*entity.User, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	objectID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	filter := bson.M{"_id": objectID}

// 	var deletedUser entity.User
// 	err2 := r.db.FindOneAndDelete(ctx, filter).Decode(&deletedUser)
// 	if err2 != nil {
// 		return nil, err
// 	}

// 	return &deletedUser, nil
// }
