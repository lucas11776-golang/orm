package statements

import (
	"strings"
	"testing"

	"github.com/lucas11776-golang/orm"
)

func TestJoinStatement(t *testing.T) {
	t.Run("TestJoinWillUseOnlyLeftJoin", func(t *testing.T) {
		statement := Join{
			Join: []*orm.JoinHolder{
				{
					Table: "images",
					Where: []interface{}{
						&orm.Where{
							Key:      "users.id",
							Operator: "=",
							Value:    "images.user_id",
						},
					},
				},
				{
					Table: "rankings",
					Where: []interface{}{
						&orm.Where{
							Key:      "users.id",
							Operator: "=",
							Value:    "rankings.user_id",
						}},
				},
			},
		}

		actual, _ := statement.Statement()
		expected := strings.Join([]string{
			"LEFT JOIN `images` ON `users`.`id` = `images`.`user_id`",
			"LEFT JOIN `rankings` ON `users`.`id` = `rankings`.`user_id`",
		}, "\r\n")

		if expected != actual {
			t.Fatalf("Expected query where to but (\r\n%s) but got (\r\n%s)", expected, actual)
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
								&orm.Where{
									Key:      "users.id",
									Operator: "=",
									Value:    "avatars.user_id",
								},
							},
						},
						"AND",
						&orm.Where{
							Key:      "avatars.group",
							Operator: "=",
							Value:    orm.Raw(10),
						},
					},
				},
				{
					Table: "user_vehicles",
					Where: []interface{}{
						&orm.Where{
							Key:      "users.id",
							Operator: "!=",
							Value:    "user_vehicles.user_id",
						}},
				},
				{
					Table: "vehicles",
					Where: []interface{}{
						&orm.Where{
							Key:      "user_vehicles.brand",
							Operator: "=",
							Value:    orm.Raw("Toyota"),
						},
						"AND",
						&orm.Where{
							Key:      "vehicles.year",
							Operator: "=",
							Value:    orm.Raw(2024),
						},
					},
				},
			},
		}

		actual, _ := statement.Statement()
		expected := strings.Join([]string{
			"LEFT JOIN `avatars` ON (`users`.`id` = `avatars`.`user_id`) AND `avatars`.`group` = ?",
			"LEFT JOIN `user_vehicles` ON `users`.`id` != `user_vehicles`.`user_id`",
			"LEFT JOIN `vehicles` ON `user_vehicles`.`brand` = ? AND `vehicles`.`year` = ?",
		}, "\r\n")

		if expected != actual {
			t.Fatalf("Expected query where to but (\r\n%s\r\n) but got (\r\n%s\r\n)", expected, actual)
		}
	})
}
