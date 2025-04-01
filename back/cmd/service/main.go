package main

import (
	"back/internal/api/user"
	"back/internal/database"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	database.ConnectDB()

	user.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
