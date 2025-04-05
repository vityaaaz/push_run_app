package models

import (
	"time"
)

type ActivityType string

const (
	ActivityTypeRun   ActivityType = "run"
	ActivityTypeBike  ActivityType = "bike"
)

type Activity struct {
	ID          string      `json:"id"`
	UserID      string      `json:"userId"`
	Type        ActivityType `json:"type"`
	StartTime   time.Time   `json:"startTime"`
	EndTime     time.Time   `json:"endTime"`
	Distance    float64     `json:"distance"`    // in kilometers
	Duration    int64       `json:"duration"`    // in seconds
	Pace        float64     `json:"pace"`        // in minutes per kilometer
	Coordinates []Coordinate `json:"coordinates"`
}

type Coordinate struct {
	Latitude  float64   `json:"latitude"`
	Longitude float64   `json:"longitude"`
	Timestamp time.Time `json:"timestamp"`
}

type CreateActivityRequest struct {
	Type        ActivityType `json:"type" binding:"required"`
	StartTime   time.Time   `json:"startTime" binding:"required"`
	EndTime     time.Time   `json:"endTime" binding:"required"`
	Distance    float64     `json:"distance" binding:"required"`
	Coordinates []Coordinate `json:"coordinates" binding:"required"`
} 