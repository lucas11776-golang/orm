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

	fmt.Printf("\r\nUsers vehicles: %s\r\n", string(data))
}
