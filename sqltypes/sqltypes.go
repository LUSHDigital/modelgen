package sqltypes

import (
	"log"
	"strings"
)

// Explain wraps the explain query results for a given table
type Explain struct {
	Field             *string
	Type              *string
	CharLength        *int64
	OctetLength       *int64
	NumericPrecision  *int64
	NumericScale      *int64
	DateTimePrecision *string
	Null              *string
	Key               *string
	Default           *string
	Extra             *string
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

var stringArrayType = SQLType{"StringArray", "StringArray"}
var intArrayType = SQLType{"IntArray", "IntArray"}
var floatArrayType = SQLType{"FloatArray", "FloatArray"}
var boolArrayType = SQLType{"BoolArray", "BoolArray"}

var dataTypes = map[string]SQLType{
	"STRING":  stringType,
	"INT":     intType,
	"FLOAT":   floatType,
	"DECIMAL": floatType,
	"BOOL":    boolType,

	"DATE":      dateType,
	"TIMESTAMP": dateType,

	"BYTES": byteSliceType,

	"JSON": jsonType,

	"STRING[]":  stringArrayType,
	"INT[]":     intArrayType,
	"FLOAT[]":   floatArrayType,
	"DECIMAL[]": floatArrayType,
	"BOOL[]":    boolArrayType,
}

// AssertType figures out which go type should be used, based on the SQL type.
func AssertType(columnType string, nullable string) string {
	nul := nullable == "YES"
	bits := strings.Split(columnType, "(")
	extractedType := bits[0]

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
