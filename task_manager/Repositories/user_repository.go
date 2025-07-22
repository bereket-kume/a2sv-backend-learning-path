package repositories

import (
	"context"
	"errors"
	"fmt"
	domain "task_manager/Domain"
	infrastructure "task_manager/Infrastructure"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryImpl struct {
	UserCollection  *mongo.Collection
	PasswordService infrastructure.PasswordService
	JWTService      infrastructure.JWTService
}

func NewUserRepository(db *mongo.Collection) *UserRepositoryImpl {
	return &UserRepositoryImpl{UserCollection: db}
}

func (r *UserRepositoryImpl) Register(user *domain.User) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	count, err := r.UserCollection.CountDocuments(ctx, primitive.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to check user collection: %w", err)
	}

	if count == 0 {
		user.Role = domain.RoleAdmin
	} else {
		user.Role = domain.RoleUser
	}

	// Password is already hashed in usecase

	res, err := r.UserCollection.InsertOne(ctx, user)
	if err != nil {
		return nil, errors.New("user registration failed")
	}

	if objId, ok := res.InsertedID.(primitive.ObjectID); ok {
		user.ID = objId
	} else {
		return nil, errors.New("failed to retrieve inserted ID")
	}

	return user, nil
}

func (r *UserRepositoryImpl) GetUserByUsername(username string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user domain.User
	err := r.UserCollection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func (r *UserRepositoryImpl) PromoteUser(id string) (*domain.User, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid ObjectId format")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{"$set": bson.M{
		"role": domain.RoleAdmin,
	}}

	res, err := r.UserCollection.UpdateByID(ctx, objId, update)
	if err != nil {
		return nil, fmt.Errorf("failed to promote user: %w", err)
	}

	if res.MatchedCount == 0 {
		return nil, errors.New("no user found with the given ID")
	}

	var updatedUser domain.User
	err = r.UserCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedUser)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve updated user: %w", err)
	}

	return &updatedUser, nil
}
