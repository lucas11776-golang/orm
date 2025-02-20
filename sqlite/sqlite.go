package sqlite

import "orm"

// TODO thinking of using turso for sqlite

type Sqlite struct{}

// Comment
func (ctx *Sqlite) Query() orm.Query {
	return &Query{}
}

type Database interface {
	Query() Query
	Database() interface{}
}
