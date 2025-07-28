package repositories

import (
	"context"
	"errors"
	domain "task_manager/Domain"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MockPasswordService struct{}

func (m *MockPasswordService) HashPassword(password string) (string, error) {
	return password, nil
}

func (m *MockPasswordService) ComparePassword(hashedPassword, plainPassword string) error {
	if hashedPassword != plainPassword {
		return errors.New("passwords do not match")
	}
	return nil
}

type UserRepositoryTestSuite struct {
	suite.Suite
	UserCollection *mongo.Collection
	Database       *mongo.Database
	client         *mongo.Client
	Repository     *UserRepository
}

func (suite *UserRepositoryTestSuite) SetupTest() {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	suite.NoError(err)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	suite.NoError(err)

	suite.client = client
	suite.Database = client.Database("task_manager_test")
	suite.UserCollection = suite.Database.Collection("users")

	// Ensure the collection is empty before each test
	err = suite.UserCollection.Drop(context.Background())
	suite.NoError(err)

	mockPasswordService := &MockPasswordService{}
	suite.Repository = &UserRepository{
		UserCollection:  suite.UserCollection,
		PasswordService: mockPasswordService,
	}
}

func (suite *UserRepositoryTestSuite) TearDownTest() {
	err := suite.Database.Drop(context.Background())
	suite.NoError(err)
	err = suite.client.Disconnect(context.Background())
	suite.NoError(err)
}

func (suite *UserRepositoryTestSuite) TestRegister() {
	user := &domain.User{
		Name:     "testuser",
		Password: "password123",
		Role:     domain.RoleUser,
	}
	adminUser := &domain.User{
		Name:     "adminuser",
		Password: "adminpassword",
		Role:     domain.RoleAdmin,
	}
	admin, _ := suite.Repository.Register(adminUser)
	suite.NotNil(admin)
	registeredUser, err := suite.Repository.Register(user)
	suite.NoError(err)

	var retrievedUser domain.User
	objectID, err := primitive.ObjectIDFromHex(registeredUser.ID)
	suite.NoError(err)
	err = suite.UserCollection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&retrievedUser)
	suite.NoError(err)
	suite.Equal(user.Name, retrievedUser.Name)
	suite.Equal(domain.RoleUser, retrievedUser.Role)
	suite.Equal(user.Password, retrievedUser.Password)
}

func (suite *UserRepositoryTestSuite) TestGetByUsername() {
	username := "testuser"
	user := &domain.User{
		Name:     username,
		Password: "password123",
		Role:     domain.RoleUser,
	}
	_, err := suite.Repository.Register(user)
	suite.NoError(err)

	retrievedUser, err := suite.Repository.GetUserByUsername(username)

	suite.NoError(err)
	suite.NotNil(retrievedUser)
	suite.Equal(username, retrievedUser.Name)
	suite.Equal(domain.RoleAdmin, retrievedUser.Role)
}

func (suite *UserRepositoryTestSuite) TestPromoteUser() {
	user := &domain.User{
		ID:       "123",
		Name:     "testuser",
		Password: "password123",
		Role:     domain.RoleUser,
	}
	_, err := suite.Repository.Register(user)
	suite.NoError(err)

	updatedUser, err := suite.Repository.PromoteUser(user.ID)
	suite.NoError(err)
	suite.NotNil(updatedUser)
	suite.Equal(domain.RoleAdmin, updatedUser.Role)
	suite.Equal(user.Name, updatedUser.Name)
	suite.Equal(user.Password, updatedUser.Password)
}

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}
