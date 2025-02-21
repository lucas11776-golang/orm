package statements

import (
	"orm"
	"strings"
	"testing"
)

func TestWhereStatement(t *testing.T) {
	t.Run("TestEmptyWhere", func(t *testing.T) {
		statement := &Where{
			Where: []interface{}{},
		}

		actual, _ := statement.Statement()
		expected := ""

		if expected != actual {
			t.Fatalf("Expected query where to but (%s) but got (%s)", expected, actual)
		}
	})

	t.Run("TestWhereWithOperation", func(t *testing.T) {
		statement := &Where{
			Where: []interface{}{
				orm.Where{"email": "jeo@gmail.com"},
				"AND",
				orm.Where{"age": orm.Where{">": 10}},
			},
		}

		actual, _ := statement.Statement()
		expected := strings.Join([]string{
			SPACE + "email = ?", SPACE + "AND", SPACE + "age > ?",
		}, "\r\n")

		if expected != actual {
			t.Fatalf("Expected query where to but (%s) but got (%s)", expected, actual)
		}
	})
}
