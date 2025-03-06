package sqlite

import (
	"orm"
	"orm/sqlite/statements"
	"strings"
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
func (ctx *QueryBuilder) Insert() (string, error) {
	stmt := []string{"INSERT INTO", strings.Join([]string{"`", ctx.Statement.Table, "`"}, "")}

	keys := []string{}
	values := []string{}

	for k, v := range ctx.Statement.Values {
		keys = append(keys, strings.Join([]string{"`", k, "`"}, ""))
		values = append(values, "?")
		ctx.Values = append(ctx.Values, v)
	}

	stmt = append(stmt, strings.Join([]string{"(", strings.Join(keys, ","), ")"}, ""))
	stmt = append(stmt, "VALUES")
	stmt = append(stmt, strings.Join([]string{"(", strings.Join(values, ","), ")"}, ""))

	return strings.Join(stmt, " "), nil
}

// Comment
func (ctx *QueryBuilder) Update() error {
	return nil
}
