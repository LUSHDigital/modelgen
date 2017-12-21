package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func migrate(cmd *cobra.Command, args []string) {
	validate()
	connect()
	tables := getTables()
	makeMigrations(tables, *output)
}

var autoincrementRegExp = regexp.MustCompile(`(?ms) AUTO_INCREMENT=[0-9]*\b`)

func makeMigrations(tables []string, dst string) {
	os.Mkdir(dst, 0777)
	now := time.Now().Unix()
	for _, table := range tables {
		// get the create statement
		row := database.QueryRow(fmt.Sprintf("SHOW CREATE TABLE `%s`", table))
		var tbl, stmt string
		row.Scan(&tbl, &stmt)

		// Create the up migration
		where := filepath.Join(dst, fmt.Sprintf("%d_create_%s.up.sql", now, tbl))
		up, err := os.Create(where)
		if err != nil {
			log.Fatal(err)
		}
		auto := autoincrementRegExp.FindString(stmt)
		stmt = strings.Replace(stmt, auto, "", 1)
		_, err = up.WriteString(stmt + ";")
		if err != nil {
			log.Fatal(err)
		}
		// Create the down migration
		where = filepath.Join(dst, fmt.Sprintf("%d_create_%s.down.sql", now, tbl))
		down, err := os.Create(where)
		if err != nil {
			log.Fatal(err)
		}
		_, err = down.WriteString(fmt.Sprintf("DROP TABLE IF EXISTS %s;", tbl))
		if err != nil {
			log.Fatal(err)
		}

		// Ensure the time increments properly
		now += 1
	}
}
