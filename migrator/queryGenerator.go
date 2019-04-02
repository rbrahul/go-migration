package migrator

import (
	"fmt"
	"strings"
)

//QueryGenerator - Blueprint of Generating SQl Query
type QueryGenerator struct {
	Table             *TableManager
	TableDefinition   string
	ToupleDefinitions []string
}

/*
FOR CREATE:
New columns are added inside create() SQL Method
and all the indexing, droping, primary, uniqe, renaming are executed as ALTER TABLE command individually
Schema::create('users', function (Blueprint $table) {
            $table->bigIncrements('id');
            $table->string('name');

            $table->dropColumn('name');
            $table->string('name');
            $table->index(['account_id']);
            $table->index(['created_at']);
            $table->string('email')->unique();
            $table->timestamp('email_verified_at')->nullable();
            $table->string('password');
            $table->rememberToken();
            $table->timestamps();
            $table->primary(["id"]);
            $table->dropPrimary(["id"]);
        });


Array
(
    [0] => create table `users` (`id` bigint unsigned not null auto_increment primary key, `name` varchar(255) not null, `name` varchar(255) not null, `email` varchar(255) not null, `email_verified_at` timestamp null, `password` varchar(255) not null, `remember_token` varchar(100) null, `created_at` timestamp null, `updated_at` timestamp null) default character set utf8mb4 collate 'utf8mb4_unicode_ci'
    [1] => alter table `users` drop `name`
    [2] => alter table `users` add index `users_account_id_index`(`account_id`)
    [3] => alter table `users` add index `users_created_at_index`(`created_at`)
    [4] => alter table `users` add primary key `users_id_primary`(`id`)
    [5] => alter table `users` drop primary key
    [6] => alter table `users` add unique `users_email_unique`(`email`)
)
*/

func createColumn(tupleInfo *TupleInfo) string {
	colunmName := fmt.Sprintf("`%s`", tupleInfo.Name)
	dataType := tupleInfo.Type
	columnLength := ""
	commentText := ""
	nullAbleText := " NOT NULL"
	unSigned := ""
	defaultValue := ""
	autoIncreament := ""
	charSet := ""
	colation := ""

	if len(tupleInfo.Name) > 0 {
		colunmName = tupleInfo.Name
	}
	if tupleInfo.Size > 0 {
		columnLength = fmt.Sprintf("(%d)", tupleInfo.Size)
	}

	if tupleInfo.Size > 0 && tupleInfo.Precision > 0 {
		columnLength = fmt.Sprintf("(%d, %d)", tupleInfo.Size, tupleInfo.Precision)
	}

	if tupleInfo.Type == ENUM && len(tupleInfo.EnumValues) > 0 {
		var modifiedValues []string
		for _, item := range tupleInfo.EnumValues {
			modifiedValues = append(modifiedValues, fmt.Sprintf("'%s'", item))
		}
		values := strings.Join(modifiedValues, ",")
		dataType = fmt.Sprintf("%s(%s)", ENUM, values)
	}

	if len(strings.TrimSpace(tupleInfo.CommentText)) > 0 {
		commentText = fmt.Sprintf(" COMMENT '%s' ", strings.TrimSpace(tupleInfo.CommentText))
	}

	if tupleInfo.IsNullable {
		nullAbleText = " NULL"
	}

	if tupleInfo.IsUnSigned {
		unSigned = " UNSIGNED"
	}

	if tupleInfo.CurrentTimeStamp {
		defaultValue = " DEFAULT CURRENT_TIMESTAMP"
	}

	if len(strings.TrimSpace(tupleInfo.DefaultValue)) > 0 {
		defaultValue = fmt.Sprintf(" DEFAULT '%s'", strings.TrimSpace(tupleInfo.DefaultValue))
	}

	if tupleInfo.IsAutoIncrement {
		autoIncreament = " AUTO_INCREMENT"
	}

	if len(strings.TrimSpace(tupleInfo.Collate)) > 0 {
		colation = fmt.Sprintf(" COLLATE %s", tupleInfo.Collate)
	}

	if len(strings.TrimSpace(tupleInfo.CharSet)) > 0 {
		charSet = fmt.Sprintf(" CHARACTER SET %s", tupleInfo.CharSet)
	}

	if len(strings.TrimSpace(tupleInfo.CharSet)) == 0 && len(strings.TrimSpace(tupleInfo.Collate)) > 0 {
		charSet = fmt.Sprintf(" CHARACTER SET %s", tupleInfo.Collate)
	}

	//TODO: Nedd to create INDEX:

	/*	ALTER TABLE `Test Table 1` ADD `a` INT NOT NULL AFTER `options`, ADD `b` INT NOT NULL AFTER `a`, ADD `c` INT NOT NULL AFTER `b`, ADD `d` INT NOT NULL AFTER `c`, ADD `e` INT NOT NULL AFTER `d`, ADD `f` INT NOT NULL AFTER `e`, ADD `g` INT NOT NULL AFTER `f`, ADD `h` INT NOT NULL AFTER `g`, ADD `i` INT NOT NULL AFTER `h`, ADD `j` INT NOT NULL AFTER `i`, ADD PRIMARY KEY (`a`, `b`), ADD INDEX (`e`), ADD INDEX (`f`), ADD UNIQUE (`c`), ADD UNIQUE (`d`), ADD FULLTEXT (`g`), ADD FULLTEXT (`h`);*/
	return fmt.Sprintf("`%s` %s%s%s%s%s%s%s%s%s", colunmName, dataType, columnLength, unSigned, charSet, colation, nullAbleText, defaultValue, autoIncreament, commentText)
}

func createCommand(tableName string, commandItem *Command) string {
	alterCommand := fmt.Sprintf("ALTER TABLE `%s` ", tableName)
	ifExist := ""
	command := ""

	if commandItem.IfExists {
		ifExist = " IF EXISTS"
	}
	switch commandItem.OperationType {
	case DropTable:
		command = fmt.Sprintf("DROP TABLE%s %s", ifExist, tableName)
	case RenameTable:
		oldName := commandItem.ToupleName[0]
		command = fmt.Sprintf("RENAME TABLE %s TO %s", oldName, commandItem.NewName)
	case DropTuple:
		columnNames := prepareKeys(commandItem.ToupleName)
		command = fmt.Sprintf("%s DROP COLUMN%s %s", alterCommand, ifExist, columnNames)
	case RenameTuple:
		oldName := commandItem.ToupleName[0]
		//TODO: Need to pass also Datatype(Size), so we need to get the current column info first
		command = fmt.Sprintf("%s CHANGE %s TO %s", alterCommand, oldName, commandItem.NewName)
	case RenameIndex:
		oldName := commandItem.ToupleName[0]
		command = fmt.Sprintf("%s RENAME INDEX %s TO %s", alterCommand, oldName, commandItem.NewName)
	case AddForeignKey:
		command = getNewForeignKeySyntax(tableName, commandItem)
	case DropForeignKey:
		columnName := commandItem.ToupleName[0]
		constraints := generateConstraints(tableName, columnName)
		command = fmt.Sprintf("%s DROP FOREIGN KEY %s", alterCommand, constraints)
	case AddPrimaryKey:
		command = fmt.Sprintf("%s ADD PRIMARY KEY(%s)", alterCommand, commandItem.ToupleName[0])
	case DropPrimaryKey:
		command = fmt.Sprintf("%s DROP PRIMARY KEY", alterCommand)
	case AddIndex:
		indexedKey := generateIndexKey(tableName, commandItem.ToupleName)
		command = fmt.Sprintf("%s ADD INDEX %s", alterCommand, indexedKey)
	case AddUnique:
		indexedKey := fmt.Sprintf("%s_unique(%s)", commandItem.ToupleName[0], commandItem.ToupleName[0])
		command = fmt.Sprintf("%s ADD UNIQUE INDEX %s", alterCommand, indexedKey)
	case DropUnique:
		indexedKey := fmt.Sprintf("%s_unique(%s)", commandItem.ToupleName[0], commandItem.ToupleName[0])
		command = fmt.Sprintf("%s DROP INDEX %s", alterCommand, indexedKey)
	case DropIndex:
		indexPath := getIndexKeyPath(tableName, commandItem.ToupleName)
		command = fmt.Sprintf("%s DROP INDEX %s", alterCommand, indexPath)
	}
	return command
}

func (qg *QueryGenerator) GenerateTableStructure() {
	qg.GenerateTupleStructure()
}

func (qg *QueryGenerator) GetIndexedColums() {
}

func (qg *QueryGenerator) GenerateTupleStructure() {
	//var tuples []string
	var commands []string
	/*for _, item := range qg.Table.Structures {
		sqlStatement := createColumn(item)
		tuples = append(tuples, sqlStatement)
	}
	*/
	for _, commandItem := range qg.Table.Commands {
		sqlStatement := createCommand(qg.Table.Schema.TableName, commandItem)
		commands = append(commands, sqlStatement)
	}

	fmt.Printf("%v", commands)
}

func (qg *QueryGenerator) PrepareQuery() {

}
