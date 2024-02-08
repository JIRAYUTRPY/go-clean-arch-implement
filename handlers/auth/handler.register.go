package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jirayutrpy/server-go/v2/entities"
	services "github.com/jirayutrpy/server-go/v2/services/auth"
)

type HttpRegisterHandler struct {
	authUsecase services.RegisterUseCase
}

func NewHttpRegisterHandler(usecase services.RegisterUseCase) *HttpRegisterHandler{
	return &HttpRegisterHandler{authUsecase: usecase}
}

func (h *HttpRegisterHandler) Register(c *fiber.Ctx) error {
	var user entities.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{ "message": err})
	}
	if err := h.authUsecase.Register(user); err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Register successfully"})
}