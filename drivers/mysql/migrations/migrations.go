package migrations

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	"github.com/lucas11776-golang/orm"
	"github.com/lucas11776-golang/orm/drivers/sql/statements"
	"github.com/lucas11776-golang/orm/migrations"
	"github.com/spf13/cast"
)

type Migration struct {
	DB *sql.DB
}

func getSchemeType(scheme orm.Scheme) (string, error) {
	switch scheme.(type) {

	case *migrations.TimeStamp:
		return "TIMESTAMP", nil

	case *migrations.Datetime:
		return "DATETIME", nil

	case *migrations.Date:
		return "DATE", nil

	case *migrations.Integer:
		return "INTEGER", nil

	case *migrations.BigInteger:
		return "BIGINT(20) UNSIGNED", nil

	case *migrations.Double:
		return "DOUBLE", nil

	case *migrations.Float:
		return "FLOAT", nil

	case *migrations.String:
		return "VARCHAR(16380)", nil

	case *migrations.Text:
		return "TEXT", nil

	case *migrations.Boolean:
		return "BOOLEAN", nil

	case *migrations.Binary:
		return "BLOB", nil

	default:
		return "", fmt.Errorf("sqlite does not support type of %s", reflect.ValueOf(scheme).Type().Name())
	}
}

// Comment
func generateStatement(scheme orm.Scheme) (string, error) {
	cType, err := getSchemeType(scheme)

	if err != nil {
		return "", err
	}

	column := scheme.Column()
	str := []string{statements.SafeKey(column.Name), cType}

	if column.PrimaryKey {
		str = append(str, "PRIMARY KEY")
	}

	if !column.Nullable {
		str = append(str, "NOT NULL")
	}

	if column.Default != nil {
		switch column.Default.(type) {
		case int, int64, float32, float64:
			str = append(str, fmt.Sprintf("DEFAULT %d", column.Default))

		case []byte:
			str = append(str, fmt.Sprintf("DEFAULT %s", string(column.Default.([]byte))))

		default:
			switch column.Default.(string) {
			case migrations.DEFAULT_CURRENT_TIMESTAMP:
				str = append(str, "DEFAULT CURRENT_TIMESTAMP")

			case migrations.DEFAULT_CURRENT_DATETIME:
				str = append(str, "DEFAULT CURRENT_TIMESTAMP")

			case migrations.DEFAULT_CURRENT_DATE:
				str = append(str, "DEFAULT CURRENT_DATE")

			default:
				str = append(str, fmt.Sprintf("DEFAULT '%s'", cast.ToString(column.Default)))
			}
		}
	}

	if column.Unique {
		str = append(str, "UNIQUE")
	}

	return strings.Join(str, " "), nil
}

// Comment
func generateColumnStatement(column orm.Scheme) (string, error) {
	switch column.(type) {

	case *migrations.Increment:
		return fmt.Sprintf("%s BIGINT(20) UNSIGNED PRIMARY KEY AUTO_INCREMENT UNIQUE", statements.SafeKey("id")), nil

	default:
		return generateStatement(column)
	}
}

// Comment
func (ctx *Migration) generateTableSchemeSQL(table *orm.TableScheme) (string, error) {
	query := []string{fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (", table.Name)}
	columns := []string{}

	for _, column := range table.Columns {
		col, err := generateColumnStatement(column)

		if err != nil {
			return "", err
		}

		columns = append(columns, fmt.Sprintf("%s%s", statements.SPACE, col))
	}

	return strings.Join(append(query, strings.Join(columns, ",\r\n"), ");"), "\r\n"), nil
}

// Comment
func (ctx *Migration) Migrate(scheme *orm.TableScheme) error {
	sql, err := ctx.generateTableSchemeSQL(scheme)

	if err != nil {
		return err
	}

	if _, err := ctx.DB.Exec(sql); err != nil {
		return err
	}

	return nil
}

// Comment
func (ctx *Migration) Drop(table string) error {
	return nil
}
