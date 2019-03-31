package migrator

import (
	"fmt"
	"strings"
)

type QueryGenerator struct {
	Table             *TableManager
	TableDefinition   string
	ToupleDefinitions []string
}

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
		colation = fmt.Sprintf(" COLLATE %", tupleInfo.Collate)
	}

	if len(strings.TrimSpace(tupleInfo.CharSet)) > 0 {
		charSet = fmt.Sprintf(" CHARACTER SET %", tupleInfo.CharSet)
	}

	if len(strings.TrimSpace(tupleInfo.CharSet)) == 0 && len(strings.TrimSpace(tupleInfo.Collate)) > 0 {
		charSet = fmt.Sprintf(" CHARACTER SET %", tupleInfo.CharSet)
	}

	return fmt.Sprintf("`%s` %s%s%s%s%s%s%s%s%s", colunmName, dataType, columnLength, unSigned, charSet, colation, nullAbleText, defaultValue, autoIncreament, commentText)
}

func (qg *QueryGenerator) GenerateTableStructure() {
	qg.GenerateTupleStructure()
}

func (qg *QueryGenerator) GenerateTupleStructure() {
	var tuples []string
	for _, item := range qg.Table.Structures {
		sqlStatement := createColumn(item)
		tuples = append(tuples, sqlStatement)
	}
	fmt.Printf("%v", tuples)
}

func (qg *QueryGenerator) PrepareQuery() {

}
