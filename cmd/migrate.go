package cmd

import (
	"database/sql"
	"github.com/LUSHDigital/modelgen/connectors"
	"github.com/LUSHDigital/modelgen/migrations"
	"github.com/LUSHDigital/modelgen/model"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Run:   migrate,
	Short: "Generate migration files from a database connection",
}

func migrate(cmd *cobra.Command, args []string) {
	var (
		connector  connectors.Connector
		connection *sql.DB
		structure  []model.EntityDescriptor
		m          []model.Migration
		err        error
	)

	validateArgs()

	username, password, host, port, err := deconstructConnectionString(*conn)
	if err != nil {
		log.Fatal(err)
	}

	connector = connectors.NewMySQL(username, password, host, port, *dbName)

	if connection, err = connector.Connect(); err != nil {
		log.Fatal(err)
	}

	if structure, err = connector.QueryStructure(connection); err != nil {
		log.Fatal(err)
	}

	if m, err = connector.QueryMigrations(connection, structure); err != nil {
		log.Fatal(err)
	}

	w := migrations.NewMigrationWriter(*output)

	if err = w.WriteMigrations(m); err != nil {
		log.Fatal(err)
	}
}
