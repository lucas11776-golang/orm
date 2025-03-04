package statements

import (
	"strings"
)

type Limit struct {
	Limit  int64
	Offset int64
	values []interface{}
}

// Comment
func (ctx *Limit) Statement() (string, error) {
	stmt := []string{}

	if ctx.Limit != 0 {
		stmt = append(stmt, "LIMIT ?")

		ctx.values = append(ctx.values, ctx.Limit)
	}

	if ctx.Limit != 0 && ctx.Offset != 0 {
		stmt = append(stmt, "OFFSET ?")

		ctx.values = append(ctx.values, ctx.Offset)
	}

	return strings.Join(stmt, " "), nil
}

// Comment
func (ctx *Limit) Values() []interface{} {
	return ctx.values
}
