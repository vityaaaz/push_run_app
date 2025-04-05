package services

import (
	"push_run_app/backend/models"
	"time"
)

type AchievementService struct {
	achievements []models.Achievement
	userAchievements []models.UserAchievement
}

func NewAchievementService() *AchievementService {
	service := &AchievementService{
		achievements: make([]models.Achievement, 0),
		userAchievements: make([]models.UserAchievement, 0),
	}
	
	// Initialize default achievements
	service.initializeDefaultAchievements()
	
	return service
}

func (s *AchievementService) initializeDefaultAchievements() {
	defaultAchievements := []models.Achievement{
		{
			ID:          "first-run",
			Name:        "Первый шаг",
			Description: "Завершите первую пробежку",
			Icon:        "🏃",
			XP:          100,
			Type:        "count",
			Target:      1,
			CreatedAt:   time.Now(),
		},
		{
			ID:          "5k-run",
			Name:        "5 километров",
			Description: "Пробегите 5 километров за одну тренировку",
			Icon:        "🏅",
			XP:          200,
			Type:        "distance",
			Target:      5,
			CreatedAt:   time.Now(),
		},
		{
			ID:          "10k-run",
			Name:        "10 километров",
			Description: "Пробегите 10 километров за одну тренировку",
			Icon:        "🎖️",
			XP:          500,
			Type:        "distance",
			Target:      10,
			CreatedAt:   time.Now(),
		},
		{
			ID:          "streak-7",
			Name:        "Неделя активности",
			Description: "Тренируйтесь 7 дней подряд",
			Icon:        "🔥",
			XP:          300,
			Type:        "streak",
			Target:      7,
			CreatedAt:   time.Now(),
		},
		{
			ID:          "total-100k",
			Name:        "100 километров",
			Description: "Пробегите в сумме 100 километров",
			Icon:        "🏆",
			XP:          1000,
			Type:        "total_distance",
			Target:      100,
			CreatedAt:   time.Now(),
		},
	}

	s.achievements = defaultAchievements
}

func (s *AchievementService) CheckAchievements(userID string, activity *models.Activity) ([]models.Achievement, error) {
	var unlockedAchievements []models.Achievement

	// Get user's current achievements
	userAchievements := s.getUserAchievements(userID)

	// Check each achievement
	for _, achievement := range s.achievements {
		// Skip if already completed
		if s.isAchievementCompleted(userAchievements, achievement.ID) {
			continue
		}

		// Get current progress
		progress := s.calculateProgress(userID, achievement, activity)

		// Check if achievement is unlocked
		if progress >= achievement.Target {
			// Create new user achievement
			userAchievement := models.UserAchievement{
				ID:            generateID(),
				UserID:        userID,
				AchievementID: achievement.ID,
				Progress:      progress,
				Completed:     true,
				CompletedAt:   time.Now(),
				CreatedAt:     time.Now(),
			}

			s.userAchievements = append(s.userAchievements, userAchievement)
			unlockedAchievements = append(unlockedAchievements, achievement)
		} else {
			// Update progress if achievement exists
			s.updateProgress(userID, achievement.ID, progress)
		}
	}

	return unlockedAchievements, nil
}

func (s *AchievementService) GetUserAchievements(userID string) ([]models.UserAchievement, error) {
	return s.getUserAchievements(userID), nil
}

func (s *AchievementService) GetAchievementDetails(achievementID string) (*models.Achievement, error) {
	for _, achievement := range s.achievements {
		if achievement.ID == achievementID {
			return &achievement, nil
		}
	}
	return nil, nil
}

// Helper functions
func (s *AchievementService) getUserAchievements(userID string) []models.UserAchievement {
	var userAchievements []models.UserAchievement
	for _, ua := range s.userAchievements {
		if ua.UserID == userID {
			userAchievements = append(userAchievements, ua)
		}
	}
	return userAchievements
}

func (s *AchievementService) isAchievementCompleted(userAchievements []models.UserAchievement, achievementID string) bool {
	for _, ua := range userAchievements {
		if ua.AchievementID == achievementID && ua.Completed {
			return true
		}
	}
	return false
}

func (s *AchievementService) calculateProgress(userID string, achievement models.Achievement, activity *models.Activity) float64 {
	// This is a simplified version. In a real app, you would need to implement
	// different calculation methods based on achievement type
	switch achievement.Type {
	case "distance":
		return activity.Distance
	case "count":
		return 1
	case "streak":
		// Implement streak calculation
		return 0
	case "total_distance":
		// Implement total distance calculation
		return 0
	default:
		return 0
	}
}

func (s *AchievementService) updateProgress(userID string, achievementID string, progress float64) {
	for i, ua := range s.userAchievements {
		if ua.UserID == userID && ua.AchievementID == achievementID {
			s.userAchievements[i].Progress = progress
			return
		}
	}

	// If achievement doesn't exist, create it
	userAchievement := models.UserAchievement{
		ID:            generateID(),
		UserID:        userID,
		AchievementID: achievementID,
		Progress:      progress,
		Completed:     false,
		CreatedAt:     time.Now(),
	}

	s.userAchievements = append(s.userAchievements, userAchievement)
} 