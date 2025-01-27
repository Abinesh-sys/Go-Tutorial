package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"Go-Tutorial/database"
	"Go-Tutorial/models"
	"Go-Tutorial/routes"
	"fmt"
)

func main() {

	fmt.Println("Program Started")

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
		ViewsLayout:  "layout",
	})

	database.Connect()

	models.Migrate(database.DB)

	routes.SetupRoutes(app)
	app.Use(logger.New())   
	app.Use(recover.New())  
	
	app.Listen(":8080")
	fmt.Println("Server Started")
}
