package routes

import (
	"github.com/gofiber/fiber/v2"
	adapters "github.com/jirayutrpy/server-go/v2/database/serve"
	handlers "github.com/jirayutrpy/server-go/v2/handlers/serve"
	"github.com/jirayutrpy/server-go/v2/middlewares"
	services "github.com/jirayutrpy/server-go/v2/services/serve"
	"gorm.io/gorm"
)

func InitServeRoutes(db *gorm.DB, c *fiber.App) error{

	getServeRepo := adapters.NewGormGetServeRepository(db)
	getServeService := services.NewGetServeService(getServeRepo)
	getServeHandler := handlers.NewHttpGetServeHandler(getServeService)

	createServeRepo := adapters.NewGormCreateServeRepository(db)
	createServeService := services.NewCreateServeService(createServeRepo)
	createServeHandler := handlers.NewHttpCreateServeHandler(createServeService)

	updateServeRepo := adapters.NewGormUpdateServeRepository(db)
	updateServeService := services.NewUpdateServeService(updateServeRepo)
	updateServeHandler := handlers.NewHttpUpdateServeHandler(updateServeService)

	groupRoute := c.Group("/api/v1/serve")
	groupRoute.Use(middlewares.HttpAuthorization)
	groupRoute.Get("/", getServeHandler.Gets)
	groupRoute.Get("/byuserid/:userId", getServeHandler.GetByUserId)
	groupRoute.Post("/", createServeHandler.Create)
	groupRoute.Patch("/", updateServeHandler.Update)
	return nil
}