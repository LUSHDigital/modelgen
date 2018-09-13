package cmd

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"

	"github.com/LUSHDigital/modelgen/connectors"
	"github.com/LUSHDigital/modelgen/model"
)

func init() {
	rootCmd.AddCommand(generateCmd)
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Run:   generate,
	Short: "Generate models from a database connection",
}

func generate(cmd *cobra.Command, args []string) {
	var (
		connection *sql.DB
		connector  connectors.Connector
		structure  []model.EntityDescriptor
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

	if err = connector.FillTemplates(connection, structure, *output, *pkgName); err != nil {
		log.Fatal(err)
	}
}
