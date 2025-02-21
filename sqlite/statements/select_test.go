package statements

import (
	"strings"
	"testing"
)

func TestSelectStatement(t *testing.T) {
	t.Run("TestEmptySelect", func(t *testing.T) {
		statement := &Select{}

		actual, _ := statement.Statement()
		expected := strings.Join([]string{
			"SELECT", SPACE + "*", "FROM",
		}, "\r\n")

		if expected != actual {
			t.Fatalf("Expected select query to but (%s) but got (%s)", expected, actual)
		}
	})

	t.Run("TestSelectFields", func(t *testing.T) {
		statement := &Select{
			Select: []interface{}{"id", "email"},
		}

		actual, _ := statement.Statement()
		expected := strings.Join([]string{
			"SELECT",
			SPACE + "`id`, `email`",
			"FROM",
		}, "\r\n")

		if expected != actual {
			t.Fatalf("Expected select query to but (%s) but got (%s)", expected, actual)
		}
	})
}
