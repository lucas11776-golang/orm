package mysql

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lucas11776-golang/orm"
	"github.com/lucas11776-golang/orm/types"
)

const (
	TestingDatabaseName     string = "orm_golang_testing"
	TestingDatabasePassword string = ""
	TestingDatabasePort     int    = 3306
	TestingDatabaseProtocol string = "tcp"
)

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

type Credentials struct {
	Host     string
	User     string
	Password string
	Database string
	Port     int16
	SSL      bool
}

// Comment
func Connect(credentials *Credentials) *sql.DB {
	// db, err := sql.Open("mysql", "user:password@/dbname")

	db, err := sql.Open("mysql", "mysql://root:@localhost:3306/orm_golang_testing")

	if err != nil {
		panic(err)
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
