package http

import (
	"github.com/asnur/vocagame-be-interview/internal/inbound/http/user"
	"go.uber.org/dig"
)

type (
	Inbound struct {
		dig.In

		User user.Controller
	}
)
