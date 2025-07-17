package statements

import (
	"sort"
	"strings"

	"github.com/lucas11776-golang/orm"
)

type Update struct {
	Table        string
	Where        []interface{}
	UpdateValues orm.Values
	values       []interface{}
}

// Comment
func (ctx *Update) keys(values orm.Values) (keys []string) {
	for key := range values {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	return keys
}

// Comment
func (ctx *Update) generateSetStatement() string {
	set := []string{}

	for _, key := range ctx.keys(ctx.UpdateValues) {
		set = append(set, strings.Join([]string{SafeKey(key), "?"}, " = "))
		ctx.values = append(ctx.values, ctx.UpdateValues[key])
	}

	return SPACE + strings.Join(set, ", ")
}

// Comment
func (ctx *Update) generateWhereStatement() (string, error) {
	where := &Where{Where: ctx.Where}

	statement, err := where.Statement()

	if err != nil {
		return "", err
	}

	ctx.values = append(ctx.values, where.Values()...)

	return statement, err
}

// comment
func (ctx *Update) Statement() (string, error) {
	if ctx.UpdateValues == nil {
		return "", ErrInsertEmptyValues
	}

	if len(ctx.UpdateValues) == 0 {
		return "", ErrInsertEmptyValues
	}

	stmt := []string{
		"UPDATE",
		SPACE + SafeKey(ctx.Table),
		"SET",
		ctx.generateSetStatement(),
	}

	where, err := ctx.generateWhereStatement()

	if err != nil {
		return "", err
	}

	stmt = append(stmt, where)

	return strings.Join(stmt, "\r\n"), nil
}

// Comment
func (ctx *Update) Values() []interface{} {
	return ctx.values
}
