package sqlite

import (
	"orm/sqlite/statements"
)

type QueryValues []interface{}

type QueryBuilder struct {
	Query  *Query
	Values QueryValues
}

// Comment
func (ctx *QueryBuilder) SelectStatement() (string, error) {
	statement := &statements.Select{
		Select: ctx.Query._select,
	}

	return statement.Statement()
}

// Comment
func (ctx *QueryBuilder) WhereStatement() (string, error) {
	statement := &statements.Where{
		Where: ctx.Query._where,
	}

	sql, err := statement.Statement()

	if err == nil {
		ctx.Values = append(ctx.Values, statement.Values...)
	}

	return sql, err
}

// Comment
func (ctx *QueryBuilder) JoinStatement() (string, error) {
	return "", nil
}

// Comment
func (ctx *QueryBuilder) Limit() string {
	return ""
}
