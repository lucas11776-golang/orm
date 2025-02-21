package statements

const SPACE = " "

type Statement interface {
	Statement() (string, error)
}
