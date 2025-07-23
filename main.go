package main

import (
	"belajargo/db"
	"belajargo/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db.InitDB()
	app := fiber.New()
	routes.RegisterRoutes(app)
	app.Listen(":3000")
}
