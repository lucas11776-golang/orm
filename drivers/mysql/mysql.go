package mysql

import (
	"database/sql"
	"fmt"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lucas11776-golang/orm"
	"github.com/lucas11776-golang/orm/drivers/mysql/migrations"
	driver "github.com/lucas11776-golang/orm/drivers/sql"
	// "github.com/lucas11776-golang/orm/drivers/mysql/migrations"
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

type TableDescription struct {
	Field   string      `column:"Field"`
	Type    string      `column:"Type"`
	Null    string      `column:"Null"`
	Key     string      `column:"Key"` // PRIM
	Default interface{} `column:"Default"`
	Extra   string      `column:"Extra"`
}

type MySQL struct {
	db *sql.DB
}

// Comment
func (ctx *MySQL) DB() *sql.DB {
	return ctx.db
}

// Comment
func (ctx *MySQL) TablePrimaryKey(table string) (key string, err error) {

	return "", nil
}

// Comment
func Connect(credentials *Credentials) orm.Database {
	db, err := sql.Open("mysql", GenerateDataSourceName(credentials))

	if err != nil {
		panic(err)
	}

	return ConnectDB(db)
}

// Comment
func ConnectDB(db *sql.DB) orm.Database {
	return driver.NewSQLDriver(&driver.DriverOptions{
		QueryBuilder: &driver.DefaultQueryBuilder{},
		Migration:    &migrations.Migration{DB: db},
		Database:     &MySQL{db: db},
	})
}

// Comment
func GenerateDataSourceName(credentials *Credentials) string {
	url := url.Values{"parseTime": []string{"true"}}

	if !credentials.SSL {
		url.Add("tls", "skip-verify")
	}

	return fmt.Sprintf(
		"%s:%s@%s(%s:%d)/%s?%s",
		credentials.User,
		credentials.Password,
		credentials.Protocol,
		credentials.Host,
		credentials.Port,
		credentials.Database,
		url.Encode(),
	)
}
