package sqlite

import (
	"orm"
)

type QueryValues []interface{}

type QueryBuilder struct {
	Statement *orm.Statement
	Values    QueryValues
}

// SELECT
// JOIN
// WHERE
// LIMIT

// Comment
func (ctx *QueryBuilder) Query() (orm.Results, error) {
	return nil, nil
}

// Comment
func (ctx *QueryBuilder) Count() (int64, error) {
	return 0, nil
}

// Comment
func (ctx *QueryBuilder) Insert() (orm.Result, error) {
	return nil, nil
}

// Comment
func (ctx *QueryBuilder) Update() error {
	return nil
}
