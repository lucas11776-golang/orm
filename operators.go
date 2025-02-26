package orm

type Operator string

const (
	AND                 = "AND"
	OR                  = "OR"
	NOT                 = "NOT"
	IS_NOT              = "IS NOT"
	EQUALS              = "="
	NOT_EQUALS          = "!="
	LESS_THEN           = "<"
	LESS_THEN_EQUALS    = "<="
	GREATER_THEN        = ">"
	GREATER_THEN_EQUALS = ">="
)
