package sqlite

import "orm"

type Query struct {
}

// Comment
func (ctx *Query) Select(s orm.Select) orm.Query {
	return ctx
}

// Comment
func (ctx *Query) Join(j orm.Join) orm.Query {
	return ctx
}

// Comment
func (ctx *Query) Where(w orm.Where) orm.Query {
	return ctx
}

// Comment
func (ctx *Query) Limit(l orm.Limit) orm.Query {
	return ctx
}

// Comment
func (ctx *Query) Offset(o int64) orm.Query {
	return ctx
}

// Comment
func (ctx *Query) Count() (int64, error) {
	return 0, nil
}

// Comment
func (ctx *Query) Get() (orm.Entity, error) {
	return nil, nil
}

// Comment
func (ctx *Query) Paginate(total int64, page int64) (*orm.Pagination, error) {
	return nil, nil
}
