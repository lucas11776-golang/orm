package sql

import (
	"strings"
	"testing"

	"github.com/lucas11776-golang/orm"
	"github.com/lucas11776-golang/orm/drivers/sql/statements"
)

func TestBuilder(t *testing.T) {
	t.Run("TestSelectQuery", func(t *testing.T) {
		builder := SQLBuilder{
			Statement: &orm.Statement{
				Table:  "users",
				Select: orm.Select{"id", "email"},
				Joins: orm.Joins{
					&orm.JoinHolder{
						Table: "images",
						Operators: []interface{}{
							&orm.Where{
								Key:      "users.id",
								Operator: "=",
								Value:    "images.user_id",
							},
						},
					},
				},
				Where: []interface{}{&orm.Where{
					Key:      "users.role",
					Operator: "=",
					Value:    1,
				}},
				OrderBy: orm.OrderBy{Columns: "users.id", Order: orm.DESC},
				Limit:   50,
				Offset:  100,
			},
			QueryBuilder: &DefaultQueryBuilder{},
		}

		expected := strings.Join([]string{
			"SELECT",
			statements.SPACE + "`id`, `email`",
			"FROM",
			statements.SPACE + "`users`",
			"LEFT JOIN `images` ON `users`.`id` = `images`.`user_id`",
			"WHERE",
			statements.SPACE + "`users`.`role` = ?",
			"ORDER BY `users`.`id` DESC",
			"LIMIT ? OFFSET ?",
		}, "\r\n")

		actual, values, _ := builder.Query()

		if expected != actual {
			t.Fatalf("Expected query statement to be (%s) but got (%s)", expected, actual)
		}

		if len(values) != 3 {
			t.Fatalf("Expected values size to be (%d) but got (%d)", 3, len(values))
		}
	})

	t.Run("TestCount", func(t *testing.T) {
		builder := SQLBuilder{
			Statement: &orm.Statement{
				Table:  "users",
				Select: orm.Select{"id", "email"},
				Joins: orm.Joins{
					&orm.JoinHolder{
						Table: "images",
						Operators: []interface{}{
							&orm.Where{
								Key:      "users.id",
								Operator: "=",
								Value:    "images.user_id",
							},
						},
					},
				},
				Where: []interface{}{&orm.Where{
					Key:      "users.role",
					Operator: "=",
					Value:    1,
				}},
				OrderBy: orm.OrderBy{Columns: "users.id", Order: orm.DESC},
				Limit:   50,
				Offset:  100,
			},
			QueryBuilder: &DefaultQueryBuilder{},
		}

		expected := strings.Join([]string{
			"SELECT",
			statements.SPACE + "COUNT(*) AS `total`",
			"FROM",
			statements.SPACE + "`users`",
			"LEFT JOIN `images` ON `users`.`id` = `images`.`user_id`",
			"WHERE",
			statements.SPACE + "`users`.`role` = ?",
		}, "\r\n")

		actual, values, _ := builder.Count()

		if expected != actual {
			t.Fatalf("Expected query statement to be (%s) but got (%s)", expected, actual)
		}

		if len(values) != 1 {
			t.Fatalf("Expected values size to be (%d) but got (%d)", 3, len(values))
		}
	})

	t.Run("TestInsert", func(t *testing.T) {
		builder := SQLBuilder{
			Statement: &orm.Statement{
				Table: "users",
				Values: orm.Values{
					"first_name": "Jeo",
					"last_name":  "Doe",
					"email":      "jeo@doe.com",
				},
			},
			QueryBuilder: &DefaultQueryBuilder{},
		}

		actual, values, _ := builder.Insert()
		expected := "INSERT INTO `users`(`email`, `first_name`, `last_name`) VALUES(?, ?, ?);"

		if expected != actual {
			t.Fatalf("Expected insert query to be (%s) but got (%s)", expected, actual)
		}

		if len(values) != 3 {
			t.Fatalf("Expected values size to be (%d) but got (%d)", 2, len(values))
		}
	})

	t.Run("TestUpdate", func(t *testing.T) {
		builder := SQLBuilder{
			Statement: &orm.Statement{
				Table: "users",
				Where: []interface{}{&orm.Where{
					Key:      "id",
					Operator: "=",
					Value:    1,
				}},
				Values: orm.Values{
					"email": "jeo@doe.com",
				},
			},
			QueryBuilder: &DefaultQueryBuilder{},
		}

		expected := strings.Join([]string{
			"UPDATE",
			statements.SPACE + "`users`",
			"SET",
			statements.SPACE + "`email` = ?",
			"WHERE",
			statements.SPACE + "`id` = ?",
		}, "\r\n")
		actual, values, _ := builder.Update()

		if expected != actual {
			t.Fatalf("Expected update query to be (%s) but got (%s)", expected, actual)
		}

		if len(values) != 2 {
			t.Fatalf("Expected values size to be (%d) but got (%d)", 2, len(values))
		}
	})

	t.Run("TestDelete", func(t *testing.T) {
		builder := SQLBuilder{
			Statement: &orm.Statement{
				Table: "users",
				Where: []interface{}{&orm.Where{
					Key:      "id",
					Operator: "=",
					Value:    1,
				}},
			},
			QueryBuilder: &DefaultQueryBuilder{},
		}

		expected := strings.Join([]string{
			"DELETE FROM",
			statements.SPACE + "`users`",
			"WHERE",
			statements.SPACE + "`id` = ?;",
		}, "\r\n")
		actual, values, _ := builder.Delete()

		if expected != actual {
			t.Fatalf("Expected update query to be (%s) but got (%s)", expected, actual)
		}

		if len(values) != 1 {
			t.Fatalf("Expected values size to be (%d) but got (%d)", 1, len(values))
		}
	})
}
