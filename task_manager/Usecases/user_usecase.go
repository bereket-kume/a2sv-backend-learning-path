package usecases

import (
	"errors"
	domain "task_manager/Domain"
)

type UserUseCaseImpl struct {
	UserRepo        domain.UserRepository
	PasswordService domain.PasswordService
	JWTService      domain.JWTService
}

func NewUserUseCase(userRepo domain.UserRepository, passwordService domain.PasswordService, jwtService domain.JWTService) *UserUseCaseImpl {
	return &UserUseCaseImpl{
		UserRepo:        userRepo,
		PasswordService: passwordService,
		JWTService:      jwtService,
	}
}

func (u *UserUseCaseImpl) Login(input *domain.User) (string, error) {
	user, err := u.UserRepo.GetUserByUsername(input.Name)
	if err != nil {
		return "", errors.New("user not found")
	}
	err = u.PasswordService.ComparePassword(user.Password, input.Password)
	if err != nil {
		return "", errors.New("incorrect password")
	}
	token, err := u.JWTService.GenerateToken(user.Name, user.Role)
	if err != nil {
		return "", errors.New("failed to generate token")
	}
	return token, nil
}

func (u *UserUseCaseImpl) Register(input *domain.User) (*domain.User, error) {
	hashedPassword, err := u.PasswordService.HashPassword(input.Password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}
	input.Password = hashedPassword
	createdUser, err := u.UserRepo.Register(input)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}

func (u *UserUseCaseImpl) GetUserByUsername(username string) (*domain.User, error) {
	return u.UserRepo.GetUserByUsername(username)
}

func (u *UserUseCaseImpl) PromoteUser(id string) (*domain.User, error) {
	return u.UserRepo.PromoteUser(id)
}
