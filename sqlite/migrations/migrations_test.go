package migrations

import (
	"database/sql"
	"orm/sqlite/statements"
	"strings"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func TestMigrationStatement(t *testing.T) {
	migration := &Migration{}

	t.Run("TestTypesQuery", func(t *testing.T) {
		primaryKeyExpected := "`id` INTEGER PRIMARY KEY AUTOINCREMENT"
		primaryKeyActual, _ := migration.ColumnStatement("primary_key", "id")

		if primaryKeyExpected != primaryKeyActual {
			t.Fatalf("Expected primary key statement to be (%s) but got (%s)", primaryKeyExpected, primaryKeyActual)
		}

		datetimeCurrentExpected := "`created_at` DATETIME DEFAULT CURRENT_TIMESTAMP"
		datetimeCurrentActual, _ := migration.ColumnStatement("datetime_current", "created_at")

		if datetimeCurrentExpected != datetimeCurrentActual {
			t.Fatalf("Expected datetime statement to be (%s) but got (%s)", datetimeCurrentExpected, datetimeCurrentActual)
		}

		datetimeExpected := "`updated_at` DATETIME"
		datetimeActual, _ := migration.ColumnStatement("datetime", "updated_at")

		if datetimeExpected != datetimeActual {
			t.Fatalf("Expected datetime statement to be (%s) but got (%s)", datetimeExpected, datetimeActual)
		}

		integerExpected := "`year` INTEGER"
		integerActual, _ := migration.ColumnStatement("integer", "year")

		if integerExpected != integerActual {
			t.Fatalf("Expected integer statement to be (%s) but got (%s)", integerExpected, integerActual)
		}

		floatExpected := "`height` FLOAT"
		floatActual, _ := migration.ColumnStatement("float", "height")

		if floatExpected != floatActual {
			t.Fatalf("Expected float statement to be (%s) but got (%s)", floatExpected, floatActual)
		}

		stringExpected := "`email` VARCHAR"
		stringActual, _ := migration.ColumnStatement("string", "email")

		if stringExpected != stringActual {
			t.Fatalf("Expected string statement to be (%s) but got (%s)", stringExpected, stringActual)
		}

		textExpected := "`bio` TEXT"
		textActual, _ := migration.ColumnStatement("text", "bio")

		if textExpected != textActual {
			t.Fatalf("Expected text statement to be (%s) but got (%s)", textExpected, textActual)
		}

		booleanExpected := "`subscribed` BOOLEAN"
		booleanActual, _ := migration.ColumnStatement("boolean", "subscribed")

		if booleanExpected != booleanActual {
			t.Fatalf("Expected boolean statement to be (%s) but got (%s)", booleanExpected, booleanActual)
		}
	})

	t.Run("TestArgumentType", func(t *testing.T) {
		defaultArgExpected := "`subscribed` BOOLEAN DEFAULT false"
		defaultArgActual, _ := migration.ColumnStatement("boolean", "subscribed", "DEFAULT:false")

		if defaultArgExpected != defaultArgActual {
			t.Fatalf("Expected statement with default arg to be (%s) but got (%s)", defaultArgExpected, defaultArgActual)
		}

		notNullArgExpected := "`email` VARCHAR NOT NULL"
		notNullArgActual, _ := migration.ColumnStatement("string", "email", "not_null")

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
		queryActual, err := migration.Query(p)

		if err != nil {
			t.Fatalf("Something went wrong when trying to generate create model table: %v", err)
		}

		if queryExpected != queryActual {
			t.Fatalf("Expected model table query to be (%s) but got (%s)", queryExpected, queryActual)
		}
	})

	migration.DB.Close()
}
