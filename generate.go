package main

import (
	"bytes"
	"go/format"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"strconv"

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
	asStructs := ToStructs(tables)

	// load the model template
	modelTpl, err := box.MustBytes("model.html")
	if err != nil {
		log.Fatal("cannot load model template")
	}
	t := template.Must(template.New("model").Funcs(tmpl.FuncMap).Parse(string(modelTpl)))

	writeModels(asStructs, t)

	// copy in helpers and test suite
	copyFile("x_helpers.html", "x_helpers.go", "helpers")
	copyFile("x_helpers_test.html", "x_helpers_test.go", "helperstest")
}

func writeModels(models []tmpl.TmplStruct, t *template.Template) {
	for _, model := range models {
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
}

func getTables() (tables []string) {
	const stmt = `SELECT table_name
				  FROM information_schema.columns AS c
				  WHERE c.table_catalog = $1
      			  AND column_name = 'id'`

	rows, err := database.Query(stmt, *dbName)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		tables = append(tables, name)
	}
	return tables
}

func GetOrderFromComment(comment string) (order int) {
	if !strings.HasPrefix(comment, "modelgen") {
		return
	}
	parts := strings.SplitN(comment, ":", 2)
	if len(parts) != 2 {
		return
	}
	var err error
	if order, err = strconv.Atoi(parts[1]); err != nil {
		log.Printf("could not parse id comment [%v], make sure to only use numbers in order comments", parts[1])
		return
	}
	return
}

func ToStructs(tables []string) []tmpl.TmplStruct {
	const stmt = `SELECT c.column_name, c.column_default, c.is_nullable, c.data_type, c.character_maximum_length, c.character_octet_length, c.numeric_precision, c.numeric_scale, c.datetime_precision
	        	  FROM information_schema.columns AS c
				  WHERE c.table_name = $1`

	var explained = make(map[string][]sqltypes.Explain)
	for _, table := range tables {
		var expl []sqltypes.Explain
		rows, err := database.Query(stmt, table)
		if err != nil {
			log.Fatal(err)
		}
		for rows.Next() {
			var ex sqltypes.Explain
			if err := rows.Scan(&ex.Field, &ex.Default, &ex.Null, &ex.Type, &ex.CharLength, &ex.OctetLength, &ex.NumericPrecision, &ex.NumericScale, &ex.DateTimePrecision); err != nil {
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
