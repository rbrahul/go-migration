package migrator

type DB struct {
	Driver       string //mysql
	Host         string //localhost
	User         string //root
	Password     string //password
	DatabaseName string
	Connection   interface{}
}

func (db *DB) ExecuteQuery(query string) {
	//do whatever
}

func (db *DB) AllTables() []string {
	//do whatever
	return []string{}
}

func (db *DB) ColumnNames() []string {
	//do whatever
	return []string{}
}

func (db *DB) ColumnExists(columnName string) bool {
	//do whatever
	return true
}

func (db *DB) TableExists(tableName string) bool {
	//do whatever
	return true
}

func (db *DB) ColumnDefinition(columnName string) TupleInfo {
	//do whatever
	return TupleInfo{}
}

func (db *DB) TableDefinition(tableName string) interface{} {
	//do whatever
	return false
}

func (db *DB) DropTable(tableName string) interface{} {
	//drop
	return true
}
