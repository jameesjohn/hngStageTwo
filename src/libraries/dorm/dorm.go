package dorm

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

// Person is a Sample type that works with this.
// For app, we're using gorm.
type Person struct {
	Id        int       `json:"id" column_name:"id" type:"int" attrs:"not null auto_increment" pk:"id"`
	Name      string    `json:"name" column_name:"name" type:"varchar(255)" attrs:"not null"`
	CreatedAt time.Time `json:"created_at" column_name:"created_at" type:"datetime" attrs:"not null default now()"`
}

// Dorm stands for Database Object Relational mapping
type columnAttrs struct {
	Value     any    `json:"value"`
	DataType  string `json:"dataType"`
	IsPrimary bool   `json:"isPrimary"`
	Attrs     string `json:"attrs"`
}

type columnName string

func getTableName(model interface{}) string {
	return strings.ToLower(reflect.TypeOf(model).Name()) + "s"
}

func getColumnAttrs(model interface{}) map[columnName]columnAttrs {
	columns := map[columnName]columnAttrs{}
	valueOf := reflect.ValueOf(model)

	for i := 0; i < valueOf.NumField(); i++ {
		field := reflect.TypeOf(model).Field(i)
		fmt.Println("Value Of", reflect.ValueOf(field))

		cName := columnName(field.Tag.Get("column_name"))
		dataType := field.Tag.Get("type")
		attrs := field.Tag.Get("attrs")

		pk := columnName(field.Tag.Get("pk"))
		isPrimary := false
		if pk == cName {
			isPrimary = true
		}
		columns[cName] = columnAttrs{Attrs: attrs, DataType: dataType, IsPrimary: isPrimary}
	}

	return columns
}

func GetMigrationQuery(model interface{}) string {
	// map of [fieldName][sqlDataType]

	var columnQueries []string

	// Pluralized table name
	tableName := getTableName(model)
	columns := getColumnAttrs(model)

	for name, attrs := range columns {
		columnQueries = append(columnQueries, fmt.Sprintf("%s %s %s", name, attrs.DataType, strings.ToUpper(attrs.Attrs)))
		if attrs.IsPrimary {
			columnQueries = append(columnQueries, fmt.Sprintf("PRIMARY KEY (%s)", name))
		}
	}

	createTableQuery := fmt.Sprintf("CREATE TABLE %s (%s);", tableName, strings.Join(columnQueries, ", "))

	return createTableQuery
}
