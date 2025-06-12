package statements

import (
	"strings"

	"github.com/lucas11776-golang/orm"
	"github.com/lucas11776-golang/orm/utils/slices"
)

type OrderBy struct {
	OrderBy orm.OrderBy
	values  []interface{}
}

// Comment
func (ctx *OrderBy) Statement() (string, error) {
	stmt := []string{}

	if len(ctx.OrderBy.Columns) != 0 {

		columns := strings.Join(slices.Map(ctx.OrderBy.Columns, func(column string) string {
			return SafeKey(column)
		}), ",")

		stmt = append(stmt, strings.Join([]string{
			"ORDER BY", columns, string(ctx.OrderBy.Order),
		}, " "))
	}

	// if ctx.OrderBy[0] != nil {
	// 	stmt = append(stmt, strings.Join([]string{
	// 		"ORDER BY", columns, string(ctx.OrderBy[1].(orm.Order)),
	// 	}, " "))
	// }

	return strings.Join(stmt, " "), nil
}

// Comment
func (ctx *OrderBy) Values() []interface{} {
	return ctx.values
}
