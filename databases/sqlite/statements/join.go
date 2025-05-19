package statements

import (
	"fmt"
	"strings"

	"github.com/lucas11776-golang/orm"
)

type Join struct {
	Table  string
	Join   orm.Joins
	values []interface{}
}

type JoinGroupQueryBuilder struct {
	Joins []interface{}
}

// Comment
func (ctx *Join) value(v interface{}) string {
	switch v.(type) {
	case *orm.RawValue:
		ctx.values = append(ctx.values, v.(*orm.RawValue).Value)

		return "?"

	case string:
		return SafeKey(v.(string))

	default:
		ctx.values = append(ctx.values, v)

		return "?"
	}
}

// Comment
func (ctx *Join) where(w *orm.Where) (string, error) {
	return strings.Join([]string{SafeKey(w.Key), ctx.value(w.Value)}, fmt.Sprintf(" %s ", w.Operator)), nil
}

// Comment
func (ctx *Join) list(where []interface{}) (string, error) {
	queries := []string{}

	for _, w := range where {
		switch w.(type) {
		case *orm.Where:
			query, err := ctx.where(w.(*orm.Where))

			if err != nil {
				return "", err
			}

			queries = append(queries, query)

		case *JoinGroupQueryBuilder:
			query, err := ctx.list(w.(*JoinGroupQueryBuilder).Joins)

			if err != nil {
				return "", err
			}

			queries = append(queries, strings.Join([]string{"(", query, ")"}, ""))

		case string:
			operator := strings.ToUpper(w.(string))

			if operator != "AND" && operator != "OR" {
				return "", fmt.Errorf("where query join must be (AND, OR) not: %s", w)
			}

			queries = append(queries, operator)

		default:
			return "", fmt.Errorf("join does not support type %v", w)

		}
	}

	return strings.Join(queries, " "), nil
}

// Comment
func (ctx *Join) Statement() (string, error) {
	joins := []string{}

	for _, j := range ctx.Join {
		w, err := ctx.list(j.Operators)

		if err != nil {
			return "", err
		}

		joins = append(joins, strings.Join([]string{"LEFT JOIN", SafeKey(j.Table), "ON", w}, " "))
	}

	return strings.Join(joins, "\r\n"), nil
}

// Comment
func (ctx *Join) Values() []interface{} {
	return ctx.values
}
