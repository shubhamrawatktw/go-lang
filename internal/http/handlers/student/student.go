package student

import (
	"go-license/internal/types"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)


func New(c *fiber.Ctx) error {


        var newStudent  types.Student

		if err := c.BodyParser(&newStudent); err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": "Invalid JSON",
            })
        }

	
	if err:=validator.New().Struct(newStudent); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON",
		})
	}
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "New student",
	})
}