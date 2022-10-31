package handlers

import (
	"Skipper_cms_users/pkg/models"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
	rolesCtx            = "userRoles"
)

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Request.Header.Del("Origin")
			c.Next()
		}
	}
}

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не авторизован"})
		c.Abort()
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Пользователь не авторизован"})
		c.Abort()
		return
	}
	if len(headerParts[1]) == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Ошибка чтения токена"})
		c.Abort()
		return
	}
	userId, _, err := ParseToken(headerParts[1])
	userRoles, err := h.services.GetUserRoles(userId)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Ошибка чтения токена"})
		c.Abort()
		return
	}
	accessRoles := []string{"super_admin"}

	isAccess := false
	for _, value := range userRoles {
		for _, accessRole := range accessRoles {
			if value.Name == accessRole {
				isAccess = true
			}
		}
	}
	if isAccess {
		c.Set(userCtx, userId)
		c.Set(rolesCtx, userRoles)
	} else {
		c.JSON(http.StatusForbidden, gin.H{"error": "Ошибка доступа"})
		c.Abort()
		return
	}
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId uint          `json:"user_id"`
	Roles  []models.Role `json:"roles"`
}

const (
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
)

func ParseToken(accessToken string) (uint, []models.Role, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, nil, err
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, nil, errors.New("token claims are not of type *tokenClaims")
	}
	return claims.UserId, claims.Roles, nil
}
