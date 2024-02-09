package routes

import (
	"github.com/gofiber/fiber/v2"
	adapters "github.com/jirayutrpy/server-go/v2/database/event"
	handlers "github.com/jirayutrpy/server-go/v2/handlers/event"
	services "github.com/jirayutrpy/server-go/v2/services/event"
	"gorm.io/gorm"
)

func InitEventRoutes(db *gorm.DB, c *fiber.App) error {

	getEventRepo := adapters.NewGormGetEventRepository(db)
	getEventService := services.NewGetEventService(getEventRepo)
	getEventHandler := handlers.NewHttpGetEventHandler(getEventService)

	groupRoute := c.Group("/api/v1/event")
	groupRoute.Get("/", getEventHandler.Gets)
	

	return nil
}