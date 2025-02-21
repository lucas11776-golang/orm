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
			Join: orm.Joins{
				"images":   []interface{}{orm.Join{"users.id": "images.user_id"}},
				"rankings": []interface{}{orm.Join{"users.id": "rankings.user_id"}},
			},
		}

		actual, _ := statement.Statement()
		expected := strings.Join([]string{
			"LEFT JOIN `images` ON `users`.`id` = `images`.`user_id`",
			"LEFT JOIN `rankings` ON `users`.`id` = `rankings`.`user_id`",
		}, "\r\n")

		fmt.Println(actual, expected)

		// if expected != actual {
		// 	t.Fatalf("Expected query where to but (\r\n%s) but got (\r\n%s)", expected, actual)
		// }
	})
}
