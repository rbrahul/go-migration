package migrator

/*

type Migrator struct {
	App {
		DB:
		SCHEMA.
		QUERYGENERATOR
		TABLEMANAGER,
		MigrationService
	}
}
*/
/*Schema - Generating Migration*/
type Schema struct {
	TableName string
	CreateNew bool
	Collation string
	CharSet   string
	Comment   string
	Temporary bool
	DB        *Database
}

/*CB -call back function*/
type CB func(*TableManager)

/*Create - Create a  Schema for new table*/
func (mg *Schema) Create(tableName string, cb CB) {
	table := &TableManager{}
	mg.TableName = tableName
	mg.CreateNew = true
	if mg.DB == nil {
		db := InitDatabase()
		mg.DB = db
	}
	table.Schema = mg

	queryGenerator := &QueryGenerator{}
	queryGenerator.Table = table
	queryGenerator.Database = mg.DB
	cb(table)

	//Prepare SQL Statement
	queryGenerator.ProcessMigration()
}

/*Table - Update Schema of existing table*/
func (mg *Schema) Table(tableName string, cb CB) {
	if mg.DB == nil {
		db := InitDatabase()
		mg.DB = db
	}
	table := &TableManager{}
	mg.TableName = tableName
	mg.CreateNew = false
	table.Schema = mg

	queryGenerator := &QueryGenerator{}
	queryGenerator.Table = table
	queryGenerator.Database = mg.DB
	cb(table)

	//Prepare SQL Statement
	queryGenerator.ProcessMigration()
}
