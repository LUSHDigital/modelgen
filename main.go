package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/gobuffalo/packr"

	"github.com/spf13/cobra"

	_ "github.com/go-sql-driver/mysql"
)

var (
	output   *string
	dbName   *string
	pkgName  *string
	conn     *string
	database *sql.DB
	version  string
	box      packr.Box
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	box = packr.NewBox("./tmpl")

	rootCmd := &cobra.Command{}

	pkgName = rootCmd.PersistentFlags().StringP("package", "p", "generated_models", "name of package")
	output = rootCmd.PersistentFlags().StringP("output", "o", "generated_models", "path to package")
	dbName = rootCmd.PersistentFlags().StringP("database", "d", "", "name of database")
	conn = rootCmd.PersistentFlags().StringP("connection", "c", "", "user:pass@host:port")

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

var formatErr = errors.New("invalid connection string format")

func mkDsn(connect, dbname string) string {
	parts := strings.Split(connect, "@")
	if len(parts) < 2 {
		log.Fatal(formatErr)
	}

	credentials := strings.Split(parts[0], ":")
	if len(credentials) < 2 {
		log.Fatal(formatErr)
	}
	database := strings.Split(parts[1], ":")
	if len(database) < 2 {
		log.Fatal(formatErr)
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", credentials[0], credentials[1], database[0], database[1], dbname)
}

func connect() {
	// connect to database
	var err error
	database, err = sql.Open("mysql", mkDsn(*conn, *dbName))
	if err != nil {
		log.Fatal(err)
	}

	// check for a valid connection
	if err := database.Ping(); err != nil {
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
