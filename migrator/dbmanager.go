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
}

func (db *DB) ColumnNames() []string {
	//do whatever
}

func (db *DB) ColumnExists(columnName string) bool {
	//do whatever
}

func (db *DB) TableExists(tableName string) bool {
	//do whatever
}

func (db *DB) ColumnDefinition(columnName string) TupleInfo {
	//do whatever
}

func (db *DB) TableDefinition(tableName string) interface{} {
	//do whatever
}

func (db *DB) DropTable(tableName string) interface{} {
	//drop
}
