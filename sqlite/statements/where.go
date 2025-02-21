package statements

import (
	"fmt"
	"orm"
	"strings"
)

type Where struct {
	Where  []interface{}
	Values []interface{}
}

// comment
func (ctx *Where) operator(operator orm.Where) (string, error) {
	if len(operator) != 1 {
		return "", fmt.Errorf("Where operator error")
	}

	keys := make([]string, 0, len(operator))

	for k, v := range operator {
		keys = append(keys, k)

		ctx.Values = append(ctx.Values, v)
	}

	return keys[0], nil
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
			_where = append(_where, SPACE+strings.Join([]string{k, "?"}, " = "))

			ctx.Values = append(ctx.Values, v)
			break
		case orm.Where:
			operator, err := ctx.operator(v.(orm.Where))

			if err != nil {
				return "", err
			}

			_where = append(_where, SPACE+strings.Join([]string{k, "?"}, fmt.Sprintf(" %s ", operator)))

			break
		default:
			return "", fmt.Errorf("Where value is current not support: (%v)", v)
		}
	}

	return strings.Join(_where, ""), nil
}

// Comment
func (ctx *Where) Statement() (string, error) {
	if len(ctx.Where) == 0 {
		return "", nil
	}

	where := []string{}

	for _, v := range ctx.Where {
		switch v.(type) {
		case string:
			where = append(where, SPACE+v.(string))
			break
		case orm.Where:
			w, err := ctx.where(v.(orm.Where))

			if err != nil {
				return "", err
			}

			where = append(where, w)

			break
		default:
		}
	}

	return strings.Join(where, "\r\n"), nil
}
