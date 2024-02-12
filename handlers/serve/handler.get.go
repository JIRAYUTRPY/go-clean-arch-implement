package handlers

import (
	"strconv"

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

func (h *HttpGetServeHandler) GetByUserId(c *fiber.Ctx) error{
	id := c.Params("userId")
	if _, errorConvert := strconv.Atoi(id); errorConvert != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Query parmas invalid"})
	}
	userId, err := strconv.Atoi(id)
	data, err := h.serveUsecase.GetByUserId(userId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}
	return c.Status(fiber.StatusFound).JSON(fiber.Map{"data": data,"message": "Fetch Service Successfully"})
}