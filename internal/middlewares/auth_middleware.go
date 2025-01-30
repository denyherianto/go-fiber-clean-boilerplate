package middlewares

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"

	"github.com/denyherianto/go-fiber-clean-boilerplate/pkg/jwt"
)

func TokenValidation(c *fiber.Ctx) error {
	tokenString := jwt.ExtractToken(c)
	coreApiHostUrl := os.Getenv("CORE_API_HOST")

	// Create a Bearer string by appending string access token
	var bearer = "Bearer " + tokenString

	// Create a new request using http
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/token/verify", coreApiHostUrl), nil)

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// Return status 500
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// Return status 500
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   fiber.StatusInternalServerError,
			"message": "unauthorized, check expiration time of your token",
		})
	}

	return c.Next()
}
