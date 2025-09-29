package middleware

import (
	"net/http"
	"strings"

	"miniProject/pkg/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware memverifikasi token JWT dari header Authorization
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Ambil Header Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Token tidak tersedia"})
			return
		}

		// 2. Cek Format (Bearer <token>)
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Format token harus 'Bearer <token>'"})
			return
		}

		encodedToken := parts[1]

		// 3. Validasi Token
		token, err := utils.ValidateToken(encodedToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Token tidak valid atau kedaluwarsa"})
			return
		}

		// 4. Ambil Username dari Klaim dan setel di Context
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Klaim token tidak valid"})
			return
		}

		username := claims["username"].(string)
		
		// Setel username di context untuk digunakan di handler (misal: created_by/modified_by)
		c.Set("username", username) 

		// Lanjutkan ke handler berikutnya
		c.Next()
	}
}