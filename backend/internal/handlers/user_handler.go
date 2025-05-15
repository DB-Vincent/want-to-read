package handlers

import (
	"net/http"

	"github.com/DB-Vincent/want-to-read/internal/services"
	"github.com/gin-gonic/gin"
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
// @Router		/login [post]
func (h *UserHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user, err := h.userService.Authenticate(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := h.userService.GenerateJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// Middleware to protect routes
func (h *UserHandler) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) < 7 || authHeader[:7] != "Bearer " {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid token"})
			return
		}

		tokenStr := authHeader[7:]
		token, err := h.userService.ParseJWT(tokenStr)
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}
		c.Next()
	}
}
