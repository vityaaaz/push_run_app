package services

import (
	"push_run_app/backend/models"
	"sort"
	"time"
)

type ChallengeService struct {
	challenges        []models.Challenge
	participants      []models.ChallengeParticipant
}

func NewChallengeService() *ChallengeService {
	return &ChallengeService{
		challenges:   make([]models.Challenge, 0),
		participants: make([]models.ChallengeParticipant, 0),
	}
}

func (s *ChallengeService) CreateChallenge(userID string, req models.CreateChallengeRequest) (*models.Challenge, error) {
	challenge := models.Challenge{
		ID:          generateID(),
		Name:        req.Name,
		Description: req.Description,
		Type:        req.Type,
		Target:      req.Target,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
		CreatedBy:   userID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	s.challenges = append(s.challenges, challenge)
	return &challenge, nil
}

func (s *ChallengeService) JoinChallenge(userID string, challengeID string) error {
	// Check if challenge exists
	var challenge *models.Challenge
	for _, c := range s.challenges {
		if c.ID == challengeID {
			challenge = &c
			break
		}
	}
	if challenge == nil {
		return nil // In a real app, return proper error
	}

	// Check if already participating
	for _, p := range s.participants {
		if p.UserID == userID && p.ChallengeID == challengeID {
			return nil // In a real app, return proper error
		}
	}

	participant := models.ChallengeParticipant{
		ID:          generateID(),
		ChallengeID: challengeID,
		UserID:      userID,
		Progress:    0,
		Rank:        0,
		JoinedAt:    time.Now(),
		UpdatedAt:   time.Now(),
	}

	s.participants = append(s.participants, participant)
	return nil
}

func (s *ChallengeService) UpdateProgress(userID string, challengeID string, progress float64) error {
	for i, p := range s.participants {
		if p.UserID == userID && p.ChallengeID == challengeID {
			s.participants[i].Progress = progress
			s.participants[i].UpdatedAt = time.Now()
			s.updateRanks(challengeID)
			return nil
		}
	}
	return nil // In a real app, return proper error
}

func (s *ChallengeService) GetChallenges() ([]models.Challenge, error) {
	return s.challenges, nil
}

func (s *ChallengeService) GetChallengeDetails(challengeID string) (*models.Challenge, error) {
	for _, c := range s.challenges {
		if c.ID == challengeID {
			return &c, nil
		}
	}
	return nil, nil // In a real app, return proper error
}

func (s *ChallengeService) GetUserChallenges(userID string) ([]models.Challenge, error) {
	var userChallenges []models.Challenge
	for _, p := range s.participants {
		if p.UserID == userID {
			for _, c := range s.challenges {
				if c.ID == p.ChallengeID {
					userChallenges = append(userChallenges, c)
					break
				}
			}
		}
	}
	return userChallenges, nil
}

func (s *ChallengeService) GetChallengeParticipants(challengeID string) ([]models.ChallengeParticipant, error) {
	var participants []models.ChallengeParticipant
	for _, p := range s.participants {
		if p.ChallengeID == challengeID {
			participants = append(participants, p)
		}
	}
	return participants, nil
}

func (s *ChallengeService) updateRanks(challengeID string) {
	// Get all participants for the challenge
	var participants []models.ChallengeParticipant
	for _, p := range s.participants {
		if p.ChallengeID == challengeID {
			participants = append(participants, p)
		}
	}

	// Sort by progress in descending order
	sort.Slice(participants, func(i, j int) bool {
		return participants[i].Progress > participants[j].Progress
	})

	// Update ranks
	for i := range participants {
		participants[i].Rank = i + 1
	}

	// Update in service
	for i, p := range s.participants {
		if p.ChallengeID == challengeID {
			for _, updated := range participants {
				if updated.UserID == p.UserID {
					s.participants[i] = updated
					break
				}
			}
		}
	}
} 