package orm

import (
	"reflect"

	"github.com/lucas11776-golang/orm/types"
	"github.com/lucas11776-golang/orm/utils/cast"
	"github.com/lucas11776-golang/orm/utils/sql"
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

type Vector []float32
type Vector8 []float32
type Vector16 []float32
type Vector32 []float32
type Vector64 []float64
type VectorB1 []float32
type VectorB16 []float32

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
	panic("AndWhere not implemented")
}

// Comment
func (ctx *JoinBuilder) OrWhere(column1 string, operator string, column2 string) *JoinBuilder {
	panic("OrWhere not implemented")
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

// type OrderBy [2]interface{}
type OrderBy struct {
	Columns []string
	Order   Order
}

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
	Model      T
	Database   Database
	Connection string
	*Statement
}

// TODO: move to orm base
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
	OrderBy(column []string, order Order) QueryBuilder[T]
	Count() (int64, error)
	First() (*T, error)
	Get() ([]*T, error)
	Exists() (bool, error)
	Paginate(perPage int64, page int64) (*Pagination[*T], error)
	Insert(values Values) (*T, error)
	InsertMany(values []Values) ([]*T, error)
	Update(values Values) error
	Delete() error
}

// Comment
func (ctx *WhereGroupQueryBuilder) Where(column string, operator string, value interface{}) WhereGroupBuilder {
	if len(ctx.Group) != 0 {
		return ctx.AndWhere(column, operator, value)
	}

	ctx.Group = append(ctx.Group, &Where{
		Key:      column,
		Operator: operator,
		Value:    value,
	})

	return ctx
}

// Comment
func (ctx *WhereGroupQueryBuilder) AndWhere(column string, operator string, value interface{}) WhereGroupBuilder {
	if len(ctx.Group) == 0 {
		return ctx.Where(column, operator, value)
	}

	ctx.Group = append(ctx.Group, "AND", &Where{
		Key:      column,
		Operator: operator,
		Value:    value,
	})

	return ctx
}

// Comment
func (ctx *WhereGroupQueryBuilder) OrWhere(column string, operator string, value interface{}) WhereGroupBuilder {
	if len(ctx.Group) == 0 {
		return ctx.Where(column, operator, value)
	}

	ctx.Group = append(ctx.Group, "OR", &Where{
		Key:      column,
		Operator: operator,
		Value:    value,
	})

	return ctx
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
	if len(ctx.Statement.Where) != 0 {
		return ctx.AndWhere(column, operator, value)
	}

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
func (ctx *QueryStatement[T]) resolveWhereGroup(group WhereGroup) *WhereGroupQueryBuilder {
	builder := &WhereGroupQueryBuilder{}

	group(builder)

	return builder
}

// Comment
func (ctx *QueryStatement[T]) WhereGroup(group WhereGroup) QueryBuilder[T] {
	if len(ctx.Statement.Where) != 0 {
		return ctx.AndWhereGroup(group)
	}

	ctx.Statement.Where = append(ctx.Statement.Where, ctx.resolveWhereGroup(group))

	return ctx
}

// Comment
func (ctx *QueryStatement[T]) AndWhereGroup(group WhereGroup) QueryBuilder[T] {
	if len(ctx.Statement.Where) == 0 {
		return ctx.WhereGroup(group)
	}

	ctx.Statement.Where = append(ctx.Statement.Where, "AND", ctx.resolveWhereGroup(group))

	return ctx
}

// Comment
func (ctx *QueryStatement[T]) OrWhereGroup(group WhereGroup) QueryBuilder[T] {
	if len(ctx.Statement.Where) == 0 {
		return ctx.WhereGroup(group)
	}

	ctx.Statement.Where = append(ctx.Statement.Where, "OR", ctx.resolveWhereGroup(group))

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
func (ctx *QueryStatement[T]) OrderBy(column []string, order Order) QueryBuilder[T] {
	ctx.Statement.OrderBy = OrderBy{
		Columns: column,
		Order:   order,
	}

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
func CastModel[T any](model T, result types.Result) *T {
	zValue := reflect.Zero(reflect.TypeOf(model)).Interface().(T)
	zElem := reflect.ValueOf(&zValue).Elem()

	for i := 0; i < zElem.NumField(); i++ {
		col := zElem.Type().Field(i).Tag.Get("column")

		if col == "" {
			continue
		}

		v, ok := result[col]

		if !ok || v == nil || v == "" {
			continue
		}

		zElem.Field(i).Set(reflect.ValueOf(cast.Kind(zElem.Type().Field(i).Type.Kind(), v)))
	}

	return &zValue
}

// Comment
func (ctx *QueryStatement[T]) First() (*T, error) {
	results, err := ctx.Database.Query(ctx.Statement)

	if err != nil || len(results) == 0 {
		return nil, err
	}

	return sql.ResultsToModels(results, ctx.Model)[0], nil
}

// Comment
func (ctx *QueryStatement[T]) Get() ([]*T, error) {
	results, err := ctx.Database.Query(ctx.Statement)

	if err != nil {
		return nil, err
	}

	return sql.ResultsToModels(results, ctx.Model), nil
}

// Comment
func (ctx *QueryStatement[T]) Exists() (bool, error) {
	count, err := ctx.Count()

	if err != nil {
		return false, err
	}

	return count != 0, nil
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

	total, err := (&QueryStatement[T]{
		Model:      ctx.Model,
		Database:   ctx.Database,
		Connection: ctx.Connection,
		Statement: &Statement{
			Table: ctx.Statement.Table,
			Joins: ctx.Statement.Joins,
			Where: ctx.Statement.Where,
		},
	}).Count()

	if err != nil {
		return nil, err
	}

	return &Pagination[*T]{
		Total:   total,
		PerPage: perPage,
		Page:    page,
		Items:   sql.ResultsToModels(results, ctx.Model),
	}, nil
}

// Comment
func (ctx *QueryStatement[T]) Insert(values Values) (*T, error) {
	ctx.Values = values

	result, err := ctx.Database.Insert(ctx.Statement)

	if err != nil {
		return nil, err
	}

	return sql.ResultToModel(result, ctx.Model), nil
}

// Comment
func (ctx *QueryStatement[T]) InsertMany(values []Values) ([]*T, error) {
	results := []*T{}

	for _, value := range values {
		ctx.Values = value

		result, err := ctx.Database.Insert(ctx.Statement)

		if err != nil {
			return nil, err
		}

		results = append(results, sql.ResultToModel(result, ctx.Model))
	}

	return results, nil
}

// Comment
func (ctx *QueryStatement[T]) Update(values Values) error {
	ctx.Values = values

	return ctx.Database.Update(ctx.Statement)
}

// Comment
func (ctx *QueryStatement[T]) Delete() error {
	return ctx.Database.Delete(ctx.Statement)
}
