package statements

import (
	"orm"
	"strings"
	"testing"
)

func TestWhereStatement(t *testing.T) {
	t.Run("TestEmptyWhere", func(t *testing.T) {
		statement := &Where{
			Where: []orm.Where{},
		}

		actual, _ := statement.Statement()
		expected := ""

		if expected != actual {
			t.Fatalf("Expected query where to but (%s) but got (%s)", expected, actual)
		}
	})

	t.Run("TestWhereWithOperation", func(t *testing.T) {
		statement := &Where{
			Where: []orm.Where{
				orm.WhereMatch{"email": "jeo@gmail.com"},
				"AND",
				orm.WhereMatch{"age": orm.WhereMatch{">": 10}},
			},
		}

		actual, _ := statement.Statement()
		expected := strings.Join([]string{
			orm.SPACE + "email = ?", orm.SPACE + "AND", orm.SPACE + "age > ?",
		}, "\r\n")

		if expected != actual {
			t.Fatalf("Expected query where to but (%s) but got (%s)", expected, actual)
		}
	})
}
