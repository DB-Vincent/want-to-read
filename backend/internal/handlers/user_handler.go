package handlers

import (
	"net/http"

	"github.com/DB-Vincent/want-to-read/internal/models"
	"github.com/DB-Vincent/want-to-read/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// @Summary		Authenticate user
// @Description	Authenticate user and return JWT token
// @Tags		users
// @Produce		json
// @Param		login	body		LoginRequest	true	"Login request"
// @Success		200		{object}	map[string]string
// @Failure		400		{string}	string
// @Failure		401		{string}	string
// @Failure		500		{string}	string
// @Router		/api/login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()})
		return
	}

	userInput := &models.User{
		Username: req.Username,
		Password: req.Password,
	}

	user, err := h.userService.Authenticate(userInput)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials", "details": err.Error()})
		return
	}

	token, err := h.userService.GenerateJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// @Summary		Register user
// @Description	Register a new user
// @Tags		users
// @Produce		json
// @Param		register	body		LoginRequest	true	"Register request"
// @Success		200		{object}	map[string]interface{}
// @Failure		500		{string}	string
// @Router		/api/register [post]
func (h *UserHandler) Register(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()})
		return
	}

	var newUser models.User
	newUser.Username = req.Username
	newUser.Password = req.Password

	user, err := h.userService.Register(&newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not register user", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": user.ID, "username": user.Username})
}

// Helper to extract and validate JWT token from Authorization header
func (h *UserHandler) extractToken(c *gin.Context) (*jwt.Token, error) {
	authHeader := c.GetHeader("Authorization")
	if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
		return nil, http.ErrNoCookie // Use as a generic error
	}
	tokenStr := authHeader[7:]
	return h.userService.ParseJWT(tokenStr)
}

// Middleware to check if the user is authenticated
func (h *UserHandler) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := h.extractToken(c)
		if err != nil || token == nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid token"})
			return
		}
		c.Next()
	}
}

// Middleware to check if user is a super user
func (h *UserHandler) SuperUserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := h.extractToken(c)
		if err != nil || token == nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid token"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || claims["is_super"] != true {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Super user access required"})
			return
		}

		c.Next()
	}
}
