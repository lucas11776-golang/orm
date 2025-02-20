package sqlite

const SPACE = "	"

type QueryValues []interface{}

type QueryBuilder struct {
	Query  *Query
	Values QueryValues
}

// Comment
func (ctx *QueryBuilder) SelectStatement() (string, error) {
	return "", nil
}

// Comment
func (ctx *QueryBuilder) WhereStatement() (string, error) {
	return "", nil
}

// Comment
func (ctx *QueryBuilder) JoinStatement() (string, error) {
	return "", nil
}

// Comment
func (ctx *QueryBuilder) Limit() string {
	return ""
}
