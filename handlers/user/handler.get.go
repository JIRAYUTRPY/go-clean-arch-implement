package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	services "github.com/jirayutrpy/server-go/v2/services/user"
)

type HttpGetHandler struct {
	userUsecase services.GetUserUseCase
}

func NewHttpGetUserHandler(usecase services.GetUserUseCase) *HttpGetHandler{
	return &HttpGetHandler{userUsecase: usecase}
}

func (h *HttpGetHandler) Gets(c *fiber.Ctx) error {
	response, err := h.userUsecase.Gets()
	if err != nil {
		return  c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Fetch user success", "data": response})
}

func (h *HttpGetHandler) GetById(c *fiber.Ctx) error {
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

func (h *HttpGetHandler) GetByEmail(c *fiber.Ctx) error {
	email := c.Params("email")
	response, err := h.userUsecase.GetByEmail(email)
	if err != nil {
		return  c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Fetch user success", "data": response})
}