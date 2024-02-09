package routes

import (
	"github.com/gofiber/fiber/v2"
	adapters "github.com/jirayutrpy/server-go/v2/database/serve"
	handlers "github.com/jirayutrpy/server-go/v2/handlers/serve"
	services "github.com/jirayutrpy/server-go/v2/services/serve"
	"gorm.io/gorm"
)

func InitServeRoutes(db *gorm.DB, c *fiber.App) error{

	getServeRepo := adapters.NewGormGetServeRepository(db)
	getServeService := services.NewGetServeService(getServeRepo)
	getServeHandler := handlers.NewHttpGetServeHandler(getServeService)

	groupRoute := c.Group("/api/v1/serve")
	groupRoute.Get("/", getServeHandler.Gets)

	return nil
}