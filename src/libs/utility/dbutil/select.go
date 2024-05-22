package dbutil

import (
	"fmt"
)

func SelectFields(dbName string, fieldStruct interface{}, filterString string) string {
	fieldString := FieldString(fieldStruct)

	fullFilterString := ""
	if filterString != "" {
		fullFilterString = fmt.Sprintf(" WHERE %s", filterString)
	}

	return fmt.Sprintf("SELECT %s FROM %s%s", fieldString, dbName, fullFilterString)
}

func SelectAll(dbName string, filterString string) string {
	fullFilterString := ""
	if filterString != "" {
		fullFilterString = fmt.Sprintf(" WHERE %s", filterString)
	}

	return fmt.Sprintf("SELECT * FROM %s%s", dbName, fullFilterString)
}
