package statements

import (
	"strings"
	"testing"

	"github.com/lucas11776-golang/orm"
)

func TestUpdateStatement(t *testing.T) {
	t.Run("TestUpdateQuery", func(t *testing.T) {
		statement := &Update{
			Table: "users",
			Where: []interface{}{&orm.Where{
				Key:      "id",
				Operator: "=",
				Value:    1,
			}},
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
