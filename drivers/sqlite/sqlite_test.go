package sqlite

import (
	"testing"

	sqlTesting "github.com/lucas11776-golang/orm/drivers/sql/testing"
)

func TestSQLite(t *testing.T) {
	sqlTesting.TestSQLDatabaseBasicOperations(Connect(":memory:"), t)
}
