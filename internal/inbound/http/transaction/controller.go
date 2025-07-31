package transaction

import (
	"github.com/asnur/vocagame-be-interview/internal/usecase"
	"github.com/asnur/vocagame-be-interview/pkg/resource"
	"go.uber.org/dig"
)

type (
	Controller struct {
		dig.In

		UseCase usecase.UseCase

		Resource resource.Resource
	}
)
