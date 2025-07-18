package sql

import (
	"database/sql"
	"strings"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestSQL(t *testing.T) {
	t.Run("TestTablePrimaryKey", func(t *testing.T) {
		db, err := sql.Open("sqlite3", ":memory:")

		if err != nil {
			t.Fatal(err)
		}

		_, err = db.Exec(strings.Join([]string{
			"CREATE TABLE `users`(",
			"  `name` VARCHAR(255) NOT NULL,",
			"  `user_id` INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,",
			"  `email` VARCHAR(255) NOT NULL",
			");",
		}, "\r\n"))

		if err != nil {
			t.Fatal(err)
		}

		if err != nil {
			t.Fatal(err)
		}

		primaryKey, err := TableInfoPrimaryKey(db, "users")

		if err != nil {
			t.Fatal(err)
		}

		if primaryKey != "user_id" {
			t.Fatalf("Expected users table primary key to be (%s) but got (%s)", "user_id", primaryKey)
		}
	})

	t.Run("TestScanRowsToResultsAndResultsToModels", func(t *testing.T) {
		db, err := sql.Open("sqlite3", ":memory:")

		if err != nil {
			t.Fatal(err)
		}

		type User struct {
			ID       int64  `column:"id"`
			FullName string `column:"full_name"`
			Email    string `column:"email"`
		}

		const (
			Email    string = "jeo@doe.com"
			FullName string = "John Deo"
			ID       int64  = 1
		)

		_, err = db.Exec(strings.Join([]string{
			"CREATE TABLE `users`(",
			"  `id` INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE,",
			"  `full_name` VARCHAR(255) NOT NULL,",
			"  `email` VARCHAR(255) NOT NULL",
			");",
		}, "\r\n"))

		if err != nil {
			t.Fatal(err)
		}

		_, err = db.Exec("INSERT INTO `users`(`full_name`, `email`) VALUES(?,?)", FullName, Email)

		if err != nil {
			t.Fatal(err)
		}

		rows, err := db.Query("SELECT * FROM `users`")

		if err != nil {
			t.Fatal(err)
		}

		results, err := ScanRowsToResults(rows)

		if err != nil {
			t.Fatal(err)
		}

		if len(results) != 1 {
			t.Fatal("Expected results to have 1 record")
		}

		if results[0]["id"] != ID {
			t.Fatalf("Expected record id to be (%d) but got (%v)", ID, results[0]["id"])
		}

		if results[0]["full_name"] != FullName {
			t.Fatalf("Expected record full_name to be (%s) but got (%v)", FullName, results[0]["full_name"])
		}

		if results[0]["email"] != Email {
			t.Fatalf("Expected record email to be (%s) but got (%v)", FullName, results[0]["email"])
		}

		models := ResultsToModels(results, User{})

		if len(models) != 1 {
			t.Fatal("Expected models to have 1 record")
		}

		if models[0].ID != ID {
			t.Fatalf("Expected record id to be (%d) but got (%v)", ID, models[0].ID)
		}

		if models[0].FullName != FullName {
			t.Fatalf("Expected record full_name to be (%s) but got (%v)", FullName, models[0].FullName)
		}

		if models[0].Email != Email {
			t.Fatalf("Expected record email to be (%s) but got (%v)", FullName, models[0].Email)
		}
	})
}
