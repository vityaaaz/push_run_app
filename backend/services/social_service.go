package services

import (
	"push_run_app/backend/models"
	"sort"
	"time"
)

type SocialService struct {
	friendships []models.FriendRequest
	feedItems   []FeedItem
}

type FeedItem struct {
	ID          string    `json:"id"`
	UserID      string    `json:"userId"`
	Type        string    `json:"type"` // activity, achievement, comment
	ActivityID  string    `json:"activityId,omitempty"`
	AchievementID string  `json:"achievementId,omitempty"`
	Content     string    `json:"content"`
	CreatedAt   time.Time `json:"createdAt"`
}

func NewSocialService() *SocialService {
	return &SocialService{
		friendships: make([]models.FriendRequest, 0),
		feedItems:   make([]FeedItem, 0),
	}
}

func (s *SocialService) SendFriendRequest(fromUserID, toUserID string) (*models.FriendRequest, error) {
	// Check if request already exists
	for _, fr := range s.friendships {
		if (fr.FromUser == fromUserID && fr.ToUser == toUserID) ||
			(fr.FromUser == toUserID && fr.ToUser == fromUserID) {
			return nil, nil // In a real app, return proper error
		}
	}

	friendRequest := models.FriendRequest{
		ID:        generateID(),
		FromUser:  fromUserID,
		ToUser:    toUserID,
		Status:    "pending",
		CreatedAt: time.Now(),
	}

	s.friendships = append(s.friendships, friendRequest)
	return &friendRequest, nil
}

func (s *SocialService) AcceptFriendRequest(requestID string) error {
	for i, fr := range s.friendships {
		if fr.ID == requestID {
			s.friendships[i].Status = "accepted"
			return nil
		}
	}
	return nil // In a real app, return proper error
}

func (s *SocialService) RejectFriendRequest(requestID string) error {
	for i, fr := range s.friendships {
		if fr.ID == requestID {
			s.friendships[i].Status = "rejected"
			return nil
		}
	}
	return nil // In a real app, return proper error
}

func (s *SocialService) GetFriends(userID string) ([]string, error) {
	var friends []string
	for _, fr := range s.friendships {
		if fr.Status == "accepted" {
			if fr.FromUser == userID {
				friends = append(friends, fr.ToUser)
			} else if fr.ToUser == userID {
				friends = append(friends, fr.FromUser)
			}
		}
	}
	return friends, nil
}

func (s *SocialService) GetFriendRequests(userID string) ([]models.FriendRequest, error) {
	var requests []models.FriendRequest
	for _, fr := range s.friendships {
		if (fr.ToUser == userID || fr.FromUser == userID) && fr.Status == "pending" {
			requests = append(requests, fr)
		}
	}
	return requests, nil
}

func (s *SocialService) AddToFeed(userID string, itemType string, content string, activityID, achievementID string) error {
	feedItem := FeedItem{
		ID:           generateID(),
		UserID:       userID,
		Type:         itemType,
		ActivityID:   activityID,
		AchievementID: achievementID,
		Content:      content,
		CreatedAt:    time.Now(),
	}

	s.feedItems = append(s.feedItems, feedItem)
	return nil
}

func (s *SocialService) GetFeed(userID string, limit int) ([]FeedItem, error) {
	// Get user's friends
	friends, err := s.GetFriends(userID)
	if err != nil {
		return nil, err
	}

	// Add user's own ID to include their activities
	friends = append(friends, userID)

	// Get feed items from friends
	var feedItems []FeedItem
	for _, item := range s.feedItems {
		for _, friendID := range friends {
			if item.UserID == friendID {
				feedItems = append(feedItems, item)
			}
		}
	}

	// Sort by creation time (newest first)
	sort.Slice(feedItems, func(i, j int) bool {
		return feedItems[i].CreatedAt.After(feedItems[j].CreatedAt)
	})

	// Apply limit
	if limit > 0 && len(feedItems) > limit {
		feedItems = feedItems[:limit]
	}

	return feedItems, nil
}

func (s *SocialService) LikeActivity(userID, activityID string) error {
	// In a real app, this would update the activity's likes in the database
	return nil
}

func (s *SocialService) CommentOnActivity(userID, activityID, text string) error {
	// In a real app, this would add a comment to the activity in the database
	return nil
} 