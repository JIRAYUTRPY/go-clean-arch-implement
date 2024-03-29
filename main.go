package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/jirayutrpy/server-go/v2/entities"
	"github.com/jirayutrpy/server-go/v2/routes"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Print(err)
	}
	app := SetUpRouter()
	app.Listen(":7777")
}

func SetUpRouter() *fiber.App{
	app := fiber.New()
	app.Use(cors.New())
	db, err := SetUpdatabase();
	if err != nil {
		fmt.Print(err)
	}
	routes.InitNoauthRoutes(db, app)
	routes.InitAuthRoutes(db, app)
	routes.InitUserRoutes(db, app)
	routes.InitServeRoutes(db, app)
	routes.InitEventRoutes(db,app)
	return app
}

func SetUpdatabase() (*gorm.DB,error){
	host := os.Getenv("DB_HOST")
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		fmt.Print("Port is should be number")
	}
	database := os.Getenv("DB_NAME")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable" ,host,port,username,password,database)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil{
		return db,errors.New("Failed to connect database")
	}

	if err := db.AutoMigrate(&entities.User{}, &entities.Serve{},&entities.Event{} ); err != nil{
		return db,errors.New("Failed to migrate database")
	}
	return db, nil
}