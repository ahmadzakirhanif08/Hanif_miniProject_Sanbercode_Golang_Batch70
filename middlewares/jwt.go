package middlewares

import (
    "time"
    "github.com/dgrijalva/jwt-go"
	"fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)

var jwtSecret = []byte("UraniumPlutonium!!!")

// Claim struct untuk JWT (payload)
type UserClaims struct {
    UserID   int    `json:"user_id"`
    Username string `json:"username"`
    jwt.StandardClaims
}

func GenerateToken(userID int, username string) (string, error) {
    expirationTime := time.Now().Add(1 * time.Hour)
    
    claims := &UserClaims{
        UserID:   userID,
        Username: username,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtSecret)
    
    return tokenString, err
}

// Middleware untuk Memverifikasi Token
func JWTAuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        
        if authHeader == "" || len(authHeader) < 7 || authHeader[:6] != "Bearer" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header not provided or invalid format (Missing Bearer token)"})
            c.Abort()
            return
        }

        tokenString := authHeader[7:] 
        claims := &UserClaims{}
        
        token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
            }
            return jwtSecret, nil
        })

        // 3. Tangani Error Validasi
        if err != nil || !token.Valid {
            if err != nil && err.(*jwt.ValidationError).Errors == jwt.ValidationErrorExpired {
                 c.JSON(http.StatusUnauthorized, gin.H{"error": "Token has expired"})
            } else {
                 c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            }
            c.Abort()
            return
        }
        c.Set("user_id", claims.UserID)
        c.Set("username", claims.Username)
        c.Next()
    }
}