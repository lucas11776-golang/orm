package sql

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/lucas11776-golang/orm"
	"github.com/lucas11776-golang/orm/types"
	utils "github.com/lucas11776-golang/orm/utils/sql"
)

type QueryValues []interface{}

type QueryBuilder interface {
	Select(statement *orm.Statement) Statement
	Join(statement *orm.Statement) Statement
	Where(statement *orm.Statement) Statement
	OrderBy(statement *orm.Statement) Statement
	Limit(statement *orm.Statement) Statement
	Insert(statement *orm.Statement) Statement
	Update(statement *orm.Statement) Statement
	Delete(statement *orm.Statement) Statement
}

type Statement interface {
	Statement() (string, error)
	Values() []interface{}
}

type Database interface {
	DB() *sql.DB
	TablePrimaryKey(table string) (key string, err error)
}

type SQL struct {
	builder         QueryBuilder
	db              *sql.DB
	tablePrimaryKey func(table string) (key string, err error)
	migration       orm.Migration
}

type DriverOptions struct {
	Database     Database
	QueryBuilder QueryBuilder
	Migration    orm.Migration
}

// Comment
func NewSQLDriver(options *DriverOptions) *SQL {
	return &SQL{
		tablePrimaryKey: options.Database.TablePrimaryKey,
		builder:         options.QueryBuilder,
		db:              options.Database.DB(),
		migration:       options.Migration,
	}
}

// Comment
func (ctx *SQL) query(query string, values QueryValues) (types.Results, error) {
	stmt, err := ctx.db.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(values...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return utils.ScanRowsToResults(rows)
}

// Comment
func (ctx *SQL) Database() interface{} {
	return ctx.db
}

// Comment
func (ctx *SQL) Query(statement *orm.Statement) (types.Results, error) {
	builder := &SQLBuilder{
		QueryBuilder: &DefaultQueryBuilder{},
		Statement:    statement,
	}

	sql, values, err := builder.Query()

	if err != nil {
		return nil, err
	}

	return ctx.query(sql, values)
}

// Comment
func (ctx *SQL) Count(statement *orm.Statement) (int64, error) {
	builder := &SQLBuilder{
		QueryBuilder: &DefaultQueryBuilder{},
		Statement:    statement,
	}

	sql, values, err := builder.Count()

	if err != nil {
		return 0, err
	}

	results, err := ctx.query(sql, values)

	if err != nil {
		return 0, nil
	}

	if len(results) != 1 {
		return 0, errors.New("failed to execute count")
	}

	total, ok := results[0]["total"]

	if !ok {
		return 0, errors.New("expected count result map to have total key")
	}

	return total.(int64), nil
}

// Comment
func (ctx *SQL) Insert(statement *orm.Statement) (types.Result, error) {
	builder := &SQLBuilder{
		QueryBuilder: &DefaultQueryBuilder{},
		Statement:    statement,
	}

	sql, values, err := builder.Insert()

	if err != nil {
		return nil, err
	}

	stmt, err := ctx.db.Prepare(sql)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	insertResult, err := stmt.Exec(values...)

	if err != nil {
		return nil, err
	}

	// TODO: Cache primary key in map for performance...
	key, err := ctx.tablePrimaryKey(statement.Table)

	if err != nil {
		return nil, err
	}

	if key == "" {
		return types.Result(statement.Values), nil
	}

	id, err := insertResult.LastInsertId()

	if err != nil {
		return nil, err
	}

	builder = &SQLBuilder{
		QueryBuilder: &DefaultQueryBuilder{},
		Statement: &orm.Statement{
			Table: statement.Table,
			Where: []interface{}{&orm.Where{
				Key:      key,
				Operator: "=",
				Value:    id,
			}},
		},
	}

	sql, values, err = builder.Query()

	if err != nil {
		return nil, err
	}

	results, err := ctx.query(sql, values)

	if err != nil {
		return nil, err
	}

	if len(results) != 1 {
		return nil, fmt.Errorf("failed to get insert row")
	}

	return results[0], nil
}

// Comment
func (ctx *SQL) Update(statement *orm.Statement) error {
	builder := &SQLBuilder{
		QueryBuilder: &DefaultQueryBuilder{},
		Statement:    statement,
	}

	sql, values, err := builder.Update()

	if err != nil {
		return err
	}

	_, err = ctx.db.Exec(sql, values...)

	if err != nil {
		return err
	}

	return nil
}

// Comment
func (ctx *SQL) Delete(statement *orm.Statement) error {
	builder := &SQLBuilder{
		QueryBuilder: &DefaultQueryBuilder{},
		Statement:    statement,
	}

	sql, values, err := builder.Delete()

	if err != nil {
		return err
	}

	_, err = ctx.db.Exec(sql, values...)

	if err != nil {
		return err
	}

	return nil
}

// Comment
func (ctx *SQL) Migration() orm.Migration {
	return ctx.migration
}

// Comment
func (ctx *SQL) Close() error {
	return ctx.db.Close()
}
