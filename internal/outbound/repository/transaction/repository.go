package transaction

import "github.com/asnur/vocagame-be-interview/pkg/resource"

type (
	Repository interface {
		ICreate
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
