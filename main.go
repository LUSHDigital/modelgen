package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/spf13/cobra"

	_ "github.com/go-sql-driver/mysql"
)

var (
	output  *string
	dbName  *string
	pkgName *string
	conn    *string
	engine  *string

	database *sql.DB
	version  string
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	rootCmd := &cobra.Command{}

	pkgName = rootCmd.PersistentFlags().StringP("package", "p", "generated_models", "name of package")
	output = rootCmd.PersistentFlags().StringP("output", "o", "generated_models", "path to package")
	dbName = rootCmd.PersistentFlags().StringP("database", "d", "", "name of database")
	conn = rootCmd.PersistentFlags().StringP("connection", "c", "", "connection string to database")
	engine = rootCmd.PersistentFlags().StringP("engine", "e", "", "type of database [mysql|postgres]")

	generateCmd := &cobra.Command{
		Use:   "generate",
		Run:   generate,
		Short: "Generate models from a database connection",
	}

	migrateCmd := &cobra.Command{
		Use:   "migrate",
		Run:   migrate,
		Short: "Generate migration files from a database connection",
	}

	versionCmd := &cobra.Command{
		Use: "version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(version)
		},
		Short: "Returns the current version name",
	}

	rootCmd.AddCommand(generateCmd, migrateCmd, versionCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

var errFormat = errors.New("invalid connection string format")

func connect() {
	// connect to database
	var err error
	if database, err = sql.Open(*engine, *conn); err != nil {
		log.Fatal(err)
	}

	// check for a valid connection
	if err = database.Ping(); err != nil {
		log.Fatal(err)
	}
}

func validate() {
	if dbName == nil || *dbName == "" {
		log.Fatal("Please provide a database name")
	}
	if conn == nil || *conn == "" {
		log.Fatal("Please provide a connection string")
	}
}
