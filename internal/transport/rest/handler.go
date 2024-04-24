package rest

import (
	"net/http"

	"github.com/Alieksieiev0/auth-service/internal/services"
	"github.com/Alieksieiev0/auth-service/internal/types"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func login(service services.AuthService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user := &types.User{}
		if err := c.BodyParser(user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		if user.Username == "" && user.Password == "" {
			return c.Status(fiber.StatusBadRequest).
				JSON(fiber.Map{"error": "insufficient user data"})
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
		if err := c.BodyParser(user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		if user.Username == "" && user.Email == "" && user.Password == "" {
			return c.Status(fiber.StatusBadRequest).
				JSON(fiber.Map{"error": "insufficient user data"})
		}

		hashedPassword, err := bcrypt.GenerateFromPassword(
			[]byte(user.Password),
			bcrypt.DefaultCost,
		)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).
				JSON(fiber.Map{"error": "couldnt process password"})
		}

		user.Password = string(hashedPassword)
		if err := service.Register(c.Context(), user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		c.Status(http.StatusCreated)
		return nil
	}
}
