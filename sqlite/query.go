package sqlite

import "orm"

type Query struct {
	_select orm.Select
	_where  []interface{}
	_join   orm.Join
	_limit  orm.Limit
	_offset orm.Offset
}

// Comment
func (ctx *Query) Select(s orm.Select) orm.Query {
	return ctx
}

// Comment
func (ctx *Query) Join(table string, j orm.Join) orm.Query {
	return ctx
}

// Comment
func (ctx *Query) JoinGroup(table string, group orm.JoinGroup) orm.Query {
	return ctx
}

// Comment
func (ctx *Query) Where(w orm.Where) orm.Query {
	return ctx
}

// Comment
func (ctx *Query) AndWhere(w orm.Where) orm.Query {
	return ctx
}

// Comment
func (ctx *Query) OrWhere(w orm.Where) orm.Query {
	return ctx
}

// Comment
func (ctx *Query) WhereGroup(group orm.WhereGroup) orm.Query {
	return ctx
}

func (ctx *Query) AndWhereGroup(group orm.WhereGroup) orm.Query {
	return ctx
}

func (ctx *Query) OrWhereGroup(group orm.WhereGroup) orm.Query {
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

// Comment
func (ctx *Query) Insert(values orm.Values) (orm.Entity, error) {
	return nil, nil
}
