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
}

type Migration interface {
	// Migrate(models Models) error
	// Drop(models Models) error
	Migrate(scheme *TableScheme) error
	Drop(table string) error
}
