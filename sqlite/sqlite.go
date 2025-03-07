package sqlite

import (
	"database/sql"
	"orm/sqlite/migrations"

	_ "github.com/mattn/go-sqlite3"

	"orm"
)

type DB *sql.DB

type SQLite struct {
	DB *sql.DB
}

// Comment
func ScanRows(rows *sql.Rows) (orm.Results, error) {
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
func (ctx *SQLite) Query(statement *orm.Statement) (orm.Results, error) {
	builder := &QueryBuilder{Statement: statement}

	query, values, err := builder.Query()

	if err != nil {
		return nil, err
	}

	stmt, err := ctx.DB.Prepare(query)

	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(values...)

	if err != nil {
		return nil, err
	}

	return ScanRows(rows)
}

// Comment
func (ctx *SQLite) Count(statement *orm.Statement) (int64, error) {
	return 0, nil
}

// Comment
func (ctx *SQLite) Insert(statement *orm.Statement) (orm.Result, error) {
	return nil, nil
}

// Comment
func (ctx *SQLite) Update(values orm.Values) error {
	return nil
}

// Comment
func (ctx *SQLite) Database() interface{} {
	return ctx.DB
}

// Comment
func (ctx *SQLite) Migration() orm.Migration {
	return &migrations.Migration{
		DB: ctx.DB,
	}
}

// Comment
func Connect(source string) (orm.Database, error) {
	db, err := sql.Open("sqlite3", source)

	if err != nil {
		return nil, err
	}

	return &SQLite{
		DB: db,
	}, nil
}
