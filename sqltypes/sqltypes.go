package sqltypes

import (
	"log"
	"strconv"
	"strings"
)

// Explain wraps the explain query results for a given table
type Explain struct {
	Field   *string
	Type    *string
	Null    *string
	Key     *string
	Default *string
	Extra   *string
}

// SQLType unwraps a SQL data type
type SQLType struct {
	notNull  string
	nullable string
}

var dateType = SQLType{"time.Time", "NullTime"}
var stringType = SQLType{"string", "NullString"}
var byteSliceType = SQLType{"[]byte", "[]byte"}
var intType = SQLType{"int64", "NullInt64"}
var boolType = SQLType{"bool", "NullBool"}
var floatType = SQLType{"float64", "NullFloat64"}
var jsonType = SQLType{"RawJSON", "RawJSON"}

var dataTypes = map[string]SQLType{
	"char":    stringType,
	"varchar": stringType,

	"tinytext":   stringType,
	"text":       stringType,
	"mediumtext": stringType,
	"longtext":   stringType,

	"json": jsonType,

	"enum": stringType,
	"set":  stringType,

	//"geometry":           stringType,
	//"point":              stringType,
	//"linestring":         stringType,
	//"polygon":            stringType,
	//"multipoint":         stringType,
	//"multilinestring":    stringType,
	//"multipolygon":       stringType,
	//"geometrycollection": stringType,

	"bit":        byteSliceType,
	"binary":     byteSliceType,
	"varbinary":  byteSliceType,
	"tinyblob":   byteSliceType,
	"mediumblob": byteSliceType,
	"blob":       byteSliceType,
	"longblob":   byteSliceType,

	"tinyint_as_bool": boolType,
	"tinyint":         intType,
	"smallint":        intType,
	"mediumint":       intType,
	"int":             intType,
	"bigint":          intType,

	"float":   floatType,
	"double":  floatType,
	"decimal": floatType,

	"date":      dateType,
	"datetime":  dateType,
	"timestamp": dateType,
	"time":      stringType,
	"year":      intType,
}

// AssertType figures out which go type should be used, based on the SQL type.
func AssertType(columnType string, nullable string) string {
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

	for dataType, sqlType := range dataTypes {
		if extractedType == dataType {
			if nul {
				return sqlType.nullable
			}
			return sqlType.notNull
		}
	}
	log.Fatalf("unsupported type: %v, please raise an issue with us if you'd like to request support.", extractedType)
	return ""
}

// NeedsImport handles adding the imports for types that require them.
func NeedsImport(typ string) (imp string, ok bool) {
	switch typ {
	case "time.Time":
		return "time", true
	case "json.RawMessage":
		return "encoding/json", true
	default:
		return
	}
}
