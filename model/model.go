package model

// SQLType unwraps a SQL data type
type SQLType struct {
	NotNull  string
	Nullable string
}

var (
	TypeDate      = SQLType{"time.Time", "NullTime"}
	TypeString    = SQLType{"string", "NullString"}
	TypeByteSlice = SQLType{"[]byte", "[]byte"}
	TypeInt       = SQLType{"int64", "NullInt64"}
	TypeBool      = SQLType{"bool", "NullBool"}
	TypeFloat     = SQLType{"float64", "NullFloat64"}
	TypeJSON      = SQLType{"RawJSON", "RawJSON"}
)

// Field defines a table field template
type Field struct {
	Name       string
	Type       string
	ColumnName string
	Nullable   bool
}

type EntityDescriptor struct {
	Name      string
	TableName string
	Fields    []Field
	Imports   map[string]struct{}
	Comment   string
}

type TemplateData struct {
	Model       EntityDescriptor
	Receiver    string
	PackageName string
}

type Migration struct {
	TableName string
	Up        string
	Down      string
	Order     int
}
