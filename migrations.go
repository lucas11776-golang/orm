package orm

type TableScheme struct {
	Name    string
	Columns []Scheme
}

type Scheme interface {
	Column() *Column
}

type Column struct {
	Name       string
	Nullable   bool
	Default    interface{}
	Unique     bool
	PrimaryKey bool
	Size       int64
}

type Migration interface {
	Migrate(scheme *TableScheme) error
	Drop(table string) error
}
