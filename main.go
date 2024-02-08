package main

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jirayutrpy/server-go/v2/routes"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	app := SetUpRouter()
	err := godotenv.Load()
	if err != nil {
		fmt.Print(err)
	}
	app.Listen(":7777")
}

func SetUpRouter() *fiber.App{
	app := fiber.New()
	db, err := SetUpdatabase();
	if err != nil {
		fmt.Print(err)
	}
	routes.InitAuthRoutes(db, app)
	routes.InitUserRoutes(db, app)
	return app
}

func SetUpdatabase() (*gorm.DB,error){
	const (
		host = "localhost"
		port = 5432
		database = "fundamental_postgres"
		username = "admin"
		password = "admin_password"
	)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable" ,host,port,username,password,database)
	db, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil{
		return db,errors.New("Failed to connect database")
	}

	if err := db.AutoMigrate(); err != nil{
		return db,errors.New("Failed to migrate database")
	}
	return db, nil
}