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
func (ctx *Migration) Statement(t string, column string) (string, error) {
	switch t {
	case "primary_key":
		return strings.Join([]string{statements.SafeKey(column), "integer primary key autoincrement"}, " "), nil

	case "datetime":
		return "", nil

	case "datetime_auto":
		return "", nil

	case "integer":
		return "", nil

	case "float":
		return "", nil

	case "string":
		return "", nil

	case "text":
		return "", nil

	case "boolean":
		return "", nil

	default:
		return "", fmt.Errorf("Type of %s is not support by migration", t)
	}
}
