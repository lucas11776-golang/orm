package sqlite

import (
	"fmt"
	"testing"
)

type User struct {
	ID    int64
	Email string
}

func TestSQLite(t *testing.T) {

	db := &SQLite{}

	// user, err := db.Query(User{}).Get()

	// if err != nil {

	// }

	fmt.Println("Yes...", db)
}
