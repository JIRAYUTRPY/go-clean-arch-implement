package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jirayutrpy/server-go/v2/interfaces"
	services "github.com/jirayutrpy/server-go/v2/services/auth"
)

type HttpLoginHandler struct {
	authUsecase services.LoginUseCase
}

func NewHttpLoginHandler(usecase services.LoginUseCase) *HttpLoginHandler{
	return &HttpLoginHandler{authUsecase: usecase}
}

func (h *HttpLoginHandler) Login(c *fiber.Ctx) error {
	var loginInterface interfaces.LoginRequest
	if err := c.BodyParser(&loginInterface); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err})
	}
	data, err := h.authUsecase.Login(loginInterface)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	// c.Cookie(&fiber.Cookie{
	// 	Name: "access_token",
	// 	Value:  data.Token,
	// 	Expires: time.Now().Add(time.Hour * 72),
	// 	HTTPOnly: true,
	// })
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Login successfully", "token": data.Token})
}