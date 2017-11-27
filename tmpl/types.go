package tmpl

// TmplStructs is a collection on TmplStruct
type TmplStructs []TmplStruct
// TmplStruct defines the table data to pass to the models
type TmplStruct struct {
	Name      string
	TableName string
	Fields    []TmplField
	Imports   map[string]struct{}
}

// TmplField defines a table field template
type TmplField struct {
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
