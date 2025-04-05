package services

import (
	"push_run_app/backend/models"
	"time"
)

type ActivityService struct {
	// In a real application, this would have a database connection
	activities []models.Activity
}

func NewActivityService() *ActivityService {
	return &ActivityService{
		activities: make([]models.Activity, 0),
	}
}

func (s *ActivityService) CreateActivity(userID string, req models.CreateActivityRequest) (*models.Activity, error) {
	activity := models.Activity{
		ID:          generateID(), // In a real app, use a proper ID generator
		UserID:      userID,
		Type:        req.Type,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
		Distance:    req.Distance,
		Duration:    int64(req.EndTime.Sub(req.StartTime).Seconds()),
		Pace:        calculatePace(req.Distance, req.StartTime, req.EndTime),
		Coordinates: req.Coordinates,
	}

	s.activities = append(s.activities, activity)
	return &activity, nil
}

func (s *ActivityService) GetActivities(userID string) ([]models.Activity, error) {
	var userActivities []models.Activity
	for _, activity := range s.activities {
		if activity.UserID == userID {
			userActivities = append(userActivities, activity)
		}
	}
	return userActivities, nil
}

func (s *ActivityService) GetActivity(userID, activityID string) (*models.Activity, error) {
	for _, activity := range s.activities {
		if activity.ID == activityID && activity.UserID == userID {
			return &activity, nil
		}
	}
	return nil, nil // In a real app, return proper error
}

func (s *ActivityService) DeleteActivity(userID, activityID string) error {
	for i, activity := range s.activities {
		if activity.ID == activityID && activity.UserID == userID {
			s.activities = append(s.activities[:i], s.activities[i+1:]...)
			return nil
		}
	}
	return nil // In a real app, return proper error
}

// Helper functions
func generateID() string {
	return time.Now().Format("20060102150405")
}

func calculatePace(distance float64, startTime, endTime time.Time) float64 {
	if distance == 0 {
		return 0
	}
	duration := endTime.Sub(startTime).Minutes()
	return duration / distance
} 