package migrator

import (
	"fmt"
	"strings"
)

type mapCB func(int, string) string

func arrayStrMap(iterators []string, cb mapCB) []string {
	mappedArray := []string{}
	for indx, item := range iterators {
		mappedItem := cb(indx, item)
		mappedArray = append(mappedArray, mappedItem)
	}
	return mappedArray
}

func getNewForeignKeySyntax(tableName string, commandItem *Command) string {
	columnName := commandItem.ToupleName[0]
	references := ""
	relativeTable := ""
	onUpdate := "NO ACTION"
	onDelete := "NO ACTION"
	if len(commandItem.ForeignRelation.Referrence) > 0 {
		references = commandItem.ForeignRelation.Referrence
	}

	if len(commandItem.ForeignRelation.RelativeTableName) > 0 {
		relativeTable = commandItem.ForeignRelation.RelativeTableName
	}

	if len(commandItem.ForeignRelation.onDelete) > 0 {
		onDelete = commandItem.ForeignRelation.onDelete
	}

	if len(commandItem.ForeignRelation.onUpdate) > 0 {
		onUpdate = commandItem.ForeignRelation.onUpdate
	}

	constraints := generateConstraints(tableName, columnName)
	return fmt.Sprintf("ALTER TABLE %s ADD CONSTRAINTS %s FOREIGN KEY(%s) REFERENCES `%s`(`%s`) ON DELETE %s ON UPDATE %s", tableName, constraints, columnName, relativeTable, references, onDelete, onUpdate)
}

func generateConstraints(tableName string, columnName string) string {
	return fmt.Sprintf("%s_FK_%s", tableName, columnName)
}

func generateIndexKey(tableName string, columnNames []string) string {
	indexes := strings.Join(arrayStrMap(columnNames, func(i int, item string) string {
		return strings.ToLower(item)
	}), "_")
	keys := prepareKeys(columnNames)
	return fmt.Sprintf("%s_%s_index(%s)", tableName, indexes, keys)
}

func prepareKeys(columnNames []string) string {
	return strings.Join(arrayStrMap(columnNames, func(i int, item string) string {
		return fmt.Sprintf("`%s`", item)
	}), ",")
}
