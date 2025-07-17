package sql

import (
	"github.com/lucas11776-golang/orm"
	"github.com/lucas11776-golang/orm/types"
)

type QueryValues []interface{}

type SQLBuilder struct {
	Statement    *orm.Statement
	QueryBuilder QueryBuilder
	Values       QueryValues
}

type SQL struct {
	Builder *SQLBuilder
}

// Comment
func (ctx *SQL) Query(statement *Statement) (types.Results, error) {
	return nil, nil
}

// Comment
func (ctx *SQL) Count(statement *Statement) (int64, error) {
	return 0, nil
}

// Comment
func (ctx *SQL) Insert(statement *Statement) (types.Result, error) {
	return nil, nil
}

// Comment
func (ctx *SQL) Update(statement *Statement) error {
	return nil
}

// Comment
func (ctx *SQL) Delete(Statement *Statement) error {
	return nil
}

type Statement interface {
	Statement() (string, error)
	Values() []interface{}
}
