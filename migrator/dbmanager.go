package migrator

import (
	"database/sql"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	Driver         string //mysql
	Host           string //localhost
	User           string //root
	Password       string //password
	DatabaseName   string
	DatabasePrefix string
	Port           int
	Engine         string
	DbCon          *sql.DB
}

var DefaultEngine string
var DB *Database

func InitDatabase() *Database {
	DB := &Database{User: "root", Password: "mysql", Host: "localhost", DatabaseName: "laravel_blog"}
	if DB.Port == 0 {
		DB.Port = 3306
	}
	if len(DB.Driver) == 0 {
		DB.Driver = "mysql"
	}
	if len(DB.Engine) == 0 {
		DB.Engine = "InnoDB"
	}
	URL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True", DB.User, DB.Password, DB.Host, DB.Port, DB.DatabaseName)
	session, err := sql.Open(DB.Driver, URL)
	CheckError(err)
	DB.DbCon = session
	return DB
}

func (db *Database) ExecuteQuery(query string) error {
	_, err := db.DbCon.Query(query)
	return err
}

func (db *Database) AllTables() []string {
	rows, err := db.DbCon.Query("SHOW TABLES")
	CheckError(err)
	var tables []string
	for rows.Next() {
		var tableName string
		err = rows.Scan(&tableName)
		CheckError(err)
		tables = append(tables, tableName)
	}
	return tables
}

func (db *Database) Columns(tableName string) []TupleInfo {
	//do whatever
	//var name string
	var tuples []TupleInfo

	rows, err := db.DbCon.Query(fmt.Sprintf("SHOW FULL COLUMNS FROM %s", tableName))
	CheckError(err)
	for rows.Next() {
		var name string
		var datatypeStr string
		var collation string
		var null string
		var key string
		var defaultValue string
		var extra string
		var previliges string
		var comment string

		rows.Scan(&name, &datatypeStr, &collation, &null, &key, &defaultValue, &extra, &previliges, &comment)
		regxDataType := regexp.MustCompile(`(\w+)\(?(\d+)?,?(\d+)?\)?`)
		matchedElements := regxDataType.FindAllStringSubmatch(datatypeStr, -1)
		tuple := TupleInfo{
			Name:         name,
			Collate:      collation,
			IsNullable:   null == "YES",
			DefaultValue: defaultValue,
			CommentText:  comment,
		}
		if collation != "NULL" {
			tuple.Collate = collation
		}
		if strings.Contains(strings.ToLower(datatypeStr), "unsigned") {
			tuple.IsUnSigned = true
		}
		if strings.Contains(strings.ToLower(extra), "auto_increment") {
			tuple.IsAutoIncrement = true
		}
		if len(matchedElements[0]) > 0 {
			if len(matchedElements[0][1]) > 0 {
				tuple.Type = strings.ToUpper(matchedElements[0][1])
				if tuple.Type == "ENUM" {
					tuple.EnumValues = getENUMValus(datatypeStr)

				}
			}
			if len(matchedElements[0][2]) > 0 {
				size, err := strconv.Atoi(matchedElements[0][2])
				CheckError(err)
				tuple.Size = size
			}

			if len(matchedElements[0][3]) > 0 {
				precision, err := strconv.Atoi(matchedElements[0][3])
				CheckError(err)
				tuple.Precision = precision
			}
		}
		tuples = append(tuples, tuple)
	}

	//fmt.Printf("%v", columns)
	return tuples
}

func (db *Database) ColumnNames(tableName string) []string {
	var columnNames []string
	columns := db.Columns(tableName)
	for _, item := range columns {
		columnNames = append(columnNames, item.Name)
	}
	return columnNames
}

func (db *Database) TupleDefinition(tableName string, tupleName string) (TupleInfo, error) {
	columns := db.Columns(tableName)
	for _, item := range columns {
		if item.Name == tupleName {
			return item, nil
		}
	}

	return TupleInfo{}, fmt.Errorf("Column `%s` doesn't exist on `%s` table", tupleName, tableName)
}

func (db *Database) ColumnExists(columnName string) bool {
	//do whatever
	return true
}

func (db *Database) TableExists(tableName string) bool {
	//do whatever
	return true
}

func (db *Database) TableDefinition(tableName string) interface{} {
	//do whatever
	return false
}

func (db *Database) DropTable(tableName string) interface{} {
	//drop
	return true
}
