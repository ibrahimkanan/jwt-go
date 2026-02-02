package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"jwt-go/initializers"
	"jwt-go/models"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func RequireAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get the cookie off request
		cookie, err := c.Cookie("Authorization")
		if err != nil {
			return c.String(http.StatusUnauthorized, "Unauthorized")
		}

		// Decode/validate it
		tokenString := cookie.Value
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte(os.Getenv("SECRET")), nil
		})

		if err != nil {
			return c.String(http.StatusUnauthorized, "Unauthorized")
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Check the exp
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				return c.String(http.StatusUnauthorized, "Unauthorized")
			}

			// Find the user with token sub
			var user models.User
			initializers.DB.First(&user, claims["id"])

			if user.ID == 0 {
				return c.String(http.StatusUnauthorized, "Unauthorized")
			}

			// Attach to req
			c.Set("user", user)

			// Continue
			return next(c)
		} else {
			return c.String(http.StatusUnauthorized, "Unauthorized")
		}
	}
}
