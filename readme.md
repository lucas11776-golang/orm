# ORM
## Go ORM - Summary

This document describes a simple Object-Relational Mapping (ORM) library for Go, designed to streamline database interactions.

**Key Features:**

*   **Database Abstraction:** Provides an abstraction layer over SQL databases, allowing developers to interact with databases using Go objects and methods instead of raw SQL.
*   **Model Definition:** Enables the definition of database tables as Go structs, where struct fields map to table columns.
*   **CRUD Operations:** Simplifies common database operations (Create, Read, Update, Delete) through easy-to-use methods.
*   **Query Building:** Offers methods for building and executing database queries without writing raw SQL.
*   **Automatic Mapping:** Automatically maps between Go objects and database rows.
*   **Transactions:** Supports database transactions for ensuring data integrity.
*   **Relationship handling:** handles relations like one to one, one to many and many to many.

**Purpose:**

The primary purpose of this ORM is to simplify database interactions in Go applications. It aims to:

*   **Reduce Boilerplate:** Reduce the amount of repetitive code required for common database operations.
*   **Improve Readability:** Make database code more readable and easier to understand by using Go objects.
*   **Enhance Maintainability:** Improve the maintainability of database-related code.
*   **Boost Development Speed:** increases the development speed.
*   **Improve Scalability:** improve the scalability of the application.

By using this ORM, developers can focus on the application's logic rather than the intricacies of SQL and database interactions.


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


## Where

***Support Operators***

```go
orm.Model(User{}).Where("column", "=", "value").First()
orm.Model(User{}).Where("column", ">", "value").First()
orm.Model(User{}).Where("column", ">=", "value").First()
orm.Model(User{}).Where("column", "<", "value").First()
orm.Model(User{}).Where("column", "<=", "value").First()
orm.Model(User{}).Where("column", "!=", "value").First()
orm.Model(User{}).Where("column", "NOT", "value").First()
orm.Model(User{}).Where("column", "IS", "value").First()
orm.Model(User{}).Where("column", "IS NOT", "value").First()
orm.Model(User{}).Where("column", "LIKE", "value").First()
```

***Support Statements***

```go
orm.Model(User{}).Where("column", "=", "value").AndWhere("column", "=", "value").Get()
orm.Model(User{}).Where("column", "=", "value").OrWhere("column", "=", "value").Get()
```

***Support Group***

```go
orm.Model(User{}).Where("column", "=", "value").AndWhereGroup(func(group orm.WhereGroupBuilder) {
    group.Where("column", "=", "value")
    group.AndWhere("column", "=", "value")
    group.OrWhere("column", "=", "value")
}).Get()
```

```go
orm.Model(User{}).OrWhereGroup(func(group orm.WhereGroupBuilder) {
    group.Where("column", "=", "value")
    group.AndWhere("column", "=", "value")
    group.OrWhere("column", "=", "value")
}).Get()
```


## Limit and Offset

```go
orm.Model(User{}).Limit(10)
orm.Model(User{}).Limit(20).Offset(20)
```


## Order By

```go
orm.Model(User{}).OrderBy("column", orm.ASC)
orm.Model(User{}).OrderBy("column", orm.DESC)
```


## Count

```go
orm.Model(User{}).Where("type", "=", "cheque").Count()
```


## Pagination

```go
orm.Model(User{}).Paginate(50, 10)
```


## Update

```go
orm.Model(User{}).Where("id", "=", 1).Update(orm.Values{
    "first_name": "John",
    "last_name":  "Peterson",
})
```


## Join 

```go
orm.Model(User{}).Join("invites", "users.id", "=", "invites.user_id").Get()
```

***Join Group***

```go
orm.Model(User{}).JoinGroup("invites", func(group orm.JoinGroupBuilder) {

}).Get()
```


### Let`s look at advance join where we are join users owned vehicle

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

type UserVehicle struct {
	Connection string `connection:"sqlite"`
	Table      string `table:"user_vehicles"`
	ID         int64  `json:"id" column:"id" type:"primary_key"`
	UserID     int64  `json:"user_id" column:"user_id" type:"integer"`
	VehicleID  int64  `json:"vehicle_id" column:"vehicle_id" type:"integer"`
	Color      string `json:"color" column:"color" type:"string"`
}

type Vehicle struct {
	Connection string `connection:"sqlite"`
	ID         int64  `json:"id" column:"id" type:"primary_key"`
	Year       int64  `json:"year" column:"year" type:"integer"`
	Brand      string `json:"brand" column:"brand" type:"string"`
	Model      string `json:"model" column:"model" type:"string"`
}

type OwedVehicle struct {
	Connection string `json:"-" connection:"sqlite"`
	Table      string `json:"-" table:"users"`
	ID         int64  `json:"id" column:"id" type:"primary_key"`
	Year       int64  `json:"year" column:"year" type:"integer"`
	Brand      string `json:"brand" column:"brand" type:"string"`
	Model      string `json:"model" column:"model" type:"string"`
}

// Comment
func SetupDatabase(db orm.Database) {
	orm.DB.Add("sqlite", db)
}

// Comment
func RunMigration(db orm.Database) error {
	return db.Migration().Migrate(orm.Models{User{}, Vehicle{}, UserVehicle{}})
}

func main() {
	db := sqlite.Connect(":memory:")

	SetupDatabase(db)
	RunMigration(db)

	user, _ := orm.Model(User{}).Insert(orm.Values{
		"first_name": "Joe",
		"last_name":  "Doe",
	})

	vehicle, _ := orm.Model(Vehicle{}).Insert(orm.Values{
		"year":  2024,
		"brand": "Toyota",
		"model": "Helix",
	})

	_, _ = orm.Model(UserVehicle{}).Insert(orm.Values{
		"user_id":    user.ID,
		"vehicle_id": vehicle.ID,
		"color":      "White",
	})

	vehicles, _ := orm.Model(OwedVehicle{}).Join(
		"user_vehicles", "users.id", "=", "user_vehicles.user_id",
	).Join(
		"vehicles", "user_vehicles.vehicle_id", "=", "vehicles.id",
	).Where(
		"users.id", "=", user.ID,
	).Get()

	data, _ := json.Marshal(vehicles)

	fmt.Printf("\r\nUsers vehicles: %s", string(data))
}
```