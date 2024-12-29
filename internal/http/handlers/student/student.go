package student

import (
	"go-license/internal/storage"
	"go-license/internal/types"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func New(storage storage.Storage) fiber.Handler {

	return func(c *fiber.Ctx) error {

		var newStudent types.Student

		if err := c.BodyParser(&newStudent); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid JSON",
			})
		}

		if err := validator.New().Struct(newStudent); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid JSON",
			})
		}

		id, err := storage.CreateStudent(newStudent.Name,newStudent.Email,newStudent.Age)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err,
			})
		}
		return c.JSON(fiber.Map{
			"status":  "success",
			"id": id,
		})
	}
}
