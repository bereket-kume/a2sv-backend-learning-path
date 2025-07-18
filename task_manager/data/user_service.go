package data

import (
	"context"
	"errors"
	"fmt"
	"task_manager/database"
	"task_manager/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const (
	RoleAdmin = "admin"
	RoleUser  = "user"
)

var jwtSecret = []byte("secret123")

func Register(user models.User) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	count, err := database.UserCollection.CountDocuments(ctx, primitive.M{})

	if err != nil {
		return models.User{}, fmt.Errorf("failed to check user collection: %w", err)
	}
	if count == 0 {
		user.Role = RoleAdmin
	} else {
		user.Role = RoleUser
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, fmt.Errorf("failed to hash password: %w", err)
	}
	user.Password = string(hashedPassword)

	res, err := database.UserCollection.InsertOne(ctx, user)
	if err != nil {
		return models.User{}, errors.New("user registartion failed")
	}
	if objId, ok := res.InsertedID.(primitive.ObjectID); ok {
		user.ID = objId
	}
	return user, nil
}

func Login(input models.User) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	err := database.UserCollection.FindOne(ctx, map[string]interface{}{"username": input.Name}).Decode(&user)

	if err != nil {
		return "", errors.New("invalid username or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return "", errors.New(("invalid username or password"))
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Name,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour * 2).Unix(),
	})
	tokenString, _ := token.SignedString(jwtSecret)

	return tokenString, nil
}

func PromoteUser(id string) (models.User, error) {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.User{}, errors.New("invalid ObjectId")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{"$set": bson.M{
		"role": "admin",
	}}

	res, err := database.UserCollection.UpdateByID(ctx, objId, update)
	if err != nil {
		return models.User{}, errors.New("update failed")
	}

	if res.MatchedCount == 0 {
		return models.User{}, errors.New("no user found to promote")
	}

	var updatedUser models.User
	err = database.UserCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedUser)
	if err != nil {
		return models.User{}, errors.New("failed to retrieve updated user")
	}

	return updatedUser, nil
}
