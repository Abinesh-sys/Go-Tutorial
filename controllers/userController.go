package controllers
import (
	"github.com/gofiber/fiber/v2"
	"Go-Tutorial/models"
	"Go-Tutorial/database"
)
func GetUsers(c *fiber.Ctx) error {
	db := database.DB
	var users []models.User
	db.Find(&users)
	return c.JSON(users)
}