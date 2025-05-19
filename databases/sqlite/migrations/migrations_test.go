package migrations

import (
	"database/sql"
	"strings"
	"testing"
	"time"

	"github.com/lucas11776-golang/orm"
	"github.com/lucas11776-golang/orm/databases/sqlite/statements"

	_ "github.com/mattn/go-sqlite3"
)

func TestMigrationStatement(t *testing.T) {
	migration := &Migration{}

	t.Run("TestTypesQuery", func(t *testing.T) {
		primaryKeyExpected := "`id` INTEGER PRIMARY KEY AUTOINCREMENT"
		primaryKeyActual, _ := migration.columnStatement("id", "primary_key")

		if primaryKeyExpected != primaryKeyActual {
			t.Fatalf("Expected primary key statement to be (%s) but got (%s)", primaryKeyExpected, primaryKeyActual)
		}

		datetimeCurrentExpected := "`created_at` DATETIME DEFAULT CURRENT_TIMESTAMP"
		datetimeCurrentActual, _ := migration.columnStatement("created_at", "datetime_current")

		if datetimeCurrentExpected != datetimeCurrentActual {
			t.Fatalf("Expected datetime statement to be (%s) but got (%s)", datetimeCurrentExpected, datetimeCurrentActual)
		}

		datetimeExpected := "`updated_at` DATETIME"
		datetimeActual, _ := migration.columnStatement("updated_at", "datetime")

		if datetimeExpected != datetimeActual {
			t.Fatalf("Expected datetime statement to be (%s) but got (%s)", datetimeExpected, datetimeActual)
		}

		integerExpected := "`year` INTEGER"
		integerActual, _ := migration.columnStatement("year", "integer")

		if integerExpected != integerActual {
			t.Fatalf("Expected integer statement to be (%s) but got (%s)", integerExpected, integerActual)
		}

		floatExpected := "`height` FLOAT"
		floatActual, _ := migration.columnStatement("height", "float")

		if floatExpected != floatActual {
			t.Fatalf("Expected float statement to be (%s) but got (%s)", floatExpected, floatActual)
		}

		stringExpected := "`email` VARCHAR"
		stringActual, _ := migration.columnStatement("email", "string")

		if stringExpected != stringActual {
			t.Fatalf("Expected string statement to be (%s) but got (%s)", stringExpected, stringActual)
		}

		textExpected := "`bio` TEXT"
		textActual, _ := migration.columnStatement("bio", "text")

		if textExpected != textActual {
			t.Fatalf("Expected text statement to be (%s) but got (%s)", textExpected, textActual)
		}

		booleanExpected := "`subscribed` BOOLEAN"
		booleanActual, _ := migration.columnStatement("subscribed", "boolean")

		if booleanExpected != booleanActual {
			t.Fatalf("Expected boolean statement to be (%s) but got (%s)", booleanExpected, booleanActual)
		}
	})

	t.Run("TestArgumentType", func(t *testing.T) {
		defaultArgExpected := "`subscribed` BOOLEAN DEFAULT false"
		defaultArgActual, _ := migration.columnStatement("subscribed", "boolean", "DEFAULT:false")

		if defaultArgExpected != defaultArgActual {
			t.Fatalf("Expected statement with default arg to be (%s) but got (%s)", defaultArgExpected, defaultArgActual)
		}

		notNullArgExpected := "`email` VARCHAR NOT NULL"
		notNullArgActual, _ := migration.columnStatement("email", "string", "not_null")

		if notNullArgExpected != notNullArgActual {
			t.Fatalf("Expected statement with not null to be (%s) but got (%s)", notNullArgExpected, notNullArgActual)
		}
	})
}

func TestRunMigration(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")

	if err != nil {
		t.Fatalf("Something went wrong when trying to connect to database: %v", err)
	}

	migration := &Migration{DB: db}

	t.Run("TestMigrationQuery", func(t *testing.T) {
		type Product struct {
			Id        int64     `column:"id" type:"primary_key"`
			CreatedAt time.Time `column:"created_at" type:"datetime_current"`
			Name      string    `column:"name" type:"string,not_null"`
			Price     float64   `column:"price" type:"float,not_null"`
			InStock   int64     `column:"in_stock" type:"integer,default:0"`
		}

		queryExpected := strings.Join([]string{
			"CREATE TABLE IF NOT EXISTS products (",
			strings.Join([]string{
				statements.SPACE + "`id` INTEGER PRIMARY KEY AUTOINCREMENT",
				statements.SPACE + "`created_at` DATETIME DEFAULT CURRENT_TIMESTAMP",
				statements.SPACE + "`name` VARCHAR NOT NULL",
				statements.SPACE + "`price` FLOAT NOT NULL",
				statements.SPACE + "`in_stock` INTEGER DEFAULT 0",
			}, ",\r\n"),
			");",
		}, "\r\n")
		p := Product{}
		queryActual, err := migration.generateModelTableQuery(p)

		if err != nil {
			t.Fatalf("Something went wrong when trying to generate create model table: %v", err)
		}

		if queryExpected != queryActual {
			t.Fatalf("Expected model table query to be (%s) but got (%s)", queryExpected, queryActual)
		}
	})

	t.Run("TestInsertRecords", func(t *testing.T) {
		type User struct {
			Id    int64  `column:"id" type:"primary_key"`
			Email string `column:"email" type:"string"`
		}

		type Subscription struct {
			Id    int64  `column:"id" type:"primary_key"`
			Email string `column:"email" type:"string"`
		}

		user := User{
			Id:    1,
			Email: "jeo@doe.com",
		}

		subscription := Subscription{
			Id:    1,
			Email: user.Email,
		}

		err := migration.Migrate(orm.Models{User{}, Subscription{}})

		if err != nil {
			t.Fatalf("Something went wrong when trying to migrate table: %v", err)
		}

		_, err = db.Exec(strings.Join([]string{
			"INSERT INTO users(email) VALUES(?);",
			"INSERT INTO subscriptions(email) VALUES(?);",
		}, "\r\n"), user.Email, subscription.Email)

		if err != nil {
			t.Fatalf("Something went wrong when trying to insert data in to tables: %v", err)
		}

		// Users Table
		userRow := db.QueryRow("SELECT * FROM users WHERE id = ?", user.Id)

		userRecord := User{}

		err = userRow.Scan(&userRecord.Id, &userRecord.Email)

		if err != nil {
			t.Fatalf("Something went wrong when trying to get user: %v", err)
		}

		if user.Id != userRecord.Id {
			t.Fatalf("Expected user id to be (%d) but got (%d)", user.Id, userRecord.Id)
		}

		if user.Email != userRecord.Email {
			t.Fatalf("Expected user email to be (%s) but got (%s)", user.Email, userRecord.Email)
		}

		// Subscription Table
		subscriptionRow := db.QueryRow("SELECT * FROM subscriptions WHERE id = ?", user.Id)

		subscriptionRecord := User{}

		err = subscriptionRow.Scan(&subscriptionRecord.Id, &subscriptionRecord.Email)

		if err != nil {
			t.Fatalf("Something went wrong when trying to get subscription: %v", err)
		}

		if user.Id != subscriptionRecord.Id {
			t.Fatalf("Expected subscription id to be (%d) but got (%d)", subscription.Id, subscriptionRecord.Id)
		}

		if user.Email != subscriptionRecord.Email {
			t.Fatalf("Expected subscription email to be (%s) but got (%s)", user.Email, subscriptionRecord.Email)
		}
	})

	migration.DB.Close()
}
