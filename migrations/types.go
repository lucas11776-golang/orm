package migrations

import (
	"time"

	"github.com/lucas11776-golang/orm"
)

const (
	DEFAULT_CURRENT_TIMESTAMP = "CURRENT_TIMESTAMP"
	DEFAULT_CURRENT_DATETIME  = "CURRENT_TIME_DATE"
	DEFAULT_CURRENT_DATE      = "CURRENT_DATE"
)

type Table struct {
	Columns []orm.Scheme
}

func newColumn(name string) *orm.Column {
	return &orm.Column{
		Name:     name,
		Nullable: false,
		Unique:   false,
	}
}

/******************************************
				Increment
******************************************/

type Increment struct {
	column *orm.Column
}

// Comment
func (ctx *Increment) Column() *orm.Column {
	return ctx.column
}

// Comment
func (ctx *Table) Increment(name string) *Increment {
	ctx.Columns = append(ctx.Columns, &Increment{
		column: newColumn(name),
	})

	return ctx.Columns[len(ctx.Columns)-1].(*Increment)
}

/******************************************
				Timestamp
******************************************/

type TimeStamp struct {
	column *orm.Column
}

// Comment
func (ctx *TimeStamp) Column() *orm.Column {
	return ctx.column
}

// Comment
func (ctx *Table) TimeStamp(name string) *TimeStamp {
	ctx.Columns = append(ctx.Columns, &TimeStamp{
		column: newColumn(name),
	})

	return ctx.Columns[len(ctx.Columns)-1].(*TimeStamp)
}

// Comment
func (ctx *TimeStamp) Nullable() *TimeStamp {
	ctx.column.Nullable = true

	return ctx
}

// Comment
func (ctx *TimeStamp) Current() *TimeStamp {
	ctx.column.Default = DEFAULT_CURRENT_TIMESTAMP

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
				Datetime
******************************************/

type Datetime struct {
	column *orm.Column
}

// Comment
func (ctx *Datetime) Column() *orm.Column {
	return ctx.column
}

func (ctx *Table) Datetime(name string) *Datetime {
	ctx.Columns = append(ctx.Columns, &Datetime{
		column: newColumn(name),
	})

	return ctx.Columns[len(ctx.Columns)-1].(*Datetime)
}

// Comment
func (ctx *Datetime) Nullable() *Datetime {
	ctx.column.Nullable = true

	return ctx
}

// Comment
func (ctx *Datetime) Current() *Datetime {
	ctx.column.Default = DEFAULT_CURRENT_TIMESTAMP

	return ctx
}

// Comment
func (ctx *Datetime) Default(time *time.Time) *Datetime {
	ctx.column.Default = time

	return ctx
}

// Comment
func (ctx *Datetime) Unique() *Datetime {
	ctx.column.Unique = true

	return ctx
}

/******************************************
				Date
******************************************/

type Date struct {
	column *orm.Column
}

// Comment
func (ctx *Date) Nullable() *Date {
	ctx.column.Nullable = true

	return ctx
}

// Comment
func (ctx *Date) Column() *orm.Column {
	return ctx.column
}

// Comment
func (ctx *Table) Date(name string) *Date {
	ctx.Columns = append(ctx.Columns, &Date{
		column: newColumn(name),
	})

	return ctx.Columns[len(ctx.Columns)-1].(*Date)
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
	column *orm.Column
}

// Comment
func (ctx *Integer) Nullable() *Integer {
	ctx.column.Nullable = true

	return ctx
}

// Comment
func (ctx *Integer) Column() *orm.Column {
	return ctx.column
}

// Comment
func (ctx *Table) Integer(name string) *Integer {
	ctx.Columns = append(ctx.Columns, &Integer{
		column: newColumn(name),
	})

	return ctx.Columns[len(ctx.Columns)-1].(*Integer)
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
				BigInteger
******************************************/

type BigInteger struct {
	column *orm.Column
}

// Comment
func (ctx *BigInteger) Nullable() *BigInteger {
	ctx.column.Nullable = true

	return ctx
}

// Comment
func (ctx *BigInteger) Column() *orm.Column {
	return ctx.column
}

// Comment
func (ctx *Table) BigInteger(name string) *BigInteger {
	ctx.Columns = append(ctx.Columns, &BigInteger{
		column: newColumn(name),
	})

	return ctx.Columns[len(ctx.Columns)-1].(*BigInteger)
}

// Comment
func (ctx *BigInteger) Default(value int64) *BigInteger {
	ctx.column.Default = value
	return ctx
}

// Comment
func (ctx *BigInteger) Unique() *BigInteger {
	ctx.column.Unique = true

	return ctx
}

/******************************************
				Double
******************************************/

type Double struct {
	column *orm.Column
}

// Comment
func (ctx *Double) Nullable() *Double {
	ctx.column.Nullable = true

	return ctx
}

// Comment
func (ctx *Double) Column() *orm.Column {
	return ctx.column
}

// Comment
func (ctx *Table) Double(name string) *Double {
	ctx.Columns = append(ctx.Columns, &Double{
		column: newColumn(name),
	})

	return ctx.Columns[len(ctx.Columns)-1].(*Double)
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
	column *orm.Column
}

// Comment
func (ctx *Float) Nullable() *Float {
	ctx.column.Nullable = true

	return ctx
}

// Comment
func (ctx *Float) Column() *orm.Column {
	return ctx.column
}

// Comment
func (ctx *Table) Float(name string) *Float {
	ctx.Columns = append(ctx.Columns, &Float{
		column: newColumn(name),
	})

	return ctx.Columns[len(ctx.Columns)-1].(*Float)
}

// Comment
func (ctx *Float) Default(value float64) *Float {
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
	column *orm.Column
}

// Comment
func (ctx *String) PrimaryKey() *String {
	ctx.column.PrimaryKey = true

	return ctx
}

// Comment
func (ctx *String) Nullable() *String {
	ctx.column.Nullable = true

	return ctx
}

// Comment
func (ctx *String) Column() *orm.Column {
	return ctx.column
}

// Comment
func (ctx *Table) String(name string) *String {
	ctx.Columns = append(ctx.Columns, &String{
		column: newColumn(name),
	})

	return ctx.Columns[len(ctx.Columns)-1].(*String)
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
	column *orm.Column
}

// Comment
func (ctx *Text) Nullable() *Text {
	ctx.column.Nullable = true

	return ctx
}

// Comment
func (ctx *Text) Column() *orm.Column {
	return ctx.column
}

// Comment
func (ctx *Table) Text(name string) *Text {
	ctx.Columns = append(ctx.Columns, &Text{
		column: newColumn(name),
	})

	return ctx.Columns[len(ctx.Columns)-1].(*Text)
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
	column *orm.Column
}

// Comment
func (ctx *Boolean) Nullable() *Boolean {
	ctx.column.Nullable = true

	return ctx
}

// Comment
func (ctx *Boolean) Column() *orm.Column {
	return ctx.column
}

// Comment
func (ctx *Table) Boolean(name string) *Boolean {
	ctx.Columns = append(ctx.Columns, &Boolean{
		column: newColumn(name),
	})

	return ctx.Columns[len(ctx.Columns)-1].(*Boolean)
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

/******************************************
				Binary
******************************************/

type Binary struct {
	column *orm.Column
}

// Comment
func (ctx *Binary) Nullable() *Binary {
	ctx.column.Nullable = true

	return ctx
}

// Comment
func (ctx *Binary) Column() *orm.Column {
	return ctx.column
}

// Comment
func (ctx *Table) Binary(name string) *Binary {
	ctx.Columns = append(ctx.Columns, &Binary{
		column: newColumn(name),
	})

	return ctx.Columns[len(ctx.Columns)-1].(*Binary)
}

// Comment
func (ctx *Binary) Default(value []byte) *Binary {
	ctx.column.Default = value

	return ctx
}

// Comment
func (ctx *Binary) Unique() *Binary {
	ctx.column.Unique = true

	return ctx
}

/******************************************
				Embedding
******************************************/

// TODO: Embedding can be float or byte/integer...
type Embedding struct {
	column *orm.Column
}

// Comment
func (ctx *Embedding) Nullable() *Embedding {
	ctx.column.Nullable = true

	return ctx
}

// Comment
func (ctx *Embedding) Column() *orm.Column {
	return ctx.column
}

// Comment
func (ctx *Table) Embedding(name string, size int64) *Embedding {
	// TODO: maximum size 65536
	ctx.Columns = append(ctx.Columns, &Embedding{
		column: newColumn(name),
	})

	ctx.Columns[len(ctx.Columns)-1].(*Embedding).column.Size = size

	return ctx.Columns[len(ctx.Columns)-1].(*Embedding)
}

// Comment
func (ctx *Embedding) Default(value []byte) *Embedding {
	ctx.column.Default = value

	return ctx
}
