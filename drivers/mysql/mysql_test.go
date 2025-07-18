package mysql

import (
	"database/sql"
	"strings"
	"testing"

	mysqlTesting "github.com/lucas11776-golang/orm/drivers/mysql/testing"
)

func TestMySQL(t *testing.T) {
	t.Run("TestTablePrimaryKey", func(t *testing.T) {

		return

		db, err := sql.Open("mysql", mysqlTesting.TestingDataSourceName())

		if err != nil {
			t.Fatal(err)
		}

		mysql := &MySQL{db: db}

		_, err = db.Exec(strings.Join([]string{
			"CREATE TABLE `users`(",
			"  `full_name` VARCHAR(255) NOT NULL,",
			"  `email` VARCHAR(255) NOT NULL,",
			"  `user_id` INTEGER PRIMARY KEY AUTO_INCREMENT UNIQUE",
			");",
		}, "\r\n"))

		if err != nil {
			t.Fatal(err)
		}

		if err != nil {
			t.Fatal(err)
		}

		primaryKey, err := mysql.TablePrimaryKey("users")

		if err != nil {
			t.Fatal(err)
		}

		if primaryKey != "user_id" {
			t.Fatalf("Expected users table primary key to be (%s) but got (%s)", "user_id", primaryKey)
		}
	})
}
