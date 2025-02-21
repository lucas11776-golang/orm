package orm

type Entity interface{}

type Select []interface{}
type Join map[interface{}]interface{}
type Where map[string]interface{}
type WhereGroup func(group WhereGroupBuilder)
type Limit int64
type Offset int64

type WhereGroupBuilder interface {
	Where(w Where) WhereGroupBuilder
	AndWhere(w Where) WhereGroupBuilder
	OrWhere(w Where) WhereGroupBuilder
}

type Pagination struct {
	Total   int64    `json:"total"`
	Page    int64    `json:"Page"`
	PerPage int64    `json:"per_page"`
	Items   []Entity `json:"items"`
}

type Values map[string]interface{}

type Query interface {
	Select(s Select) Query
	Join(j Join) Query
	AndJoin(l Join) Query
	OrJoin(l Join) Query
	Where(w Where) Query
	AndWhere(w Where) Query
	OrWhere(w Where) Query
	WhereGroup(group WhereGroup) Query
	AndWhereGroup(group WhereGroup) Query
	OrWhereGroup(group WhereGroup) Query
	Limit(l Limit) Query
	Offset(o int64) Query
	Count() (int64, error)
	Get() (Entity, error)
	Paginate(total int64, page int64) (*Pagination, error)
	Insert(values Values) (Entity, error)
}
