package statements

type Limit struct {
	Limit  int64
	Offset int64
}

// Comment
func (ctx *Limit) Statement() (string, error) {
	return "", nil
}
