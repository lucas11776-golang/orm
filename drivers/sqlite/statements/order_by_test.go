package statements

import (
	"testing"

	"github.com/lucas11776-golang/orm"
)

func TestOrderByStatement(t *testing.T) {
	t.Run("TestOrderByDesc", func(t *testing.T) {
		statement := &OrderBy{
			OrderBy: orm.OrderBy{Columns: []string{"id"}, Order: orm.DESC},
		}

		expected := "ORDER BY `id` DESC"
		actual, _ := statement.Statement()

		if expected != actual {
			t.Fatalf("Expected order by statement to be (%s) but got (%s)", expected, actual)
		}
	})

	t.Run("TestOrderByASC", func(t *testing.T) {
		statement := &OrderBy{
			OrderBy: orm.OrderBy{Columns: []string{"id"}, Order: orm.ASC},
		}

		expected := "ORDER BY `id` ASC"
		actual, _ := statement.Statement()

		if expected != actual {
			t.Fatalf("Expected order by statement to be (%s) but got (%s)", expected, actual)
		}
	})

	t.Run("TestOrderByASCManyColumns", func(t *testing.T) {
		statement := &OrderBy{
			OrderBy: orm.OrderBy{Columns: []string{"created_at", "published_at"}, Order: orm.ASC},
		}

		expected := "ORDER BY `created_at`,`published_at` ASC"
		actual, _ := statement.Statement()

		if expected != actual {
			t.Fatalf("Expected order by statement to be (%s) but got (%s)", expected, actual)
		}
	})
}
