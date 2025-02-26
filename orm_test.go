package orm

import (
	"fmt"
	"testing"
)

type User struct {
	// Connection string `connection:"sqlite" table:"accounts"`
	// Table      string `table:"accounts"`
	ID    int64  `column:"id" type:"primary_key"`
	Email string `column:"email" type:"string"`
}

func TestOrm(t *testing.T) {
	user, err := Model(User{}).Select(Select{
		"*",
	}).Where(Where{
		"email": "jeo@doe.com",
	}).Limit(1).First()

	if err != nil {
		t.Fatalf("Something went wrong when trying to get users")
	}

	fmt.Println("User:", user)
}
