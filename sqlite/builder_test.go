package sqlite

import (
	"testing"
)

func TestBuilder(t *testing.T) {
	t.Run("TestSelectQuery", func(t *testing.T) {
		// statement := &orm.Statement{
		// 	Table:  "users",
		// 	Select: orm.Select{"id", "email"},
		// 	Joins: orm.Joins{
		// 		&orm.JoinHolder{
		// 			Table: "images",
		// 			Where: []interface{}{
		// 				orm.Join{"users.id": "images.user_id"},
		// 			},
		// 		},
		// 	},
		// 	Where:  []interface{}{orm.Where{"users.role": 1}},
		// 	Limit:  50,
		// 	Offset: 100,
		// }

		// expected := strings.Join([]string{
		// 	"SELECT",
		// 	statements.SPACE + "`id`, `email`",
		// 	"FROM",
		// 	statements.SPACE + "`users`" +
		// 	"LEFT JOIN `images` ON `users`.`id` = `images`.`id`",
		// 	"WHERE",
		// 	statements.SPACE + "`users`.`role` = ?",
		// 	"LIMIT ? OFFSET ?",
		// }, "\r\n")
	})
}
