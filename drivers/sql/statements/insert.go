package statements

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/lucas11776-golang/orm"
)

var (
	ErrInsertEmptyValues error = errors.New("insert operation must have values")
)

type Insert struct {
	Table        string
	InsertValues orm.Values
	values       []interface{}
}

// Comment
func (ctx *Insert) keys() (keys []string) {
	for key := range ctx.InsertValues {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	return keys
}

// Comment
func (ctx *Insert) Statement() (string, error) {
	if ctx.InsertValues == nil {
		return "", ErrInsertEmptyValues
	}

	if len(ctx.InsertValues) == 0 {
		return "", ErrInsertEmptyValues
	}

	keys := ctx.keys()
	values := []string{}

	for i, key := range keys {
		ctx.values = append(ctx.values, ctx.InsertValues[key])
		values = append(values, "?")
		keys[i] = SafeKey(key)
	}

	return fmt.Sprintf(
		"INSERT INTO %s(%s) VALUES(%s);",
		SafeKey(ctx.Table),
		strings.Join(keys, ", "),
		strings.Join(values, ", "),
	), nil
}

// Comment
func (ctx *Insert) Values() []interface{} {
	return ctx.values
}
