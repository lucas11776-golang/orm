package mysql

import (
	"database/sql"
	"fmt"
	"net/url"

	"github.com/lucas11776-golang/orm/utils/env"
	"github.com/spf13/cast"
)

// Comment
func TestingDataSourceName() string {
	credentials := map[string]interface{}{
		"Host":     env.Get("DB_MYSQL_HOST", "localhost"),
		"User":     env.Get("DB_MYSQL_USER", "root"),
		"Password": env.Get("DB_MYSQL_PASS", "password"),
		"Database": env.Get("DB_MYSQL_DATABASE", "orm_golang_testing"),
		"Port":     cast.ToInt16(env.Get("DB_MYSQL_DATABASE_PORT", "3306")),
		"SSL":      cast.ToBool(env.Get("DB_MYSQL_DATABASE_SSL", "false")),
		"Protocol": env.Get("DB_MYSQL_DATABASE_PROTOCOL", "tcp"),
	}

	url := url.Values{
		"parseTime": []string{"true"},
		"charset":   []string{"utf8mb4"},
	}

	if !credentials["SSL"].(bool) {
		url.Add("tls", "skip-verify")
	}

	return fmt.Sprintf(
		"%s:%s@%s(%s:%d)/%s?%s",
		credentials["User"],
		credentials["Password"],
		credentials["Protocol"],
		credentials["Host"],
		credentials["Port"],
		credentials["Database"],
		url.Encode(),
	)
}

// Comment
func ConnectTestingDB() *sql.DB {
	db, err := sql.Open("mysql", TestingDataSourceName())

	if err != nil {
		panic(err)
	}

	return db
}
