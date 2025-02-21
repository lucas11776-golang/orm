package statements

import (
	"fmt"
	"orm"
	"strings"
)

type Join struct {
	Table  string
	Join   orm.Joins
	Values []interface{}
}

type JoinGroupQueryBuilder struct {
	Joins []interface{}
}

// Comment
func (ctx *JoinGroupQueryBuilder) Join(j orm.Join) orm.JoinGroupBuilder {
	return ctx
}

// Comment
func (ctx *JoinGroupQueryBuilder) And(j orm.Join) orm.JoinGroupBuilder {
	return ctx
}

// Comment
func (ctx *JoinGroupQueryBuilder) Or(j orm.Join) orm.JoinGroupBuilder {
	return ctx
}

// Comment
func (ctx *JoinGroupQueryBuilder) Group(j orm.JoinGroup) orm.JoinGroupBuilder {
	return ctx
}

// Comment
func (ctx *Join) group() (string, error) {
	return "", nil
}

// Comment
func (ctx *Join) where(w []interface{}) (string, error) {
	query := []string{}

	for _, v := range w {
		switch v.(type) {
		case orm.Join:
			for k, v := range v.(orm.Join) {
				raw, ok := v.(*orm.Raw)

				if ok {
					query = append(query, strings.Join([]string{SafeKey(k), "?"}, " = "))
					ctx.Values = append(ctx.Values, raw)
					continue
				}

				_v, ok := v.(string)

				if ok {

					query = append(query, strings.Join([]string{SafeKey(k), SafeKey(_v)}, " = "))
					continue
				}

				return "", fmt.Errorf("Join does not support value: %v", v)
			}
			break

		case string:
			v := strings.ToUpper(v.(string))

			if v != "OR" && v != "AND" {
				return "", fmt.Errorf("Join operators must be (AND,OR) not (%v)", v)
			}

			query = append(query, v)
			break

		case *JoinGroupQueryBuilder:
			w, err := ctx.where(v.(*JoinGroupQueryBuilder).Joins)

			if err != nil {
				return "", err
			}

			query = append(query, strings.Join([]string{"(", w, ")"}, ""))

			break

		default:
			fmt.Println(v)

			return "", fmt.Errorf("Join where does not support value: %v", v)
		}
	}

	return strings.Join(query, " "), nil
}

// Comment
func (ctx *Join) Statement() (string, error) {
	joins := []string{}

	for _, j := range ctx.Join {
		w, err := ctx.where(j.Where)

		if err != nil {
			return "", err
		}

		joins = append(joins, strings.Join([]string{"LEFT JOIN", SafeKey(j.Table), "ON", w}, " "))
	}

	return strings.Join(joins, "\r\n"), nil
}
