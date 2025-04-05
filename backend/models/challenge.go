package models

import (
	"time"
)

type Challenge struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Type        string    `json:"type"` // distance, time, count
	Target      float64   `json:"target"`
	StartDate   time.Time `json:"startDate"`
	EndDate     time.Time `json:"endDate"`
	CreatedBy   string    `json:"createdBy"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type ChallengeParticipant struct {
	ID          string    `json:"id"`
	ChallengeID string    `json:"challengeId"`
	UserID      string    `json:"userId"`
	Progress    float64   `json:"progress"`
	Rank        int       `json:"rank"`
	JoinedAt    time.Time `json:"joinedAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type CreateChallengeRequest struct {
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"`
	Type        string    `json:"type" binding:"required"`
	Target      float64   `json:"target" binding:"required"`
	StartDate   time.Time `json:"startDate" binding:"required"`
	EndDate     time.Time `json:"endDate" binding:"required"`
}

type JoinChallengeRequest struct {
	ChallengeID string `json:"challengeId" binding:"required"`
} 