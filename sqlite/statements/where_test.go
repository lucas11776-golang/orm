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
				&orm.Where{
					Key:      "email",
					Operator: "=",
					Value:    "jeo@gmail.com",
				},
				"AND",
				&orm.Where{
					Key:      "age",
					Operator: ">",
					Value:    10,
				},
			},
		}

		actual, err := statement.Statement()

		if err != nil {
			t.Errorf("Failed to build query: %v", err)
		}

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
					Group: []interface{}{&orm.Where{
						Key:      "year",
						Operator: "BETWEEN",
						Value:    []int{2007, 2023},
					}},
				},
				"OR",
				&orm.Where{
					Key:      "title",
					Operator: "LIKE",
					Value:    "lord of the rings",
				},
			},
		}

		actual, err := statement.Statement()

		if err != nil {
			t.Errorf("Failed to build query: %v", err)
		}

		expected := strings.Join([]string{
			"WHERE",
			SPACE + "(",
			SPACE + SPACE + "`year` BETWEEN ? AND ?",
			SPACE + ")",
			SPACE + "OR",
			SPACE + "`title` LIKE \"%?%\"",
		}, "\r\n")

		if expected != actual {
			t.Fatalf("Expected query where to be (%s) but got (%s)", expected, actual)
		}
	})
}
