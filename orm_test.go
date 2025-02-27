package orm

import (
	"fmt"
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

	db := &TestingMemoryDB{}

	DB.Add(connection, db)

	t.Run("TestFirst", func(t *testing.T) {
		db.NextResults(Results{
			map[string]interface{}{"id": 1, "email": "jeo@deo.com"},
		})

		user, _ := Model(User{}).First()

		if user == nil {
			t.Fatalf("User was not found")
		}

		if user.ID != 1 {
			t.Fatalf("Expected id to be (%d) but got (%d)", 1, user.ID)
		}

		if user.ID != 1 {
			t.Fatalf("Expected email to be (%s) but got (%s)", "jeo@deo.com", user.Email)
		}

		fmt.Println("User:", user)
	})

	DB.Remove(connection)
}

type TestingMemoryDB struct {
	next interface{}
}

// Comment
func (ctx *TestingMemoryDB) NextResults(result Results) *TestingMemoryDB {
	ctx.next = result

	return ctx
}

// Comment
func (ctx *TestingMemoryDB) Query(statement *Statement) (Results, error) {
	results := ctx.next

	ctx.next = nil

	return results.(Results), nil
}

// Comment
func (ctx *TestingMemoryDB) Insert(statement *Statement) (Result, error) {
	return ctx.next.(Result), nil
}

// Comment
func (ctx *TestingMemoryDB) Count(statement *Statement) (int64, error) {
	return ctx.next.(int64), nil
}

// Comment
func (ctx *TestingMemoryDB) Database() interface{} {
	return ctx
}

// Comment
func (ctx *TestingMemoryDB) Migration() Migration {
	return nil
}
