package migrator

/*Schema - Generating Migration*/
type Schema struct {
	TableName string
	CreateNew bool
	Collation string
	CharSet   string
	Temporary bool
	Database  *Database
}

/*CB -call back function*/
type CB func(*TableManager)

/*Create - Create a  Schema for new table*/
func (mg *Schema) Create(tableName string, cb CB) {
	table := &TableManager{}
	mg.TableName = tableName
	mg.CreateNew = true
	table.Schema = mg

	queryGenerator := &QueryGenerator{Database: mg.Database}
	queryGenerator.Table = table
	cb(table)

	//Prepare SQL Statement
	queryGenerator.GenerateTableStructure()
}

/*Table - Update Schema of existing table*/
func (mg *Schema) Table(tableName string, cb CB) {
	table := &TableManager{}
	mg.TableName = tableName
	mg.CreateNew = false
	table.Schema = mg

	queryGenerator := &QueryGenerator{Database: mg.Database}
	queryGenerator.Table = table
	cb(table)

	//Prepare SQL Statement
	queryGenerator.GenerateTableStructure()
}
