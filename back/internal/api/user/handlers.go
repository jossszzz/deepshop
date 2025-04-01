package user

import (
	"back/internal/database"
	"back/internal/models/user"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	var users []user.User
	database.DB.Find(&users)
	return c.JSON(users)
}

func CreateUser(c *fiber.Ctx) error {
	return nil
}

func UpdateUser(c *fiber.Ctx) error {
	return nil
}

func DeleteUser(c *fiber.Ctx) error {
	return nil
}
