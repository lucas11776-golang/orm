# ORM

### Golang ORM is a simple database layer


***Supported Database Types***

- SQLite


<!-- ## Let's get started with golang ORM -->


## Migrations

### First we has to run migration below is a simple example of migration

```go
package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/lucas11776-golang/orm"
	"github.com/lucas11776-golang/orm/sqlite"
)

type User struct {
	Connection string    `connection:"sqlite"`
	ID         int64     `json:"id" column:"id" type:"primary_key"`
	CreatedAt  time.Time `json:"created_at" column:"created_at" type:"datetime_current"`
	FirstName  string    `json:"first_name" column:"first_name" type:"string"`
	LastName   string    `json:"last_name" column:"last_name" type:"string"`
}

func main() {
	// SQLite Connection
	sqlite := sqlite.Connect(":memory:")

	// Add SQLite to the global connections
	orm.DB.Add("sqlite", sqlite)

	// run migrations
	err := sqlite.Migration().Migrate(orm.Models{User{}})

	if err != nil {
		panic(err)
	}

	user, err := orm.Model(User{}).Insert(orm.Values{
		"first_name": "Jeo",
		"last_name":  "Doe",
	})

	if err != nil {
		panic(err)
	}

	data, _ := json.Marshal(user)

	for i := 1; i < 9; i++ {
		fmt.Printf("\r\n%s\r\n", string(data))
	}
}
```

***ORM Supported Types***
- PRIMARY_KEY       - 
- TIMESTAMP         - 
- TIMESTAMP_CURRENT - 
- DATETIME          -  
- DATETIME_CURRENT  -
- INTEGER           -
- FLOAT             -
- STRING            -
- TEXT              -


## Insert

### Now let get started with insert some data in database

```go
package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/lucas11776-golang/orm"
	"github.com/lucas11776-golang/orm/sqlite"
)

type User struct {
	Connection string    `connection:"sqlite"`
	ID         int64     `json:"id" column:"id" type:"primary_key"`
	CreatedAt  time.Time `json:"created_at" column:"created_at" type:"datetime_current"`
	FirstName  string    `json:"first_name" column:"first_name" type:"string"`
	LastName   string    `json:"last_name" column:"last_name" type:"string"`
}

// Comment
func SetupDatabase(db orm.Database) {
	orm.DB.Add("sqlite", db)
}

// Comment
func RunMigration(db orm.Database) error {
	return db.Migration().Migrate(orm.Models{User{}})
}

func main() {
	// SQLite Database
	db := sqlite.Connect(":memory:")

	// SetUp Database
	SetupDatabase(db)
	RunMigration(db)

	user, err := orm.Model(User{}).Insert(orm.Values{
		"first_name": "Jeo",
		"last_name":  "Doe",
	})

	if err != nil {
		panic(err)
	}

	data, _ := json.Marshal(user)

	for i := 1; i < 9; i++ {
		fmt.Printf("\r\n%s\r\n", string(data))
	}
}
```