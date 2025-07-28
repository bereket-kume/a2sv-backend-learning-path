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

type UserRepository struct {
	UserCollection  *mongo.Collection
	PasswordService infrastructure.PasswordService
	JWTService      infrastructure.JWTService
}

func NewUserRepository(db *mongo.Collection) *UserRepository {
	return &UserRepository{UserCollection: db}
}

func (r *UserRepository) Register(user *domain.User) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	count, err := r.UserCollection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to check user collection: %w", err)
	}

	if count == 0 {
		user.Role = domain.RoleAdmin
	} else {
		user.Role = domain.RoleUser
	}

	hashedPassword, err := r.PasswordService.HashPassword(user.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}
	user.Password = hashedPassword

	res, err := r.UserCollection.InsertOne(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("user registration failed: %w", err)
	}

	if objId, ok := res.InsertedID.(primitive.ObjectID); ok {
		user.ID = objId.Hex()
	} else {
		return nil, errors.New("failed to retrieve inserted ID")
	}

	return user, nil
}

func (r *UserRepository) GetUserByUsername(username string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user domain.User
	err := r.UserCollection.FindOne(ctx, bson.M{"name": username}).Decode(&user)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func (r *UserRepository) PromoteUser(id string) (*domain.User, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid ObjectId format")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": objId}
	update := bson.M{"$set": bson.M{"role": domain.RoleAdmin}}
	res, err := r.UserCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, fmt.Errorf("failed to promote user: %w", err)
	}

	if res.MatchedCount == 0 {
		return nil, errors.New("no user found with the given ID")
	}

	var updatedUser domain.User
	err = r.UserCollection.FindOne(ctx, filter).Decode(&updatedUser)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve updated user: %w", err)
	}

	return &updatedUser, nil
}
