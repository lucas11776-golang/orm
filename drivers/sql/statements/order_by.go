package statements

import (
	"strings"

	"github.com/lucas11776-golang/orm"
)

type OrderBy struct {
	OrderBy orm.OrderBy
	values  []interface{}
}

// Comment
func (ctx *OrderBy) Statement() (string, error) {
	stmt := []string{}

	if len(ctx.OrderBy.Columns) != 0 {

		// switch

		stmt = append(stmt, strings.Join([]string{
			"ORDER BY", SafeKey(ctx.OrderBy.Columns), ctx.OrderBy.Order.(string),
		}, " "))
	}

	return strings.Join(stmt, " "), nil
}

// Comment
func (ctx *OrderBy) Values() []interface{} {
	return ctx.values
}
