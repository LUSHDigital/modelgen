package generation

// TmplStruct defines the table data to pass to the models
type TmplStruct struct {
	Name      string
	TableName string
	Fields    []TmplField
	Imports   map[string]struct{}
}

// TmplField defines a table field template
type Field struct {
	Name       string
	Type       string
	ColumnName string
	Nullable   bool
}

// StructTmplData defines the top level struct data to pass to the models
type StructTmplData struct {
	Model       TmplStruct
	Receiver    string
	PackageName string
}
