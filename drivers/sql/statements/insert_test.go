package statements

import "testing"

func TestInsertStatement(t *testing.T) {
	t.Run("TestInsertWithEmptyValue", func(t *testing.T) {
		statement := &Insert{
			Table:        "products",
			InsertValues: map[string]interface{}{},
		}

		if _, err := statement.Statement(); err == nil {
			t.Fatal("Expected to error if insert does not have values")
		}
	})

	t.Run("TestInsert", func(t *testing.T) {
		statement := &Insert{
			Table: "products",
			InsertValues: map[string]interface{}{
				"name":     "Power Bank - (3000mAh)",
				"brand":    "Samsung",
				"price":    490,
				"in_stock": 5,
			},
		}

		actual, _ := statement.Statement()
		expected := "INSERT INTO `products`(`brand`, `in_stock`, `name`, `price`) VALUES(?, ?, ?, ?);"

		if expected != actual {
			t.Fatalf("Expected insert query to be (%s) but got (%s)", actual, expected)
		}

		if size := len(statement.Values()); size != 4 {
			t.Fatalf("Expected values len to be (%d) but got (%d)", 4, size)
		}
	})
}
