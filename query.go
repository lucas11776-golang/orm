package orm

type Entity interface{}

type Select []interface{}
type Join map[interface{}]interface{}
type Where map[interface{}]interface{}
type Limit int64

type Pagination struct {
	Total   int64    `json:"total"`
	Page    int64    `json:"Page"`
	PerPage int64    `json:"per_page"`
	Items   []Entity `json:"items"`
}

type Query interface {
	Select(s Select) Query
	Join(j Join) Query
	Where(w Where) Query
	Limit(l Limit) Query
	Offset(o int64) Query
	Count() (int64, error)
	Get() (Entity, error)
	Paginate(total int64, page int64) (*Pagination, error)
}
