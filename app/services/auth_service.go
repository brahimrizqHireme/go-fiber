package services

import (
	"github.com/brahimrizqHireme/go-fiber/app/errors"
	"github.com/brahimrizqHireme/go-fiber/app/models"
	"github.com/brahimrizqHireme/go-fiber/app/repositories"
	"github.com/brahimrizqHireme/go-fiber/app/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	userRepo *repositories.UserRepository
	jwt      *utils.JWT
}

func NewAuthService() *AuthService {
	return &AuthService{
		userRepo: repositories.NewUserRepository(),
		jwt:      utils.NewJWT(),
	}
}

func (s *AuthService) RegisterUser(user models.User) error {
	existingUser, err := s.userRepo.FindByEmail(user.Email)
	if err != nil && err.Error() != "record not found" {
		return err
	}
	if existingUser != nil {
		return errors.ErrUserAlreadyExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	err = s.userRepo.Create(&user)
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.ErrInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.ErrInvalidCredentials
	}

	token, err := s.jwt.GenerateNewAccessToken(user.ID.String())
	if err != nil {
		return "", err
	}

	return token, nil
}
