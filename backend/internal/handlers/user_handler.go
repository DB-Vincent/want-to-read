package handlers

import (
	"net/http"
	"strconv"

	"github.com/DB-Vincent/want-to-read/internal/models"
	"github.com/DB-Vincent/want-to-read/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
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

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type EditUserRequest struct {
	Username string `json:"username"`
	IsSuper  bool   `json:"is_super"`
}

//	@Summary		Authenticate user
//	@Description	Authenticate user and return JWT token
//	@Tags			users
//	@Produce		json
//	@Param			login	body		LoginRequest	true	"Login request"
//	@Success		200		{object}	map[string]string
//	@Failure		400		{string}	string
//	@Failure		401		{string}	string
//	@Failure		500		{string}	string
//	@Router			/api/login [post]
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

//	@Summary		Register user
//	@Description	Register a new user
//	@Tags			users
//	@Produce		json
//	@Param			register	body		LoginRequest	true	"Register request"
//	@Success		200			{object}	map[string]interface{}
//	@Failure		500			{string}	string
//	@Router			/api/register [post]
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

//	@Summary		List users
//	@Description	List all users
//	@Tags			users
//	@Produce		json
//	@Success		200	{array}		models.User
//	@Failure		500	{string}	string
//	@Router			/api/users [get]
func (h *UserHandler) ListUsers(c *gin.Context) {
	users, err := h.userService.ListUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve users", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

//	@Summary		Change password
//	@Description	Change user password
//	@Tags			users
//	@Produce		json
//	@Param			change_password	body		ChangePasswordRequest	true	"Change password request"
//	@Success		200				{object}	map[string]string
//	@Failure		400				{string}	string
//	@Failure		401				{string}	string
//	@Failure		500				{string}	string
//	@Router			/api/change_password [post]
func (h *UserHandler) ChangePassword(c *gin.Context) {
	var req ChangePasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()})
		return
	}

	token, err := h.extractToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing or invalid token", "details": err.Error()})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		return
	}
	userID, ok := claims["user_id"].(float64)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		return
	}
	dbUser, err := h.userService.GetUserByID(uint(userID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve user", "details": err.Error()})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(req.OldPassword)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials", "details": err.Error()})
		return
	}

	if err := h.userService.ChangePassword(&models.User{
		ID:       dbUser.ID,
		Username: dbUser.Username,
		Password: req.NewPassword,
		IsSuper:  dbUser.IsSuper,
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not change password", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password changed successfully"})
}

//	@Summary		Edit user
//	@Description	Edit user details
//	@Tags			users
//	@Produce		json
//	@Param			id			path		int				true	"User ID"
//	@Param			edit_user	body		EditUserRequest	true	"Edit user request"
//	@Success		200			{object}	map[string]string
//	@Failure		400			{string}	string
//	@Failure		401			{string}	string
//	@Failure		404			{string}	string
//	@Failure		500			{string}	string
//	@Router			/api/user/{id} [patch]
func (h *UserHandler) EditUser(c *gin.Context) {
	var req EditUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request", "details": err.Error()})
		return
	}

	idParam := c.Param("id")
	idUint64, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID", "details": err.Error()})
		return
	}

	userId := uint(idUint64)
	user, err := h.userService.GetUserByID(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Could not retrieve user", "details": err.Error()})
		return
	}

	user.Username = req.Username
	user.IsSuper = req.IsSuper

	if _, err := h.userService.UpdateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update user", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
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
