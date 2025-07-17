package statements

import (
	"strings"
	"testing"
	"time"

	"github.com/lucas11776-golang/orm"
)

// builder := QueryBuilder{
// 			Statement: &orm.Statement{
// 				Table: "users",
// 				Where: []interface{}{&orm.Where{
// 					Key:      "id",
// 					Operator: "=",
// 					Value:    1,
// 				}},
// 			},
// 			Builder: &DefaultQueryBuilder{},
// 		}

// 		expected := strings.Join([]string{
// 			"DELETE FROM",
// 			statements.SPACE + "`users`",
// 			"WHERE",
// 			statements.SPACE + "`id` = ?",
// 		}, "\r\n")
// 		actual, values, _ := builder.Delete()

// 		if expected != actual {
// 			t.Fatalf("Expected update query to be (%s) but got (%s)", expected, actual)
// 		}

// 		if len(values) != 1 {
// 			t.Fatalf("Expected values size to be (%d) but got (%d)", 1, len(values))
// 		}

func TestDeleteStatement(t *testing.T) {
	t.Run("TestDeleteRecord", func(t *testing.T) {
		statement := &Delete{
			Table: "users",
			Where: []interface{}{
				&orm.Where{
					Key:      "status",
					Operator: orm.EQUALS,
					Value:    "DONE",
				},
				"AND",
				&orm.Where{
					Key:      "timestamp",
					Operator: orm.LESS_THEN,
					Value:    time.Now().UnixMilli(),
				},
			},
		}

		expected := strings.Join([]string{
			"DELETE FROM",
			SPACE + "`users`",
			"WHERE",
			SPACE + "`status` = ?",
			SPACE + "AND",
			SPACE + "`timestamp` < ?;",
		}, "\r\n")

		actual, _ := statement.Statement()

		if expected != actual {
			t.Fatalf("Expected insert query to be (%s) but got (%s)", expected, actual)
		}

		if size := len(statement.Values()); size != 2 {
			t.Fatalf("Expected values len to be (%d) but got (%d)", 4, size)
		}
	})
}
