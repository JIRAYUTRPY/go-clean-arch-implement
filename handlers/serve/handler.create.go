package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jirayutrpy/server-go/v2/entities"
	services "github.com/jirayutrpy/server-go/v2/services/serve"
)

type HttpCreateServeHandler struct {
	serveUsecase services.CreateServeUseCase
}

func NewHttpCreateServeHandler(usecase services.CreateServeUseCase) *HttpCreateServeHandler{
	return &HttpCreateServeHandler{serveUsecase: usecase}
}

func (h *HttpCreateServeHandler) Create(c *fiber.Ctx) error {
	var data entities.Serve
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err})
	}
	err := h.serveUsecase.Create(data)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return c.Status(fiber.StatusFound).JSON(fiber.Map{"message": "Created serve success"})
}