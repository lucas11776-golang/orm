package statements

import (
	"orm"
	"testing"
)

func TestOrderByStatement(t *testing.T) {
	t.Run("TestOrderByDesc", func(t *testing.T) {
		statement := &OrderBy{
			OrderBy: orm.OrderBy{"id", orm.DESC},
		}

		expected := "ORDER BY `id` DESC"
		actual, _ := statement.Statement()

		if expected != actual {
			t.Fatalf("Expected order by statement to be (%s) but got (%s)", expected, actual)
		}
	})

	t.Run("TestOrderByASC", func(t *testing.T) {
		statement := &OrderBy{
			OrderBy: orm.OrderBy{"id", orm.ASC},
		}

		expected := "ORDER BY `id` ASC"
		actual, _ := statement.Statement()

		if expected != actual {
			t.Fatalf("Expected order by statement to be (%s) but got (%s)", expected, actual)
		}
	})
}
