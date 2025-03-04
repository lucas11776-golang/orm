package statements

import (
	"orm"
	"strings"
	"testing"
)

func TestOrderByStatement(t *testing.T) {
	t.Run("TestOrderByDesc", func(t *testing.T) {
		statement := &OrderBy{
			Order: orm.DESC,
		}

		expected := strings.Join([]string{"ORDER BY", string(statement.Order)}, " ")
		actual, _ := statement.Statement()

		if expected != actual {
			t.Fatalf("Expected order by statement to be (%s) but got (%s)", expected, actual)
		}
	})

	t.Run("TestOrderByASC", func(t *testing.T) {
		statement := &OrderBy{
			Order: orm.ASC,
		}

		expected := strings.Join([]string{"ORDER BY", string(statement.Order)}, " ")
		actual, _ := statement.Statement()

		if expected != actual {
			t.Fatalf("Expected order by statement to be (%s) but got (%s)", expected, actual)
		}
	})
}
