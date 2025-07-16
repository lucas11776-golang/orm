package sqlite

import (
	"strings"

	"github.com/lucas11776-golang/orm"
	"github.com/lucas11776-golang/orm/drivers/sql/statements"
	"github.com/lucas11776-golang/orm/utils/slices"
)

type QueryValues []interface{}

type QueryBuilder struct {
	Statement *orm.Statement
	Values    QueryValues
}

type Statement interface {
	Statement() (string, error)
	Values() []interface{}
}

// Comment
func (ctx *QueryBuilder) queryStatementBuild(statements []Statement) (string, error) {
	queries := []string{}

	for _, stmt := range statements {
		qry, err := stmt.Statement()

		if err != nil {
			return "", err
		}

		queries = append(queries, qry)

		ctx.Values = append(ctx.Values, stmt.Values()...)
	}

	queries = slices.Filter(queries, func(item string) bool {
		return item == ""
	})

	return strings.Join(queries, "\r\n"), nil
}

// Comment
func (ctx *QueryBuilder) Query() (string, QueryValues, error) {
	query, err := ctx.queryStatementBuild([]Statement{
		&statements.Select{
			Table:  ctx.Statement.Table,
			Select: ctx.Statement.Select,
		},
		&statements.Join{
			Table: ctx.Statement.Table,
			Join:  ctx.Statement.Joins,
		},
		&statements.Where{
			Where: ctx.Statement.Where,
		},
		&statements.OrderBy{
			OrderBy: ctx.Statement.OrderBy,
		},
		&statements.Limit{
			Limit:  ctx.Statement.Limit,
			Offset: ctx.Statement.Offset,
		},
	})

	if err != nil {
		return "", nil, err
	}

	return query, ctx.Values, nil
}

// Comment
func (ctx *QueryBuilder) Count() (string, QueryValues, error) {
	query, err := ctx.queryStatementBuild([]Statement{
		&statements.Select{
			Table:  ctx.Statement.Table,
			Select: orm.Select{orm.COUNT{"*", "total"}},
		},
		&statements.Join{
			Table: ctx.Statement.Table,
			Join:  ctx.Statement.Joins,
		},
		&statements.Where{
			Where: ctx.Statement.Where,
		},
		&statements.OrderBy{
			OrderBy: ctx.Statement.OrderBy,
		},
		&statements.Limit{
			Limit:  ctx.Statement.Limit,
			Offset: ctx.Statement.Offset,
		},
	})

	if err != nil {
		return "", nil, err
	}

	return query, ctx.Values, nil
}

// Comment
func (ctx *QueryBuilder) Insert() (string, QueryValues, error) {
	stmt := []string{"INSERT INTO", statements.SafeKey(ctx.Statement.Table)}

	keys := []string{}
	values := []string{}

	for k, v := range ctx.Statement.Values {
		keys = append(keys, statements.SafeKey(k))
		values = append(values, "?")
		ctx.Values = append(ctx.Values, v)
	}

	stmt = append(stmt, strings.Join([]string{"(", strings.Join(keys, ", "), ")"}, ""))
	stmt = append(stmt, "VALUES")
	stmt = append(stmt, strings.Join([]string{"(", strings.Join(values, ", "), ")"}, ""))

	return strings.Join(stmt, " "), ctx.Values, nil
}

// Comment
func (ctx *QueryBuilder) Update() (string, QueryValues, error) {
	stmt := []string{
		"UPDATE",
		statements.SPACE + statements.SafeKey(ctx.Statement.Table),
		"SET",
	}

	fields := []string{}

	for k, v := range ctx.Statement.Values {
		fields = append(fields, strings.Join([]string{statements.SafeKey(k), "?"}, " = "))
		ctx.Values = append(ctx.Values, v)
	}

	stmt = append(stmt, statements.SPACE+strings.Join(fields, ", "))

	where := &statements.Where{Where: ctx.Statement.Where}
	whereStmt, err := where.Statement()

	if err != nil {
		return "", nil, err
	}

	stmt = append(stmt, whereStmt)

	ctx.Values = append(ctx.Values, where.Values()...)

	return strings.Join(stmt, "\r\n"), ctx.Values, nil
}

// Comment
func (ctx *QueryBuilder) Delete() (string, QueryValues, error) {
	stmt := &statements.Where{
		Where: ctx.Statement.Where,
	}

	query, err := stmt.Statement()

	if err != nil {
		return "", nil, err
	}

	return strings.Join([]string{
		"DELETE FROM",
		statements.SPACE + statements.SafeKey(ctx.Statement.Table),
		query,
	}, "\r\n"), stmt.Values(), nil
}
