package statements

import "testing"

func TestStatements(t *testing.T) {
	t.Run("TestSqlSaveValue", func(t *testing.T) {
		if SafeKey("users") != "`users`" {
			t.Fatalf("Expected safe key to be (%s) but got (%s)", "`users`", SafeKey("users"))
		}

		if SafeKey("products.id") != "`products`.`id`" {
			t.Fatalf("Expected safe key to be (%s) but got (%s)", "`products`.`id`", SafeKey("products.id"))
		}
	})
}
