package statements

import (
	"strings"
	"testing"

	"github.com/lucas11776-golang/orm"
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
}

func TestSelectOperators(t *testing.T) {
	t.Run("TestAs", func(t *testing.T) {
		statement := &Select{
			Table:  "users",
			Select: []interface{}{orm.AS{"email", "account"}},
		}

		actual, _ := statement.Statement()
		expected := strings.Join([]string{
			"SELECT",
			SPACE + "`email` AS `account`",
			"FROM",
			SPACE + "`users`",
		}, "\r\n")

		if expected != actual {
			t.Fatalf("Expected select query to but (%s) but got (%s)", expected, actual)
		}
	})

	t.Run("TestSum", func(t *testing.T) {
		statement := &Select{
			Table:  "users",
			Select: []interface{}{orm.SUM{"amount", "balance"}},
		}

		actual, _ := statement.Statement()
		expected := strings.Join([]string{
			"SELECT",
			SPACE + "SUM(`amount`) AS `balance`",
			"FROM",
			SPACE + "`users`",
		}, "\r\n")

		if expected != actual {
			t.Fatalf("Expected select query to but (%s) but got (%s)", expected, actual)
		}
	})

	t.Run("TestSum", func(t *testing.T) {
		statement := &Select{
			Table:  "users",
			Select: []interface{}{orm.COUNT{"id", "total"}},
		}

		actual, _ := statement.Statement()
		expected := strings.Join([]string{
			"SELECT",
			SPACE + "COUNT(`id`) AS `total`",
			"FROM",
			SPACE + "`users`",
		}, "\r\n")

		if expected != actual {
			t.Fatalf("Expected select query to but (%s) but got (%s)", expected, actual)
		}
	})
}
