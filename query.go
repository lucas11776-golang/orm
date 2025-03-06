package orm

import (
	"orm/utils/cast"
	"reflect"
)

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
	Table string
	Where []interface{}
}

type Select []interface{}
type Join map[string]interface{}
type Joins []*JoinHolder
type JoinGroup func(group JoinGroupBuilder)
type Where map[string]interface{}
type WhereGroup func(group WhereGroupBuilder)
type Limit int64
type Offset int64
type OrderBy [2]interface{}

type JoinGroupBuilder interface {
	Join(j Join) JoinGroupBuilder
	And(j Join) JoinGroupBuilder
	Or(j Join) JoinGroupBuilder
	Group(group JoinGroup) JoinGroupBuilder
}

type WhereGroupBuilder interface {
	Where(w Where) WhereGroupBuilder
	AndWhere(w Where) WhereGroupBuilder
	OrWhere(w Where) WhereGroupBuilder
}

type Pagination[T any] struct {
	Total   int64 `json:"total"`
	Page    int64 `json:"Page"`
	PerPage int64 `json:"per_page"`
	Items   []T   `json:"items"`
}

type Statement struct {
	Table   string
	Select  Select
	Joins   Joins
	Where   []interface{}
	Limit   int64
	Offset  int64
	OrderBy OrderBy
	Values  Values
}

type QueryStatement[T any] struct {
	Model    T
	Database Database
	*Statement
}

type Order string

type QueryBuilder[T any] interface {
	Select(s Select) QueryBuilder[T]
	Join(table string, j Join) QueryBuilder[T]
	JoinGroup(table string, group JoinGroup) QueryBuilder[T]
	Where(w Where) QueryBuilder[T]
	AndWhere(w Where) QueryBuilder[T]
	OrWhere(w Where) QueryBuilder[T]
	WhereGroup(group WhereGroup) QueryBuilder[T]
	AndWhereGroup(group WhereGroup) QueryBuilder[T]
	OrWhereGroup(group WhereGroup) QueryBuilder[T]
	Limit(l Limit) QueryBuilder[T]
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
	return ctx
}

// Comment
func (ctx *QueryStatement[T]) Join(table string, j Join) QueryBuilder[T] {
	return ctx
}

// Comment
func (ctx *QueryStatement[T]) JoinGroup(table string, group JoinGroup) QueryBuilder[T] {
	return ctx
}

// Comment
func (ctx *QueryStatement[T]) Where(w Where) QueryBuilder[T] {
	return ctx
}

// Comment
func (ctx *QueryStatement[T]) AndWhere(w Where) QueryBuilder[T] {
	return ctx
}

// Comment
func (ctx *QueryStatement[T]) OrWhere(w Where) QueryBuilder[T] {
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
func (ctx *QueryStatement[T]) Limit(l Limit) QueryBuilder[T] {
	return ctx
}

// Comment
func (ctx *QueryStatement[T]) Offset(o int64) QueryBuilder[T] {
	return ctx
}

// Comment
func (ctx *QueryStatement[T]) OrderBy(column string, order Order) QueryBuilder[T] {
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

		if col == "" {
			continue
		}

		v, ok := raw[col]

		if !ok {
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
	results, err := ctx.Database.Query(ctx.Statement)

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
	result, err := ctx.Database.Insert(ctx.Statement)

	if err != nil {
		return nil, err
	}

	return ctx.result(result), nil
}

// Comment
func (ctx *QueryStatement[T]) Update(values Values) error {
	return ctx.Database.Update(values)
}
