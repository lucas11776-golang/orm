package statements

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/lucas11776-golang/orm"
)

type Where struct {
	Where  []interface{}
	values []interface{}
}

// TODO: move to orm base
type WhereGroupQueryBuilder struct {
	Group []interface{}
}

// Comment
func (ctx *WhereGroupQueryBuilder) Where(column string, operator string, value interface{}) orm.WhereGroupBuilder {
	if len(ctx.Group) != 0 {
		return ctx.AndWhere(column, operator, value)
	}

	ctx.Group = append(ctx.Group, &orm.Where{
		Key:      column,
		Operator: operator,
		Value:    value,
	})

	return ctx
}

// Comment
func (ctx *WhereGroupQueryBuilder) AndWhere(column string, operator string, value interface{}) orm.WhereGroupBuilder {
	if len(ctx.Group) == 0 {
		return ctx.Where(column, operator, value)
	}

	ctx.Group = append(ctx.Group, "AND", &orm.Where{
		Key:      column,
		Operator: operator,
		Value:    value,
	})

	return ctx
}

// Comment
func (ctx *WhereGroupQueryBuilder) OrWhere(column string, operator string, value interface{}) orm.WhereGroupBuilder {
	if len(ctx.Group) == 0 {
		return ctx.Where(column, operator, value)
	}

	ctx.Group = append(ctx.Group, "OR", &orm.Where{
		Key:      column,
		Operator: operator,
		Value:    value,
	})

	return ctx
}

// Comment
func castArray[T any](array T) []interface{} {
	v := reflect.ValueOf(array)
	arr := []interface{}{}

	for i := 0; i < v.Len(); i++ {
		arr = append(arr, v.Index(i).Interface())
	}

	return arr
}

// Comment
func (ctx *Where) where(where *orm.Where) (string, error) {
	switch strings.ToUpper(where.Operator) {
	case orm.EQUALS, orm.NOT_EQUALS, orm.NOT, orm.IS_NOT, orm.LESS_THEN,
		orm.LESS_THEN_EQUALS, orm.GREATER_THEN, orm.GREATER_THEN_EQUALS:
		ctx.values = append(ctx.values, where.Value)

		return strings.Join([]string{SafeKey(where.Key), "?"}, fmt.Sprintf(" %s ", where.Operator)), nil

	case "LIKE":
		ctx.values = append(ctx.values, where.Value)

		return strings.Join([]string{SafeKey(where.Key), "\"%?%\""}, fmt.Sprintf(" %s ", where.Operator)), nil

	case "BETWEEN":
		v := castArray(where.Value)

		if len(v) != 2 {
			return "", fmt.Errorf("Where between operator must be array of 2 values: %v", where.Value)
		}

		ctx.values = append(ctx.values, v[0], v[1])

		return strings.Join([]string{SafeKey(where.Key), strings.Join([]string{"?", "?"}, " AND ")}, fmt.Sprintf(" %s ", where.Operator)), nil

	default:
		return "", fmt.Errorf("Where operate is not supported: %v", where.Operator)
	}
}

// comment
func (ctx *Where) list(where []interface{}) (string, error) {
	stmt := []string{}

	for _, w := range where {
		switch w.(type) {
		case string:
			stmt = append(stmt, strings.Join([]string{SPACE, w.(string)}, ""))

		case *orm.Where:
			w, err := ctx.where(w.(*orm.Where))

			if err != nil {
				return "", err
			}

			stmt = append(stmt, strings.Join([]string{SPACE, w}, ""))

		case *WhereGroupQueryBuilder:
			query, err := ctx.list(w.(*WhereGroupQueryBuilder).Group)

			if err != nil {
				return "", err
			}

			stmt = append(stmt, strings.Join([]string{SPACE + "(", SPACE + query, SPACE + ")"}, "\r\n"))

		default:
			return "", fmt.Errorf("Where operate is not supported: %v", w)

		}
	}

	return strings.Join(stmt, "\r\n"), nil
}

// Comment
func (ctx *Where) Statement() (string, error) {
	if len(ctx.Where) == 0 {
		return "", nil
	}

	stmt, err := ctx.list(ctx.Where)

	if err != nil {
		return "", err
	}

	return strings.Join([]string{"WHERE", stmt}, "\r\n"), nil
}

// Comment
func (ctx *Where) Values() []interface{} {
	return ctx.values
}
