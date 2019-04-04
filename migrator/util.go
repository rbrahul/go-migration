package migrator

import (
	"fmt"
	"regexp"
	"strings"
)

type mapCB func(int, string) string

func CheckError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

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

func getIndexKeyPath(tableName string, columnNames []string) string {
	indexes := strings.Join(arrayStrMap(columnNames, func(i int, item string) string {
		return strings.ToLower(item)
	}), "_")
	return fmt.Sprintf("%s_%s_index", tableName, indexes)
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

func ENUMValus(datatypeStr string) []string {
	regxDataType := regexp.MustCompile(`(\w+)\((.*)\)`)
	matchedElements := regxDataType.FindAllStringSubmatch(datatypeStr, -1)
	if len(matchedElements[0]) > 0 {
		if len(matchedElements[0][2]) > 0 {
			var values = strings.Split(matchedElements[0][2], ",")
			return arrayStrMap(values, func(i int, item string) string {
				return strings.Replace(item, "'", "", -1)
			})
		}
	}
	return []string{}
}

func findColumnChanges(oldTuple TupleInfo, newTuple TupleInfo) TupleInfo {
	changedTuple := oldTuple
	if oldTuple.Size != newTuple.Size {
		changedTuple.Size = newTuple.Size
	}

	if oldTuple.Precision != newTuple.Precision && newTuple.Precision > 0 {
		changedTuple.Precision = newTuple.Precision
	}

	if oldTuple.Type != newTuple.Type {
		changedTuple.Type = newTuple.Type
	}

	changedTuple.EnumValues = newTuple.EnumValues

	if oldTuple.CommentText != newTuple.CommentText && len(newTuple.CommentText) > 0 {
		changedTuple.CommentText = newTuple.CommentText
	}

	if oldTuple.DefaultValue != newTuple.DefaultValue && len(newTuple.DefaultValue) > 0 {
		changedTuple.DefaultValue = newTuple.DefaultValue
	}

	if oldTuple.Collate != newTuple.Collate && len(newTuple.Collate) > 0 {
		changedTuple.Collate = newTuple.Collate
	}

	if oldTuple.CharSet != newTuple.CharSet && len(newTuple.CharSet) > 0 {
		changedTuple.Collate = newTuple.Collate
	}

	if oldTuple.IsUnique != newTuple.IsUnique && newTuple.IsUnique {
		changedTuple.IsUnique = newTuple.IsUnique
	}

	if oldTuple.IsAutoIncrement != newTuple.IsAutoIncrement && newTuple.IsAutoIncrement {
		changedTuple.IsAutoIncrement = newTuple.IsAutoIncrement
	}

	if oldTuple.IsNullable != newTuple.IsNullable && newTuple.IsNullable {
		changedTuple.IsNullable = newTuple.IsNullable
	}

	if oldTuple.IsUnSigned != newTuple.IsUnSigned && newTuple.IsUnSigned {
		changedTuple.IsUnSigned = newTuple.IsUnSigned
	}

	changedTuple.ChangeOnly = newTuple.ChangeOnly
	return changedTuple
}
