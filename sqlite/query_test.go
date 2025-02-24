package sqlite

import (
	"orm"
	"testing"
)

func TestQuery(t *testing.T) {

	query := &Query{}

	_, err := query.Select(orm.Select{
		"id", "first_name", "last_name", "email", "role",
	}).WhereGroup(func(group orm.WhereGroupBuilder) {
		group.Where(orm.Where{"role": orm.Where{">=": 1}})
	}).AndWhere(orm.Where{
		"subscribed": true,
	}).Paginate(30, 1)

	if err != nil {
		// Handle query/database error
	}

	// fmt.Println("", subscribedUsers)
}
