package statements

import (
	"fmt"
	"orm"
	"strings"
)

type Where struct {
	Where  []orm.Where
	Values []interface{}
}

// comment
func (ctx *Where) operator(operator orm.WhereMatch) (string, error) {
	if len(operator) != 1 {
		return "", fmt.Errorf("W operator")
	}

	keys := make([]string, 0, len(operator))

	for k, v := range operator {
		keys = append(keys, k)

		ctx.Values = append(ctx.Values, v)
	}

	return keys[0], nil
}

// Comment
func (ctx *Where) where(where orm.WhereMatch) (string, error) {
	if len(where) != 1 {
		return "", nil
	}

	_where := []string{}

	for k, v := range where {
		switch v.(type) {
		case int, int8, int16, int32, int64, string, float32, float64, []byte:
			_where = append(_where, orm.SPACE+strings.Join([]string{k, "?"}, " = "))

			ctx.Values = append(ctx.Values, v)
			break
		case orm.WhereMatch:
			operator, err := ctx.operator(v.(orm.WhereMatch))

			if err != nil {
				return "", err
			}

			_where = append(_where, orm.SPACE+strings.Join([]string{k, "?"}, fmt.Sprintf(" %s ", operator)))

			break
		default:
			return "", fmt.Errorf("where")
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
			where = append(where, orm.SPACE+v.(string))
			break
		case orm.WhereMatch:
			w, err := ctx.where(v.(orm.WhereMatch))

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
