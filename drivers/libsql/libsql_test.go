package libsql

import (
	"testing"

	driverTesting "github.com/lucas11776-golang/orm/drivers/sql/testing"
)

func TestLibSQL(t *testing.T) {
	t.Run("TestBasicSQLOperations", func(t *testing.T) {
		driverTesting.TestSQLDatabaseBasicOperations(Connect(":memory:"), t)
	})
}
