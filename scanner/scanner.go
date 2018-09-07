package scanner

import (
	"github.com/nicklanng/modelgen/model"
)

type Scanner interface {
	Connect() error
	QueryStructure() ([]model.EntityDescriptor, error)
}
