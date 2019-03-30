package migrator

import "fmt"

/*TableManager - Contains blue print of table */
type TableManager struct {
	Schema     *Schema
	Structures []*TupleInfo
	Commands   []*Command
}

/*TupleInfo - Contains blue print of Table Column */
type TupleInfo struct {
	Name             string
	Type             string
	Size             int
	Precision        int
	CharSet          string
	Collate          string
	IsAutoIncrement  bool
	CommentText      string
	DefaultValue     interface{}
	CurrentTimeStamp bool
	EnumValues       []string
	IsUnSigned       bool
	IsPrimary        bool
	IsIndexed        bool
	IsUnique         bool
	IsNullable       bool
	IsForeignKey     bool
	ForeignRelation  *ForeignInfo
}

type Command struct {
	OperationType   string
	ToupleName      []string
	ForeignRelation *ForeignInfo
	NewColumnName   string
}

type ForeignInfo struct {
	ToupleName        string
	Referrence        string
	RelativeTableName string
	onUpdate          string
	onDelete          string
}

func (fr *ForeignInfo) Referrences(key string) *ForeignInfo {
	fr.Referrence = key
	return fr
}

func (fr *ForeignInfo) On(name string) *ForeignInfo {
	fr.RelativeTableName = name
	return fr
}

func (fr *ForeignInfo) OnDelete(value string) *ForeignInfo {
	fr.onDelete = value
	return fr
}

func (fr *ForeignInfo) OnUpdate(value string) *ForeignInfo {
	fr.onUpdate = value
	return fr
}

func (tm *TableManager) FindExistingTuple() {
	for _, Item := range tm.Structures {
		fmt.Printf("%v", Item)
	}
}

/*--------------------------------------*
/*Touple Manager - Utility               *
*---------------------------------------*/
func initializeTouple(tm *TableManager, name string, datatype string) *TupleInfo {
	tuple := &TupleInfo{Name: name, Type: datatype}
	tm.Structures = append(tm.Structures, tuple)
	return tuple
}

func (tpl *TupleInfo) UseCurrent() *TupleInfo {
	tpl.CurrentTimeStamp = true
	return tpl
}

func (tpl *TupleInfo) Comment(value string) *TupleInfo {
	tpl.CommentText = value
	return tpl
}

func (tpl *TupleInfo) Default(value interface{}) *TupleInfo {
	tpl.DefaultValue = value
	return tpl
}

func (tpl *TupleInfo) Unsigned() *TupleInfo {
	tpl.IsUnSigned = true
	return tpl
}

func (tpl *TupleInfo) Autoincrement() *TupleInfo {
	tpl.IsAutoIncrement = true
	return tpl
}

func (tpl *TupleInfo) Charset(charset string) *TupleInfo {
	tpl.CharSet = charset
	return tpl
}

func (tpl *TupleInfo) Collation(charset string) *TupleInfo {
	tpl.Collate = charset
	return tpl
}

func (tpl *TupleInfo) Nullable() *TupleInfo {
	tpl.IsNullable = true
	return tpl
}

func (tpl *TupleInfo) SetSize(size int, precison ...int) *TupleInfo {
	tpl.Size = size

	if len(precison) > 0 {
		tpl.Precision = precison[0]
	}
	return tpl
}

/*--------------------------------------*
/*Table Manager - Utility               *
*---------------------------------------*/

func (tm *TableManager) String(name string, size ...int) *TupleInfo {
	tuple := initializeTouple(tm, name, VARCHAR)
	if len(size) > 0 {
		tuple.Size = size[0]
	}
	return tuple
}

func (tm *TableManager) TinyInteger(name string) *TupleInfo {
	return initializeTouple(tm, name, TINYINT)
}

func (tm *TableManager) SmallInteger(name string) *TupleInfo {
	return initializeTouple(tm, name, SMALLINT)
}

func (tm *TableManager) UnsignedSmallInteger(name string) *TupleInfo {
	return initializeTouple(tm, name, SMALLINT).Unsigned()
}

func (tm *TableManager) Integer(name string) *TupleInfo {
	return initializeTouple(tm, name, INTEGER)
}

func (tm *TableManager) UnsignedInteger(name string) *TupleInfo {
	return initializeTouple(tm, name, INTEGER).Unsigned()
}

func (tm *TableManager) MediumInteger(name string) *TupleInfo {
	return initializeTouple(tm, name, MEDIUMINT)
}

func (tm *TableManager) UnsignedMediumInteger(name string) *TupleInfo {
	return initializeTouple(tm, name, MEDIUMINT).Unsigned()
}

func (tm *TableManager) BigInteger(name string) *TupleInfo {
	return initializeTouple(tm, name, BIGINT)
}

func (tm *TableManager) UnsignedBigInteger(name string) *TupleInfo {
	return initializeTouple(tm, name, BIGINT).Unsigned()
}

func (tm *TableManager) Increments(name string) *TupleInfo {
	return initializeTouple(tm, name, INTEGER).Unsigned().Autoincrement()
}

func (tm *TableManager) SmallIncrements(name string) *TupleInfo {
	return initializeTouple(tm, name, SMALLINT).Unsigned().Autoincrement()
}

func (tm *TableManager) MediumIncrements(name string) *TupleInfo {
	return initializeTouple(tm, name, MEDIUMINT).Unsigned().Autoincrement()
}

func (tm *TableManager) Binary(name string) *TupleInfo {
	return initializeTouple(tm, name, BLOB)
}

func (tm *TableManager) Boolean(name string) *TupleInfo {
	return initializeTouple(tm, name, BOOLEAN)
}

func (tm *TableManager) Date(name string) *TupleInfo {
	return initializeTouple(tm, name, DATE)
}

func (tm *TableManager) DateTime(name string) *TupleInfo {
	return initializeTouple(tm, name, DATETIME)
}

func (tm *TableManager) Timestamp(name string) *TupleInfo {
	return initializeTouple(tm, name, TIMESTAMP)
}

func (tm *TableManager) Timestamps() {
	tm.Timestamp("created_at").UseCurrent()
	tm.Timestamp("updated_at").UseCurrent()
}

func (tm *TableManager) Time(name string) *TupleInfo {
	return initializeTouple(tm, name, TIME)
}

func (tm *TableManager) Year(name string) *TupleInfo {
	return initializeTouple(tm, name, YEAR)
}

func (tm *TableManager) Decimal(name string, size int, precision int) *TupleInfo {
	return initializeTouple(tm, name, DECIMAL).SetSize(size, precision)
}

func (tm *TableManager) UnsignedDecimal(name string, size int, precision int) *TupleInfo {
	return initializeTouple(tm, name, DECIMAL).Unsigned().SetSize(size, precision)
}

func (tm *TableManager) Float(name string, size int, precision int) *TupleInfo {
	return initializeTouple(tm, name, FLOAT).SetSize(size, precision)
}

func (tm *TableManager) Char(name string) *TupleInfo {
	return initializeTouple(tm, name, CHAR)
}

func (tm *TableManager) Text(name string) *TupleInfo {
	return initializeTouple(tm, name, TEXT)
}

func (tm *TableManager) MediumText(name string) *TupleInfo {
	return initializeTouple(tm, name, MEDIUMTEXT)
}

func (tm *TableManager) Json(name string) *TupleInfo {
	return initializeTouple(tm, name, JSON)
}

func (tm *TableManager) ENUM(name string, values []string) *TupleInfo {
	tuple := initializeTouple(tm, name, ENUM)
	tuple.EnumValues = values
	return tuple
}

func addCommand(tm *TableManager, touples []string, commandType string) *Command {
	command := &Command{}
	command.OperationType = commandType
	command.ToupleName = touples
	tm.Commands = append(tm.Commands, command)
	return command
}

//COMMANDS - all kind of INDEX related operation

func (tm *TableManager) Drop() *Command {
	return addCommand(tm, []string{}, DropTable)
}

func (tm *TableManager) DropColumn(name string) {
	addCommand(tm, []string{name}, DropTuple)
}

func (tm *TableManager) DropColumns(names []string) {
	addCommand(tm, names, DropTuple)
}

func (tm *TableManager) DropTimeStamps() {
	tm.DropColumns([]string{"created_at", "deleted_at"})
}

func (tm *TableManager) RenameColumns(oldName string, newName string) {
	command := addCommand(tm, []string{oldName}, RenameTuple)
	command.NewColumnName = newName
}

func (tm *TableManager) DropPrimary(name string) {
	addCommand(tm, []string{name}, DropPrimaryKey)
}

func (tm *TableManager) DropPrimaries(names []string) {
	addCommand(tm, names, DropPrimaryKey)
}

func (tm *TableManager) DropIndex(name string) {
	addCommand(tm, []string{name}, DropIndex)
}

func (tm *TableManager) DropIndexes(names []string) {
	addCommand(tm, names, DropIndex)
}

func (tm *TableManager) DropUnique(name string) {
	addCommand(tm, []string{name}, DropUnique)
}

func (tm *TableManager) DropForeign(name string) {
	addCommand(tm, []string{name}, DropForeignKey)
}

func (tm *TableManager) Foreign(name string) *ForeignInfo {
	command := addCommand(tm, []string{name}, AddForeignKey)
	foreign := &ForeignInfo{}
	command.ForeignRelation = foreign
	return foreign
}

func (tm *TableManager) AllCommands() {
	for _, Item := range tm.Commands {
		fmt.Printf("Touple", Item.ToupleName)
		fmt.Printf("Foreign %v", Item.ForeignRelation)
	}
}
