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

type WhereGroupQueryBuilder struct {
	Group []interface{}
}

// Comment
func (ctx *WhereGroupQueryBuilder) Where(column string, operator string, value interface{}) orm.WhereGroupBuilder {
	if len(ctx.Group) != 0 {
		return ctx.AndWhere(column, operator, value)
	}

	ctx.Group = append(ctx.Group, orm.Where{column: orm.Where{operator: value}})

	return ctx
}

// Comment
func (ctx *WhereGroupQueryBuilder) AndWhere(column string, operator string, value interface{}) orm.WhereGroupBuilder {
	if len(ctx.Group) == 0 {
		return ctx.Where(column, operator, value)
	}

	ctx.Group = append(ctx.Group, "AND", orm.Where{column: orm.Where{operator: value}})

	return ctx
}

// Comment
func (ctx *WhereGroupQueryBuilder) OrWhere(column string, operator string, value interface{}) orm.WhereGroupBuilder {
	if len(ctx.Group) == 0 {
		return ctx.Where(column, operator, value)
	}

	ctx.Group = append(ctx.Group, "OR", orm.Where{column: orm.Where{operator: value}})

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

// comment
func (ctx *Where) operator(where orm.Where) (string, error) {
	if len(where) != 1 {
		return "", fmt.Errorf("invalid where query must has one key value: %v", where)
	}

	var operator string
	var value interface{}

	for k, v := range where {
		operator = strings.ToUpper(k)
		value = v
	}

	switch operator {
	case orm.EQUALS, orm.NOT_EQUALS, orm.NOT, orm.IS_NOT, orm.LESS_THEN,
		orm.LESS_THEN_EQUALS, orm.GREATER_THEN, orm.GREATER_THEN_EQUALS:
		ctx.values = append(ctx.values, value)

		return strings.Join([]string{operator, "?"}, " "), nil

	case "LIKE":
		ctx.values = append(ctx.values, value)

		return strings.Join([]string{operator, "\"%?%\""}, " "), nil

	case "BETWEEN":
		v := castArray(value)

		if len(v) != 2 {
			return "", fmt.Errorf("Where between operator must be array of 2 values: %v", value)
		}

		ctx.values = append(ctx.values, v...)

		return strings.Join([]string{"BETWEEN", "?", "AND", "?"}, " "), nil

	default:
		return "", fmt.Errorf("Where operate is not supported: %v", operator)
	}
}

// Comment
func (ctx *Where) where(where orm.Where) (string, error) {
	if len(where) > 1 {
		return "", fmt.Errorf("Where statement only support one value in map: %v", where)
	}

	_where := []string{}

	for k, v := range where {
		switch v.(type) {
		case int, int8, int16, int32, int64, string, float32, float64:
			_where = append(_where, strings.Join([]string{SafeKey(k), "?"}, " = "))

			ctx.values = append(ctx.values, v)
			break

		case orm.Where:
			operator, err := ctx.operator(v.(orm.Where))

			if err != nil {
				return "", err
			}

			_where = append(_where, strings.Join([]string{SafeKey(k), operator}, " "))
			break

		default:
			return "", fmt.Errorf("Where value is current not support: (%v)", v)
		}
	}

	return strings.Join(_where, ""), nil
}

// Comment
func (ctx *Where) whereType(v interface{}) (string, error) {
	switch v.(type) {
	case string:
		return strings.Join([]string{SPACE, v.(string)}, ""), nil

	case orm.Where:
		w, err := ctx.where(v.(orm.Where))

		if err != nil {
			return "", err
		}

		return strings.Join([]string{SPACE, w}, ""), nil

	case *WhereGroupQueryBuilder:
		w, err := ctx.whereList(v.(*WhereGroupQueryBuilder).Group)

		if err != nil {
			return "", err
		}

		return strings.Join([]string{SPACE + "(", SPACE + w, SPACE + ")"}, "\r\n"), nil

	default:
		return "", fmt.Errorf("Where query value is not supported: %v", v)
	}
}

// comment
func (ctx *Where) whereList(where []interface{}) (string, error) {
	stmt := []string{}

	for _, v := range where {
		vR, err := ctx.whereType(v)

		if err != nil {
			return "", err
		}

		stmt = append(stmt, vR)
	}

	return strings.Join(stmt, "\r\n"), nil
}

// Comment
func (ctx *Where) Statement() (string, error) {
	if len(ctx.Where) == 0 {
		return "", nil
	}

	stmt, err := ctx.whereList(ctx.Where)

	if err != nil {
		return "", err
	}

	return strings.Join([]string{"WHERE", stmt}, "\r\n"), nil
}

// Comment
func (ctx *Where) Values() []interface{} {
	return ctx.values
}
