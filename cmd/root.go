package cmd

import (
	"errors"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	output  *string
	dbName  *string
	pkgName *string
	conn    *string
)

var (
	errFormat = errors.New("invalid connection string format")
	rootCmd   = &cobra.Command{}
)

func init() {
	pkgName = rootCmd.PersistentFlags().StringP("package", "p", "generated_models", "name of package")
	output = rootCmd.PersistentFlags().StringP("output", "o", "generated_models", "path to package")
	dbName = rootCmd.PersistentFlags().StringP("database", "d", "", "name of database")
	conn = rootCmd.PersistentFlags().StringP("connection", "c", "", "user:pass@host:port")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func validateArgs() {
	if dbName == nil || *dbName == "" {
		log.Fatal("Please provide a database name")
	}
	if conn == nil || *conn == "" {
		log.Fatal("Please provide a connection string")
	}
}

func deconstructConnectionString(connect string) (username, password, host, port string, err error) {
	parts := strings.Split(connect, "@")
	if len(parts) < 2 {
		return "", "", "", "", errFormat
	}

	credentials := strings.Split(parts[0], ":")
	if len(credentials) < 2 {
		return "", "", "", "", errFormat
	}

	database := strings.Split(parts[1], ":")
	if len(database) < 2 {
		return "", "", "", "", errFormat
	}

	return credentials[0], credentials[1], database[0], database[1], nil
}
