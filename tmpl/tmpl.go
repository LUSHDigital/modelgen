package tmpl

import (
	"fmt"
	"strings"
	"text/template"
)

var FuncMap = template.FuncMap{
	"insert_fields":      GetInsertFields,
	"insert_values":      GetInsertValues,
	"insert_args":        GetInsertArgs,
	"scan_fields":        GetScanFields,
	"update_args":        GetUpdateArgs,
	"update_values":      GetUpdateValues,
	"update_values_size": GetUpdateValuesLength,
}

func GetInsertFields(fields []TmplField) string {
	var parts []string
	for _, fl := range fields {
		if fl.ColumnName == "id" || fl.ColumnName == "created_at" {
			continue
		}
		parts = append(parts, `"`+fl.ColumnName+`"`)
	}
	return strings.Join(parts, ", ")
}

func GetInsertValues(fields []TmplField) string {
	var parts []string
	i := 1
	for _, fl := range fields {
		switch fl.ColumnName {
		case "id", "created_at":
			continue
		default:
			parts = append(parts, fmt.Sprintf("$%d", i))
			i++
		}
	}
	return strings.Join(parts, ", ")
}

func GetInsertArgs(m StructTmplData) string {
	var parts []string
	for _, fl := range m.Model.Fields {
		switch fl.Name {
		case "ID", "CreatedAt":
			continue
		}
		parts = append(parts, fmt.Sprintf(`%s.%s`, m.Receiver, fl.Name))
	}
	return strings.Join(parts, ", ")
}

func GetScanFields(m StructTmplData) string {
	var parts []string
	for _, fl := range m.Model.Fields {
		parts = append(parts, fmt.Sprintf(`&%s.%s`, m.Receiver, fl.Name))
	}
	return strings.Join(parts, ", ")
}

func GetUpdateArgs(m StructTmplData) string {
	var parts []string
	for _, fl := range m.Model.Fields {
		switch fl.Name {
		case "ID", "CreatedAt", "UpdatedAt":
			continue
		}
		parts = append(parts, fmt.Sprintf(`%s.%s`, m.Receiver, fl.Name))
	}
	return strings.Join(parts, ", ")
}

func GetUpdateValues(m StructTmplData) string {
	var parts []string
	for i, fl := range m.Model.Fields {
		switch fl.Name {
		case "ID", "CreatedAt":
			continue
		case "UpdatedAt":
			parts = append(parts, fmt.Sprintf(`"%s"=now()`, fl.ColumnName))
		default:
			parts = append(parts, fmt.Sprintf(`"%s"=$%d`, fl.ColumnName, i))
		}
	}
	return strings.Join(parts, ", ")
}

func GetUpdateValuesLength(m StructTmplData) string {
	return fmt.Sprintf("$%d", len(m.Model.Fields))
}
