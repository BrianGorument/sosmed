package users

import (
	"errors"
	"fmt"
	"sosmed/shared/utils"

	"github.com/go-playground/validator/v10"
)

// userService struct
type userService struct {
	repo IUserRepository
}

// NewUserService (Dependency Injection)
func NewUserService(repo IUserRepository) IUserService {
	return &userService{repo}
}

// Implementasi UserService
func (s *userService) RegisterUser(req CreateUserRequest) (*UserResponse, error) {
	existingUsers, _ := s.repo.FindAll()
	for _, u := range existingUsers {
		if u.Email == req.Email {
			return nil, errors.New("email already registered")
		}
	}

	hashedPassword, _ := utils.HashPassword(req.Password)
	validate := validator.New()
	validEmail := req.Email

	err := validate.Struct(validEmail)
	if err != nil {
		fmt.Println("Invalid email:", err)
		return nil, errors.New("invalid Email Format")
	}
	user := User{
		Username:  req.Username,
		Email:     validEmail,
		Password:  hashedPassword,
		Role_code: req.Role_code,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}
	err = s.repo.Create(&user)
	if err != nil {
		return nil, err
	}

	return &UserResponse{ID: user.ID, Username: user.Username, Email: user.Email}, nil
}

func (s *userService) GetAllUsers() ([]UserResponse, error) {
	users, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	var userResponses []UserResponse
	for _, user := range users {
		userResponses = append(userResponses, UserResponse{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		})
	}

	return userResponses, nil
}
