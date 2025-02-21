package statements

import "orm"

type Join struct {
	Table string
	Join  orm.Joins
}

// Comment
func (ctx *Join) Statement() (string, error) {
	return "", nil
}
