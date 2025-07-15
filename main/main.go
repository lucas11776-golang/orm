package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/lucas11776-golang/orm"
	"github.com/lucas11776-golang/orm/databases/sqlite"
	"github.com/lucas11776-golang/orm/migrations"
)

type User struct {
	Connection string    `json:"-" connection:"sqlite" table:"users"`
	ID         int64     `json:"id" column:"id" type:"primary_key"`
	CreatedAt  time.Time `json:"created_at" column:"created_at" type:"datetime_current"`
	FirstName  string    `json:"first_name" column:"first_name" type:"string"`
	LastName   string    `json:"last_name" column:"last_name" type:"string"`
}

func (ctx *User) Up() {
	migrations.Create("sqlite", "users", func(table *migrations.Table) {
		table.Increment("id")
		table.TimeStamp("created_at").Current()
		table.String("first_name").Nullable()
		table.String("last_name").Nullable()
		table.String("email")
	})
}

func (ctx *User) Down() {
	migrations.Drop("sqlite", "users")
}

// Comment
func SetupDatabase() {
	db := sqlite.Connect(":memory:")

	orm.DB.Add("sqlite", db)

	migrations.Migrations(&User{}).Up()
}

func main() {
	SetupDatabase()

	user, err := orm.Model(User{}).Insert(orm.Values{
		"first_name": "Joe",
		"last_name":  "Doe",
	})

	if err != nil {
		panic(err)
	}

	data, _ := json.Marshal(user)

	fmt.Println("DATA:", string(data))
}
