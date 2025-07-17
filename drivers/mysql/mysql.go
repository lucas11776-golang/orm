package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lucas11776-golang/orm"

	// "github.com/lucas11776-golang/orm/drivers/mysql/migrations"
	"github.com/lucas11776-golang/orm/drivers/mysql/migrations"
	sqlDriver "github.com/lucas11776-golang/orm/drivers/sql"
	"github.com/lucas11776-golang/orm/drivers/sql/statements"
	"github.com/lucas11776-golang/orm/types"
	utils "github.com/lucas11776-golang/orm/utils/sql"
)

type Credentials struct {
	Host     string
	User     string
	Password string
	Database string
	Port     int16
	SSL      bool
	Protocol string
}

type MySQL struct {
	DB *sql.DB
}

type TableInfo struct {
	CID          int    `column:"cid"`
	Name         string `column:"name"`
	Type         string `column:"type"`
	NotNull      bool   `column:"notnull"`
	DefaultValue string `column:"dflt_value"`
	PrimaryKey   bool   `column:"pk"`
}

// Comment
func (ctx *MySQL) query(query string, values sqlDriver.QueryValues) (types.Results, error) {

	stmt, err := ctx.DB.Prepare(query)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(values...)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return utils.ScanRowsToResults(rows)
}

// Comment
func (ctx *MySQL) Query(statement *orm.Statement) (types.Results, error) {
	builder := &sqlDriver.QueryBuilder{Statement: statement}

	sql, values, err := builder.Query()

	if err != nil {
		return nil, err
	}

	return ctx.query(sql, values)
}

// Comment
func (ctx *MySQL) Count(statement *orm.Statement) (int64, error) {
	builder := &sqlDriver.QueryBuilder{Statement: statement}

	sql, values, err := builder.Count()

	if err != nil {
		return 0, err
	}

	results, err := ctx.query(sql, values)

	if err != nil {
		return 0, nil
	}

	if len(results) != 1 {
		return 0, errors.New("failed to execute count")
	}

	total, ok := results[0]["total"]

	if !ok {
		return 0, errors.New("expected count result map to have total key")
	}

	return total.(int64), nil
}

// Comment
func (ctx *MySQL) getPrimaryKey(table string) (string, error) {
	rows, err := ctx.DB.Query(fmt.Sprintf("PRAGMA table_info(%s);", statements.SafeKey(table)))

	if err != nil {
		return "", err
	}

	columns, err := utils.ScanRowsToModels(rows, TableInfo{})

	if err != nil {
		return "", err
	}

	for _, column := range columns {
		if column.PrimaryKey {
			return column.Name, nil
		}
	}

	return "", nil

}

// Comment
func (ctx *MySQL) Insert(statement *orm.Statement) (types.Result, error) {
	builder := &sqlDriver.QueryBuilder{Statement: statement}

	sql, values, err := builder.Insert()

	if err != nil {
		return nil, err
	}

	stmt, err := ctx.DB.Prepare(sql)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	insertResult, err := stmt.Exec(values...)

	if err != nil {
		return nil, err
	}

	key, err := ctx.getPrimaryKey(statement.Table)

	if err != nil {
		return nil, err
	}

	if key == "" {
		return types.Result(statement.Values), nil
	}

	id, err := insertResult.LastInsertId()

	if err != nil {
		return nil, err
	}

	builder = &sqlDriver.QueryBuilder{Statement: &orm.Statement{
		Table: statement.Table,
		Where: []interface{}{&orm.Where{
			Key:      key,
			Operator: "=",
			Value:    id,
		}},
	}}

	sql, values, err = builder.Query()

	if err != nil {
		return nil, err
	}

	results, err := ctx.query(sql, values)

	if err != nil {
		return nil, err
	}

	if len(results) != 1 {
		return nil, fmt.Errorf("failed to get insert row")
	}

	return results[0], nil
}

// Comment
func (ctx *MySQL) Update(statement *orm.Statement) error {
	builder := &sqlDriver.QueryBuilder{Statement: statement}

	sql, values, err := builder.Update()

	if err != nil {
		return err
	}

	_, err = ctx.DB.Exec(sql, values...)

	if err != nil {
		return err
	}

	return nil
}

// Comment
func (ctx *MySQL) Delete(statement *orm.Statement) error {
	builder := &sqlDriver.QueryBuilder{Statement: statement}

	sql, values, err := builder.Delete()

	if err != nil {
		return err
	}

	_, err = ctx.DB.Exec(sql, values...)

	if err != nil {
		return err
	}

	return nil
}

// Comment
func (ctx *MySQL) Database() interface{} {
	return ctx.DB
}

// Comment
func (ctx *MySQL) Migration() orm.Migration {
	return &migrations.Migration{DB: ctx.DB}
}

// Comment
func (ctx *MySQL) Close() error {
	return ctx.DB.Close()
}

/*********************************************************************************/
// // Comment
// func (ctx *MySQL) Query(statement *orm.Statement) (types.Results, error) {
// 	return nil, nil
// }

// // Comment
// func (ctx *MySQL) Count(statement *orm.Statement) (int64, error) {
// 	return 0, nil
// }

// // Comment
// func (ctx *MySQL) Insert(statement *orm.Statement) (types.Result, error) {
// 	return nil, nil
// }

// // Comment
// func (ctx *MySQL) Update(statement *orm.Statement) error {
// 	return nil
// }

// // Comment
// func (ctx *MySQL) Delete(Statement *orm.Statement) error {
// 	return nil
// }

// // Comment
// func (ctx *MySQL) Database() interface{} {
// 	return nil
// }

// // Comment
// func (ctx *MySQL) Migration() orm.Migration {
// 	return nil
// }

// // Comment
// func (ctx *MySQL) Close() error {
// 	return ctx.DB.Close()
// }

// Comment
func Connect(credentials *Credentials) orm.Database {
	db, err := sql.Open("mysql", GenerateDataSourceName(credentials))

	if err != nil {
		panic(err)
	}

	return &MySQL{DB: db}
}

// Comment
func GenerateDataSourceName(credentials *Credentials) string {
	url := url.Values{"parseTime": []string{"true"}}

	if !credentials.SSL {
		url.Add("tls", "skip-verify")
	}

	return fmt.Sprintf(
		"%s:%s@%s(%s:%d)/%s?%s",
		credentials.User,
		credentials.Password,
		credentials.Protocol,
		credentials.Host,
		credentials.Port,
		credentials.Database,
		url.Encode(),
	)
}
