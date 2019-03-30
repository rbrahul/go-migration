package migrator

/*SMALLINT - Touple Data type*/
const (
	SMALLINT  = "SMALLINT"
	TINYINT   = "TINYINT"
	MEDIUMINT = "MEDIUMINT"
	INTEGER   = "INTEGER"
	BIGINT    = "BIGINT"
	DECIMAL   = "DECIMAL"
	DOUBLE    = "DOUBLE"
	FLOAT     = "FLOAT"

	DATE      = "DATE"
	DATETIME  = "DATETIME"
	TIME      = "TIME"
	TIMESTAMP = "TIMESTAMP"
	YEAR      = "YEAR"

	CHAR       = "CHAR"
	VARCHAR    = "VARCHAR"
	BINARY     = "BINARY"
	VARBINARY  = "VARBINARY"
	BLOB       = "BLOB"
	TINYBLOB   = "TINYBLOB"
	MEDIUMBLOB = "MEDIUMBLOB"
	TEXT       = "TEXT"
	MEDIUMTEXT = "MEDIUMTEXT"
	LONGTEXT   = "LONGTEXT"
	TINYTEXT   = "TINYTEXT"
	ENUM       = "ENUM"
	BOOLEAN    = "BOOLEAN"
	JSON       = "JSON"
)

//DropTable - drop a table operation
const (
	DropTable      = "DROP TABLE"
	DropTuple      = "DROP COLUMN"
	RenameTuple    = "RENAME COLUMN"
	AddForeignKey  = "FOREIGN KEY"
	DropForeignKey = "DROP FOREIGN KEY"
	DropIndex      = "DROP INDEX"
	DropPrimaryKey = "DROP PRIMARY KEY"
	DropUnique     = DropIndex
)
