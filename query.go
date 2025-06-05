package orm

import (
	"reflect"

	"github.com/lucas11776-golang/orm/utils/cast"
)

type Order string

const (
	ASC  Order = "ASC"
	DESC Order = "DESC"
)

type Entity interface{}

type Values map[string]interface{}

type RawValue struct {
	Value interface{}
}

type JoinHolder struct {
	Table     string
	Operators []interface{}
}

type JoinGroupBuilder interface {
	Where(column1 string, operator string, column2 string) JoinGroupBuilder
	AndWhere(column1 string, operator string, column2 string) JoinGroupBuilder
	OrWhere(column1 string, operator string, column2 string) JoinGroupBuilder
}

type JoinBuilder struct {
	Operators []interface{}
}

// Comment
func (ctx *JoinBuilder) Where(column1 string, operator string, column2 string) *JoinBuilder {
	if len(ctx.Operators) >= 1 {
		return ctx.AndWhere(column1, operator, column2)
	}

	return ctx
}

// Comment
func (ctx *JoinBuilder) AndWhere(column1 string, operator string, column2 string) *JoinBuilder {
	return ctx
}

// Comment
func (ctx *JoinBuilder) OrWhere(column1 string, operator string, column2 string) *JoinBuilder {
	return ctx
}

type Select []interface{}
type Join map[string]interface{}
type Joins []*JoinHolder
type JoinGroup func(group *JoinBuilder)

type Where struct {
	Key      string
	Operator string
	Value    interface{}
}

type WhereGroup func(group WhereGroupBuilder)
type Limit int64
type Offset int64
type OrderBy [2]interface{}

type WhereGroupBuilder interface {
	Where(column string, operator string, value interface{}) WhereGroupBuilder
	AndWhere(column string, operator string, value interface{}) WhereGroupBuilder
	OrWhere(column string, operator string, value interface{}) WhereGroupBuilder
}

type Pagination[T any] struct {
	Total   int64 `json:"total"`
	Page    int64 `json:"page"`
	PerPage int64 `json:"per_page"`
	Items   []T   `json:"items"`
}

type Statement struct {
	Table      string
	Select     Select
	Joins      Joins
	Where      []interface{}
	Limit      int64
	Offset     int64
	OrderBy    OrderBy
	Values     Values
	PrimaryKey string
}

type QueryStatement[T any] struct {
	Model      T
	Database   Database
	Connection string
	*Statement
}

type WhereGroupQueryBuilder struct {
	Group []interface{}
}

type QueryBuilder[T any] interface {
	Select(s Select) QueryBuilder[T]
	Join(table string, column string, operator string, tableColumn string) QueryBuilder[T]
	JoinGroup(table string, builder func(group JoinGroupBuilder)) QueryBuilder[T]
	Where(column string, operator string, value interface{}) QueryBuilder[T]
	AndWhere(column string, operator string, value interface{}) QueryBuilder[T]
	OrWhere(column string, operator string, value interface{}) QueryBuilder[T]
	WhereGroup(group WhereGroup) QueryBuilder[T]
	AndWhereGroup(group WhereGroup) QueryBuilder[T]
	OrWhereGroup(group WhereGroup) QueryBuilder[T]
	Limit(l int64) QueryBuilder[T]
	Offset(o int64) QueryBuilder[T]
	OrderBy(column string, order Order) QueryBuilder[T]
	Count() (int64, error)
	First() (*T, error)
	Get() ([]*T, error)
	Paginate(perPage int64, page int64) (*Pagination[*T], error)
	Insert(values Values) (*T, error)
	Update(values Values) error
}

// Comment
func Raw(v interface{}) *RawValue {
	return &RawValue{Value: v}
}

// Comment
func (ctx *QueryStatement[T]) Select(s Select) QueryBuilder[T] {
	ctx.Statement.Select = s

	return ctx
}

// Comment
func (ctx *QueryStatement[T]) Join(table string, column string, operator string, value string) QueryBuilder[T] {
	ctx.Statement.Joins = append(ctx.Statement.Joins, &JoinHolder{
		Table: table,
		Operators: []interface{}{&Where{
			Key:      column,
			Operator: operator,
			Value:    value,
		}},
	})

	return ctx
}

// Comment
func (ctx *QueryStatement[T]) JoinGroup(table string, builder func(group JoinGroupBuilder)) QueryBuilder[T] {
	return ctx
}

// Comment
func (ctx *QueryStatement[T]) Where(column string, operator string, value interface{}) QueryBuilder[T] {
	ctx.Statement.Where = append(ctx.Statement.Where, &Where{
		Key:      column,
		Operator: operator,
		Value:    value,
	})

	return ctx
}

// Comment
func (ctx *QueryStatement[T]) AndWhere(column string, operator string, value interface{}) QueryBuilder[T] {
	if len(ctx.Statement.Where) == 0 {
		return ctx.Where(column, operator, value)
	}

	ctx.Statement.Where = append(ctx.Statement.Where, "AND", &Where{
		Key:      column,
		Operator: operator,
		Value:    value,
	})

	return ctx
}

// Comment
func (ctx *QueryStatement[T]) OrWhere(column string, operator string, value interface{}) QueryBuilder[T] {
	if len(ctx.Statement.Where) == 0 {
		return ctx.Where(column, operator, value)
	}

	ctx.Statement.Where = append(ctx.Statement.Where, "OR", &Where{
		Key:      column,
		Operator: operator,
		Value:    value,
	})

	return ctx
}

// Comment
func (ctx *QueryStatement[T]) WhereGroup(group WhereGroup) QueryBuilder[T] {
	return ctx
}

// Comment
func (ctx *QueryStatement[T]) AndWhereGroup(group WhereGroup) QueryBuilder[T] {
	return ctx
}

// Comment
func (ctx *QueryStatement[T]) OrWhereGroup(group WhereGroup) QueryBuilder[T] {
	return ctx
}

// Comment
func (ctx *QueryStatement[T]) Limit(l int64) QueryBuilder[T] {
	ctx.Statement.Limit = l

	return ctx
}

// Comment
func (ctx *QueryStatement[T]) Offset(o int64) QueryBuilder[T] {
	ctx.Statement.Offset = o

	return ctx
}

// Comment
func (ctx *QueryStatement[T]) OrderBy(column string, order Order) QueryBuilder[T] {
	ctx.Statement.OrderBy = OrderBy{column, order}

	return ctx
}

// Comment
func (ctx *QueryStatement[T]) Count() (int64, error) {
	result, err := ctx.Database.Count(ctx.Statement)

	if err != nil {
		return 0, err
	}

	return result, nil
}

// Comment
func (ctx *QueryStatement[T]) result(raw Result) *T {
	zValue := reflect.Zero(reflect.TypeOf(ctx.Model)).Interface().(T)
	zElem := reflect.ValueOf(&zValue).Elem()

	for i := 0; i < zElem.NumField(); i++ {
		col := zElem.Type().Field(i).Tag.Get("column")

		_, connection := zElem.Type().Field(i).Tag.Lookup("connection")

		if connection {
			zElem.Field(i).Set(reflect.ValueOf(ctx.Connection))
		}

		if col == "" {
			continue
		}

		v, ok := raw[col]

		if !ok || v == nil || v == "" {
			continue
		}

		zElem.Field(i).Set(reflect.ValueOf(cast.Kind(zElem.Type().Field(i).Type.Kind(), v)))
	}

	return &zValue
}

// Comment
func (ctx *QueryStatement[T]) results(raws Results) []*T {
	results := []*T{}

	for _, result := range raws {
		results = append(results, ctx.result(result))
	}

	return results
}

// Comment
func (ctx *QueryStatement[T]) First() (*T, error) {
	results, err := ctx.Database.Query(ctx.Statement)

	if err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return nil, nil
	}

	return ctx.results(results)[0], nil
}

// Comment
func (ctx *QueryStatement[T]) Get() ([]*T, error) {
	results, err := ctx.Database.Query(ctx.Statement)

	if err != nil {
		return nil, err
	}

	return ctx.results(results), nil
}

// Comment
func (ctx *QueryStatement[T]) Paginate(perPage int64, page int64) (*Pagination[*T], error) {
	if page > 1 {
		ctx.Offset(perPage * (page - 1))
	}

	results, err := ctx.Database.Query(ctx.Limit(perPage).(*QueryStatement[T]).Statement)

	if err != nil {
		return nil, err
	}

	total, err := ctx.Count()

	if err != nil {
		return nil, err
	}

	return &Pagination[*T]{
		Total:   total,
		PerPage: perPage,
		Page:    page,
		Items:   ctx.results(results),
	}, nil
}

// Comment
func (ctx *QueryStatement[T]) Insert(values Values) (*T, error) {
	ctx.Values = values

	result, err := ctx.Database.Insert(ctx.Statement)

	if err != nil {
		return nil, err
	}

	return ctx.result(result), nil
}

// Comment
func (ctx *QueryStatement[T]) Update(values Values) error {
	ctx.Values = values

	return ctx.Database.Update(ctx.Statement)
}
