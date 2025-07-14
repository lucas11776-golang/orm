package migration

import "testing"

func TestSelectStatement(t *testing.T) {
	// TODO: Find best way to group migrations tables...
	Migrate := func(connection string, table string, callback func(table *Table)) error {
		return nil
	}

	Migrate("sqlite", "users", func(table *Table) {
		table.PrimaryKey("id")
		table.TimeStamp("created_at").Current()
		table.String("first_name").Nullable()
		table.String("last_name").Nullable()
		table.String("email").Unique()
		table.String("country").Default("ZAR")
		table.String("password")
	})
}
