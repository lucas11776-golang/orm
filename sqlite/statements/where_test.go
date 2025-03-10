package statements

import (
	"strings"
	"testing"

	"github.com/lucas11776-golang/orm"
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
			"WHERE",
			SPACE + "`email` = ?",
			SPACE + "AND",
			SPACE + "`age` > ?",
		}, "\r\n")

		if expected != actual {
			t.Fatalf("Expected query where to but (%s) but got (%s)", expected, actual)
		}
	})

	t.Run("TestWhereWithGroup", func(t *testing.T) {
		statement := &Where{
			Where: []interface{}{
				&WhereGroupQueryBuilder{
					Group: []interface{}{orm.Where{"year": orm.Where{"BETWEEN": []int{2007, 2023}}}},
				},
				"OR",
				orm.Where{"title": orm.Where{"LIKE": "lord of the rings"}},
			},
		}

		actual, _ := statement.Statement()
		expected := strings.Join([]string{
			"WHERE",
			SPACE + "(",
			SPACE + SPACE + "`year` BETWEEN ? AND ?",
			SPACE + ")",
			SPACE + "OR",
			SPACE + "`title` LIKE \"%?%\"",
		}, "\r\n")

		if expected != actual {
			t.Fatalf("Expected query where to but (\r\n%s) but got (\r\n%s)", expected, actual)
		}
	})
}
