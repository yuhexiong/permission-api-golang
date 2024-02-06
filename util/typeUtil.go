package util

import "reflect"

func GetPointer[T any](value T) *T {
	return &value
}

// 檢驗這個 interface 是不是 nil
func IsInterfaceValueNil(i interface{}) bool {
	return i == nil || reflect.ValueOf(i).IsNil()
}
