package orm

import (
	"math/rand"
	"testing"
)

func TestOrm(t *testing.T) {
	t.Run("TestGetOptions", func(t *testing.T) {
		type User struct {
			Connection string `connection:"sqlite" table:"accounts"`
		}

		userOptions := getOptions(User{})

		if userOptions.connection != "sqlite" {
			t.Fatalf("Expected connection to be (%s) but got (%s)", "sqlite", userOptions.connection)
		}

		if userOptions.table != "accounts" {
			t.Fatalf("Expected table to be (%s) but got (%s)", "accounts", userOptions.table)
		}

		type Product struct{}

		productOptions := getOptions(Product{})

		if productOptions.connection != DefaultDatabaseName {
			t.Fatalf("Expected connection to be (%s) but got (%s)", DefaultDatabaseName, productOptions.connection)
		}

		if productOptions.table != "products" {
			t.Fatalf("Expected table to be (%s) but got (%s)", "products", productOptions.table)
		}
	})

}

func TestOrmQuery(t *testing.T) {
	const connection = "memory"

	type User struct {
		Connection string `connection:"memory"`
		ID         int64  `column:"id" type:"primary_key"`
		Email      string `column:"email" type:"string"`
	}

	db := &mockDB{}

	DB.Add(connection, db)

	t.Run("TestCount", func(t *testing.T) {
		accounts := int64(rand.Float32() * 10000)

		db.nextResults(accounts)

		result, _ := Model(User{}).Where("account_type", "=", "savings").Count()

		if accounts != result {
			t.Fatalf("Expected number of accounts to be (%d) but got (%d)", accounts, result)
		}
	})

	t.Run("TestFirst", func(t *testing.T) {
		users := Results{
			map[string]interface{}{"id": int64(1), "email": "jeo@deo.com"},
		}

		db.nextResults(users)

		result, _ := Model(User{}).First()

		if result == nil {
			t.Fatalf("User was not found")
		}

		if result.ID != users[0]["id"] {
			t.Fatalf("Expected id to be (%d) but got (%d)", users[0]["id"], result.ID)
		}

		if result.Email != users[0]["email"] {
			t.Fatalf("Expected email to be (%s) but got (%s)", users[0]["email"], result.Email)
		}
	})

	t.Run("TestGet", func(t *testing.T) {
		users := Results{
			map[string]interface{}{"id": int64(1), "email": "jeo@deo.com"},
			map[string]interface{}{"id": int64(2), "email": "jane@deo.com"},
		}

		db.nextResults(users)

		results, _ := Model(User{}).Get()

		if results == nil {
			t.Fatalf("Users was not found")
		}

		for i := 0; i < len(users); i++ {
			if results[i].ID != users[i]["id"] {
				t.Fatalf("Expected id  in index %d  to be (%d) but got (%d)", i, users[i]["id"], results[i].ID)
			}

			if results[i].Email != users[i]["email"] {
				t.Fatalf("Expected email in index %d to be (%s) but got (%s)", i, users[i]["email"], results[i].Email)
			}
		}
	})

	t.Run("TestPagination", func(t *testing.T) {
		total := int64(rand.Float32() * 1000000)
		perPage := int64(50)
		page := int64(4)
		users := Results{
			map[string]interface{}{"id": int64(1), "email": "jeo@deo.com"},
			map[string]interface{}{"id": int64(2), "email": "jane@deo.com"},
		}

		db.nextResults(users)
		db.nextResults(total)

		results, _ := Model(User{}).Where("email", "LIKE", "@deo.com").Paginate(perPage, page)

		if results == nil {
			t.Fatalf("Users was not found")
		}

		if results.Total != total {
			t.Fatalf("Expected total results count to be (%d) but got (%d)", total, results.Total)
		}

		if results.PerPage != perPage {
			t.Fatalf("Expected per page results to be (%d) but got (%d)", perPage, results.PerPage)
		}

		if results.Page != page {
			t.Fatalf("Expected current page to be (%d) but got (%d)", page, results.Page)
		}

		for i := 0; i < len(users); i++ {
			if results.Items[i].ID != users[i]["id"] {
				t.Fatalf("Expected id  in index %d  to be (%d) but got (%d)", i, users[i]["id"], results.Items[i].ID)
			}

			if results.Items[i].Email != users[i]["email"] {
				t.Fatalf("Expected email in index %d to be (%s) but got (%s)", i, users[i]["email"], results.Items[i].Email)
			}
		}
	})

	t.Run("TestUpdate", func(t *testing.T) {
		db.nextResults(nil)

		err := Model(User{}).Where("id", "=", 1).Update(Values{"email": "john@doe.com"})

		if err != nil {
			t.Fatalf("Something went wrong when trying to update record: %v", err)
		}
	})

	t.Run("TestInsert", func(t *testing.T) {
		update := Values{"email": "john@deo.com"}
		record := Result{"id": int64(1), "email": update["email"]}

		db.nextResults(record)

		result, _ := Model(User{}).Insert(update)

		if result == nil {
			t.Fatalf("Failed to insert user record")
		}

		if result.ID != record["id"] {
			t.Fatalf("Expected id to be (%d) but got (%d)", record["id"], result.ID)
		}

		if result.Email != record["email"] {
			t.Fatalf("Expected email to be (%s) but got (%s)", record["email"], result.Email)
		}
	})

	t.Run("TestDelete", func(t *testing.T) {
		update := Values{"email": "john@deo.com"}
		record := Result{"id": int64(1), "email": update["email"]}

		db.nextResults(record)

		result, _ := Model(User{}).Insert(update)

		if result == nil {
			t.Fatalf("Failed to insert user record")
		}

		db.nextResults(int64(1))

		_ = Model(User{}).Where("id", "=", 1).Delete()

		count, _ := Model(User{}).Count()

		if count != 0 {
			t.Fatal("Expected users table to be empty")
		}
	})

	DB.Remove(connection)
}

type mockDB struct {
	next []interface{}
}

// Comment
func (ctx *mockDB) nextResults(result interface{}) *mockDB {
	ctx.next = append(ctx.next, result)

	return ctx
}

// Comment
func (ctx *mockDB) unshift() interface{} {
	if len(ctx.next) == 0 {
		return nil
	}

	v := ctx.next[0]

	ctx.next = ctx.next[1:]

	return v
}

// Comment
func (ctx *mockDB) Query(statement *Statement) (Results, error) {
	return ctx.unshift().(Results), nil
}

// Comment
func (ctx *mockDB) Insert(statement *Statement) (Result, error) {
	return ctx.unshift().(Result), nil
}

// Comment
func (ctx *mockDB) Count(statement *Statement) (int64, error) {
	return ctx.unshift().(int64), nil
}

// Comment
func (ctx *mockDB) Update(Statement *Statement) error {
	err, ok := ctx.unshift().(error)

	if !ok {
		return nil
	}

	return err
}

// Comment
func (ctx *mockDB) Delete(Statement *Statement) error {
	ctx.unshift()

	ctx.nextResults(int64(0))

	return nil
}

// Comment
func (ctx *mockDB) Database() interface{} {
	return ctx
}

// Comment
func (ctx *mockDB) Migration() Migration {
	return nil
}
