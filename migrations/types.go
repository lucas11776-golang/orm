package migrations

import "time"

const (
	DEFAULT_CURRENT_TIME = "CURRENT_TIMESTAMP"
)

type Table struct {
	columns []Type
}

type Type interface {
	Column() *Column
}

type Column struct {
	Name     string
	Nullable bool
	Default  interface{}
	Unique   bool
}

func newColumn(name string) *Column {
	return &Column{
		Name:     name,
		Nullable: false,
		Unique:   false,
	}
}

/******************************************
				Increment
******************************************/

type Increment struct {
	column *Column
}

// Comment
func (ctx *Increment) Column() *Column {
	return ctx.column
}

func (ctx *Table) Increment(name string) *Increment {
	ctx.columns = append(ctx.columns, &Increment{
		column: newColumn(name),
	})

	return ctx.columns[len(ctx.columns)-1].(*Increment)
}

/******************************************
				Timestamp
******************************************/

type TimeStamp struct {
	column *Column
}

// Comment
func (ctx *TimeStamp) Column() *Column {
	return ctx.column
}

func (ctx *Table) TimeStamp(name string) *TimeStamp {
	ctx.columns = append(ctx.columns, &TimeStamp{
		column: newColumn(name),
	})

	return ctx.columns[len(ctx.columns)-1].(*TimeStamp)
}

// Comment
func (ctx *TimeStamp) Nullable() *TimeStamp {
	ctx.column.Nullable = true

	return ctx
}

// Comment
func (ctx *TimeStamp) Current() *TimeStamp {
	ctx.column.Default = DEFAULT_CURRENT_TIME

	return ctx
}

// Comment
func (ctx *TimeStamp) Default(time *time.Time) *TimeStamp {
	ctx.column.Default = time

	return ctx
}

// Comment
func (ctx *TimeStamp) Unique() *TimeStamp {
	ctx.column.Unique = true

	return ctx
}

/******************************************
				Date
******************************************/

type Date struct {
	column *Column
}

// Comment
func (ctx *Date) Nullable() *Date {
	ctx.column.Nullable = true

	return ctx
}

// Comment
func (ctx *Date) Column() *Column {
	return ctx.column
}

func (ctx *Table) Date(name string) *Date {
	ctx.columns = append(ctx.columns, &Date{
		column: newColumn(name),
	})

	return ctx.columns[len(ctx.columns)-1].(*Date)
}

// Comment
func (ctx *Date) Default(date *time.Time) *Date {
	ctx.column.Default = date

	return ctx
}

// Comment
func (ctx *Date) Unique() *Date {
	ctx.column.Unique = true

	return ctx
}

/******************************************
				Integer
******************************************/

type Integer struct {
	column *Column
}

// Comment
func (ctx *Integer) Nullable() *Integer {
	ctx.column.Nullable = true

	return ctx
}

// Comment
func (ctx *Integer) Column() *Column {
	return ctx.column
}

func (ctx *Table) Integer(name string) *Integer {
	ctx.columns = append(ctx.columns, &Integer{
		column: newColumn(name),
	})

	return ctx.columns[len(ctx.columns)-1].(*Integer)
}

// Comment
func (ctx *Integer) Default(value int) *Integer {
	ctx.column.Default = value
	return ctx
}

// Comment
func (ctx *Integer) Unique() *Integer {
	ctx.column.Unique = true

	return ctx
}

/******************************************
				Double
******************************************/

type Double struct {
	column *Column
}

// Comment
func (ctx *Double) Nullable() *Double {
	ctx.column.Nullable = true

	return ctx
}

// Comment
func (ctx *Double) Column() *Column {
	return ctx.column
}

func (ctx *Table) Double(name string) *Double {
	ctx.columns = append(ctx.columns, &Double{
		column: newColumn(name),
	})

	return ctx.columns[len(ctx.columns)-1].(*Double)
}

// Comment
func (ctx *Double) Default(value int64) *Double {
	ctx.column.Default = value

	return ctx
}

// Comment
func (ctx *Double) Unique() *Double {
	ctx.column.Unique = true

	return ctx
}

/******************************************
				Float
******************************************/

type Float struct {
	column *Column
}

// Comment
func (ctx *Float) Nullable() *Float {
	ctx.column.Nullable = true

	return ctx
}

// Comment
func (ctx *Float) Column() *Column {
	return ctx.column
}

func (ctx *Table) Float(name string) *Float {
	ctx.columns = append(ctx.columns, &Float{
		column: newColumn(name),
	})

	return ctx.columns[len(ctx.columns)-1].(*Float)
}

// Comment
func (ctx *Float) Default(value int64) *Float {
	ctx.column.Default = value

	return ctx
}

// Comment
func (ctx *Float) Unique() *Float {
	ctx.column.Unique = true

	return ctx
}

/******************************************
				String
******************************************/

type String struct {
	column *Column
}

// Comment
func (ctx *String) Nullable() *String {
	ctx.column.Nullable = true

	return ctx
}

// Comment
func (ctx *String) Column() *Column {
	return ctx.column
}

func (ctx *Table) String(name string) *String {
	ctx.columns = append(ctx.columns, &String{
		column: newColumn(name),
	})

	return ctx.columns[len(ctx.columns)-1].(*String)
}

// Comment
func (ctx *String) Default(value string) *String {
	ctx.column.Default = value

	return ctx
}

// Comment
func (ctx *String) Unique() *String {
	ctx.column.Unique = true

	return ctx
}

/******************************************
				Text
******************************************/

type Text struct {
	column *Column
}

// Comment
func (ctx *Text) Nullable() *Text {
	ctx.column.Nullable = true

	return ctx
}

// Comment
func (ctx *Text) Column() *Column {
	return ctx.column
}

func (ctx *Table) Text(name string) *Text {
	ctx.columns = append(ctx.columns, &Text{
		column: newColumn(name),
	})

	return ctx.columns[len(ctx.columns)-1].(*Text)
}

// Comment
func (ctx *Text) Default(value string) *Text {
	ctx.column.Default = value

	return ctx
}

// Comment
func (ctx *Text) Unique() *Text {
	ctx.column.Unique = true

	return ctx
}

/******************************************
				Boolean
******************************************/

type Boolean struct {
	column *Column
}

// Comment
func (ctx *Boolean) Nullable() *Boolean {
	ctx.column.Nullable = true

	return ctx
}

// Comment
func (ctx *Boolean) Column() *Column {
	return ctx.column
}

func (ctx *Table) Boolean(name string) *Boolean {
	ctx.columns = append(ctx.columns, &Boolean{
		column: newColumn(name),
	})

	return ctx.columns[len(ctx.columns)-1].(*Boolean)
}

// Comment
func (ctx *Boolean) Default(value bool) *Boolean {
	ctx.column.Default = value

	return ctx
}

// Comment
func (ctx *Boolean) Unique() *Boolean {
	ctx.column.Unique = true

	return ctx
}

// type Schema func(connection string, table string, callback func(table *Table))

type Schema struct {
	Table   string
	Columns []Type
}
