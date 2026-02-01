package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ugenius/backend/internal/models"
	"github.com/ugenius/backend/internal/utils"
	"github.com/ugenius/backend/pkg/response"
)

const (
	AuthorizationHeader = "Authorization"
	BearerPrefix        = "Bearer "
	ContextUserID       = "user_id"
	ContextUserEmail    = "user_email"
	ContextUserRole     = "user_role"
	ContextCampusID     = "campus_id"
)

// AuthMiddleware validates JWT tokens and sets user info in context
func AuthMiddleware(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader(AuthorizationHeader)
		if authHeader == "" {
			response.Unauthorized(c, "Authorization header required")
			c.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, BearerPrefix) {
			response.Unauthorized(c, "Invalid authorization format")
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, BearerPrefix)
		claims, err := utils.ValidateToken(tokenString, jwtSecret)
		if err != nil {
			if err == utils.ErrExpiredToken {
				response.Unauthorized(c, "Token has expired")
			} else {
				response.Unauthorized(c, "Invalid token")
			}
			c.Abort()
			return
		}

		// Set user info in context
		c.Set(ContextUserID, claims.UserID)
		c.Set(ContextUserEmail, claims.Email)
		c.Set(ContextUserRole, claims.Role)
		if claims.CampusID != nil {
			c.Set(ContextCampusID, *claims.CampusID)
		}

		c.Next()
	}
}

// RequireRole middleware checks if user has required role
func RequireRole(roles ...models.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get(ContextUserRole)
		if !exists {
			response.Unauthorized(c, "User role not found")
			c.Abort()
			return
		}

		role := userRole.(models.Role)
		for _, allowedRole := range roles {
			if role == allowedRole {
				c.Next()
				return
			}
		}

		response.Forbidden(c, "Insufficient permissions")
		c.Abort()
	}
}

// RequireAdmin middleware ensures user is a campus admin or super admin
func RequireAdmin() gin.HandlerFunc {
	return RequireRole(models.RoleCampusAdmin, models.RoleSuperAdmin)
}

// RequireSuperAdmin middleware ensures user is a super admin
func RequireSuperAdmin() gin.HandlerFunc {
	return RequireRole(models.RoleSuperAdmin)
}

// GetUserID gets the user ID from context
func GetUserID(c *gin.Context) (uuid.UUID, bool) {
	userID, exists := c.Get(ContextUserID)
	if !exists {
		return uuid.Nil, false
	}
	return userID.(uuid.UUID), true
}

// GetUserRole gets the user role from context
func GetUserRole(c *gin.Context) (models.Role, bool) {
	role, exists := c.Get(ContextUserRole)
	if !exists {
		return "", false
	}
	return role.(models.Role), true
}

// GetCampusID gets the campus ID from context
func GetCampusID(c *gin.Context) (uuid.UUID, bool) {
	campusID, exists := c.Get(ContextCampusID)
	if !exists {
		return uuid.Nil, false
	}
	return campusID.(uuid.UUID), true
}

// OptionalAuth middleware tries to authenticate but doesn't require it
func OptionalAuth(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader(AuthorizationHeader)
		if authHeader == "" || !strings.HasPrefix(authHeader, BearerPrefix) {
			c.Next()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, BearerPrefix)
		claims, err := utils.ValidateToken(tokenString, jwtSecret)
		if err == nil {
			c.Set(ContextUserID, claims.UserID)
			c.Set(ContextUserEmail, claims.Email)
			c.Set(ContextUserRole, claims.Role)
			if claims.CampusID != nil {
				c.Set(ContextCampusID, *claims.CampusID)
			}
		}

		c.Next()
	}
}

