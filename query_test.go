package orm

import (
	"database/sql"
	"fmt"
	"orm/utils/cast"
	"reflect"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
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

	cols, err := rows.Columns()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		record := reflect.Zero(reflect.TypeOf(entity)).Interface().(T)
		v := make([]any, len(cols))
		maps := make([]interface{}, len(v))
		vMap := map[string]*interface{}{}

		for i := 0; i < len(maps); i++ {
			vMap[cols[i]] = &v[i]

			maps[i] = &v[i]
		}

		rows.Scan(maps...)

		vElem := reflect.ValueOf(&record).Elem()

		for i := 0; i < vElem.NumField(); i++ {
			tag := vElem.Type().Field(i).Tag.Get("column")

			if tag == "" {
				continue
			}

			v, ok := vMap[tag]

			if !ok {
				continue
			}

			vElem.Field(i).Set(reflect.ValueOf(cast.Kind(vElem.Field(i).Type().Kind(), *v)))
		}

		items = append(items, record)
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
		t.Fatalf("Something went wrong when trying to create table: %v", err)
	}

	_, err = db.Exec(`
		INSERT INTO "msisdns" ("msisdn", "name", "province", "number_of_children", "agreed_terms") VALUES ('253846568785', 'Paulo Maculuve', 'Maputo', '2', '1');
		INSERT INTO "msisdns" ("msisdn", "name", "province", "number_of_children", "agreed_terms") VALUES ('258843127837', 'Comfy', 'Maputo', '1', '1');
	`)

	if err != nil {
		t.Fatalf("Something went wrong when trying insert record: %v", err)
	}

	rows, err := db.Query("SELECT * FROM msisdns ORDER BY id ASC")

	if err != nil {
		t.Fatalf("Something went wrong when trying to get records: %v", err)
	}

	msisdns, err := RowsScan(rows, Msisdn{})

	if err != nil {
		t.Fatalf("Something when wrong when trying to scan rows from database: %v", err)
	}

	if len(msisdns) != 2 {
		t.Fatalf("Expected msisdns to have total of (%d) items but got (%d)", 2, len(msisdns))
	}

	msisdn := msisdns[0]

	if msisdn.ID != 1 {
		t.Fatalf("Expected msisdns first record id to be (%d) but got (%d)", 1, msisdn.ID)
	}

	if msisdn.Msisdn != "253846568785" {
		t.Fatalf("Expected msisdns first record msisdn to be (%s) but got (%s)", "253846568785", msisdn.Msisdn)
	}

	if msisdn.AgreedTerms != true {
		t.Fatalf("Expected msisdns first record agreed terms to be (%v) but got (%v)", true, msisdn.AgreedTerms)
	}

	db.Close()
}
