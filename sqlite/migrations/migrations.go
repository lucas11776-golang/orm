package migrations

import (
	"database/sql"
	"fmt"
	"orm/sqlite/statements"
	"strings"
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
				return "", fmt.Errorf("Default must have a value")
			}

			ars = append(ars, strings.Join([]string{"DEFAULT", argArr[1]}, " "))

			break

		default:
			return "", fmt.Errorf("Argument of %s is not supported in migrations", argArr[0])
		}
	}

	return strings.Join(ars, " "), nil
}

// Comment
func (ctx *Migration) types(t string) (string, error) {
	switch strings.ToUpper(t) {
	case "PRIMARY_KEY":
		return "INTEGER PRIMARY KEY AUTOINCREMENT", nil

	case "DATETIME":
		return "DATETIME", nil

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
		return "", fmt.Errorf("Type of %s is not support by migration", t)
	}
}

// Comment
func (ctx *Migration) Statement(t string, column string, args ...string) (string, error) {
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
