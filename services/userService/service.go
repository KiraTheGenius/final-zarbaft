package userService

import (
	"bankai/models"
	"bankai/repository/userRepository"
	"bankai/utils"
)

type UserService interface {
	CreateUser(user *models.User) error
	GetUser(username string) (*models.User, error)
	DeleteUser(username string) error
	GetSlider(locale string, tag string) ([]map[string]interface{}, error)
	GetCollection(locale string, slug string) (*models.Collection, error)
}

func NewUserService(userRepo userRepository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

type userService struct {
	userRepository userRepository.UserRepository
}

func (s *userService) CreateUser(user *models.User) error {
	hashed, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashed
	return s.userRepository.CreateUser(user)
}

func (s *userService) GetUser(username string) (*models.User, error) {
	return s.userRepository.GetUserByUsername(username)
}

func (s *userService) GetSlider(locale string, tag string) ([]map[string]interface{}, error) {
	return s.userRepository.GetSlider(locale, tag)
}

func (s *userService) GetCollection(locale string, slug string) (*models.Collection, error) {
	return s.userRepository.GetCollection(locale, slug)
}

func (s *userService) DeleteUser(username string) error {
	return nil
}
