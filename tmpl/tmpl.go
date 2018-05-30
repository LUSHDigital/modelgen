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
		switch fl.ColumnName {
		case "id", "updated_at":
			continue
		default:
			parts = append(parts, `"`+fl.ColumnName+`"`)
		}
	}
	return strings.Join(parts, ", ")
}

func GetInsertValues(fields []TmplField) string {
	var parts []string
	i := 1
	for _, fl := range fields {
		switch fl.ColumnName {
		case "id", "updated_at":
			continue
		case "created_at":
			parts = append(parts, "now()")
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
		case "ID", "CreatedAt", "UpdatedAt":
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
	i := 1
	for _, fl := range m.Model.Fields {
		switch fl.ColumnName {
		case "id", "created_at":
			continue
		case "updated_at":
			parts = append(parts, fmt.Sprintf(`"%s"=now()`, fl.ColumnName))
		default:
			parts = append(parts, fmt.Sprintf(`"%s"=$%d`, fl.ColumnName, i))
			i++
		}
	}
	return strings.Join(parts, ", ")
}

func GetUpdateValuesLength(m StructTmplData) string {
	count := 0
	for _, field := range m.Model.Fields {
		switch field.ColumnName {
		case "created_at", "updated_at":
			continue
		default:
			count++
		}
	}
	return fmt.Sprintf("$%d", count)
}
