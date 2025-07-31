package inbound

import (
	"github.com/asnur/vocagame-be-interview/internal/inbound/http"
	"go.uber.org/dig"
)

type (
	Inbound struct {
		dig.In

		Http http.Inbound
	}
)
