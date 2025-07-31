package injection

import (
	"github.com/asnur/vocagame-be-interview/pkg/config"
	"github.com/gofiber/fiber/v2"
)

type (
	Server struct {
		*fiber.App
		config config.ServerConfig
	}
)

func NewServer(config config.ServerConfig) Server {
	app := fiber.New(fiber.Config{})

	return Server{App: app, config: config}
}

func (s Server) Start() error {
	return s.Listen(s.config.Host + ":" + s.config.Port)
}
