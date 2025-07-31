package shared

import "github.com/asnur/vocagame-be-interview/pkg/resource"

type (
	UseCase interface {
		IHashPassword
		ICheckPassword
		IAuthToken
	}

	usecase struct {
		resource resource.Resource
	}
)

func New(
	resource resource.Resource,
) UseCase {
	return &usecase{
		resource: resource,
	}
}
