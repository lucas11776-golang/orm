package sqlite

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/lucas11776-golang/orm/drivers/sql/statements"
	"github.com/lucas11776-golang/orm/drivers/sqlite/migrations"
	"github.com/lucas11776-golang/orm/types"

	_ "github.com/tursodatabase/go-libsql"

	"github.com/lucas11776-golang/orm"
	utils "github.com/lucas11776-golang/orm/utils/sql"
)

type SQLite struct {
	DB *sql.DB
}

type TableInfo struct {
	CID          int    `column:"cid"`
	Name         string `column:"name"`
	Type         string `column:"type"`
	NotNull      bool   `column:"notnull"`
	DefaultValue string `column:"dflt_value"`
	PrimaryKey   bool   `column:"pk"`
}

// Comment
func (ctx *SQLite) query(query string, values QueryValues) (types.Results, error) {

	stmt, err := ctx.DB.Prepare(query)

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
func (ctx *SQLite) Query(statement *orm.Statement) (types.Results, error) {
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
func (ctx *SQLite) getPrimaryKey(table string) (string, error) {
	rows, err := ctx.DB.Query(fmt.Sprintf("PRAGMA table_info(%s);", statements.SafeKey(table)))

	if err != nil {
		return "", err
	}

	columns, err := utils.ScanRowsToModels(rows, TableInfo{})

	if err != nil {
		return "", err
	}

	for _, column := range columns {
		if column.PrimaryKey {
			return column.Name, nil
		}
	}

	return "", nil

}

// Comment
func (ctx *SQLite) Insert(statement *orm.Statement) (types.Result, error) {
	builder := &QueryBuilder{Statement: statement}

	sql, values, err := builder.Insert()

	if err != nil {
		return nil, err
	}

	stmt, err := ctx.DB.Prepare(sql)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	insertResult, err := stmt.Exec(values...)

	if err != nil {
		return nil, err
	}

	key, err := ctx.getPrimaryKey(statement.Table)

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

	builder = &QueryBuilder{Statement: &orm.Statement{
		Table: statement.Table,
		Where: []interface{}{&orm.Where{
			Key:      key,
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
		return nil, fmt.Errorf("failed to get insert row")
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

	_, err = ctx.DB.Exec(sql, values...)

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

	_, err = ctx.DB.Exec(sql, values...)

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
func (ctx *SQLite) Close() error {
	return ctx.DB.Close()
}

// Comment
func Connect(source string) orm.Database {
	db, err := sql.Open("libsql", source)

	if err != nil {
		panic(err)
	}

	var mode string
	if err := db.QueryRow("PRAGMA journal_mode = WAL").Scan(&mode); err != nil {
		panic(fmt.Sprintf("failed to set journal_mode: %v", err))
	}

	var timeout int64
	if err := db.QueryRow("PRAGMA busy_timeout = 10000").Scan(&timeout); err != nil {
		panic(fmt.Sprintf("failed to set busy_timeout: %v", err))
	}

	return &SQLite{DB: db}
}
