package migration

import "time"

type Type interface {
	Column() *Column
}

type Column struct {
	Name     string
	Nullable bool
	Default  interface{}
	Unique   bool
}

/******************************************
				PRIMARY KEY
******************************************/

type PrimaryKey struct {
	column *Column
}

// Comment
func (ctx *PrimaryKey) Column() *Column {
	return ctx.column
}

func (ctx *Table) PrimaryKey(name string) *PrimaryKey {
	ctx.columns = append(ctx.columns, &PrimaryKey{
		column: &Column{
			Name: name,
		},
	})
	return nil
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
		column: &Column{
			Name: name,
		},
	})
	return nil
}

// Comment
func (ctx *TimeStamp) Nullable() *TimeStamp {
	return ctx
}

// Comment
func (ctx *TimeStamp) Current() *TimeStamp {
	return ctx
}

// Comment
func (ctx *TimeStamp) Default(time *time.Time) *TimeStamp {
	return ctx
}

// Comment
func (ctx *TimeStamp) Unique() *TimeStamp {
	return ctx
}

/******************************************
				Datetime
******************************************/

type Datetime struct {
	column *Column
}

// Comment
func (ctx *Datetime) Column() *Column {
	return ctx.column
}

// Comment
func (ctx *Datetime) Nullable() *Datetime {
	return ctx
}

func (ctx *Table) Datetime(name string) *Datetime {
	ctx.columns = append(ctx.columns, &Datetime{
		column: &Column{
			Name: name,
		},
	})
	return nil
}

// Comment
func (ctx *Datetime) Default(time *time.Time) *Datetime {
	return ctx
}

// Comment
func (ctx *Datetime) Unique() *Datetime {
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
	return ctx
}

// Comment
func (ctx *Date) Column() *Column {
	return ctx.column
}

func (ctx *Table) Date(name string) *Date {
	ctx.columns = append(ctx.columns, &Date{
		column: &Column{
			Name: name,
		},
	})
	return nil
}

// Comment
func (ctx *Date) Default(time *time.Time) *Date {
	return ctx
}

// Comment
func (ctx *Date) Unique() *Date {
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
	return ctx
}

// Comment
func (ctx *Integer) Column() *Column {
	return ctx.column
}

func (ctx *Table) Integer(name string) *Integer {
	ctx.columns = append(ctx.columns, &Integer{
		column: &Column{
			Name: name,
		},
	})
	return nil
}

// Comment
func (ctx *Integer) Default(time *time.Time) *Integer {
	return ctx
}

// Comment
func (ctx *Integer) Unique() *Integer {
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
	return ctx
}

// Comment
func (ctx *Double) Column() *Column {
	return ctx.column
}

func (ctx *Table) Double(name string) *Double {
	ctx.columns = append(ctx.columns, &Double{
		column: &Column{
			Name: name,
		},
	})
	return nil
}

// Comment
func (ctx *Double) Default(value int64) *Double {
	return ctx
}

// Comment
func (ctx *Double) Unique() *Double {
	return ctx
}

/******************************************
				Double
******************************************/

type Float struct {
	column *Column
}

// Comment
func (ctx *Float) Nullable() *Float {
	return ctx
}

// Comment
func (ctx *Float) Column() *Column {
	return ctx.column
}

func (ctx *Table) Float(name string) *Float {
	ctx.columns = append(ctx.columns, &Float{
		column: &Column{
			Name: name,
		},
	})
	return nil
}

// Comment
func (ctx *Float) Default(value Float) *Float {
	return ctx
}

// Comment
func (ctx *Float) Unique() *Float {
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
	return ctx
}

// Comment
func (ctx *String) Column() *Column {
	return ctx.column
}

func (ctx *Table) String(name string) *String {
	ctx.columns = append(ctx.columns, &String{
		column: &Column{
			Name: name,
		},
	})
	return ctx.columns[len(ctx.columns)-1].(*String)
}

// Comment
func (ctx *String) Default(value string) *String {
	return ctx
}

// Comment
func (ctx *String) Unique() *String {
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
	return ctx
}

// Comment
func (ctx *Text) Column() *Column {
	return ctx.column
}

func (ctx *Table) Text(name string) *Text {
	ctx.columns = append(ctx.columns, &Text{
		column: &Column{
			Name: name,
		},
	})
	return ctx.columns[len(ctx.columns)-1].(*Text)
}

// Comment
func (ctx *Text) Default(time *time.Time) *Text {
	return ctx
}

// Comment
func (ctx *Text) Unique() *Text {
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
	return ctx
}

// Comment
func (ctx *Boolean) Column() *Column {
	return ctx.column
}

func (ctx *Table) Boolean(name string) *Boolean {
	ctx.columns = append(ctx.columns, &Boolean{
		column: &Column{
			Name: name,
		},
	})
	return ctx.columns[len(ctx.columns)-1].(*Boolean)
}

// Comment
func (ctx *Boolean) Default(time *time.Time) *Boolean {
	return ctx
}

// Comment
func (ctx *Boolean) Unique() *Boolean {
	return ctx
}

type Table struct {
	columns []Type
}

// type Schema func(connection string, table string, callback func(table *Table))

type Schema struct {
	Table   string
	Columns []Type
}
