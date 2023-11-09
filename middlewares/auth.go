package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Kunniii/go_gin_gorm_test/internal"
	"github.com/Kunniii/go_gin_gorm_test/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func CheckAuth(context *gin.Context) {
	authorization := context.GetHeader("Authorization")
	if authorization == "" {
		fmt.Println("Auth Empty")
		context.AbortWithStatus(http.StatusUnauthorized)
	}

	token, err := jwt.Parse(authorization, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		fmt.Println("Parse Error")
		context.AbortWithStatus(http.StatusUnauthorized)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// check expires time
		if time.Now().Unix() > claims["expires"].(int64) {
			fmt.Println("Token Expired")
			context.AbortWithStatus(http.StatusUnauthorized)
		}

		var user models.User
		if err := internal.DB.First(&user, "email = ?", claims["email"]).Error; err != nil {
			fmt.Println("User Not found")
			context.AbortWithStatus(http.StatusUnauthorized)
		}

		context.Set("userId", user.ID)
		context.Set("userEmail", user.ID)

		context.Next()
	} else {
		fmt.Println("Claims not OK")
		context.AbortWithStatus(http.StatusUnauthorized)
	}
}
