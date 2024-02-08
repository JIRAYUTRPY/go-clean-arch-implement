package routes

import (
	"github.com/gofiber/fiber/v2"
	adapters "github.com/jirayutrpy/server-go/v2/database/user"
	handlers "github.com/jirayutrpy/server-go/v2/handlers/user"
	services "github.com/jirayutrpy/server-go/v2/services/user"
	"gorm.io/gorm"
)

func InitUserRoutes(db *gorm.DB, c *fiber.App) error {

	getByIdRepo := adapters.NewGormGetByIDRepository(db)
	getByIdService := services.NewGetByIdService(getByIdRepo)
	getByIdHandler := handlers.NewHttpGetByIdHandler(getByIdService)

	groupRoute := c.Group("/api/v1/user")
	groupRoute.Get("/:id", getByIdHandler.GetById)
	return nil
}