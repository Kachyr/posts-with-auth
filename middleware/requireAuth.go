package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Kachyr/crud/initializers"
	"github.com/Kachyr/crud/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {

	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		println("Error: Unauthorized, ", err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, validationErr := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if validationErr != nil {
		println("Error: Unauthorized, ", err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		expTime, ok := claims["exp"].(float64)
		if isExpired := float64(time.Now().Unix()) >= expTime; isExpired && ok {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		var user models.User
		initializers.DB.First(&user, claims["sub"])
		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Set("user", user)

		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}
