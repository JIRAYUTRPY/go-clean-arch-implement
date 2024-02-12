package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	services "github.com/jirayutrpy/server-go/v2/services/event"
)

type HttpGetEventHandler struct {
	eventUsecase services.GetEventUsecase
}

func NewHttpGetEventHandler(usecase services.GetEventUsecase) *HttpGetEventHandler {
	return &HttpGetEventHandler{eventUsecase: usecase}
}

func (h *HttpGetEventHandler) Gets(c *fiber.Ctx) error {
	_, err := h.eventUsecase.Gets()
	if err != nil {
		return  c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return c.Status(fiber.StatusFound).JSON(fiber.Map{"message": "Fetch event success"})
}

func (h *HttpGetEventHandler) GetsByUserId(c *fiber.Ctx) error {
	userId := c.Params("userId")
	intUserId, err := strconv.Atoi(userId);
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}
	data, err := h.eventUsecase.GetByUserId(intUserId)
	if err != nil {
		return  c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return c.Status(fiber.StatusFound).JSON(fiber.Map{"data": data, "message": "Fetch event success"})
}