package statements

import (
	"orm"
	"strings"
)

type OrderBy struct {
	Order orm.Order
}

// Comment
func (ctx *OrderBy) Statement() (string, error) {
	stmt := []string{}

	if ctx.Order != "" {
		stmt = append(stmt, strings.Join([]string{"ORDER BY", string(ctx.Order)}, " "))
	}

	return strings.Join(stmt, " "), nil
}
