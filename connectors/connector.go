package connectors

import (
	"database/sql"

	"github.com/LUSHDigital/modelgen/model"
)

type Connector interface {
	Connect() (*sql.DB, error)
	QueryStructure(*sql.DB) ([]model.EntityDescriptor, error)
	FillTemplates(*sql.DB, []model.EntityDescriptor, string, string) error
	QueryMigrations(*sql.DB, []model.EntityDescriptor) ([]model.Migration, error)
}
