package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jirayutrpy/server-go/v2/interfaces"
	services "github.com/jirayutrpy/server-go/v2/services/serve"
)

type HttpUpdateServeHandler struct {
	serveUsecase services.UpdateServeUseCase
}

func NewHttpUpdateServeHandler(usecase services.UpdateServeUseCase) *HttpUpdateServeHandler{
	return &HttpUpdateServeHandler{serveUsecase: usecase}
}

func (h *HttpUpdateServeHandler) Update(c *fiber.Ctx) error{
	var data interfaces.ServeUpdateRequest
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	if err := h.serveUsecase.Update(data); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Update data sucessfully",
	})
}