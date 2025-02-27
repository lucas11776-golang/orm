package statements

import "orm"

type Update struct {
	Where  []interface{}
	Values orm.Values
}

// comment
func (ctx *Update) Statement() (string, error) {
	return "", nil
}
