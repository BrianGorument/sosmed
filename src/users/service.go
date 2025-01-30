package users

import "errors"

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
	// Cek apakah email sudah ada
	existingUsers, _ := s.repo.FindAll()
	for _, u := range existingUsers {
		if u.Email == req.Email {
			return nil, errors.New("email already registered")
		}
	}

	// Simpan user baru
	user := User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password, // Hashing password seharusnya di sini
	}
	err := s.repo.Create(&user)
	if err != nil {
		return nil, err
	}

	// Kembalikan response
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
