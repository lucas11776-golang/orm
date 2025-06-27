package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/lucas11776-golang/orm/databases/sqlite/migrations"

	_ "github.com/tursodatabase/go-libsql"

	"github.com/lucas11776-golang/orm"
)

const Timeout time.Duration = 10 * time.Second

type SQLite struct {
	DB *sql.DB
}

// Comment
func (ctx *SQLite) scan(rows *sql.Rows) (orm.Results, error) {
	results := orm.Results{}

	cols, err := rows.Columns()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		v := make([]any, len(cols))
		maps := make([]interface{}, len(v))
		vMap := map[string]interface{}{}

		for i := 0; i < len(maps); i++ {
			maps[i] = &v[i]
		}

		rows.Scan(maps...)

		for i, v := range v {
			vMap[cols[i]] = v
		}

		results = append(results, vMap)
	}

	return results, nil
}

// Comment
func (ctx *SQLite) query(sql string, values QueryValues) (orm.Results, error) {
	stmt, err := ctx.DB.Prepare(sql)

	if err != nil {
		return nil, err
	}

	ctxt, cancel := context.WithTimeout(context.Background(), Timeout)

	defer cancel()

	rows, err := stmt.QueryContext(ctxt, values...)

	if err != nil {
		return nil, err
	}

	return ctx.scan(rows)
}

// Comment
func (ctx *SQLite) Query(statement *orm.Statement) (orm.Results, error) {
	builder := &QueryBuilder{Statement: statement}

	sql, values, err := builder.Query()

	if err != nil {
		return nil, err
	}

	return ctx.query(sql, values)
}

// Comment
func (ctx *SQLite) Count(statement *orm.Statement) (int64, error) {
	builder := &QueryBuilder{Statement: statement}

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
func (ctx *SQLite) Insert(statement *orm.Statement) (orm.Result, error) {
	builder := &QueryBuilder{Statement: statement}

	sql, values, err := builder.Insert()

	if err != nil {
		return nil, err
	}

	stmt, err := ctx.DB.Prepare(sql)

	if err != nil {
		return nil, err
	}

	ctxt, cancel := context.WithTimeout(context.Background(), Timeout)

	defer cancel()

	exec, err := stmt.ExecContext(ctxt, values...)

	if err != nil {
		return nil, err
	}

	id, err := exec.LastInsertId()

	if err != nil {
		return nil, err
	}

	builder = &QueryBuilder{Statement: &orm.Statement{
		Table: statement.Table,
		Where: []interface{}{&orm.Where{
			Key:      statement.PrimaryKey,
			Operator: "=",
			Value:    id,
		}},
	}}

	sql, values, err = builder.Query()

	if err != nil {
		return nil, err
	}

	results, err := ctx.query(sql, values)

	if err != nil {
		return nil, err
	}

	if len(results) != 1 {
		return nil, fmt.Errorf("failed to get insert result")
	}

	return results[0], nil
}

// Comment
func (ctx *SQLite) Update(statement *orm.Statement) error {
	builder := &QueryBuilder{Statement: statement}

	sql, values, err := builder.Update()

	if err != nil {
		return err
	}

	ctxt, cancel := context.WithTimeout(context.Background(), Timeout)

	defer cancel()

	_, err = ctx.DB.ExecContext(ctxt, sql, values...)

	if err != nil {
		return err
	}

	return nil
}

// Comment
func (ctx *SQLite) Delete(statement *orm.Statement) error {
	builder := &QueryBuilder{Statement: statement}

	sql, values, err := builder.Delete()

	if err != nil {
		return err
	}

	ctxt, cancel := context.WithTimeout(context.Background(), Timeout)

	defer cancel()

	_, err = ctx.DB.ExecContext(ctxt, sql, values...)

	if err != nil {
		return err
	}

	return nil
}

// Comment
func (ctx *SQLite) Database() interface{} {
	return ctx.DB
}

// Comment
func (ctx *SQLite) Migration() orm.Migration {
	return &migrations.Migration{DB: ctx.DB}
}

// Comment
func Connect(source string) orm.Database {
	db, err := sql.Open("libsql", source)

	if err != nil {
		panic(err)
	}

	return &SQLite{DB: db}
}
