package statements

import (
	"fmt"
	"strings"
)

type Delete struct {
	Table  string
	Where  []interface{}
	values []interface{}
}

// Comment
func (ctx *Delete) generateWhere() (string, error) {
	statement := &Where{Where: ctx.Where}

	where, err := statement.Statement()

	if err != nil {
		return "", nil
	}

	ctx.values = append(ctx.values, statement.Values()...)

	return where, nil
}

// Comment
func (ctx *Delete) Statement() (string, error) {
	statement := []string{
		"DELETE FROM",
		SPACE + SafeKey(ctx.Table),
	}

	where, err := ctx.generateWhere()

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s;", strings.Join(append(statement, where), "\r\n")), nil
}

// Comment
func (ctx *Delete) Values() []interface{} {
	return ctx.values
}
