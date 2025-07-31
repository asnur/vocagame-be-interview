package user

import (
	"github.com/asnur/vocagame-be-interview/internal/outbound"
	"github.com/asnur/vocagame-be-interview/internal/usecase/shared"
	"github.com/asnur/vocagame-be-interview/pkg/resource"
)

type (
	UseCase interface {
		IRegister
		ILogin
	}

	usecase struct {
		outbound.Outbound
		shared   shared.UseCase
		resource resource.Resource
	}
)

func New(
	shared shared.UseCase,
	resource resource.Resource,
	outbound outbound.Outbound,
) UseCase {
	return &usecase{
		Outbound: outbound,
		shared:   shared,
		resource: resource,
	}
}
