package sqlite

import (
	"testing"

	"github.com/lucas11776-golang/orm/drivers/sql"
)

func TestSQLite(t *testing.T) {
	sql.TestSQLDatabaseBasicOperations(Connect(":memory:"), t)
}
