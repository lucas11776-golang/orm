package orm

const DefaultDatabaseName = "default"

var DATABASES = map[string]Database{}

type Database interface {
	Query() Query
	Database() interface{}
}
