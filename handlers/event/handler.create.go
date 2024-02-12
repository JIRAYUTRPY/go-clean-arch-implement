package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jirayutrpy/server-go/v2/entities"
	services "github.com/jirayutrpy/server-go/v2/services/event"
)

type HttpCraeteEventHandler struct {
	eventUsecase services.CreateEventUseCase
}

func NewHttpCreatEventHandler(usecase services.CreateEventUseCase) *HttpCraeteEventHandler{
	return &HttpCraeteEventHandler{eventUsecase: usecase}
}

func (h *HttpCraeteEventHandler) Create(c *fiber.Ctx) error {
	var data entities.Event
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{ "message": err})
	}
	if err := h.eventUsecase.Create(data); err != nil {
		return  c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return c.Status(fiber.StatusFound).JSON(fiber.Map{"message": "Create event success"})
}