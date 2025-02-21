package statements

import (
	"fmt"
	"orm"
	"reflect"
	"strings"
)

type Where struct {
	Where  []interface{}
	Values []interface{}
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
		return "", fmt.Errorf("Where operator error")
	}

	var operator string
	var value interface{}

	for k, v := range where {
		operator = strings.ToUpper(k)
		value = v
	}

	switch operator {
	case "<", "<=", ">", ">=", "=", "!=", "NOT", "IS", "IS NOT":
		ctx.Values = append(ctx.Values, value)

		return strings.Join([]string{operator, "?"}, " "), nil

	case "LIKE":
		ctx.Values = append(ctx.Values, value)

		return strings.Join([]string{operator, "\"%?%\""}, " "), nil

	case "BETWEEN":
		v := castArray(value)

		if len(v) != 2 {
			return "", fmt.Errorf("Where between operator must be array of 2 values: %v", value)
		}

		ctx.Values = append(ctx.Values, v...)

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
			_where = append(_where, strings.Join([]string{k, "?"}, " = "))

			ctx.Values = append(ctx.Values, v)
			break

		case orm.Where:
			operator, err := ctx.operator(v.(orm.Where))

			if err != nil {
				return "", err
			}

			_where = append(_where, strings.Join([]string{k, operator}, " "))
			break

		default:
			return "", fmt.Errorf("Where value is current not support: (%v)", v)
		}
	}

	return strings.Join(_where, ""), nil
}

// comment
func (ctx *Where) whereList(where []interface{}) (string, error) {
	_where := []string{}

	for _, v := range where {
		switch v.(type) {
		case string:
			_where = append(_where, strings.Join([]string{SPACE, v.(string)}, ""))
			break

		case orm.Where:
			w, err := ctx.where(v.(orm.Where))

			if err != nil {
				return "", err
			}

			_where = append(_where, strings.Join([]string{SPACE, w}, ""))
			break

		case *WhereGroupQueryBuilder:
			w, err := ctx.whereList(v.(*WhereGroupQueryBuilder).Group)

			if err != nil {
				return "", err
			}

			_where = append(_where, strings.Join([]string{SPACE + "(", SPACE + w, SPACE + ")"}, "\r\n"))
			break

		default:
			return "", fmt.Errorf("Where query value is not supported: %v", v)
		}
	}

	return strings.Join(_where, "\r\n"), nil
}

// Comment
func (ctx *Where) Statement() (string, error) {
	if len(ctx.Where) == 0 {
		return "", nil
	}

	return ctx.whereList(ctx.Where)
}
