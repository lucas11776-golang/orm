package sql

import (
	"strings"

	"github.com/lucas11776-golang/orm"
	"github.com/lucas11776-golang/orm/utils/slices"
)

type SQLBuilder struct {
	Statement    *orm.Statement
	QueryBuilder QueryBuilder
}

// Comment
func (ctx *SQLBuilder) build(statements []Statement) (string, QueryValues, error) {
	query := []string{}
	values := QueryValues{}

	for _, stmt := range statements {
		statement, err := stmt.Statement()

		if err != nil {
			return "", nil, err
		}

		query = append(query, statement)
		values = append(values, stmt.Values()...)
	}

	query = slices.Filter(query, func(item string) bool {
		return item == ""
	})

	return strings.Join(query, "\r\n"), values, nil
}

// Comment
func (ctx *SQLBuilder) Query() (string, QueryValues, error) {
	return ctx.build([]Statement{
		ctx.QueryBuilder.Select(ctx.Statement),
		ctx.QueryBuilder.Join(ctx.Statement),
		ctx.QueryBuilder.Where(ctx.Statement),
		ctx.QueryBuilder.OrderBy(ctx.Statement),
		ctx.QueryBuilder.Limit(ctx.Statement),
	})
}

// Comment
func (ctx *SQLBuilder) Count() (string, QueryValues, error) {
	return ctx.build([]Statement{
		ctx.QueryBuilder.Select(func() *orm.Statement {
			ctx.Statement.Select = orm.Select{orm.COUNT{"*", "total"}}

			return ctx.Statement
		}()),
		ctx.QueryBuilder.Join(ctx.Statement),
		ctx.QueryBuilder.Where(ctx.Statement),
	})
}

// Comment
func (ctx *SQLBuilder) Insert() (string, QueryValues, error) {
	return ctx.build([]Statement{
		ctx.QueryBuilder.Insert(ctx.Statement),
	})
}

// Comment
func (ctx *SQLBuilder) Update() (string, QueryValues, error) {
	return ctx.build([]Statement{
		ctx.QueryBuilder.Update(ctx.Statement),
	})
}

// Comment
func (ctx *SQLBuilder) Delete() (string, QueryValues, error) {
	return ctx.build([]Statement{
		ctx.QueryBuilder.Delete(ctx.Statement),
	})
}
