package statements

import (
	"strings"
	"testing"

	"github.com/lucas11776-golang/orm"
)

func TestUpdateStatement(t *testing.T) {
	t.Run("TestUpdateEmptyValues", func(t *testing.T) {
		statement := &Update{
			Table: "users",
			Where: []interface{}{&orm.Where{
				Key:      "id",
				Operator: "=",
				Value:    1,
			}},
			UpdateValues: orm.Values{},
		}

		if _, err := statement.Statement(); err == nil {
			t.Fatal("Expected to error if update does not have values")
		}
	})

	t.Run("TestUpdate", func(t *testing.T) {
		statement := &Update{
			Table: "users",
			Where: []interface{}{&orm.Where{
				Key:      "id",
				Operator: "=",
				Value:    1,
			}},
			UpdateValues: orm.Values{
				"name":  "John",
				"email": "john@deo.com",
			},
		}

		expected := strings.Join([]string{
			"UPDATE",
			SPACE + "`users`",
			"SET",
			SPACE + "`email` = ?, `name` = ?",
			"WHERE",
			SPACE + "`id` = ?",
		}, "\r\n")
		actual, _ := statement.Statement()

		if expected != actual {
			t.Fatalf("Expected update query to be (%s) but got (%s)", expected, actual)
		}

		if size := len(statement.Values()); size != 3 {
			t.Fatalf("Expected values len to be (%d) but got (%d)", 3, size)
		}
	})
}
