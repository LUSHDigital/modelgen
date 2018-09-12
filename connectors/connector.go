package connectors

import (
	"database/sql"

	"github.com/nicklanng/modelgen/model"
)

type Connector interface {
	Connect() (*sql.DB, error)
	QueryStructure(*sql.DB) ([]model.EntityDescriptor, error)
	FillTemplates(*sql.DB, []model.EntityDescriptor, string, string) error
}
