package rest

import (
	"fmt"
	"net/http"

	"github.com/Alieksieiev0/auth-service/internal/services"
	"github.com/Alieksieiev0/auth-service/internal/types"
	"github.com/gofiber/fiber/v2"
)

func login(service services.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := &types.User{}
		if err := c.BodyParser(user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		userToken, err := service.Login(c.Context(), user)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		return c.Status(http.StatusOK).JSON(userToken)
	}
}

func register(service services.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := &types.User{}
		fmt.Println(string(c.Body()))
		if err := c.BodyParser(user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		fmt.Println(user)

		if err := service.Register(c.Context(), user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		c.Status(http.StatusOK)
		return nil
	}
}
