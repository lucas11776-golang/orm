package mysql

import (
	"database/sql"
	"fmt"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lucas11776-golang/orm"
	"github.com/lucas11776-golang/orm/types"
	"github.com/lucas11776-golang/orm/utils/env"
	"github.com/spf13/cast"
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

// Comment
func GetTestingCredentials() *Credentials {
	return &Credentials{
		Host:     env.Get("DB_MYSQL_HOST", "localhost"),
		User:     env.Get("DB_MYSQL_USER", "root"),
		Password: env.Get("DB_MYSQL_PASS", "password"),
		Database: env.Get("DB_MYSQL_DATABASE", "orm_golang_testing"),
		Port:     cast.ToInt16(env.Get("DB_MYSQL_DATABASE_PORT", "3306")),
		SSL:      cast.ToBool(env.Get("DB_MYSQL_DATABASE_SSL", "false")),
		Protocol: env.Get("DB_MYSQL_DATABASE_PROTOCOL", "tcp"),
	}
}

// Comment
func GetTestingDataSourceName() string {
	return GenerateDataSourceName(GetTestingCredentials())
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
func GenerateDataSourceName(cred *Credentials) string {
	url := url.Values{}

	url.Add("parseTime", "true")

	if !cred.SSL {
		url.Add("tls", "skip-verify")
	}

	return fmt.Sprintf("%s:%s@%s(%s:%d)/%s?%s", cred.User, cred.Password, cred.Protocol, cred.Host, cred.Port, cred.Database, url.Encode())
}

// Comment
func Connect(credentials *Credentials) *sql.DB {
	db, err := sql.Open("mysql", GenerateDataSourceName(credentials))

	if err != nil {
		panic(err)
	}

	return db
}
