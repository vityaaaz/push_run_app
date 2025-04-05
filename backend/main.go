package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"push_run_app/backend/models"
	"push_run_app/backend/services"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

var (
	userService       = services.NewUserService()
	activityService   = services.NewActivityService()
	achievementService = services.NewAchievementService()
	socialService     = services.NewSocialService()
	challengeService  = services.NewChallengeService()
	mapService        = services.NewMapService(os.Getenv("YANDEX_MAPS_API_KEY"))
)

func main() {
	router := gin.Default()

	// Configure CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// API routes
	api := router.Group("/api")
	{
		// Auth routes
		auth := api.Group("/auth")
		{
			auth.POST("/login", handleLogin)
			auth.POST("/register", handleRegister)
		}

		// Activity routes
		activities := api.Group("/activities")
		{
			activities.GET("/", handleGetActivities)
			activities.POST("/", handleCreateActivity)
			activities.GET("/:id", handleGetActivity)
			activities.DELETE("/:id", handleDeleteActivity)
			activities.POST("/:id/like", handleLikeActivity)
			activities.POST("/:id/comment", handleCommentActivity)
		}

		// User routes
		users := api.Group("/users")
		{
			users.GET("/profile", handleGetProfile)
			users.PUT("/profile", handleUpdateProfile)
			users.GET("/achievements", handleGetUserAchievements)
			users.GET("/stats", handleGetUserStats)
		}

		// Social routes
		social := api.Group("/social")
		{
			social.GET("/feed", handleGetFeed)
			social.POST("/friends/request", handleSendFriendRequest)
			social.POST("/friends/accept", handleAcceptFriendRequest)
			social.POST("/friends/reject", handleRejectFriendRequest)
			social.GET("/friends", handleGetFriends)
			social.GET("/friends/requests", handleGetFriendRequests)
		}

		// Challenge routes
		challenges := api.Group("/challenges")
		{
			challenges.GET("/", handleGetChallenges)
			challenges.POST("/", handleCreateChallenge)
			challenges.GET("/:id", handleGetChallengeDetails)
			challenges.POST("/:id/join", handleJoinChallenge)
			challenges.GET("/:id/participants", handleGetChallengeParticipants)
			challenges.GET("/user", handleGetUserChallenges)
		}

		// Map routes
		maps := api.Group("/maps")
		{
			maps.POST("/route", handleGetRouteInfo)
			maps.POST("/elevation", handleGetElevation)
			maps.GET("/weather", handleGetWeather)
		}
	}

	// Start server
	log.Println("Server starting on :8080...")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

// Handler functions
func handleLogin(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := userService.Login(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	if response == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, response)
}

func handleRegister(c *gin.Context) {
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := userService.Register(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func handleGetActivities(c *gin.Context) {
	// In a real app, get userID from JWT token
	userID := "dummy-user-id"
	activities, err := activityService.GetActivities(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, activities)
}

func handleCreateActivity(c *gin.Context) {
	var req models.CreateActivityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// In a real app, get userID from JWT token
	userID := "dummy-user-id"
	activity, err := activityService.CreateActivity(userID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusCreated, activity)
}

func handleGetActivity(c *gin.Context) {
	activityID := c.Param("id")
	// In a real app, get userID from JWT token
	userID := "dummy-user-id"

	activity, err := activityService.GetActivity(userID, activityID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	if activity == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Activity not found"})
		return
	}

	c.JSON(http.StatusOK, activity)
}

func handleDeleteActivity(c *gin.Context) {
	activityID := c.Param("id")
	// In a real app, get userID from JWT token
	userID := "dummy-user-id"

	err := activityService.DeleteActivity(userID, activityID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.Status(http.StatusNoContent)
}

func handleGetProfile(c *gin.Context) {
	// In a real app, get userID from JWT token
	userID := "dummy-user-id"
	user, err := userService.GetProfile(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func handleUpdateProfile(c *gin.Context) {
	var req models.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// In a real app, get userID from JWT token
	userID := "dummy-user-id"
	user, err := userService.UpdateProfile(userID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// New handler functions for social features
func handleGetFeed(c *gin.Context) {
	// In a real app, get userID from JWT token
	userID := "dummy-user-id"
	limit := 20 // Default limit

	feedItems, err := socialService.GetFeed(userID, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, feedItems)
}

func handleSendFriendRequest(c *gin.Context) {
	var req struct {
		ToUserID string `json:"toUserId" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// In a real app, get userID from JWT token
	fromUserID := "dummy-user-id"
	friendRequest, err := socialService.SendFriendRequest(fromUserID, req.ToUserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusCreated, friendRequest)
}

func handleAcceptFriendRequest(c *gin.Context) {
	var req struct {
		RequestID string `json:"requestId" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := socialService.AcceptFriendRequest(req.RequestID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.Status(http.StatusOK)
}

func handleRejectFriendRequest(c *gin.Context) {
	var req struct {
		RequestID string `json:"requestId" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := socialService.RejectFriendRequest(req.RequestID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.Status(http.StatusOK)
}

func handleGetFriends(c *gin.Context) {
	// In a real app, get userID from JWT token
	userID := "dummy-user-id"
	friends, err := socialService.GetFriends(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, friends)
}

func handleGetFriendRequests(c *gin.Context) {
	// In a real app, get userID from JWT token
	userID := "dummy-user-id"
	requests, err := socialService.GetFriendRequests(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, requests)
}

func handleGetUserAchievements(c *gin.Context) {
	// In a real app, get userID from JWT token
	userID := "dummy-user-id"
	achievements, err := achievementService.GetUserAchievements(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, achievements)
}

func handleGetUserStats(c *gin.Context) {
	// In a real app, get userID from JWT token
	userID := "dummy-user-id"
	stats, err := activityService.GetUserStats(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, stats)
}

func handleLikeActivity(c *gin.Context) {
	activityID := c.Param("id")
	// In a real app, get userID from JWT token
	userID := "dummy-user-id"

	err := socialService.LikeActivity(userID, activityID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.Status(http.StatusOK)
}

func handleCommentActivity(c *gin.Context) {
	activityID := c.Param("id")
	var req struct {
		Text string `json:"text" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// In a real app, get userID from JWT token
	userID := "dummy-user-id"
	err := socialService.CommentOnActivity(userID, activityID, req.Text)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.Status(http.StatusOK)
}

// New handler functions for challenges
func handleGetChallenges(c *gin.Context) {
	challenges, err := challengeService.GetChallenges()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, challenges)
}

func handleCreateChallenge(c *gin.Context) {
	var req models.CreateChallengeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// In a real app, get userID from JWT token
	userID := "dummy-user-id"
	challenge, err := challengeService.CreateChallenge(userID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusCreated, challenge)
}

func handleGetChallengeDetails(c *gin.Context) {
	challengeID := c.Param("id")
	challenge, err := challengeService.GetChallengeDetails(challengeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	if challenge == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Challenge not found"})
		return
	}

	c.JSON(http.StatusOK, challenge)
}

func handleJoinChallenge(c *gin.Context) {
	challengeID := c.Param("id")
	var req models.JoinChallengeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// In a real app, get userID from JWT token
	userID := "dummy-user-id"
	err := challengeService.JoinChallenge(userID, challengeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.Status(http.StatusOK)
}

func handleGetChallengeParticipants(c *gin.Context) {
	challengeID := c.Param("id")
	participants, err := challengeService.GetChallengeParticipants(challengeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, participants)
}

func handleGetUserChallenges(c *gin.Context) {
	// In a real app, get userID from JWT token
	userID := "dummy-user-id"
	challenges, err := challengeService.GetUserChallenges(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, challenges)
}

// New handler functions for map features
func handleGetRouteInfo(c *gin.Context) {
	var coordinates []models.Coordinate
	if err := c.ShouldBindJSON(&coordinates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	routeInfo, err := mapService.GetRouteInfo(coordinates)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get route info"})
		return
	}

	c.JSON(http.StatusOK, routeInfo)
}

func handleGetElevation(c *gin.Context) {
	var coordinates []models.Coordinate
	if err := c.ShouldBindJSON(&coordinates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	elevations, err := mapService.GetElevation(coordinates)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get elevation data"})
		return
	}

	c.JSON(http.StatusOK, elevations)
}

func handleGetWeather(c *gin.Context) {
	var coordinate models.Coordinate
	if err := c.ShouldBindJSON(&coordinate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	weather, err := mapService.GetWeather(coordinate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get weather data"})
		return
	}

	c.JSON(http.StatusOK, weather)
} 