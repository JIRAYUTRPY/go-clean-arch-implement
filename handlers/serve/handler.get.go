package handlers

import (
	"github.com/gofiber/fiber/v2"
	services "github.com/jirayutrpy/server-go/v2/services/serve"
)

type HttpGetServeHandler struct {
	serveUsecase services.GetServeUseCase
}

func NewHttpGetServeHandler(usecase services.GetServeUseCase) *HttpGetServeHandler{
	return &HttpGetServeHandler{serveUsecase: usecase}
}

func (h *HttpGetServeHandler) Gets(c *fiber.Ctx) error {
	err := h.serveUsecase.Gets()
	if err != nil {
		return  c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return c.Status(fiber.StatusFound).JSON(fiber.Map{"message": "Fetch serve success"})
}