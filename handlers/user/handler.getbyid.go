package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	services "github.com/jirayutrpy/server-go/v2/services/user"
)

type HttpGetByIdHandler struct {
	userUsecase services.GetByIdUseCase
}

func NewHttpGetByIdHandler(usecase services.GetByIdUseCase) *HttpGetByIdHandler{
	return &HttpGetByIdHandler{userUsecase: usecase}
}

func (h *HttpGetByIdHandler) GetById(c *fiber.Ctx) error {
	idString := c.Params("id")
	id, err := strconv.Atoi(idString)
	if  err != nil {
		return  c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err})
	}
	response, err := h.userUsecase.GetById(uint(id))
	if err != nil {
		return  c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Fetch user success", "data": response})
}