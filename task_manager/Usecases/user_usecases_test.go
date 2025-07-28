package usecases

import (
	domain "task_manager/Domain"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) Register(user *domain.User) (*domain.User, error) {
	args := m.Called(user)
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *UserRepositoryMock) Login(user *domain.User) (string, error) {
	args := m.Called(user)
	return args.String(0), args.Error(1)
}

func (m *UserRepositoryMock) GetUserByUsername(username string) (*domain.User, error) {
	args := m.Called(username)
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *UserRepositoryMock) PromoteUser(id string) (*domain.User, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.User), args.Error(1)
}

type PasswordServiceMock struct {
	mock.Mock
}

func (m *PasswordServiceMock) HashPassword(password string) (string, error) {
	args := m.Called(password)
	return args.String(0), args.Error(1)
}

func (m *PasswordServiceMock) ComparePassword(hashedPassword, password string) error {
	args := m.Called(hashedPassword, password)
	return args.Error(0)
}

type JWTServiceMock struct {
	mock.Mock
}

func (m *JWTServiceMock) GenerateToken(userID string, role string) (string, error) {
	args := m.Called(userID, role)
	return args.String(0), args.Error(1)
}

type UserUseCaseTestSuite struct {
	suite.Suite
	UserRepoMock        *UserRepositoryMock
	PasswordServiceMock *PasswordServiceMock
	JWTServiceMock      *JWTServiceMock
	UserUseCase         *UserUseCase
}

func (suite *UserUseCaseTestSuite) SetupTest() {
	suite.UserRepoMock = new(UserRepositoryMock)
	suite.PasswordServiceMock = new(PasswordServiceMock)
	suite.JWTServiceMock = new(JWTServiceMock)
	suite.UserUseCase = NewUserUseCase(suite.UserRepoMock, suite.PasswordServiceMock, suite.JWTServiceMock)
}

func (suite *UserUseCaseTestSuite) TestRegister() {
	testUser := &domain.User{
		Name:     "testUser",
		Password: "password123",
		Role:     "user",
	}
	suite.PasswordServiceMock.On("HashPassword", testUser.Password).Return("hashedPassword", nil)
	suite.UserRepoMock.On("Register", testUser).Return(testUser, nil)
	createdUser, err := suite.UserUseCase.Register(testUser)

	suite.NoError(err)
	suite.Equal(testUser.Name, createdUser.Name)
	suite.UserRepoMock.AssertExpectations(suite.T())
	suite.PasswordServiceMock.AssertExpectations(suite.T())
}

func (suite *UserUseCaseTestSuite) TestLogin() {
	testUser := &domain.User{
		Name:     "testUser",
		Password: "password123",
		Role:     "user",
	}
	suite.UserRepoMock.On("GetUserByUsername", testUser.Name).Return(testUser, nil)
	suite.PasswordServiceMock.On("ComparePassword", testUser.Password, testUser.Password).Return(nil)
	suite.JWTServiceMock.On("GenerateToken", testUser.Name, testUser.Role).Return("token123", nil)
	token, err := suite.UserUseCase.Login(testUser)
	suite.NoError(err)
	suite.Equal("token123", token)
	suite.UserRepoMock.AssertExpectations(suite.T())
	suite.PasswordServiceMock.AssertExpectations(suite.T())
	suite.JWTServiceMock.AssertExpectations(suite.T())
}

func (suite *UserUseCaseTestSuite) TestGetByUsername() {
	testUser := &domain.User{
		Name:     "testuser",
		Password: "password123",
		Role:     "user",
	}
	suite.UserRepoMock.On("GetUserByUsername", testUser.Name).Return(testUser, nil)
	user, err := suite.UserUseCase.GetUserByUsername(testUser.Name)

	suite.NoError(err)
	suite.Equal(user.Name, testUser.Name)
	suite.Equal(testUser, user)
	suite.UserRepoMock.AssertExpectations(suite.T())
}

func (suite *UserUseCaseTestSuite) TestPromoteUser() {
	testuser := &domain.User{
		ID:       "123",
		Name:     "testuser",
		Password: "password123",
		Role:     "user",
	}
	updatedUser := &domain.User{
		ID:       "123",
		Name:     "testuser",
		Password: "password123",
		Role:     "admin",
	}
	suite.UserRepoMock.On("PromoteUser", testuser.ID).Return(updatedUser, nil)
	user, err := suite.UserUseCase.PromoteUser(testuser.ID)

	suite.NoError(err)
	suite.Equal("admin", user.Role)
	suite.Equal(updatedUser, user)
	suite.UserRepoMock.AssertExpectations(suite.T())
}

func TestUserUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(UserUseCaseTestSuite))
}
