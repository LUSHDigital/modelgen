package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"

	_ "github.com/go-sql-driver/mysql"
)

var (
	output   *string
	dbName   *string
	pkgName  *string
	conn     *string
	database *sql.DB
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	rootCmd := &cobra.Command{}

	pkgName = rootCmd.PersistentFlags().StringP("package", "p", "generated_models", "name of package")
	output = rootCmd.PersistentFlags().StringP("output", "o", "generated_models", "path to package")
	dbName = rootCmd.PersistentFlags().StringP("database", "d", "", "name of database")
	conn = rootCmd.PersistentFlags().StringP("connection", "c", "", "user:pass@host:port")

	generateCmd := &cobra.Command{
		Use: "generate",
		Run: generate,
	}

	migrateCmd := &cobra.Command{
		Use: "migrate",
		Run: migrate,
	}

	rootCmd.AddCommand(generateCmd, migrateCmd)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

var formatErr = errors.New("Invalid connection string format")

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
	if dbName == nil {
		log.Fatal("Please provide a database name")
	}
	if conn == nil {
		log.Fatal("Please provide a connection string")
	}
}
