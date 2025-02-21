package statements

import (
	"fmt"
	"strings"
)

const SPACE = " "

type Statement interface {
	Statement() (string, error)
}

// Comment
func SafeKey(v string) string {
	vs := strings.Split(strings.ReplaceAll(v, "`", ""), ".")

	for i, v := range vs {
		vs[i] = fmt.Sprintf("`%s`", v)
	}

	return strings.Join(vs, ".")
}
