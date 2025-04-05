package models

import (
	"time"
)

type User struct {
	ID            string    `json:"id"`
	Email         string    `json:"email"`
	Password      string    `json:"-"` // Password hash, not exposed in JSON
	FirstName     string    `json:"firstName"`
	LastName      string    `json:"lastName"`
	Username      string    `json:"username"`
	Avatar        string    `json:"avatar"`
	Bio           string    `json:"bio"`
	TotalDistance float64   `json:"totalDistance"`
	TotalTime     int64     `json:"totalTime"`
	Level         int       `json:"level"`
	XP            int       `json:"xp"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

type RegisterRequest struct {
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6"`
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Username  string `json:"username" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

type UpdateProfileRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
	Bio       string `json:"bio"`
	Avatar    string `json:"avatar"`
}

type FriendRequest struct {
	ID        string    `json:"id"`
	FromUser  string    `json:"fromUser"`
	ToUser    string    `json:"toUser"`
	Status    string    `json:"status"` // pending, accepted, rejected
	CreatedAt time.Time `json:"createdAt"`
}

type Achievement struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Icon        string    `json:"icon"`
	XP          int       `json:"xp"`
	Type        string    `json:"type"` // distance, time, streak, etc.
	Target      float64   `json:"target"`
	CreatedAt   time.Time `json:"createdAt"`
}

type UserAchievement struct {
	ID           string    `json:"id"`
	UserID       string    `json:"userId"`
	AchievementID string   `json:"achievementId"`
	Progress     float64   `json:"progress"`
	Completed    bool      `json:"completed"`
	CompletedAt  time.Time `json:"completedAt"`
	CreatedAt    time.Time `json:"createdAt"`
} 