package sql

import (
	"database/sql"
	"fmt"
	"reflect"

	"github.com/lucas11776-golang/orm/types"
	"github.com/lucas11776-golang/orm/utils/cast"
)

type TableInfo struct {
	CID          int    `column:"cid"`
	Name         string `column:"name"`
	Type         string `column:"type"`
	NotNull      bool   `column:"notnull"`
	DefaultValue string `column:"dflt_value"`
	PrimaryKey   bool   `column:"pk"`
}

type TablePrimaryKey func(table string) (key string, err error)

type PrimaryKeyCache struct {
	cache    map[string]string
	callback TablePrimaryKey
}

// Comment
func NewPrimaryKeyCache(callback TablePrimaryKey) *PrimaryKeyCache {
	return &PrimaryKeyCache{
		callback: callback,
		cache:    make(map[string]string),
	}
}

// Comment
func (ctx *PrimaryKeyCache) TablePrimaryKey(table string) (key string, err error) {
	key, ok := ctx.cache[table]

	if ok {
		return key, nil
	}

	key, err = ctx.callback(table)

	if err != nil {
		return "", err
	}

	ctx.cache[table] = key

	return key, nil
}

// Comment
func TableInfoPrimaryKey(db *sql.DB, table string) (string, error) {
	rows, err := db.Query(fmt.Sprintf("PRAGMA table_info(`%s`);", table))

	if err != nil {
		return "", err
	}

	columns, err := ScanRowsToModels(rows, TableInfo{})

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
func ScanRowToResult(row *sql.Rows, columns []string) types.Result {
	v := make([]any, len(columns))
	maps := make([]interface{}, len(v))
	vMap := map[string]interface{}{}

	for i := 0; i < len(maps); i++ {
		maps[i] = &v[i]
	}

	row.Scan(maps...)

	for i, v := range v {
		vMap[columns[i]] = v
	}

	return vMap
}

// Comment
func ScanRowsToResults(rows *sql.Rows) (types.Results, error) {
	results := types.Results{}

	columns, err := rows.Columns()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		results = append(results, ScanRowToResult(rows, columns))
	}

	return results, nil
}

// Comment
func ScanRowsToModels[Model any](rows *sql.Rows, model Model) ([]*Model, error) {
	results, err := ScanRowsToResults(rows)

	if err != nil {
		return nil, err
	}

	return ResultsToModels(results, model), nil
}

// Comment
func ResultToModel[Model any](result types.Result, model Model) *Model {
	modelElements := reflect.ValueOf(&model).Elem()

	for i := 0; i < modelElements.NumField(); i++ {
		col := modelElements.Type().Field(i).Tag.Get("column")

		if col == "" {
			continue
		}

		v, ok := result[col]

		if !ok || v == nil || v == "" {
			continue
		}

		modelElements.Field(i).Set(reflect.ValueOf(cast.Kind(modelElements.Type().Field(i).Type.Kind(), v)))
	}

	return &model
}

// Comment
func ResultsToModels[Model any](results types.Results, model Model) []*Model {
	models := []*Model{}

	for _, result := range results {
		models = append(models, ResultToModel(result, reflect.Zero(reflect.TypeOf(model)).Interface().(Model)))
	}

	return models
}
