package routes
import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"Go-Tutorial/controllers"
)

func SetupRoutes(app *fiber.App) {
	
	fmt.Println("Routes")
	app.Get("/", controllers.GetHome)        
	app.Get("/users", controllers.GetUsers)       
}