package migrations

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	"github.com/lucas11776-golang/orm"
	"github.com/lucas11776-golang/orm/databases/sqlite/statements"
	str "github.com/lucas11776-golang/orm/utils/strings"
)

type Migration struct {
	DB *sql.DB
}

// Comment
func (ctx *Migration) args(args []string) (string, error) {
	ars := []string{}

	for _, arg := range args {
		argArr := strings.Split(arg, ":")

		switch strings.ToUpper(argArr[0]) {
		case "DEFAULT":
			if len(argArr) != 2 {
				return "", fmt.Errorf("default must have a value")
			}

			ars = append(ars, strings.Join([]string{"DEFAULT", argArr[1]}, " "))

		case "NOT_NULL":
			ars = append(ars, "NOT NULL")

		default:
			return "", fmt.Errorf("argument of %s is not supported in migrations", argArr[0])
		}
	}

	return strings.Join(ars, " "), nil
}

// Comment
func (ctx *Migration) types(t string) (string, error) {
	switch strings.ToUpper(t) {
	case "PRIMARY_KEY":
		return "INTEGER PRIMARY KEY AUTOINCREMENT", nil

	case "TIMESTAMP":
		return "TIMESTAMP", nil

	case "TIMESTAMP_CURRENT":
		return "TIMESTAMP DEFAULT CURRENT_TIMESTAMP", nil

	case "DATETIME":
		return "DATETIME", nil

	case "DATE":
		return "DATE", nil

	case "DATETIME_CURRENT":
		return "DATETIME DEFAULT CURRENT_TIMESTAMP", nil

	case "INTEGER":
		return "INTEGER", nil

	case "FLOAT":
		return "FLOAT", nil

	case "STRING":
		return "VARCHAR", nil

	case "TEXT":
		return "TEXT", nil

	case "BOOLEAN":
		return "BOOLEAN", nil

	default:
		return "", fmt.Errorf("type of %s is not support by migration", t)
	}
}

// Comment
func (ctx *Migration) columnStatement(column string, t string, args ...string) (string, error) {
	tp, err := ctx.types(t)

	if err != nil {
		return "", err
	}

	ars, err := ctx.args(args)

	if err != nil {
		return "", nil
	}

	return strings.Trim(strings.Join([]string{statements.SafeKey(column), tp, ars}, " "), " "), nil
}

// Comment
func (ctx *Migration) table(name string) string {
	return str.Plural(strings.ToLower(name))
}

// Comment
func (ctx *Migration) generateModelTableQuery(model interface{}) (string, error) {
	queries := []string{}

	if reflect.ValueOf(model).Type().Kind() != reflect.Struct {
		return "", fmt.Errorf("type of model (%v) is not a (%s)", model, reflect.Struct)
	}

	mVal := reflect.ValueOf(model)
	table := ctx.table(mVal.Type().Name())

	for i := 0; i < mVal.NumField(); i++ {
		tag := mVal.Type().Field(i).Tag

		col := tag.Get("column")

		if t := tag.Get("table"); t != "" {
			table = t
		}

		if col == "" {
			continue
		}

		tp := strings.Split(tag.Get("type"), ",")

		if len(tp) == 0 {
			return "", fmt.Errorf("type is required for column %s", statements.SafeKey(col))
		}

		stmt, err := ctx.columnStatement(col, tp[0], tp[1:]...)

		if err != nil {
			return "", err
		}

		queries = append(queries, strings.Join([]string{statements.SPACE, stmt}, ""))
	}

	return strings.Join([]string{
		fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (", table),
		strings.Join(queries, ",\r\n"),
		");",
	}, "\r\n"), nil
}

// Comment
func (ctx *Migration) modelsTablesQueries(models orm.Models) (string, error) {
	queries := []string{}

	for _, m := range models {
		query, err := ctx.generateModelTableQuery(m)

		if err != nil {
			return "", err
		}

		queries = append(queries, query)
	}

	return strings.Join(queries, "\r\n\r\n"), nil
}

// Comment
func (ctx *Migration) Migrate(models orm.Models) error {
	query, err := ctx.modelsTablesQueries(models)

	if err != nil {
		return err
	}

	_, err = ctx.DB.Exec(query)

	return err
}

// Comment
func (ctx *Migration) Truncate(models orm.Models) error {
	return nil
}
