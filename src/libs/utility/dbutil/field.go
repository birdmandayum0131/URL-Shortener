package dbutil

import (
	"reflect"
	"strings"
)

func FieldString(fieldStruct interface{}) string {
	fields := reflect.ValueOf(fieldStruct)
	types := fields.Type()
	fieldArray := make([]string, 0)
	// * iterate all fields, if not empty, append to filterStrings
	for i := 0; i < fields.NumField(); i++ {
		if !fields.Field(i).IsZero() {
			fieldArray = append(fieldArray, types.Field(i).Tag.Get("db"))
		}
	}
	return strings.Join(fieldArray, ", ")
}
