package services

import (
	"push_run_app/backend/models"
	"time"
)

type UserService struct {
	// In a real application, this would have a database connection
	users []models.User
}

func NewUserService() *UserService {
	return &UserService{
		users: make([]models.User, 0),
	}
}

func (s *UserService) Register(req models.RegisterRequest) (*models.User, error) {
	// In a real app, hash the password and check for existing user
	user := models.User{
		ID:        generateID(),
		Email:     req.Email,
		Password:  req.Password, // In a real app, this would be hashed
		FirstName: req.FirstName,
		LastName:  req.LastName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	s.users = append(s.users, user)
	return &user, nil
}

func (s *UserService) Login(req models.LoginRequest) (*models.LoginResponse, error) {
	// In a real app, verify password hash and generate JWT token
	for _, user := range s.users {
		if user.Email == req.Email && user.Password == req.Password {
			return &models.LoginResponse{
				Token: "dummy-token", // In a real app, generate JWT
				User:  user,
			}, nil
		}
	}
	return nil, nil // In a real app, return proper error
}

func (s *UserService) GetProfile(userID string) (*models.User, error) {
	for _, user := range s.users {
		if user.ID == userID {
			return &user, nil
		}
	}
	return nil, nil // In a real app, return proper error
}

func (s *UserService) UpdateProfile(userID string, req models.UpdateProfileRequest) (*models.User, error) {
	for i, user := range s.users {
		if user.ID == userID {
			if req.FirstName != "" {
				s.users[i].FirstName = req.FirstName
			}
			if req.LastName != "" {
				s.users[i].LastName = req.LastName
			}
			s.users[i].UpdatedAt = time.Now()
			return &s.users[i], nil
		}
	}
	return nil, nil // In a real app, return proper error
}

// Helper function
func generateID() string {
	return time.Now().Format("20060102150405")
} 