package sql

import (
	"database/sql"
	"strings"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

// func TestScanRows(t *testing.T) {
// 	sqlite := Connect(":memory:")

// 	db := sqlite.Database().(*sql.DB)

// 	_, err := db.Exec(`CREATE TABLE "msisdns" (
// 		"id" integer primary key autoincrement not null,
// 		"created_at" datetime default CURRENT_TIMESTAMP,
// 		"updated_at" datetime default CURRENT_TIMESTAMP,
// 		"msisdn" varchar not null,
// 		"name" varchar,
// 		"province" varchar,
// 		"number_of_children" integer,
// 		"agreed_terms" tinyint(1) not null default 0
// 		)`)

// 	if err != nil {
// 		t.Fatalf("Something went wrong when trying to create table: %v", err)
// 	}

// 	_, err = db.Exec(`
// 		INSERT INTO "msisdns" ("msisdn", "name", "province", "number_of_children", "agreed_terms") VALUES ('258843127837', 'Comfy', 'Maputo', 1, 1);
// 	`)

// 	if err != nil {
// 		t.Fatalf("Something went wrong when trying insert record: %v", err)
// 	}

// 	rows, err := db.Query("SELECT * FROM msisdns ORDER BY id ASC")

// 	if err != nil {
// 		t.Fatalf("Something went wrong when trying to get records: %v", err)
// 	}

// 	results, err := sqlite.(*SQLite).scan(rows)

// 	if err != nil {
// 		t.Fatalf("Something when wrong when trying to scan rows from database: %v", err)
// 	}

// 	if len(results) != 1 {
// 		t.Fatalf("Expected msisdns to have total of (%d) items but got (%d)", 1, len(results))
// 	}

// 	result := results[0]

// 	if result["id"] != int64(1) {
// 		t.Fatalf("Expected msisdns first record id to be (%d) but got (%d)", 1, result["id"])
// 	}

// 	if result["msisdn"] != "258843127837" {
// 		t.Fatalf("Expected msisdns first record msisdn to be (%s) but got (%s)", "258843127837", result["msisdn"])
// 	}

// 	if result["agreed_terms"] != int64(1) {
// 		t.Fatalf("Expected msisdns first record agreed terms to be (%v) but got (%v)", int64(1), result["agreed_terms"])
// 	}

// 	db.Close()
// }

func TestSQL(t *testing.T) {
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
