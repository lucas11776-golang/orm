package orm

import (
	"log"
	"reflect"
	"strings"

	"github.com/lucas11776-golang/orm/types"
	str "github.com/lucas11776-golang/orm/utils/strings"
)

const DefaultDatabaseName = "default"

type db map[string]Database

type Models []interface{}

// type Result map[string]interface{}

// type Results []Result

type options struct {
	connection string
	table      string
	key        string
}

type Database interface {
	Query(statement *Statement) (types.Results, error)
	Count(statement *Statement) (int64, error)
	Insert(statement *Statement) (types.Result, error)
	Update(statement *Statement) error
	Delete(Statement *Statement) error
	Database() interface{}
	Migration() Migration
}

var DB = db{}

// Comment
func (ctx *db) Database(name string) Database {
	db, ok := DB[name]

	if !ok {
		return db
	}

	return db
}

// Comment
func TableName(model interface{}) string {
	vType := reflect.ValueOf(model)

	if vType.Kind() != reflect.Struct {
		log.Fatalf("Model is not a struct: %v", model)
	}

	return str.Plural(strings.ToLower(vType.Type().Name()))
}

// Comment
func getOptions(model interface{}) *options {
	opt := &options{
		connection: DefaultDatabaseName,
		table:      TableName(model),
	}

	vType := reflect.ValueOf(model).Type()

	for i := 0; i < vType.NumField(); i++ {
		tag := vType.Field(i).Tag

		if tag.Get("connection") != "" {
			opt.connection = tag.Get("connection")
		}

		if tag.Get("table") != "" {
			opt.table = tag.Get("table")
		}

		if strings.ToUpper(tag.Get("type")) == "PRIMARY_KEY" {
			opt.key = tag.Get("column")
		}
	}

	if opt.connection == "" {
		opt.connection = DefaultDatabaseName
	}

	return opt
}

// Comment
func Model[T any](model T) QueryBuilder[T] {
	if reflect.ValueOf(model).Type().Kind() != reflect.Struct {
		log.Fatalf("Model is not type of struct: %v", model)
	}

	options := getOptions(model)
	database, ok := DB[options.connection]

	if !ok {
		log.Fatalf("Connection %s does not exist", options.connection)
	}

	return &QueryStatement[T]{
		Model:      model,
		Database:   database,
		Connection: options.connection,
		Statement: &Statement{
			Table: options.table,
		},
	}
}

// Comment
func (ctx *db) Add(name string, database Database) {
	DB[name] = database
}

// Comment
func (ctx *db) Remove(name string) {
	delete(DB, name)
}
