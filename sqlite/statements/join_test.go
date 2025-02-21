package statements

import (
	"fmt"
	"orm"
	"strings"
	"testing"
)

func TestJoin(t *testing.T) {
	t.Run("TestJoinWillUseOnlyLeftJoin", func(t *testing.T) {
		statement := Join{
			Join: []*orm.JoinHolder{
				{
					Table: "images",
					Where: []interface{}{orm.Join{"users.id": "images.user_id"}},
				},
				{
					Table: "rankings",
					Where: []interface{}{orm.Join{"users.id": "rankings.user_id"}},
				},
			},
		}

		actual, _ := statement.Statement()
		expected := strings.Join([]string{
			"LEFT JOIN `images` ON `users`.`id` = `images`.`user_id`",
			"LEFT JOIN `rankings` ON `users`.`id` = `rankings`.`user_id`",
		}, "\r\n")

		// fmt.Println(actual, expected)
		fmt.Println()

		if expected != actual {
			t.Fatalf("Expected query where to but (%s) but got (%s)", expected, actual)
		}
	})

	t.Run("TestJoinWithGroup", func(t *testing.T) {
		statement := Join{
			Join: []*orm.JoinHolder{
				{
					Table: "avatars",
					Where: []interface{}{
						&JoinGroupQueryBuilder{
							Joins: []interface{}{
								orm.Join{"users.id": "avatars.user_id"},
							},
						},
						"AND",
						orm.Join{"avatars.group": orm.RawValue(10)},
					},
				},
				{
					Table: "rankings",
					Where: []interface{}{orm.Join{"users.id": "rankings.user_id"}},
				},
			},
		}

		actual, _ := statement.Statement()
		expected := strings.Join([]string{
			"LEFT JOIN `avatars` ON (`users`.`id` = `avatars`.`user_id`) AND `avatars`.`group` = ?",
			"LEFT JOIN `rankings` ON `users`.`id` = `rankings`.`user_id`",
		}, "\r\n")

		if expected != actual {
			t.Fatalf("Expected query where to but (\r\n%s\r\n) but got (\r\n%s\r\n)", expected, actual)
		}
	})
}
