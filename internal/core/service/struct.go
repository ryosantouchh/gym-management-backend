package service

import (
	"fmt"
	"reflect"
)

func StructToMap(obj interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	v := reflect.ValueOf(obj)
	t := v.Type()
	fmt.Println(v)
	fmt.Println(t)

	fmt.Println(v.Field(1))
	// for i := 0; i < t.NumField(); i++ {
	// 	field := t.Field(i)
	// 	value := v.Field(i)
	//
	// 	if field.PkgPath != "" {
	// 		continue
	// 	}
	//
	// 	result[field.Name] = value.Interface()
	// }

	return result
}
