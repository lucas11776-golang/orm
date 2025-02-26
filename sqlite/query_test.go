package sqlite

import (
	"fmt"
	"testing"
)

func TestQuery(t *testing.T) {
	type User struct {
		ID    int64
		Email string
	}

	// query := &Query{}

	// users, err := query.Select(orm.Select{
	// 	"id", "first_name", "last_name", "email", "role",
	// }).WhereGroup(func(group orm.WhereGroupBuilder) {
	// 	group.Where(orm.Where{"role": orm.Where{">=": 1}})
	// }).AndWhere(orm.Where{
	// 	"subscribed": true,
	// }).Paginate(30, 1)

	// if err != nil || users == nil {
	// 	// Handle query/database error
	// 	return
	// }

	fmt.Println("", "")
}
