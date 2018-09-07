package cmd

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"

	"github.com/nicklanng/modelgen/scanner"
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
		err error
		in  scanner.Scanner
		// out generator.Generator
	)

	validateArgs()

	username, password, host, port, err := deconstructConnectionString(*conn)
	if err != nil {
		log.Fatal(err)
	}

	// input
	in = scanner.NewMySQL(username, password, host, port, *dbName)

	err = in.Connect()
	if err != nil {
		log.Fatal(err)
	}

	structure, err := in.QueryStructure()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(structure)

	// output
	// out = generator.NewMySQL(structure)
	// out.FillTemplates()
}
