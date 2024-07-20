package _struct

import (
	"reflect"
)

// 从一个结构体切片中获取指定字段并生成新的结构体切片
func extractFields(src interface{}, dstType reflect.Type, fields []string) interface{} {
	srcValue := reflect.ValueOf(src)
	dstSlice := reflect.MakeSlice(reflect.SliceOf(dstType), srcValue.Len(), srcValue.Len())

	for i := 0; i < srcValue.Len(); i++ {
		dstStruct := reflect.New(dstType).Elem()
		srcStruct := srcValue.Index(i)

		for _, field := range fields {
			dstField := dstStruct.FieldByName(field)
			srcField := srcStruct.FieldByName(field)
			dstField.Set(srcField)
		}
		dstSlice.Index(i).Set(dstStruct)
	}

	return dstSlice.Interface()
}
