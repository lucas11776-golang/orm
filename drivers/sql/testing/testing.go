package testing

import (
	"database/sql"
	"testing"
	"time"

	"github.com/lucas11776-golang/orm"
	"github.com/lucas11776-golang/orm/migrations"
	utils "github.com/lucas11776-golang/orm/utils/sql"
)

// Comment
func TestSQLDatabaseBasicOperations(db orm.Database, t *testing.T) {
	type User struct {
		ID        int64     `json:"id" column:"id"`
		CreatedAt time.Time `json:"created_at" column:"created_at"`
		Email     string    `json:"email" column:"email"`
	}

	userTableScheme := func() *orm.TableScheme {
		return &orm.TableScheme{
			Name: "users",
			Columns: func() []orm.Scheme {
				table := migrations.Table{}

				table.Increment("id")
				table.TimeStamp("created_at").Current()
				table.String("email")

				return table.Columns
			}(),
		}
	}

	t.Run("TestQuery", func(t *testing.T) {
		err := db.Migration().Migrate(userTableScheme())

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
			Where: []interface{}{&orm.Where{
				Key:      "id",
				Operator: "=",
				Value:    user.ID,
			}},
		})

		if err != nil {
			t.Fatalf("Query failed: %v", err)
		}

		models := utils.ResultsToModels(results, User{})

		if len(models) != 1 {
			t.Fatalf("Expected query result to be (%d) but got (%d)", 1, len(models))
		}

		if models[0].ID != user.ID {
			t.Fatalf("Expected user id to be (%d) but got (%d)", user.ID, models[0].ID)
		}

		if models[0].Email != user.Email {
			t.Fatalf("Expected user email to be (%s) but got (%s)", user.Email, models[0].Email)
		}

		if err := db.Migration().Drop("users"); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("TestCount", func(t *testing.T) {
		err := db.Migration().Migrate(userTableScheme())

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
			Where: []interface{}{&orm.Where{
				Key:      "email",
				Operator: "=",
				Value:    "jane@deo.com",
			}},
		})

		if err != nil {
			t.Fatalf("Failed to execute count: %v", nonExistingUserCount)
		}

		if nonExistingUserCount != int64(0) {
			t.Fatalf("Expected count results to be (%d) but got (%d)", 0, nonExistingUserCount)
		}

		existingUserCount, err := db.Count(&orm.Statement{
			Table: "users",
			Where: []interface{}{&orm.Where{
				Key:      "email",
				Operator: "=",
				Value:    user.Email,
			}},
		})

		if err != nil {
			t.Fatalf("Failed to execute count: %v", existingUserCount)
		}

		if existingUserCount != int64(1) {
			t.Fatalf("Expected count results to be (%d) but got (%d)", 1, existingUserCount)
		}

		if err := db.Migration().Drop("users"); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("TestInsert", func(t *testing.T) {
		err := db.Migration().Migrate(userTableScheme())

		if err != nil {
			t.Fatalf("Database migration failed: %v", err)
		}

		values := orm.Values{"email": "john@doe.com"}

		result, err := db.Insert(&orm.Statement{
			Table:  "users",
			Values: values,
		})

		model := utils.ResultToModel(result, User{})

		if err != nil {
			t.Fatalf("Failed insert data: %v", err)
		}

		if model.ID != int64(1) {
			t.Fatalf("Expected insert user id to be (%d) but got (%d)", 1, model.ID)
		}

		if model.Email != values["email"] {
			t.Fatalf("Expected insert user email to be (%s) but got (%s)", values["email"], model.Email)
		}

		if err := db.Migration().Drop("users"); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("TestUpdate", func(t *testing.T) {
		err := db.Migration().Migrate(userTableScheme())

		if err != nil {
			t.Fatalf("Database migration failed: %v", err)
		}

		values := orm.Values{"email": "john@doe.com"}

		_, err = db.Insert(&orm.Statement{
			Table:  "users",
			Values: values,
		})

		if err != nil {
			t.Fatalf("Failed to insert: %v", err)
		}

		update := orm.Values{"email": "james@doe.com"}

		err = db.Update(&orm.Statement{
			Table: "users",
			Where: []interface{}{&orm.Where{
				Key:      "id",
				Operator: "=",
				Value:    1,
			}},
			Values: update,
		})

		if err != nil {
			t.Fatalf("Failed to updated: %v", err)
		}

		results, err := db.Query(&orm.Statement{
			Table: "users",
			Where: []interface{}{&orm.Where{
				Key:      "email",
				Operator: "=",
				Value:    update["email"],
			}},
			Values: update,
		})

		models := utils.ResultsToModels(results, User{})

		if err != nil {
			t.Fatalf("Failed query users: %v", err)
		}

		if len(models) != 1 {
			t.Fatalf("Expected users result to be (%d) but got (%d)", 1, len(models))
		}

		if models[0].Email != update["email"] {
			t.Fatalf("Expected updated user email to but (%s) but got (%s)", update["email"], models[0].Email)
		}

		if err := db.Migration().Drop("users"); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("TestDelete", func(t *testing.T) {
		err := db.Migration().Migrate(userTableScheme())

		if err != nil {
			t.Fatalf("Database migration failed: %v", err)
		}

		values := orm.Values{"email": "john@doe.com"}

		_, err = db.Insert(&orm.Statement{
			Table:  "users",
			Values: values,
		})

		if err != nil {
			t.Fatalf("Failed to insert: %v", err)
		}

		delete := orm.Values{"email": "james@doe.com"}

		err = db.Delete(&orm.Statement{
			Table: "users",
			Where: []interface{}{&orm.Where{
				Key:      "id",
				Operator: "=",
				Value:    1,
			}},
			Values: delete,
		})

		if err != nil {
			t.Fatalf("Failed to delete: %v", err)
		}

		count, err := db.Count(&orm.Statement{
			Table:  "users",
			Values: delete,
		})

		if err != nil {
			t.Fatalf("Failed to count: %v", err)
		}

		if count != 0 {
			t.Fatalf("Expected users table to be empty.")
		}

		if err := db.Migration().Drop("users"); err != nil {
			t.Fatal(err)
		}
	})
}
