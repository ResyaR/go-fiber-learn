package handlers

import (
	"belajargo/db"
	"belajargo/models"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	db.DB.Find(&users)
	return c.JSON(users)
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).SendString("Invalid JSON")
	}
	db.DB.Create(&user)
	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if err := db.DB.First(&user, id).Error; err != nil {
		return c.Status(404).SendString("User not found")
	}
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).SendString("Invalid JSON")
	}
	db.DB.Save(&user)
	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if err := db.DB.First(&user, id).Error; err != nil {
		return c.Status(404).SendString("User not found")
	}
	db.DB.Delete(&user)
	return c.SendStatus(204)
}
