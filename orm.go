package orm

import (
	"log"
	"reflect"
	"strings"

	"github.com/lucas11776-golang/orm/types"
	str "github.com/lucas11776-golang/orm/utils/strings"
)

const (
	DefaultDatabaseName = "default"
)

type db map[string]Database

type modelOptions struct {
	Connection string
	Table      string
}

type Database interface {
	Query(statement *Statement) (types.Results, error)
	Count(statement *Statement) (int64, error)
	Insert(statement *Statement) (types.Result, error)
	Update(statement *Statement) error
	Delete(Statement *Statement) error
	Database() interface{}
	Migration() Migration
	Close() error
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
func (ctx *db) Add(name string, database Database) {
	if name == "" {
		DB[DefaultDatabaseName] = database

		return
	}

	DB[name] = database
}

// Comment
func (ctx *db) Remove(name string) {
	delete(DB, name)
}

// Comment
func generateModelTableName(model interface{}) string {
	vType := reflect.ValueOf(model)

	if vType.Kind() != reflect.Struct {
		log.Fatalf("Model is not a struct: %v", model)
	}

	return str.Plural(strings.ToLower(vType.Type().Name()))
}

// Comment
func getModelOptions(model interface{}) *modelOptions {
	opt := &modelOptions{
		Connection: DefaultDatabaseName,
		Table:      generateModelTableName(model),
	}

	vType := reflect.ValueOf(model).Type()

	for i := 0; i < vType.NumField(); i++ {
		tag := vType.Field(i).Tag

		if tag.Get("connection") != "" {
			opt.Connection = tag.Get("connection")
		}

		if tag.Get("table") != "" {
			opt.Table = tag.Get("table")
		}
	}

	return opt
}

// Comment
func Model[T any](model T) QueryBuilder[T] {
	if reflect.ValueOf(model).Type().Kind() != reflect.Struct {
		log.Fatalf("model is not type of struct: %v", model)
	}

	options := getModelOptions(model)
	database, ok := DB[options.Connection]

	if !ok {
		log.Fatalf("connection %s does not exist", options.Connection)
	}

	return &QueryStatement[T]{
		Model:      model,
		Database:   database,
		Connection: options.Connection,
		Statement: &Statement{
			Table: options.Table,
		},
	}
}
