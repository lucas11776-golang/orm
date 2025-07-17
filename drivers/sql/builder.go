package sql

import (
	"strings"

	"github.com/lucas11776-golang/orm"
	"github.com/lucas11776-golang/orm/drivers/sql/statements"
	"github.com/lucas11776-golang/orm/utils/slices"
)

type QueryValues []interface{}

type QueryBuilder struct {
	Statement *orm.Statement

	Builder SQLQueryBuilder

	Values QueryValues
}

type Statement interface {
	Statement() (string, error)
	Values() []interface{}
}

type SQLQueryBuilder interface {
	Select(statement *orm.Statement) Statement
	Join(statement *orm.Statement) Statement
	Where(statement *orm.Statement) Statement
	OrderBy(statement *orm.Statement) Statement
	Limit(statement *orm.Statement) Statement
	Insert(statement *orm.Statement) Statement
	// TablePrimaryKey(statement *orm.Statement) (string, error)
}

type DefaultQueryBuilder struct{}

// Comment
func (ctx *DefaultQueryBuilder) Select(statement *orm.Statement) Statement {
	return &statements.Select{
		Table:  statement.Table,
		Select: statement.Select,
	}
}

// Comment
func (ctx *DefaultQueryBuilder) Join(statement *orm.Statement) Statement {
	return &statements.Join{
		Table: statement.Table,
		Join:  statement.Joins,
	}
}

// Comment
func (ctx *DefaultQueryBuilder) Where(statement *orm.Statement) Statement {
	return &statements.Where{
		Where: statement.Where,
	}
}

// Comment
func (ctx *DefaultQueryBuilder) OrderBy(statement *orm.Statement) Statement {
	return &statements.OrderBy{
		OrderBy: statement.OrderBy,
	}
}

// Comment
func (ctx *DefaultQueryBuilder) Limit(statement *orm.Statement) Statement {
	return &statements.Limit{
		Limit:  statement.Limit,
		Offset: statement.Offset,
	}
}

// Comment
func (ctx *DefaultQueryBuilder) Insert(statement *orm.Statement) Statement {
	return &statements.Insert{
		Table:        statement.Table,
		InsertValues: statement.Values,
	}
}

// Comment
func (ctx *QueryBuilder) build(statements []Statement) (string, error) {
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
	query, err := ctx.build([]Statement{
		ctx.Builder.Select(ctx.Statement),
		ctx.Builder.Join(ctx.Statement),
		ctx.Builder.Where(ctx.Statement),
		ctx.Builder.OrderBy(ctx.Statement),
		ctx.Builder.Limit(ctx.Statement),
	})

	if err != nil {
		return "", nil, err
	}

	return query, ctx.Values, nil
}

// Comment
func (ctx *QueryBuilder) Count() (string, QueryValues, error) {
	query, err := ctx.build([]Statement{
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
	query, err := ctx.build([]Statement{ctx.Builder.Insert(ctx.Statement)})

	if err != nil {
		return "", nil, err
	}

	return query, ctx.Values, nil
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
