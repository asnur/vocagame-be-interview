package wallet_balance

import "github.com/asnur/vocagame-be-interview/pkg/resource"

type (
	Repository interface {
		IUpdate
		IGet
	}

	repository struct {
		resource resource.Resource
	}
)

func New(resource resource.Resource) Repository {
	return &repository{
		resource: resource,
	}
}
