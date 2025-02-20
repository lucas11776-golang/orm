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
func (ctx *Select) Statement() (string, error) {
	if len(ctx.Select) == 0 {
		ctx.Select = append(ctx.Select, "*")
	}

	_select := []string{}

	for _, v := range ctx.Select {
		switch v.(type) {
		case string:
			_select = append(_select, v.(string))
			break
		default:
			return "", fmt.Errorf("Unsupported select type (%v)", v)
		}
	}

	return strings.Join([]string{
		"SELECT", orm.SPACE + strings.Join(_select, ", "), "FROM"}, "\r\n",
	), nil
}
