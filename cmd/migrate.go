package cmd

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"path/filepath"
// 	"regexp"
// 	"strings"
// 	"time"

// 	"github.com/spf13/cobra"
// 	"sort"
// )

// func migrate(cmd *cobra.Command, args []string) {
// 	validate()
// 	connect()
// 	tables := getTables()
// 	makeMigrations(tables, *output)
// }

// var autoincrementRegExp = regexp.MustCompile(`(?ms) AUTO_INCREMENT=[0-9]*\b`)

// // if the folder you are trying to output the migrations in already exists
// // archive will move the previous migrations into a timestamped archived folder
// // so you can do easy rollbacks
// func archive(folder string) {
// 	f, err := os.Stat(folder)
// 	if err != nil || !f.IsDir() {
// 		return
// 	}
// 	archive := fmt.Sprintf("%s_%s", folder, time.Now().Format("2006_01_02_15_04_05"))
// 	if err := os.Rename(folder, archive); err != nil {
// 		log.Fatalf("cannot archive %s folder", folder)
// 	}
// }

// type statement struct {
// 	tbl   string
// 	stmt  string
// 	order int
// }
// type statements []statement

// func makeMigrations(tables map[string]string, dst string) {
// 	archive(dst)
// 	os.Mkdir(dst, 0777)
// 	now := time.Now().Unix()

// 	var sts statements

// 	for table, comment := range tables {
// 		// get the create statement
// 		row := database.QueryRow(fmt.Sprintf("SHOW CREATE TABLE `%s`", table))
// 		var tbl, stmt string
// 		row.Scan(&tbl, &stmt)
// 		order := GetOrderFromComment(comment)
// 		st := statement{tbl, stmt, order}
// 		sts = append(sts, st)
// 	}
// 	sort.Slice(sts, func(i, j int) bool {
// 		return sts[i].order < sts[j].order
// 	})

// 	for _, st := range sts {
// 		// Create the up migration
// 		where := filepath.Join(dst, fmt.Sprintf("%d_create_%s.up.sql", now, st.tbl))
// 		up, err := os.Create(where)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		auto := autoincrementRegExp.FindString(st.stmt)
// 		st.stmt = strings.Replace(st.stmt, auto, "", 1)
// 		_, err = up.WriteString(st.stmt + ";")
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		// Create the down migration
// 		where = filepath.Join(dst, fmt.Sprintf("%d_create_%s.down.sql", now, st.tbl))
// 		down, err := os.Create(where)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		_, err = down.WriteString(fmt.Sprintf("DROP TABLE IF EXISTS %s;", st.tbl))
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		// Ensure the time increments properly
// 		now += 1
// 	}
// }
