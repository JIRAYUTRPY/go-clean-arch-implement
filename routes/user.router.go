package routes

import (
	"github.com/gofiber/fiber/v2"
	adapters "github.com/jirayutrpy/server-go/v2/database/user"
	handlers "github.com/jirayutrpy/server-go/v2/handlers/user"
	services "github.com/jirayutrpy/server-go/v2/services/user"
	"gorm.io/gorm"
)

func InitUserRoutes(db *gorm.DB, c *fiber.App) error {

	getUserRepo := adapters.NewGormGetUserRepository(db)
	getUserService := services.NewGetUserService(getUserRepo)
	getUserHandler := handlers.NewHttpGetUserHandler(getUserService)

	groupRoute := c.Group("/api/v1/user")
	groupRoute.Get("/", getUserHandler.Gets)
	groupRoute.Get("/:id", getUserHandler.GetById)
	groupRoute.Get("/:email/email", getUserHandler.GetByEmail)
	return nil
}