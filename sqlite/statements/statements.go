package statements

import "orm"

const SPACE = " "

type WhereGroupQueryBuilder struct {
	Group []interface{}
}

// Comment
func (ctx *WhereGroupQueryBuilder) Where(w orm.Where) orm.WhereGroupBuilder {
	return ctx
}

// Comment
func (ctx *WhereGroupQueryBuilder) AndWhere(w orm.Where) orm.WhereGroupBuilder {
	return ctx
}

// Comment
func (ctx *WhereGroupQueryBuilder) OrWhere(w orm.Where) orm.WhereGroupBuilder {
	return ctx
}
