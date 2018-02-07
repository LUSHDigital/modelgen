package main

import (
	"bytes"
	"go/format"
	"html/template"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/LUSHDigital/modelgen/sqlfmt"
	"github.com/LUSHDigital/modelgen/sqltypes"
	"github.com/LUSHDigital/modelgen/tmpl"
	"github.com/spf13/cobra"
)

func generate(cmd *cobra.Command, args []string) {
	validate()
	connect()

	// get the list of tables from the database
	tables := getTables()
	if len(tables) == 0 {
		log.Fatal("No tables to read")
	}

	// make structs from tables
	asStructs := toStructs(tables)

	// load the model template
	modelTpl, err := box.MustBytes("model.html")
	if err != nil {
		log.Fatal("cannot load model template")
	}
	t := template.Must(template.New("model").Funcs(tmpl.FuncMap).Parse(string(modelTpl)))

	// write the models to disk
	for _, model := range asStructs {
		writeModel(model, t)
	}

	// copy in helpers and test suite
	copyFile("x_helpers.html", "x_helpers.go", "helpers")
	copyFile("x_helpers_test.html", "x_helpers_test.go", "helperstest")
}

func writeModel(model tmpl.TmplStruct, t *template.Template) {
	m := tmpl.StructTmplData{
		Model:       model,
		Receiver:    strings.ToLower(string(model.Name[0])),
		PackageName: *pkgName,
	}

	buf := new(bytes.Buffer)
	err := t.Execute(buf, m)
	if err != nil {
		log.Fatal(err)
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}

	buf = bytes.NewBuffer(formatted)

	out := *output
	os.Mkdir(out, 0777)

	p := filepath.Join(out, model.TableName)
	f, err := os.Create(p + ".go")
	if err != nil {
		log.Fatal(err)
	}
	buf.WriteTo(f)
	f.Close()
}

func getTables() (tables []string) {
	const stmt = `SELECT table_name
				  FROM information_schema.columns AS c
				  WHERE c.column_key = "PRI"
				  AND c.table_schema = ?
      			  AND column_name = "id"`

	rows, err := database.Query(stmt, *dbName)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var a string
		if err := rows.Scan(&a); err != nil {
			log.Fatal(err)
		}
		tables = append(tables, a)
	}
	return tables
}

func toStructs(tables []string) []tmpl.TmplStruct {
	var explained = make(map[string][]sqltypes.Explain)
	for _, table := range tables {
		var expl []sqltypes.Explain
		rows, err := database.Query("EXPLAIN " + table)
		if err != nil {
			log.Fatal(err)
		}
		for rows.Next() {
			var ex sqltypes.Explain
			if err := rows.Scan(&ex.Field, &ex.Type, &ex.Null, &ex.Key, &ex.Default, &ex.Extra); err != nil {
				log.Fatal(err)
			}
			expl = append(expl, ex)
		}
		rows.Close()
		explained[table] = expl
	}

	var structStore tmpl.TmplStructs
	for k, explain := range explained {
		t := tmpl.TmplStruct{
			Name:      sqlfmt.ToPascalCase(k),
			TableName: k,
			Imports:   make(map[string]struct{}),
		}

		for _, expl := range explain {
			f := tmpl.TmplField{
				Name:       sqlfmt.ToPascalCase(*expl.Field),
				Type:       sqltypes.AssertType(*expl.Type, *expl.Null),
				ColumnName: strings.ToLower(*expl.Field),
				Nullable:   *expl.Null == "YES",
			}
			t.Fields = append(t.Fields, f)
			if imp, ok := sqltypes.NeedsImport(f.Type); ok {
				t.Imports[imp] = struct{}{}
			}
		}
		structStore = append(structStore, t)
	}

	return structStore
}

func copyFile(src, dst, templateName string) {
	dbFile, err := box.MustBytes(src)
	if err != nil {
		log.Fatalf("cannot retrieve template file: %v", err)
	}

	t := template.Must(template.New(templateName).Parse(string(dbFile)))
	buf := new(bytes.Buffer)
	err = t.Execute(buf, map[string]string{
		"PackageName": *pkgName,
	})
	if err != nil {
		log.Fatal(err)
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatal(err)
	}
	buf = bytes.NewBuffer(formatted)

	if err != nil {
		log.Fatal("cannot copy file")
	}
	out := filepath.Join(*output, dst)
	to, err := os.Create(out)
	if err != nil {
		log.Fatal(err)
	}
	defer to.Close()

	_, err = io.Copy(to, buf)
	if err != nil {
		log.Fatal(err)
	}
}
