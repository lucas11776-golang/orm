package statements

import (
	"fmt"
	"strings"

	"github.com/lucas11776-golang/orm"
)

type Join struct {
	Table  string
	Join   orm.Joins
	values []interface{}
}

type JoinGroupQueryBuilder struct {
	Joins []interface{}
}

// Comment
func (ctx *Join) rawValue(key string, raw *orm.RawValue) string {
	ctx.values = append(ctx.values, raw.Value)

	return strings.Join([]string{SafeKey(key), "?"}, " = ")
}

// comment
func (ctx *Join) whereOperator(key string, where orm.Where) (string, error) {
	if len(where) != 1 {
		return "", fmt.Errorf("join where must have one map key value pair: %v", where)
	}

	keys := make([]string, 0, len(where))

	for k := range where {
		keys = append(keys, k)
	}

	k := keys[0]
	v := where[k]

	switch v.(type) {
	case *orm.RawValue:
		ctx.values = append(ctx.values, v.(*orm.RawValue).Value)

		return strings.Join([]string{SafeKey(key), "?"}, fmt.Sprintf(" %s ", strings.Trim(k, " "))), nil

	case string:
		return strings.Join([]string{SafeKey(key), SafeKey(v.(string))}, fmt.Sprintf(" %s ", strings.Trim(k, " "))), nil

	default:
		return "", fmt.Errorf("type value of the join is not supported: %v", v)
	}
}

// Comment
// func (ctx *Join) joinOperator(join orm.Join) (string, error) {

// }

// Comment
func (ctx *Join) where(w []interface{}) (string, error) {
	query := []string{}

	for _, v := range w {
		switch v.(type) {
		case orm.Join:

			for k, v := range v.(orm.Join) {
				switch v.(type) {
				case *orm.Values:
					query = append(query, strings.Join([]string{SafeKey(k), "?"}, " = "))

					ctx.values = append(ctx.values, v.(*orm.RawValue).Value)

					break

				case orm.Where:
					where, err := ctx.whereOperator(k, v.(orm.Where))

					if err != nil {
						return "", err
					}

					query = append(query, where)

					break

				case string:
					query = append(query, strings.Join([]string{SafeKey(k), SafeKey(v.(string))}, " = "))

					break

				case *orm.RawValue:
					query = append(query, ctx.rawValue(k, v.(*orm.RawValue)))

					break

				default:

					fmt.Println("V", v)

					return "", fmt.Errorf("the value of join where is not supported: %v", v)
				}
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

// Comment
func (ctx *Join) Values() []interface{} {
	return ctx.values
}
