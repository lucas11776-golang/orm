package orm

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"orm/utils/cast"
	"reflect"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
	// _ "github.com/tursodatabase/go-libsql"
)

type Msisdn struct {
	ID               int64     `json:"id" type:"primary_key" column:"id" `
	CreatedAt        time.Time `json:"created_at" type:"datetime" column:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" type:"datetime" column:"updated_at"`
	Msisdn           string    `json:"msisdn" type:"string" column:"msisdn"`
	Name             string    `json:"name" type:"text" column:"name"`
	Province         string    `json:"province" type:"string" column:"province"`
	NumberOfChildren int8      `json:"number_of_children" type:"integer" column:"number_of_children"`
	AgreedTerms      bool      `json:"agreed_terms" type:"bool" column:"agreed_terms"`
}

// Comment
func (ctx *Msisdn) Update() bool {
	fmt.Println("Updated Called")

	return true
}

// Comment
func RowsScan[T any](rows *sql.Rows, entity T) ([]T, error) {
	items := []T{}

	columns, err := rows.Columns()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		item := reflect.Zero(reflect.TypeOf(entity)).Interface().(T)
		values := make([]any, len(columns))
		vMaps := map[string]*interface{}{}

		maps := make([]interface{}, len(values))

		for i := 0; i < len(maps); i++ {
			vMaps[columns[i]] = &values[i]

			maps[i] = &values[i]
		}

		rows.Scan(maps...)

		vt := reflect.ValueOf(&item).Elem()

		for i := 0; i < vt.NumField(); i++ {
			t := vt.Type().Field(i).Tag.Get("column")

			if t == "" {
				continue
			}

			v, ok := vMaps[t]

			if !ok {
				continue
			}

			vt.Field(i).Set(reflect.ValueOf(cast.CastValue(vt.Field(i).Type(), *v)))
		}

		items = append(items, item)
	}

	return items, nil
}

func TestQuery(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")

	if err != nil {
		t.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE "msisdns" (
		"id" integer primary key autoincrement not null,
		"created_at" datetime default CURRENT_TIMESTAMP,
		"updated_at" datetime default CURRENT_TIMESTAMP,
		"msisdn" varchar not null,
		"name" varchar,
		"province" varchar,
		"number_of_children" integer,
		"agreed_terms" tinyint(1) not null default 0
		)`)

	if err != nil {
		t.Fatal(err)
	}

	_, err = db.Exec(`
		INSERT INTO "main"."msisdns" ("id", "created_at", "updated_at", "msisdn", "name", "province", "number_of_children", "agreed_terms") VALUES ('1', '2023-11-08 04:19:07', '2023-11-11 20:26:56', '253846568785', 'Paulo Maculuve', 'Maputo', '2', '1');
		INSERT INTO "main"."msisdns" ("id", "created_at", "updated_at", "msisdn", "name", "province", "number_of_children", "agreed_terms") VALUES ('3', '2023-11-08 07:40:18', '2023-11-09 17:22:20', '258843127837', 'Comfy', 'Maputo', '1', '1');
		INSERT INTO "main"."msisdns" ("id", "created_at", "updated_at", "msisdn", "name", "province", "number_of_children", "agreed_terms") VALUES ('4', '2023-11-08 07:40:36', '2023-11-08 07:40:36', '2582373240894', '', '', '', '0');
		INSERT INTO "main"."msisdns" ("id", "created_at", "updated_at", "msisdn", "name", "province", "number_of_children", "agreed_terms") VALUES ('5', '2023-11-08 09:48:16', '2023-11-08 09:48:16', '253943240784', '', '', '', '0');
	`)

	if err != nil {
		t.Fatal(err)
	}

	rows, err := db.Query("SELECT * FROM msisdns WHERE name != '' AND agreed_terms IS true LIMIT 1000")

	if err != nil {
		t.Fatal(err)
	}

	items, err := RowsScan(rows, Msisdn{})

	if err != nil {
		t.Fatal(err)
	}

	j, err := json.Marshal(items)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("JSON: ", string(j))
}
