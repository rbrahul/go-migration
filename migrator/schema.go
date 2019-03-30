package migrator

/*Schema - Generating Migration*/
type Schema struct {
	TableName string
	createNew bool
	Collation string
	CharSet   string
	Temporary bool
}

/*CB -call back function*/
type CB func(*TableManager)

/*Create - Create a  Schema new table*/
func (mg *Schema) Create(tableName string, cb CB) {
	table := &TableManager{}
	mg.TableName = tableName
	mg.createNew = true
	table.Schema = mg
	cb(table)
}
