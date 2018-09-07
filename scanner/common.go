package scanner

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nicklanng/modelgen/model"
)

// mapType figures out which go type should be used, based on the SQL type.
func mapType(columnType string, nullable string, typeMap map[string]model.SQLType) (string, error) {
	nul := nullable == "YES"

	bits := strings.Split(columnType, "(")
	extractedType := bits[0]
	var extractedLength int
	if len(bits) > 1 {
		idx := strings.Index(bits[1], ")")
		inner := bits[1][:idx]

		// ignoring error here because if any length cannot be found,
		// this does not constitute an error, but an expected outcome
		extractedLength, _ = strconv.Atoi(inner)
		if extractedType == "tinyint" && extractedLength == 1 {
			extractedType = "tinyint_as_bool"
		}
	}

	for dataType, sqlType := range typeMap {
		if extractedType == dataType {
			if nul {
				return sqlType.Nullable, nil
			}
			return sqlType.NotNull, nil
		}
	}
	return "", fmt.Errorf("unsupported type: %v, please raise an issue with us if you'd like to request support", extractedType)
}

func needsImport(typ string) (imp string, ok bool) {
	switch typ {
	case "time.Time":
		return "time", true
	case "json.RawMessage":
		return "encoding/json", true
	default:
		return
	}
}

// toPascalCase take a snake_case string and converts it to PascalCase
func toPascalCase(field string) string {
	parts := strings.Split(field, "_")
	for i := 0; i < len(parts); i++ {
		parts[i] = shouldCapitilize(parts[i])
	}
	field = strings.Join(parts, "_")
	field = strings.Replace(field, "_", " ", -1)
	field = strings.Title(field)
	field = strings.Replace(field, " ", "", -1)
	return field
}

// shouldCapitilize defines if the acronym should be capitalized
func shouldCapitilize(word string) string {
	for _, acc := range acronyms {
		if strings.ToLower(word) == acc {
			return strings.ToUpper(word)
		}
	}
	return word
}
