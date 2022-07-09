package routes

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"

	"github.com/bratushkadan/My-Fancy-Project/internal/common"
)

var noUserProvidedError = errors.New("No 'user' query parameter was provided.")
var internalServerError = errors.New("Internal server error.")

func IssueToken(c *fiber.Ctx) error {
	c.Set("Content-Type", "application/json")

	user := c.Query("user")
	if user == "" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(noUserProvidedError)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": "bar",
		"time": time.Now().Unix(),
	})

	tokenString, err := token.SignedString(common.HmacAuthSecret)

	if err != nil {
		fmt.Fprint(os.Stderr, err)
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(internalServerError)
	}

	fmt.Println(tokenString)

	c.Status(fiber.StatusOK)
	return c.JSON(struct {
		Token string `json:"token"`
	}{
		Token: tokenString,
	})
}
