package routes

import (
	"github.com/gofiber/fiber/v2"
	eventAdapters "github.com/jirayutrpy/server-go/v2/database/event"
	adapters "github.com/jirayutrpy/server-go/v2/database/serve"
	eventHandlers "github.com/jirayutrpy/server-go/v2/handlers/event"
	handlers "github.com/jirayutrpy/server-go/v2/handlers/serve"
	eventServices "github.com/jirayutrpy/server-go/v2/services/event"
	services "github.com/jirayutrpy/server-go/v2/services/serve"
	"gorm.io/gorm"
)

func InitNoauthRoutes(db *gorm.DB, c *fiber.App) error {
	getServeRepo := adapters.NewGormGetServeRepository(db)
	getServeService := services.NewGetServeService(getServeRepo)
	getServeHandler := handlers.NewHttpGetServeHandler(getServeService)

	createEventRepo := eventAdapters.NewGormCreateEventRepository(db)
	createEventService := eventServices.NewCreateEventService(createEventRepo)
	createEventHandler := eventHandlers.NewHttpCreatEventHandler(createEventService)

	groupRoute := c.Group("/api/v1/noauth")
	groupRoute.Get("/withouttoken/:userId", getServeHandler.GetByUserId)
	groupRoute.Post("/createEvent", createEventHandler.Create)
	return nil
}
