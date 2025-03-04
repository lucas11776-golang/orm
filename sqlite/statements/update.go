package statements

import (
	"orm"
	"strings"
)

type Update struct {
	Table  string
	Where  []interface{}
	Update orm.Values
	Values []interface{}
}

// comment
func (ctx *Update) Statement() (string, error) {
	stmt := []string{
		"UPDATE",
		SPACE + SafeKey(ctx.Table),
		"SET",
	}

	set := []string{}

	for k, v := range ctx.Update {
		set = append(set, strings.Join([]string{SafeKey(k), "?"}, " = "))

		ctx.Values = append(ctx.Values, v)
	}

	stmt = append(stmt, SPACE+strings.Join(set, ", "))

	where := &Where{
		Where: ctx.Where,
	}

	stmtWhere, err := where.Statement()

	if err != nil {
		return "", err
	}

	stmt = append(stmt, stmtWhere)

	return strings.Join(stmt, "\r\n"), nil
}

/**
UPDATE table_name
SET column1 = value1, column2 = value2, ...
WHERE condition;
**/
