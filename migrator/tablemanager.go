package migrator

import "fmt"

//OperationQueue - Contains all the structure changes both ColumnDefinition and ALTER Commands
type OperationQueue struct {
	OperationType string
	Command       *Command
	TupleInfo     *TupleInfo
}

/*TableManager - Contains blue print of table */
type TableManager struct {
	Schema     *Schema
	Structures []*OperationQueue
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
	DefaultValue     string
	CurrentTimeStamp bool
	EnumValues       []string
	IsUnSigned       bool
	IsPrimary        bool
	IsIndexed        bool
	IsUnique         bool
	IsNullable       bool
	//TODO: IsFirst and After
	IsFirst         bool
	AfterTupleName  string
	ForeignRelation *ForeignInfo
	ChangeOnly      bool
}

type Command struct {
	OperationType   string
	ToupleName      []string
	ForeignRelation *ForeignInfo
	NewName         string
	IfExists        bool
}

type ForeignInfo struct {
	ToupleName        string
	Referrence        string
	RelativeTableName string
	onUpdate          string
	onDelete          string
}

var query *QueryGenerator

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
	operation := &OperationQueue{OperationType: defineTuple, TupleInfo: tuple}
	tm.Structures = append(tm.Structures, operation)
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

func (tpl *TupleInfo) Default(value string) *TupleInfo {
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

func (tpl *TupleInfo) Primary() *TupleInfo {
	tpl.IsPrimary = true
	return tpl
}

func (tpl *TupleInfo) Unique() *TupleInfo {
	tpl.IsUnique = true
	return tpl
}

func (tpl *TupleInfo) Change() {
	tpl.ChangeOnly = true
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
	command := &Command{ToupleName: touples, OperationType: commandType}
	operation := &OperationQueue{OperationType: defineCommand, Command: command}
	tm.Structures = append(tm.Structures, operation)
	return command
}

//COMMANDS - all kind of INDEX related operation

func (tm *TableManager) Index(index []string) {
	addCommand(tm, index, AddIndex)
}

func (tm *TableManager) Drop(string) {
	addCommand(tm, []string{}, DropTable)
}

func (tm *TableManager) DropIfExists() {
	command := addCommand(tm, []string{}, DropTable)
	command.IfExists = true
}

func (tm *TableManager) DropColumn(name string) {
	addCommand(tm, []string{name}, DropTuple)
}

func (tm *TableManager) DropColumnIfExists(name string) {
	command := addCommand(tm, []string{name}, DropTuple)
	command.IfExists = true
}

func (tm *TableManager) DropColumns(names []string) {
	addCommand(tm, names, DropTuple)
}

func (tm *TableManager) DropTimeStamps() {
	tm.DropColumns([]string{"created_at", "deleted_at"})
}

func (tm *TableManager) RenameColumn(oldName string, newName string) {
	command := addCommand(tm, []string{oldName}, RenameTuple)
	command.NewName = newName
}

func (tm *TableManager) RenameTable(oldName string, newName string) {
	command := addCommand(tm, []string{oldName}, RenameTable)
	command.NewName = newName
}

func (tm *TableManager) RenameIndex(oldName string, newName string) {
	command := addCommand(tm, []string{oldName}, RenameIndex)
	command.NewName = newName
}

func (tm *TableManager) DropPrimary(names []string) {
	addCommand(tm, names, DropPrimaryKey)
}

func (tm *TableManager) DropIndex(name []string) {
	addCommand(tm, name, DropIndex)
}

func (tm *TableManager) DropIndexes(names []string) {
	addCommand(tm, names, DropIndex)
}

func (tm *TableManager) DropUnique(name []string) {
	addCommand(tm, name, DropUnique)
}

func (tm *TableManager) DropForeign(name string) {
	addCommand(tm, []string{name}, DropForeignKey)
}

func (tm *TableManager) Primary(name []string) {
	addCommand(tm, name, AddPrimaryKey)
}

func (tm *TableManager) Foreign(name string) *ForeignInfo {
	command := addCommand(tm, []string{name}, AddForeignKey)
	foreign := &ForeignInfo{}
	command.ForeignRelation = foreign
	return foreign
}

func (tm *TableManager) AllCommands() {
}
