package sqlite

import (
	"orm"
	"orm/sqlite/statements"
	"strings"
	"testing"
)

func TestBuilder(t *testing.T) {
	t.Run("TestSelectQuery", func(t *testing.T) {
		builder := QueryBuilder{
			Statement: &orm.Statement{
				Table:  "users",
				Select: orm.Select{"id", "email"},
				Joins: orm.Joins{
					&orm.JoinHolder{
						Table: "images",
						Where: []interface{}{
							orm.Join{"users.id": "images.user_id"},
						},
					},
				},
				Where:   []interface{}{orm.Where{"users.role": 1}},
				OrderBy: orm.OrderBy{"users.id", orm.DESC},
				Limit:   50,
				Offset:  100,
			},
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
		builder := QueryBuilder{
			Statement: &orm.Statement{
				Table:  "users",
				Select: orm.Select{"id", "email"},
				Joins: orm.Joins{
					&orm.JoinHolder{
						Table: "images",
						Where: []interface{}{
							orm.Join{"users.id": "images.user_id"},
						},
					},
				},
				Where:   []interface{}{orm.Where{"users.role": 1}},
				OrderBy: orm.OrderBy{"users.id", orm.DESC},
				Limit:   50,
				Offset:  100,
			},
		}

		expected := strings.Join([]string{
			"SELECT",
			statements.SPACE + "COUNT(*) AS `total`",
			"FROM",
			statements.SPACE + "`users`",
			"LEFT JOIN `images` ON `users`.`id` = `images`.`user_id`",
			"WHERE",
			statements.SPACE + "`users`.`role` = ?",
			"ORDER BY `users`.`id` DESC",
			"LIMIT ? OFFSET ?",
		}, "\r\n")
		actual, values, _ := builder.Count()

		if expected != actual {
			t.Fatalf("Expected query statement to be (%s) but got (%s)", expected, actual)
		}

		if len(values) != 3 {
			t.Fatalf("Expected values size to be (%d) but got (%d)", 3, len(values))
		}
	})

	t.Run("TestInsert", func(t *testing.T) {
		builder := QueryBuilder{
			Statement: &orm.Statement{
				Table: "subscriptions",
				Values: orm.Values{
					"email": "jeo@doe.com",
				},
			},
		}

		expected := "INSERT INTO `subscriptions` (`email`) VALUES (?)"
		actual, _ := builder.Insert()

		if expected != actual {
			t.Fatalf("Expected insert query to be (%s) but got (%s)", expected, actual)
		}

		if len(builder.Values) != 1 {
			t.Fatalf("Expected values size to be (%d) but got (%d)", 1, len(builder.Values))
		}
	})
}
