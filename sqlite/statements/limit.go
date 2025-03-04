package statements

import (
	"strings"
)

type Limit struct {
	Limit  int64
	Offset int64
	Values []interface{}
}

// Comment
func (ctx *Limit) Statement() (string, error) {
	stmt := []string{}

	if ctx.Limit != 0 {
		stmt = append(stmt, "LIMIT ?")

		ctx.Values = append(ctx.Values, ctx.Limit)
	}

	if ctx.Limit != 0 && ctx.Offset != 0 {
		stmt = append(stmt, "OFFSET ?")

		ctx.Values = append(ctx.Values, ctx.Offset)
	}

	return strings.Join(stmt, " "), nil
}
