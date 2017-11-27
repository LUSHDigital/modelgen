package tmpl

import (
	"fmt"
	"html/template"
	"strings"
)

var FuncMap = template.FuncMap{
	"insert_fields": GetInsertFields,
	"insert_values": GetInsertValues,
	"insert_args":   GetInsertArgs,
	"scan_fields":   GetScanFields,
	"update_args":   GetUpdateArgs,
	"update_values": GetUpdateValues,
}

func GetInsertFields(fields []TmplField) string {
	var parts []string
	for _, fl := range fields {
		if fl.ColumnName == "id" {
			continue
		}
		parts = append(parts, "`"+fl.ColumnName+"`")
	}
	return strings.Join(parts, ", ")
}

func GetInsertValues(fields []TmplField) string {
	var parts []string
	for _, fl := range fields {
		switch fl.ColumnName {
		case "id":
			continue
		case "created_at":
			parts = append(parts, "NOW()")
			continue
		default:
			parts = append(parts, "?")
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
		parts = append(parts, fmt.Sprintf("%s.%s", m.Receiver, fl.Name))
	}
	return strings.Join(parts, ", ")
}

func GetScanFields(m StructTmplData) template.HTML {
	var parts []string
	for _, fl := range m.Model.Fields {
		parts = append(parts, fmt.Sprintf("&%s.%s", m.Receiver, fl.Name))
	}
	return template.HTML(strings.Join(parts, ", "))
}

func GetUpdateArgs(m StructTmplData) template.HTML {
	var parts []string
	for _, fl := range m.Model.Fields {
		switch fl.Name {
		case "ID", "CreatedAt":
			continue
		case "UpdatedAt":
			if fl.Nullable {
				parts = append(parts, "ToNullTime(time.Now())")
				continue
			}
			parts = append(parts, "time.Now()")
			continue
		}
		parts = append(parts, fmt.Sprintf("%s.%s", m.Receiver, fl.Name))
	}
	return template.HTML(strings.Join(parts, ", "))
}

func GetUpdateValues(m StructTmplData) string {
	var parts []string
	for _, fl := range m.Model.Fields {
		switch fl.Name {
		case "ID", "CreatedAt":
			continue
		default:
			parts = append(parts, fmt.Sprintf("`%s`=?", fl.ColumnName))
		}
	}
	return strings.Join(parts, ", ")
}
