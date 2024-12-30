package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/shubhamrwtktw/lic/internal/models"
	"github.com/shubhamrwtktw/lic/internal/services"
)

func CreateUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid body"})
	}

	result, err := services.CreateUser(user)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created",
		"id":      result.InsertedID,
	})

}
