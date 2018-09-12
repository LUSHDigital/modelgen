package templates

import (
	"fmt"
	"html/template"
	"strings"

	"github.com/gobuffalo/packr"
	"github.com/nicklanng/modelgen/model"
)

var (
	box     packr.Box
	funcMap = template.FuncMap{
		"insert_fields":       GetInsertFields,
		"insert_values":       GetInsertValues,
		"insert_args":         GetInsertArgs,
		"scan_fields":         GetScanFields,
		"update_args":         GetUpdateArgs,
		"update_values":       GetUpdateValues,
		"upsert_fields":       GetUpsertFields,
		"upsert_values":       GetUpsertValues,
		"upsert_on_duplicate": GetUpsertOnDuplicate,
		"upsert_args":         GetUpsertArgs,
	}
)

func init() {
	box = packr.NewBox("../templates")
}

func GetInsertFields(fields []model.Field) string {
	var parts []string
	for _, fl := range fields {
		if fl.ColumnName == "id" {
			continue
		}
		parts = append(parts, "`"+fl.ColumnName+"`")
	}
	return strings.Join(parts, ", ")
}

func GetInsertValues(fields []model.Field) string {
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

func GetInsertArgs(m model.TemplateData) string {
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

func GetScanFields(m model.TemplateData) template.HTML {
	var parts []string
	for _, fl := range m.Model.Fields {
		parts = append(parts, fmt.Sprintf("&%s.%s", m.Receiver, fl.Name))
	}
	return template.HTML(strings.Join(parts, ", "))
}

func GetUpdateArgs(m model.TemplateData) template.HTML {
	var parts []string
	for _, fl := range m.Model.Fields {
		switch fl.Name {
		case "ID", "CreatedAt", "UpdatedAt":
			continue
		}
		parts = append(parts, fmt.Sprintf("%s.%s", m.Receiver, fl.Name))
	}
	return template.HTML(strings.Join(parts, ", "))
}

func GetUpdateValues(m model.TemplateData) string {
	var parts []string
	for _, fl := range m.Model.Fields {
		switch fl.Name {
		case "ID", "CreatedAt":
			continue
		case "UpdatedAt":
			parts = append(parts, fmt.Sprintf("`%s`=UTC_TIMESTAMP()", fl.ColumnName))
		default:
			parts = append(parts, fmt.Sprintf("`%s`=?", fl.ColumnName))
		}
	}
	return strings.Join(parts, ", ")
}

func GetUpsertFields(fields []model.Field) string {
	var parts []string
	for _, fl := range fields {
		parts = append(parts, "`"+fl.ColumnName+"`")
	}
	return strings.Join(parts, ", ")
}

func GetUpsertValues(fields []model.Field) string {
	var parts []string
	for _, fl := range fields {
		switch fl.ColumnName {
		case "created_at":
			parts = append(parts, "NOW()")
			continue
		default:
			parts = append(parts, "?")
		}
	}
	return strings.Join(parts, ", ")
}

func GetUpsertOnDuplicate(m model.TemplateData) string {
	var parts []string
	for _, fl := range m.Model.Fields {
		switch fl.Name {
		case "CreatedAt":
			continue
		case "ID":
			parts = append(parts, fmt.Sprintf("`%s`=LAST_INSERT_ID(`%s`)", fl.ColumnName, fl.ColumnName))
		case "UpdatedAt":
			parts = append(parts, fmt.Sprintf("`%s`=UTC_TIMESTAMP()", fl.ColumnName))
		default:
			parts = append(parts, fmt.Sprintf("`%s`=VALUES(`%s`)", fl.ColumnName, fl.ColumnName))
		}
	}
	return strings.Join(parts, ", ")
}

func GetUpsertArgs(m model.TemplateData) string {
	var parts []string
	for _, fl := range m.Model.Fields {
		switch fl.Name {
		case "CreatedAt":
			continue
		}
		parts = append(parts, fmt.Sprintf("%s.%s", m.Receiver, fl.Name))
	}
	return strings.Join(parts, ", ")
}
