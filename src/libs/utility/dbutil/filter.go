package dbutil

import (
	"fmt"
	"reflect"
	"strings"
)

func FilterString(fieldStruct interface{}) string {
	fields := reflect.ValueOf(fieldStruct)
	types := fields.Type()
	filterStrings := make([]string, 0)
	// * iterate all fields, if not empty, append to filterStrings
	for i := 0; i < fields.NumField(); i++ {
		if !fields.Field(i).IsZero() {
			var compareString string
			if types.Field(i).Type == reflect.TypeOf("") {
				compareString = "%s = '%s'"
			} else {
				compareString = "%s = %s"
			}
			filterStrings = append(filterStrings, fmt.Sprintf(compareString, types.Field(i).Tag.Get("db"), fields.Field(i)))
		}
	}
	
	return strings.Join(filterStrings, " AND ")
}
