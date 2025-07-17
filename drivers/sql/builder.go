package sql

import (
	"github.com/lucas11776-golang/orm"
	"github.com/lucas11776-golang/orm/drivers/sql/statements"
)

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
func (ctx *DefaultQueryBuilder) Update(statement *orm.Statement) Statement {
	return &statements.Update{
		Table:        statement.Table,
		Where:        statement.Where,
		UpdateValues: statement.Values,
	}
}

// Comment
func (ctx *DefaultQueryBuilder) Delete(statement *orm.Statement) Statement {
	return &statements.Delete{
		Table: statement.Table,
		Where: statement.Where,
	}
}
