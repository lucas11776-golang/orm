package orm

const DefaultDatabaseName = "default"

type DB map[string]Database

type Model struct {
}

type Database interface {
	Query() Query
	Database() interface{}
}

var Databases = DB{}
