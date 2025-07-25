package migrations

import (
	"database/sql"
	"fmt"
	"strings"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lucas11776-golang/orm"
	mysqlTesting "github.com/lucas11776-golang/orm/drivers/mysql/testing"
	"github.com/lucas11776-golang/orm/drivers/sql/statements"
	"github.com/lucas11776-golang/orm/migrations"
)

func TestMigrationStatementColumnBuilder(t *testing.T) {
	t.Run("TestColumnTypes", func(t *testing.T) {
		t.Run("TestIncrement", func(t *testing.T) {
			expected := fmt.Sprintf("%s BIGINT(20) UNSIGNED PRIMARY KEY AUTO_INCREMENT UNIQUE", statements.SafeKey("id"))

			if actual, _ := generateColumnStatement((&migrations.Table{}).Increment("id")); expected != actual {
				t.Fatalf("expected column statement to be (%s) but got (%s)", expected, actual)
			}
		})

		t.Run("TestTimestamp", func(t *testing.T) {
			expected := fmt.Sprintf("%s TIMESTAMP NOT NULL", statements.SafeKey("created_at"))

			if actual, _ := generateColumnStatement((&migrations.Table{}).TimeStamp("created_at")); expected != actual {
				t.Fatalf("expected column statement to be (%s) but got (%s)", expected, actual)
			}
		})

		t.Run("TestDatetime", func(t *testing.T) {
			expected := fmt.Sprintf("%s DATETIME NOT NULL", statements.SafeKey("birth_of_birth"))

			if actual, _ := generateColumnStatement((&migrations.Table{}).Datetime("birth_of_birth")); expected != actual {
				t.Fatalf("expected column statement to be (%s) but got (%s)", expected, actual)
			}
		})

		t.Run("TestDate", func(t *testing.T) {
			expected := fmt.Sprintf("%s DATE NOT NULL", statements.SafeKey("expires"))

			if actual, _ := generateColumnStatement((&migrations.Table{}).Date("expires")); expected != actual {
				t.Fatalf("expected column statement to be (%s) but got (%s)", expected, actual)
			}
		})

		t.Run("TestInteger", func(t *testing.T) {
			expected := fmt.Sprintf("%s INTEGER NOT NULL", statements.SafeKey("units"))

			if actual, _ := generateColumnStatement((&migrations.Table{}).Integer("units")); expected != actual {
				t.Fatalf("expected column statement to be (%s) but got (%s)", expected, actual)
			}
		})

		t.Run("TestBigInteger", func(t *testing.T) {
			expected := fmt.Sprintf("%s BIGINT(20) UNSIGNED NOT NULL", statements.SafeKey("user_id"))

			if actual, _ := generateColumnStatement((&migrations.Table{}).BigInteger("user_id")); expected != actual {
				t.Fatalf("expected column statement to be (%s) but got (%s)", expected, actual)
			}
		})

		t.Run("TestDouble", func(t *testing.T) {
			expected := fmt.Sprintf("%s DOUBLE NOT NULL", statements.SafeKey("distance"))

			if actual, _ := generateColumnStatement((&migrations.Table{}).Double("distance")); expected != actual {
				t.Fatalf("expected column statement to be (%s) but got (%s)", expected, actual)
			}
		})

		t.Run("TestFloat", func(t *testing.T) {
			expected := fmt.Sprintf("%s FLOAT NOT NULL", statements.SafeKey("amount"))

			if actual, _ := generateColumnStatement((&migrations.Table{}).Float("amount")); expected != actual {
				t.Fatalf("expected column statement to be (%s) but got (%s)", expected, actual)
			}
		})

		t.Run("TestString", func(t *testing.T) {
			expected := fmt.Sprintf("%s VARCHAR(16380) NOT NULL", statements.SafeKey("email"))

			if actual, _ := generateColumnStatement((&migrations.Table{}).String("email")); expected != actual {
				t.Fatalf("expected column statement to be (%s) but got (%s)", expected, actual)
			}
		})

		t.Run("TestText", func(t *testing.T) {
			expected := fmt.Sprintf("%s TEXT NOT NULL", statements.SafeKey("content"))

			if actual, _ := generateColumnStatement((&migrations.Table{}).Text("content")); expected != actual {
				t.Fatalf("expected column statement to be (%s) but got (%s)", expected, actual)
			}
		})

		t.Run("TestBoolean", func(t *testing.T) {
			expected := fmt.Sprintf("%s BOOLEAN NOT NULL", statements.SafeKey("active"))

			if actual, _ := generateColumnStatement((&migrations.Table{}).Boolean("active")); expected != actual {
				t.Fatalf("expected column statement to be (%s) but got (%s)", expected, actual)
			}
		})

		t.Run("TestBinary", func(t *testing.T) {
			expected := fmt.Sprintf("%s BLOB NOT NULL", statements.SafeKey("document"))

			if actual, _ := generateColumnStatement((&migrations.Table{}).Binary("document")); expected != actual {
				t.Fatalf("expected column statement to be (%s) but got (%s)", expected, actual)
			}
		})
	})

	t.Run("TestColumnOptions", func(t *testing.T) {
		t.Run("TestIncrement", func(t *testing.T) {
			expected := fmt.Sprintf("%s BIGINT(20) UNSIGNED PRIMARY KEY AUTO_INCREMENT UNIQUE", statements.SafeKey("id"))

			if actual, _ := generateColumnStatement((&migrations.Table{}).Increment("id")); expected != actual {
				t.Fatalf("expected column statement to be (%s) but got (%s)", expected, actual)
			}
		})

		t.Run("TestNullable", func(t *testing.T) {
			expected := fmt.Sprintf("%s DATETIME", statements.SafeKey("notification_time"))
			actual, _ := generateColumnStatement((&migrations.Table{}).Datetime("notification_time").Nullable())

			if expected != actual {
				t.Fatalf("expected column statement to be (%s) but got (%s)", expected, actual)
			}
		})

		t.Run("TestPrimaryKey", func(t *testing.T) {
			expected := fmt.Sprintf("%s VARCHAR(16380) PRIMARY KEY NOT NULL", statements.SafeKey("uuid"))
			actual, _ := generateColumnStatement((&migrations.Table{}).String("uuid").PrimaryKey())

			if expected != actual {
				t.Fatalf("expected column statement to be (%s) but got (%s)", expected, actual)
			}
		})

		t.Run("TestUnique", func(t *testing.T) {
			expected := fmt.Sprintf("%s VARCHAR(16380) NOT NULL UNIQUE", statements.SafeKey("email"))
			actual, _ := generateColumnStatement((&migrations.Table{}).String("email").Unique())

			if expected != actual {
				t.Fatalf("expected column statement to be (%s) but got (%s)", expected, actual)
			}
		})

		t.Run("TestDefault", func(t *testing.T) {
			t.Run("TestDefaultString", func(t *testing.T) {
				expected := fmt.Sprintf("%s VARCHAR(16380) NOT NULL DEFAULT 'jeo@doe.com'", statements.SafeKey("email"))
				actual, _ := generateColumnStatement((&migrations.Table{}).String("email").Default("jeo@doe.com"))

				if expected != actual {
					t.Fatalf("expected column statement to be (%s) but got (%s)", expected, actual)
				}
			})

			t.Run("TestDefaultSpecialDefaultWork", func(t *testing.T) {
				expected := fmt.Sprintf("%s TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP", statements.SafeKey("created_at"))
				actual, _ := generateColumnStatement((&migrations.Table{}).TimeStamp("created_at").Current())

				if expected != actual {
					t.Fatalf("expected column statement to be (%s) but got (%s)", expected, actual)
				}
			})
		})

		t.Run("TestOptionOrder", func(t *testing.T) {
			// TODO: check options order
		})
	})
}

func TestRunMigration(t *testing.T) {
	t.Run("TestMigrateQueryGenerator", func(t *testing.T) {
		db, err := sql.Open("mysql", mysqlTesting.TestingDataSourceName())

		if err != nil {
			t.Fatalf("Something went wrong when trying to connect to database: %v", err)
		}

		migration := &Migration{DB: db}

		queryExpected := strings.Join([]string{
			"CREATE TABLE IF NOT EXISTS products (",
			strings.Join([]string{
				statements.SPACE + "`id` BIGINT(20) UNSIGNED PRIMARY KEY AUTO_INCREMENT UNIQUE",
				statements.SPACE + "`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP",
				statements.SPACE + "`name` VARCHAR(16380) NOT NULL",
				statements.SPACE + "`price` FLOAT NOT NULL",
				statements.SPACE + "`in_stock` INTEGER DEFAULT 0",
			}, ",\r\n"),
			");",
		}, "\r\n")

		table := migrations.Table{}

		table.Increment("id")
		table.TimeStamp("created_at").Current()
		table.String("name")
		table.Float("price")
		table.Integer("in_stock").Nullable().Default(0)

		queryActual, err := migration.generateTableSchemeSQL(&orm.TableScheme{
			Name:    "products",
			Columns: table.Columns,
		})

		if err != nil {
			t.Fatalf("Something went wrong when trying to generate create model table: %v", err)
		}

		if queryExpected != queryActual {
			t.Fatalf("Expected model table query to be (%s) but got (%s)", queryExpected, queryActual)
		}

		migration.DB.Close()
	})

	t.Run("TestMigrateAndDropTable", func(t *testing.T) {
		db, err := sql.Open("mysql", mysqlTesting.TestingDataSourceName())

		if err != nil {
			t.Fatalf("Something went wrong when trying to connect to database: %v", err)
		}

		migration := &Migration{DB: db}

		type User struct {
			Connection string    `json:"-" connection:"sqlite"`
			ID         int64     `json:"id"`
			CreatedAt  time.Time `json:"created_at"`
			Email      string    `json:"email"`
		}

		table := migrations.Table{}

		table.Increment("id")
		table.TimeStamp("created_at").Current()
		table.String("email")

		err = migration.Migrate(&orm.TableScheme{
			Name:    "users",
			Columns: table.Columns,
		})

		if err != nil {
			t.Fatal(err)
		}

		email := "jeo@doe.com"

		result, err := db.Exec("INSERT INTO `users`(`email`) VALUES(?)", email)

		if err != nil {
			t.Fatal(err)
		}

		id, err := result.LastInsertId()

		if err != nil {
			t.Fatal(err)
		}

		row := db.QueryRow("SELECT `id`, `created_at`, `email` FROM `users` WHERE `id` = ?", id)

		if err := row.Err(); err != nil {
			t.Fatal(err)
		}

		user := &User{}

		if err = row.Scan(&user.ID, &user.CreatedAt, &user.Email); err != nil {
			t.Fatal(err)
		}

		if user.ID != 1 {
			t.Fatalf("Expected id to be (%d) but got (%d)", 1, user.ID)
		}

		if user.Email != email {
			t.Fatalf("Expected email to be (%s) but got (%s)", email, user.Email)
		}

		if err := migration.Drop("users"); err != nil {
			t.Fatal(err)
		}

		if _, err := db.Exec("INSERT INTO `users`(`email`) VALUES(?)", email); err == nil {
			t.Fatal("Expected insert user to return error because table has been dropped.")
		}

		migration.DB.Close()
	})

}
