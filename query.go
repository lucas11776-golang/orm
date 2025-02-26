package orm

type Entity interface{}

type Values map[string]interface{}

type Raw struct {
	Value interface{}
}

// Comment
func RawValue(v interface{}) *Raw {
	return &Raw{Value: v}
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
	Items   []*T  `json:"items"`
}

type Statement struct {
	Model      interface{}
	Connection string
	SELECT     Select
	JOINS      Joins
	WHERE      []interface{}
	LIMIT      Limit
	OFFSET     Offset
	UPDATA     Values
}

type QueryStatement[T any] struct {
	*Statement
}

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
	Count() (int64, error)
	First() (*T, error)
	Get() ([]*T, error)
	Paginate(total int64, page int64) (*Pagination[*T], error)
	Insert(values Values) (*T, error)
	Update(values Values) (*T, error)
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
func (ctx *QueryStatement[T]) Count() (int64, error) {
	return 0, nil
}

// Comment
func (ctx *QueryStatement[T]) First() (*T, error) {
	return nil, nil
}

// Comment
func (ctx *QueryStatement[T]) Get() ([]*T, error) {
	return nil, nil
}

// Comment
func (ctx *QueryStatement[T]) Paginate(total int64, page int64) (*Pagination[*T], error) {
	return nil, nil
}

// Comment
func (ctx *QueryStatement[T]) Insert(values Values) (*T, error) {
	return nil, nil
}

// Comment
func (ctx *QueryStatement[T]) Update(Values Values) (*T, error) {
	return nil, nil
}
