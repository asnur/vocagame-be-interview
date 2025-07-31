package resource

import (
	"github.com/asnur/vocagame-be-interview/pkg/config"
	"github.com/asnur/vocagame-be-interview/pkg/resource/injection"
	"go.uber.org/dig"
)

type Resource struct {
	dig.In

	// Config
	AppConfig config.AppConfig

	ServerConfig config.ServerConfig

	PostgresConfig config.PostgresConfig

	TokenConfig config.TokenConfig

	// Resources
	Server injection.Server

	Postgres injection.SQL

	Logger injection.Logger

	Jwt injection.Jwt
}
