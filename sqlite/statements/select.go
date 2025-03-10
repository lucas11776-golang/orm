package statements

import (
	"fmt"
	"strings"

	"github.com/lucas11776-golang/orm"
)

type Select struct {
	Table  string
	Select orm.Select
	values []interface{}
}

// Comment
func (ctx *Select) operator(v interface{}) (string, error) {
	switch v.(type) {
	case string:
		if v != "*" {
			v = SafeKey(v.(string))
		}

		return v.(string), nil

	case orm.AS:
		return strings.Join([]string{SafeKey(v.(orm.AS)[0]), "AS", SafeKey(v.(orm.AS)[1])}, " "), nil

	case orm.SUM:
		return strings.Join([]string{
			strings.Join([]string{"SUM(", SafeKey(v.(orm.SUM)[0]), ")"}, ""), "AS", SafeKey(v.(orm.SUM)[1]),
		}, " "), nil

	case orm.COUNT:
		var field string

		if v.(orm.COUNT)[0] == "*" {
			field = "*"
		} else {
			field = SafeKey(v.(orm.COUNT)[0])
		}

		return strings.Join([]string{
			strings.Join([]string{"COUNT(", field, ")"}, ""),
			"AS",
			SafeKey(v.(orm.COUNT)[1]),
		}, " "), nil

	default:
		return "", fmt.Errorf("unsupported select type (%v)", v)
	}
}

// Comment
func (ctx *Select) Statement() (string, error) {
	if len(ctx.Select) == 0 {
		ctx.Select = append(ctx.Select, "*")
	}

	slt := []string{}

	for _, v := range ctx.Select {
		field, err := ctx.operator(v)

		if err != nil {
			return "", err
		}

		slt = append(slt, field)
	}

	return strings.Join([]string{
		"SELECT",
		SPACE + strings.Join(slt, ", "),
		"FROM",
		SPACE + SafeKey(ctx.Table)}, "\r\n",
	), nil
}

// Comment
func (ctx *Select) Values() []interface{} {
	return ctx.values
}
