package statements

import (
	"orm"
	"strings"
	"testing"
)

func TestSelectStatement(t *testing.T) {
	t.Run("TestEmptySelect", func(t *testing.T) {
		statement := &Select{
			Table: "users",
		}

		actual, _ := statement.Statement()
		expected := strings.Join([]string{
			"SELECT",
			SPACE + "*",
			"FROM",
			SPACE + "`users`",
		}, "\r\n")

		if expected != actual {
			t.Fatalf("Expected select query to but (%s) but got (%s)", expected, actual)
		}
	})

	t.Run("TestSelectFields", func(t *testing.T) {
		statement := &Select{
			Table:  "users",
			Select: []interface{}{"id", "email"},
		}

		actual, _ := statement.Statement()
		expected := strings.Join([]string{
			"SELECT",
			SPACE + "`id`, `email`",
			"FROM",
			SPACE + "`users`",
		}, "\r\n")

		if expected != actual {
			t.Fatalf("Expected select query to but (%s) but got (%s)", expected, actual)
		}
	})

	t.Run("TestSelectOperators", func(t *testing.T) {
		statement := &Select{
			Table:  "users",
			Select: []interface{}{"id", orm.AS{"email", "account"}, orm.SUM{"amount", "balance"}},
		}

		actual, _ := statement.Statement()
		expected := strings.Join([]string{
			"SELECT",
			SPACE + "`id`, `email` AS `account`, SUM(`amount`) AS `balance`",
			"FROM",
			SPACE + "`users`",
		}, "\r\n")

		if expected != actual {
			t.Fatalf("Expected select query to but (%s) but got (%s)", expected, actual)
		}

	})
}
