package infrastructure

import "golang.org/x/crypto/bcrypt"

type PasswordService interface {
	HashPassword(password string) (string, error)
	ComparePassword(hashedPassword, password string) error
}

type PasswordServiceImpl struct{}

func NewPasswordService() *PasswordServiceImpl {
	return &PasswordServiceImpl{}
}

func (p *PasswordServiceImpl) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (p *PasswordServiceImpl) ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
