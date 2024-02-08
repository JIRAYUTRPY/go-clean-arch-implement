package routes

import (
	"github.com/gofiber/fiber/v2"
	adapters "github.com/jirayutrpy/server-go/v2/database/auth"
	handlers "github.com/jirayutrpy/server-go/v2/handlers/auth"
	services "github.com/jirayutrpy/server-go/v2/services/auth"
	"gorm.io/gorm"
)

func InitAuthRoutes(db *gorm.DB,c *fiber.App) error {
	
	loginRepo := adapters.NewGormLoginRepository(db)
	loginService := services.NewLoginService(loginRepo)
	loginHandler := handlers.NewHttpLoginHandler(loginService)

	registerRepo := adapters.NewGormRegisterRepository(db)
	registerService := services.NewRegisterService(registerRepo)
	registerHandler := handlers.NewHttpRegisterHandler(registerService)

	groupRoute := c.Group("/api/v1")
	groupRoute.Post("/login", loginHandler.Login)
	groupRoute.Post("/register", registerHandler.Register)
	groupRoute.Get("/",func(c *fiber.Ctx) error {
		return c.SendString("Reservation system api")
	})
	return nil
}