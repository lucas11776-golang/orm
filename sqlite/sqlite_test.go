package sqlite

import (
	"database/sql"
	"testing"
	"time"

	"github.com/lucas11776-golang/orm"
)

type User struct {
	ID    int64
	Email string
}

func TestScanRows(t *testing.T) {
	sqlite, err := Connect(":memory:")

	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	db := sqlite.Database().(*sql.DB)

	_, err = db.Exec(`CREATE TABLE "msisdns" (
		"id" integer primary key autoincrement not null,
		"created_at" datetime default CURRENT_TIMESTAMP,
		"updated_at" datetime default CURRENT_TIMESTAMP,
		"msisdn" varchar not null,
		"name" varchar,
		"province" varchar,
		"number_of_children" integer,
		"agreed_terms" tinyint(1) not null default 0
		)`)

	if err != nil {
		t.Fatalf("Something went wrong when trying to create table: %v", err)
	}

	_, err = db.Exec(`
		INSERT INTO "msisdns" ("msisdn", "name", "province", "number_of_children", "agreed_terms") VALUES ('253846568785', 'Paulo', 'Maputo', '2', '1');
		INSERT INTO "msisdns" ("msisdn", "name", "province", "number_of_children", "agreed_terms") VALUES ('258843127837', 'Comfy', 'Maputo', '1', '1');
	`)

	if err != nil {
		t.Fatalf("Something went wrong when trying insert record: %v", err)
	}

	rows, err := db.Query("SELECT * FROM msisdns ORDER BY id ASC")

	if err != nil {
		t.Fatalf("Something went wrong when trying to get records: %v", err)
	}

	results, err := sqlite.(*SQLite).scan(rows)

	if err != nil {
		t.Fatalf("Something when wrong when trying to scan rows from database: %v", err)
	}

	if len(results) != 2 {
		t.Fatalf("Expected msisdns to have total of (%d) items but got (%d)", 2, len(results))
	}

	result := results[0]

	if result["id"] != int64(1) {
		t.Fatalf("Expected msisdns first record id to be (%d) but got (%d)", 1, result["id"])
	}

	if result["msisdn"] != "253846568785" {
		t.Fatalf("Expected msisdns first record msisdn to be (%s) but got (%s)", "253846568785", result["msisdn"])
	}

	if result["agreed_terms"] != int64(1) {
		t.Fatalf("Expected msisdns first record agreed terms to be (%v) but got (%v)", int64(1), result["agreed_terms"])
	}

	db.Close()
}

func TestSQLite(t *testing.T) {
	type User struct {
		ID        int64     `json:"id" column:"id" type:"primary_key"`
		CreatedAt time.Time `json:"created_at" column:"created_at" type:"datetime_current"`
		Email     string    `json:"email" column:"email" type:"string"`
	}

	t.Run("TestQuery", func(t *testing.T) {
		db, err := Connect(":memory:")

		if err != nil {
			t.Fatalf("Database connection failed: %v", err)
		}

		err = db.Migration().Migrate(orm.Models{User{}})

		if err != nil {
			t.Fatalf("Database migration failed: %v", err)
		}

		user := &User{
			ID:    1,
			Email: "jeo@doe.com",
		}

		stmt, err := db.Database().(*sql.DB).Prepare("INSERT INTO `users`(`email`) VALUES(?)")

		if err != nil {
			t.Fatalf("Failed to prepare statement: %v", err)
		}

		_, err = stmt.Exec(user.Email)

		if err != nil {
			t.Fatalf("Failed to execute query: %v", err)
		}

		results, err := db.Query(&orm.Statement{
			Table:  "users",
			Select: orm.Select{"id", "email"},
			Where:  []interface{}{orm.Where{"id": user.ID}},
		})

		if err != nil {
			t.Fatalf("Query failed: %v", err)
		}

		if len(results) != 1 {
			t.Fatalf("Expected query result to be (%d) but got (%d)", 1, len(results))
		}

		if results[0]["id"] != user.ID {
			t.Fatalf("Expected user id to be (%d) but got (%d)", user.ID, results[0]["id"])
		}

		if results[0]["email"] != user.Email {
			t.Fatalf("Expected user email to be (%s) but got (%s)", user.Email, results[0]["email"])
		}

	})

	t.Run("TestCount", func(t *testing.T) {
		db, err := Connect(":memory:")

		if err != nil {
			t.Fatalf("Database connection failed: %v", err)
		}

		err = db.Migration().Migrate(orm.Models{User{}})

		if err != nil {
			t.Fatalf("Database migration failed: %v", err)
		}

		user := &User{
			Email: "jeo@doe.com",
		}

		stmt, err := db.Database().(*sql.DB).Prepare("INSERT INTO `users`(`email`) VALUES(?)")

		if err != nil {
			t.Fatalf("Failed to prepare statement: %v", err)
		}

		_, err = stmt.Exec(user.Email)

		if err != nil {
			t.Fatalf("Failed to execute query %v:", err)
		}

		nonExistingUserCount, err := db.Count(&orm.Statement{
			Table: "users",
			Where: []interface{}{orm.Where{"email": "jane@deo.com"}},
		})

		if err != nil {
			t.Fatalf("Failed to execute count: %v", nonExistingUserCount)
		}

		if nonExistingUserCount != int64(0) {
			t.Fatalf("Expected count results to be (%d) but got (%d)", 0, nonExistingUserCount)
		}

		existingUserCount, err := db.Count(&orm.Statement{
			Table: "users",
			Where: []interface{}{orm.Where{"email": user.Email}},
		})

		if err != nil {
			t.Fatalf("Failed to execute count: %v", existingUserCount)
		}

		if existingUserCount != int64(1) {
			t.Fatalf("Expected count results to be (%d) but got (%d)", 1, existingUserCount)
		}
	})

	t.Run("TestInsert", func(t *testing.T) {
		db, err := Connect(":memory:")

		if err != nil {
			t.Fatalf("Database connection failed: %v", err)
		}

		err = db.Migration().Migrate(orm.Models{User{}})

		if err != nil {
			t.Fatalf("Database migration failed: %v", err)
		}

		values := orm.Values{"email": "john@doe.com"}

		result, err := db.Insert(&orm.Statement{
			Table:      "users",
			Values:     values,
			PrimaryKey: "id",
		})

		if err != nil {
			t.Fatalf("Failed insert data: %v", err)
		}

		if result["id"] != int64(1) {
			t.Fatalf("Expected insert user id to be (%d) but got (%d)", 1, result["id"])
		}

		if result["email"] != values["email"] {
			t.Fatalf("Expected insert user id to be (%d) but got (%d)", 1, result["id"])
		}
	})

	t.Run("TestUpdate", func(t *testing.T) {
		db, err := Connect(":memory:")

		if err != nil {
			t.Fatalf("Database connection failed: %v", err)
		}

		err = db.Migration().Migrate(orm.Models{User{}})

		if err != nil {
			t.Fatalf("Database migration failed: %v", err)
		}

		values := orm.Values{"email": "john@doe.com"}

		_, err = db.Insert(&orm.Statement{
			Table:      "users",
			Values:     values,
			PrimaryKey: "id",
		})

		if err != nil {
			t.Fatalf("Failed to insert: %v", err)
		}

		updateUser := orm.Values{"email": "james@doe.com"}

		err = db.Update(&orm.Statement{
			Table:  "users",
			Where:  []interface{}{orm.Where{"id": orm.Where{"=": 1}}},
			Values: updateUser,
		})

		if err != nil {
			t.Fatalf("Failed to updated: %v", err)
		}

		users, err := db.Query(&orm.Statement{
			Table:  "users",
			Where:  []interface{}{orm.Where{"email": orm.Where{"=": updateUser["email"]}}},
			Values: updateUser,
		})

		if err != nil {
			t.Fatalf("Failed query users: %v", err)
		}

		if len(users) != 1 {
			t.Fatalf("Expected users result to be (%d) but got (%d)", 1, len(users))
		}

		if users[0]["email"] != updateUser["email"] {
			t.Fatalf("Expected updated user email to but (%s) but got (%s)", updateUser["email"], users[0]["email"])
		}
	})
}
