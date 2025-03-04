package statements

import (
	"orm"
	"strings"
	"testing"
)

func TestUpdateStatement(t *testing.T) {
	t.Run("TestUpdateQuery", func(t *testing.T) {
		statement := &Update{
			Table:  "users",
			Where:  []interface{}{orm.Where{"id": 1}},
			Update: orm.Values{"name": "John"},
		}

		expected := strings.Join([]string{
			"UPDATE",
			SPACE + "`users`",
			"SET",
			SPACE + "`name` = ?",
			"WHERE",
			SPACE + "`id` = ?",
		}, "\r\n")
		actual, _ := statement.Statement()

		if expected != actual {
			t.Fatalf("Expected update query to be (%s) but got (%s)", expected, actual)
		}
	})
}

/**
UPDATE table_name
SET column1 = value1, column2 = value2, ...
WHERE condition;
**/
