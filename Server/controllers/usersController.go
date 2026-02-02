package controllers

// usersController.go
import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"jwt-go/initializers"
	"jwt-go/models"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type jwtCustomClaims struct {
	Username string    `json:"username"`
	ID       uint      `json:"id"`
	Expiry   time.Time `json:"expiry"`
	jwt.RegisteredClaims
}

var secret = []byte(os.Getenv("SECRET"))

func createToken(user models.User) (string, error) {
	claims := jwtCustomClaims{
		Username: user.Username,
		ID:       user.ID,
		Expiry:   time.Now().Add(time.Hour * 24 * 30),
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "go-blog",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)

}

func Signup(c echo.Context) error {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid request body",
		})
	}

	hashedPassword, err := hashPassword(body.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to hash password",
		})
	}

	user := models.User{
		Username: body.Username,
		Password: hashedPassword,
	}

	if err := initializers.DB.Create(&user).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key") || strings.Contains(err.Error(), "unique constraint") {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"error": "Username already exists",
			})
		}
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Failed to create user",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "User created successfully",
	})
}

func Login(c echo.Context) error {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.Bind(&body); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid request body",
		})
	}

	var user models.User
	if err := initializers.DB.Where("username = ?", body.Username).First(&user).Error; err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid credentials",
		})
	}

	if err := checkPassword(body.Password, user.Password); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid credentials",
		})
	}

	tokenString, err := createToken(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Failed to generate token",
		})
	}

	c.SetCookie(&http.Cookie{
		Name:     "Authorization",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour * 24 * 30),
		SameSite: http.SameSiteLaxMode,
		HttpOnly: true,
	})

	return c.JSON(http.StatusOK, echo.Map{
		"message": "User logged in successfully",
	})
}

func Logout(c echo.Context) error {
	c.SetCookie(&http.Cookie{
		Name:     "Authorization",
		Value:    "",
		Path:     "/",
		Expires:  time.Now().Add(-time.Hour),
		MaxAge:   -1,
		SameSite: http.SameSiteLaxMode,
		HttpOnly: true,
		Secure:   true,
	})
	return c.JSON(http.StatusOK, echo.Map{
		"message": "User logged out successfully",
	})
}

func Validate(c echo.Context) error {
	// Get the user from the context
	user, _ := c.Get("user").(models.User)
	return c.JSON(http.StatusOK, echo.Map{
		"message": "You are logged in",
		"user":    user,
	})
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hashedPassword), nil
}

func checkPassword(password string, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err
}
