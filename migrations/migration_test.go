package migrations

import (
	"testing"
	"time"

	"github.com/lucas11776-golang/orm"
	"github.com/lucas11776-golang/orm/databases/sqlite"
)

type user struct {
	Connection string    `json:"-" connection:"sqlite" table:"users"`
	ID         int64     `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
}

func (ctx *user) Up() {
	Create("sqlite", "users", func(table *Table) {
		table.Increment("id")
		table.TimeStamp("created_at").Current()
		table.String("first_name").Nullable()
		table.String("last_name").Nullable()
		table.String("email").Unique()
		table.String("password")
	})
}

func (ctx *user) Down() {
	Drop("sqlite", "users")
}

func TestMigration(t *testing.T) {
	t.Run("TestCreateTable", func(t *testing.T) {
		orm.DB.Add("sqlite", sqlite.Connect(":memory:"))

		Migrations(&user{}).Up()

		// TODO: Insert and check if value ID in database...

		orm.DB.Remove("sqlite")
	})

	t.Run("TestDropTable", func(t *testing.T) {
		orm.DB.Add("sqlite", sqlite.Connect(":memory:"))

		migrator := Migrations(&user{})

		migrator.Up()
		migrator.Down()

		// TODO: Insert should error with table does not exits...

		orm.DB.Remove("sqlite")

	})
}
