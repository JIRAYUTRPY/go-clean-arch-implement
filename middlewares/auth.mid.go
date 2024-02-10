package middlewares

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func HttpAuthorization(c *fiber.Ctx) error {
	headers := c.GetReqHeaders()
	bearer := headers["Authorization"]
	if bearer == nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	stringWithBearer := bearer[0]
	token := strings.Split(stringWithBearer, " ")
	if token == nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	pureToken := token[1]
	jwtSecretKey := os.Getenv("SECRET_KEY")
	jwtToken, jwtErr := jwt.ParseWithClaims(pureToken, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})
	if jwtErr != nil && !jwtToken.Valid {
		fmt.Println(jwtSecretKey, jwtErr, jwtToken.Signature)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "token invalid",
		})
	}
	claims := jwtToken.Claims.(jwt.MapClaims)
	exp := claims["exp"]
	expInt := exp.(float64)
	if time.Now().Hour() - time.Unix(int64(expInt), 0).Hour() > 3 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "token timeout",
		})
	}
	return c.Next()
}