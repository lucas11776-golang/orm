package migrations

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

// Model Example
type User struct {
	Id              int64  `column:"id" type:"primary_key"`
	CreatedAt       int64  `column:"created_at" type:"datetime_current"`
	YearDateOfBirth int    `column:"year_date_of_birth" type:"datetime,default:2008"`
	Email           string `column:"email" type:"string,not_null"`
	Name            string `column:"Name" type:"string,not_null"`
	Bio             int    `column:"bio" type:"text,default:'Biography'"`
	Password        int    `column:"password" type:"string,not_null"`
}

func TestMigration(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")

	if err != nil {
		t.Fatalf("Something went wrong when trying to connect to database: %v", err)
	}

	migration := &Migration{DB: db}

	t.Run("TestTypesQuery", func(t *testing.T) {
		primaryKeyExpected := "`id` INTEGER PRIMARY KEY AUTOINCREMENT"
		primaryKeyActual, _ := migration.Statement("primary_key", "id")

		if primaryKeyExpected != primaryKeyActual {
			t.Fatalf("Expected primary key statement to be (%s) but got (%s)", primaryKeyExpected, primaryKeyActual)
		}

		datetimeCurrentExpected := "`created_at` DATETIME DEFAULT CURRENT_TIMESTAMP"
		datetimeCurrentActual, _ := migration.Statement("datetime_current", "created_at")

		if datetimeCurrentExpected != datetimeCurrentActual {
			t.Fatalf("Expected datetime statement to be (%s) but got (%s)", datetimeCurrentExpected, datetimeCurrentActual)
		}

		datetimeExpected := "`updated_at` DATETIME"
		datetimeActual, _ := migration.Statement("datetime", "updated_at")

		if datetimeExpected != datetimeActual {
			t.Fatalf("Expected datetime statement to be (%s) but got (%s)", datetimeExpected, datetimeActual)
		}

		integerExpected := "`year` INTEGER"
		integerActual, _ := migration.Statement("integer", "year")

		if integerExpected != integerActual {
			t.Fatalf("Expected integer statement to be (%s) but got (%s)", integerExpected, integerActual)
		}

		floatExpected := "`height` FLOAT"
		floatActual, _ := migration.Statement("float", "height")

		if floatExpected != floatActual {
			t.Fatalf("Expected float statement to be (%s) but got (%s)", floatExpected, floatActual)
		}

		stringExpected := "`email` VARCHAR"
		stringActual, _ := migration.Statement("string", "email")

		if stringExpected != stringActual {
			t.Fatalf("Expected string statement to be (%s) but got (%s)", stringExpected, stringActual)
		}

		textExpected := "`bio` TEXT"
		textActual, _ := migration.Statement("text", "bio")

		if textExpected != textActual {
			t.Fatalf("Expected text statement to be (%s) but got (%s)", textExpected, textActual)
		}

		booleanExpected := "`subscribed` BOOLEAN"
		booleanActual, _ := migration.Statement("boolean", "subscribed")

		if booleanExpected != booleanActual {
			t.Fatalf("Expected boolean statement to be (%s) but got (%s)", booleanExpected, booleanActual)
		}
	})

	t.Run("TestArgumentType", func(t *testing.T) {
		booleanExpected := "`subscribed` BOOLEAN DEFAULT false"
		booleanActual, _ := migration.Statement("boolean", "subscribed", "DEFAULT:false")

		if booleanExpected != booleanActual {
			t.Fatalf("Expected boolean statement to be (%s) but got (%s)", booleanExpected, booleanActual)
		}
	})

	db.Close()
}
