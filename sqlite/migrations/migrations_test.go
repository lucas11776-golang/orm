package migrations

import (
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestMigration(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")

	if err != nil {
		t.Fatalf("Something went wrong when trying to connect to database: %v", err)
	}

	migration := &Migration{DB: db}

	t.Run("TestTypesQuery", func(t *testing.T) {
		primaryKeyExpected := "`id` integer primary key autoincrement"
		primaryKeyActual, _ := migration.Statement("primary_key", "id")

		if primaryKeyExpected != primaryKeyActual {
			t.Fatalf("Expected primary key statement to be (%s) but got (%s)", primaryKeyExpected, primaryKeyActual)
		}

		// datetimeExpected := "`created_at` datetime"
		// datetimeActual, _ := migration.Statement("datetime", "created_at")

		// if datetimeExpected != datetimeActual {
		// 	t.Fatalf("Expected datetime statement to be (%s) but got (%s)", datetimeExpected, datetimeActual)
		// }

		// datetimeAutoExpected := "`created_at` datetime"
		// datetimeAutoActual, _ := migration.Statement("datetime", "created_at")

		// if datetimeAutoExpected != datetimeAutoActual {
		// 	t.Fatalf("Expected datetime statement to be (%s) but got (%s)", datetimeExpected, datetimeActual)
		// }

		// integerExpected := "`created_at` datetime"
		// integerActual, _ := migration.Statement("datetime", "created_at")

		// if integerExpected != integerActual {
		// 	t.Fatalf("Expected datetime statement to be (%s) but got (%s)", datetimeExpected, datetimeActual)
		// }

		// floatExpected := "`created_at` datetime"
		// floatActual, _ := migration.Statement("datetime", "created_at")

		// if floatExpected != floatActual {
		// 	t.Fatalf("Expected datetime statement to be (%s) but got (%s)", datetimeExpected, datetimeActual)
		// }

		// stringExpected := "`created_at` datetime"
		// stringActual, _ := migration.Statement("datetime", "created_at")

		// if stringExpected != stringActual {
		// 	t.Fatalf("Expected datetime statement to be (%s) but got (%s)", datetimeExpected, datetimeActual)
		// }

		// textExpected := "`created_at` datetime"
		// textActual, _ := migration.Statement("datetime", "created_at")

		// if textExpected != textActual {
		// 	t.Fatalf("Expected datetime statement to be (%s) but got (%s)", datetimeExpected, datetimeActual)
		// }

		// booleanExpected := "`created_at` datetime"
		// booleanActual, _ := migration.Statement("datetime", "created_at")

		// if booleanExpected != booleanActual {
		// 	t.Fatalf("Expected datetime statement to be (%s) but got (%s)", datetimeExpected, datetimeActual)
		// }
	})

	db.Close()
}
