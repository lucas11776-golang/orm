package statements

import (
	"fmt"
	"orm"
	"strings"
)

type Select struct {
	Select orm.Select
	Values []interface{}
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

	default:
		return "", fmt.Errorf("Unsupported select type (%v)", v)
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
		"FROM"}, "\r\n",
	), nil
}
