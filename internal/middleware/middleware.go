package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gitlab.com/p9359/backend-prob/febry-go/internal/helper"
)

type Auth struct {
	Uuid  string
	Name  string
	email string
	phone string
}

func AuthHandler(auth AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			errors := gin.H{"error": "Authorization Header Invalid"}
			response := helper.APIResponse(http.StatusUnauthorized, false, "Unauthorize", nil, nil, errors)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		var tokenString string
		arrayHeader := strings.Split(authHeader, " ")
		if len(arrayHeader) == 2 {
			tokenString = arrayHeader[1]
		}

		token, err := auth.ValidateOnlineToken(tokenString)
		if err != nil {
			errors := gin.H{"error": "Authorization Header Invalid"}
			response := helper.APIResponse(http.StatusUnauthorized, false, "Unauthorize", nil, nil, errors)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		_, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			err := gin.H{"errors": "Cannot claim token."}
			response := helper.APIResponse(http.StatusUnauthorized, false, "Unauthorized", nil, nil, err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		payload := Auth{}

		c.Set("payload", payload)
		c.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache_Control, X-Requested-With")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	}
}
