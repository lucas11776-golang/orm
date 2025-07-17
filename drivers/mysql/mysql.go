package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lucas11776-golang/orm"
	"github.com/lucas11776-golang/orm/types"
)

type Credentials struct {
	Host     string
	User     string
	Password string
	Database string
	Port     int16
	SSL      bool
	Protocol string
}

type MySQL struct {
	DB *sql.DB
}

// Comment
func (ctx *MySQL) Query(statement *orm.Statement) (types.Results, error) {
	return nil, nil
}

// Comment
func (ctx *MySQL) Count(statement *orm.Statement) (int64, error) {
	return 0, nil
}

// Comment
func (ctx *MySQL) Insert(statement *orm.Statement) (types.Result, error) {
	return nil, nil
}

// Comment
func (ctx *MySQL) Update(statement *orm.Statement) error {
	return nil
}

// Comment
func (ctx *MySQL) Delete(Statement *orm.Statement) error {
	return nil
}

// Comment
func (ctx *MySQL) Database() interface{} {
	return nil
}

// Comment
func (ctx *MySQL) Migration() orm.Migration {
	return nil
}

// Comment
func (ctx *MySQL) Close() error {
	return ctx.DB.Close()
}

// Comment
func Connect(credentials *Credentials) orm.Database {
	db, err := sql.Open("mysql", GenerateDataSourceName(credentials))

	if err != nil {
		panic(err)
	}

	return &MySQL{DB: db}
}
