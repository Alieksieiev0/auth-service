package rest

import (
	"github.com/Alieksieiev0/auth-service/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type RESTServer struct {
	app  *fiber.App
	addr string
}

func NewServer(app *fiber.App, addr string) *RESTServer {
	return &RESTServer{
		app:  app,
		addr: addr,
	}
}

func (us *RESTServer) Start(service services.AuthService) error {
	us.app.Use(logger.New(logger.Config{
		Format: "${time} | ${status} | ${latency} | ${method} | ${path} | ${error}\nResponse Body: ${resBody}\n",
	}))
	us.app.Use(cors.New())

	us.app.Post("/login", login(service))
	us.app.Post("/register", register(service))

	return us.app.Listen(us.addr)
}
