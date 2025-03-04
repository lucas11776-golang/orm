package statements

import (
	"orm"
	"strings"
)

type OrderBy struct {
	OrderBy orm.OrderBy
	values  []interface{}
}

// Comment
func (ctx *OrderBy) Statement() (string, error) {
	stmt := []string{}

	if ctx.OrderBy[0] != nil {
		stmt = append(stmt, strings.Join([]string{
			"ORDER BY", SafeKey(ctx.OrderBy[0].(string)), string(ctx.OrderBy[1].(orm.Order)),
		}, " "))
	}

	return strings.Join(stmt, " "), nil
}

// Comment
func (ctx *OrderBy) Values() []interface{} {
	return ctx.values
}
