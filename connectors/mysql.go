package connectors

import (
	"database/sql"
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/nicklanng/modelgen/model"
	"github.com/nicklanng/modelgen/sqlfmt"
	"github.com/nicklanng/modelgen/templates"
)

type mySQLExplain struct {
	Field   *string
	Type    *string
	Null    *string
	Key     *string
	Default *string
	Extra   *string
}

var mysqlDataTypes = map[string]model.SQLType{
	"char":    model.TypeString,
	"varchar": model.TypeString,

	"tinytext":   model.TypeString,
	"text":       model.TypeString,
	"mediumtext": model.TypeString,
	"longtext":   model.TypeString,

	"json": model.TypeJSON,

	"enum": model.TypeString,
	"set":  model.TypeString,

	"bit":        model.TypeByteSlice,
	"binary":     model.TypeByteSlice,
	"varbinary":  model.TypeByteSlice,
	"tinyblob":   model.TypeByteSlice,
	"mediumblob": model.TypeByteSlice,
	"blob":       model.TypeByteSlice,
	"longblob":   model.TypeByteSlice,

	"tinyint_as_bool": model.TypeBool,
	"tinyint":         model.TypeInt,
	"smallint":        model.TypeInt,
	"mediumint":       model.TypeInt,
	"int":             model.TypeInt,
	"bigint":          model.TypeInt,

	"float":   model.TypeFloat,
	"double":  model.TypeFloat,
	"decimal": model.TypeFloat,

	"date":      model.TypeDate,
	"datetime":  model.TypeDate,
	"timestamp": model.TypeDate,
	"time":      model.TypeString,
	"year":      model.TypeInt,
}

type MySQL struct {
	username string
	password string
	host     string
	port     string
	database string
}

func NewMySQL(username, password, host, port, database string) *MySQL {
	return &MySQL{
		username: username,
		password: password,
		host:     host,
		port:     port,
		database: database,
	}
}

func (m *MySQL) Connect() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", m.username, m.password, m.host, m.port, m.database)

	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// check for a valid connection
	if err := conn.Ping(); err != nil {
		return nil, err
	}

	return conn, nil
}

func (m *MySQL) QueryStructure(conn *sql.DB) ([]model.EntityDescriptor, error) {
	tables, err := m.queryTables(conn)
	if err != nil {
		return nil, err
	}

	if len(tables) == 0 {
		return nil, errors.New("no tables to read")
	}

	var structs []model.EntityDescriptor
	for tableName, comment := range tables {
		explanations, err := m.explainTable(conn, tableName)
		if err != nil {
			return nil, err
		}

		s, err := m.parseExplanation(tableName, comment, explanations)
		if err != nil {
			return nil, err
		}

		structs = append(structs, s)
	}

	return structs, nil
}

func (m *MySQL) queryTables(conn *sql.DB) (map[string]string, error) {
	tables := make(map[string]string)

	stmt := `SELECT table_name, column_comment
			 FROM information_schema.columns AS c
			 WHERE c.column_key = "PRI"
			 AND c.table_schema = ?
		     AND column_name = "id"`

	rows, err := conn.Query(stmt, m.database)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var name string
		var comment string
		if err := rows.Scan(&name, &comment); err != nil {
			return nil, err
		}
		tables[name] = comment
	}

	return tables, nil
}

func (m *MySQL) explainTable(conn *sql.DB, table string) ([]mySQLExplain, error) {
	var tableExplanations []mySQLExplain

	rows, err := conn.Query("EXPLAIN " + table)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var explanation mySQLExplain
		if err := rows.Scan(&explanation.Field, &explanation.Type, &explanation.Null, &explanation.Key, &explanation.Default, &explanation.Extra); err != nil {
			return nil, err
		}
		tableExplanations = append(tableExplanations, explanation)
	}

	return tableExplanations, nil
}

func (m *MySQL) parseExplanation(table, comment string, explanations []mySQLExplain) (model.EntityDescriptor, error) {
	t := model.EntityDescriptor{
		Name:      sqlfmt.ToPascalCase(table),
		TableName: table,
		Imports:   make(map[string]struct{}),
		Comment:   comment,
	}

	for _, expl := range explanations {
		typ, err := sqlfmt.MapType(*expl.Type, *expl.Null, mysqlDataTypes)
		if err != nil {
			return model.EntityDescriptor{}, err
		}

		f := model.Field{
			Name:       sqlfmt.ToPascalCase(*expl.Field),
			Type:       typ,
			ColumnName: strings.ToLower(*expl.Field),
			Nullable:   *expl.Null == "YES",
		}

		t.Fields = append(t.Fields, f)
		if imp, ok := sqlfmt.NeedsImport(f.Type); ok {
			t.Imports[imp] = struct{}{}
		}
	}

	return t, nil
}

func (m *MySQL) FillTemplates(conn *sql.DB, models []model.EntityDescriptor, outputPath, packageName string) error {
	var err error

	w := templates.NewTemplateWriter("mysql", outputPath, packageName)

	if err = w.WriteModels(models); err != nil {
		return err
	}

	if err = w.WriteHelpers(); err != nil {
		return err
	}

	if err = w.WriteHelperTests(); err != nil {
		return err
	}

	return nil
}

func (m *MySQL) QueryMigrations(conn *sql.DB, models []model.EntityDescriptor) ([]model.Migration, error) {
	var (
		table      string
		upQuery    string
		downQuery  string
		row        *sql.Row
		order      int
		migrations []model.Migration
		err        error
	)

	migrations = []model.Migration{}

	for i := range models {
		if row = conn.QueryRow("SHOW CREATE TABLE " + models[i].TableName); err != nil {
			return nil, err
		}
		row.Scan(&table, &upQuery)

		if order, err = sqlfmt.GetOrderFromComment(models[i].Comment); err != nil {
			return nil, err
		}

		downQuery = fmt.Sprintf("DROP TABLE IF EXISTS `%s`", models[i].TableName)

		migrations = append(migrations, model.Migration{
			TableName: models[i].TableName,
			Up:        upQuery,
			Down:      downQuery,
			Order:     order,
		})
	}

	sort.Slice(migrations, func(i, j int) bool {
		return migrations[i].Order < migrations[j].Order
	})

	return migrations, nil
}
