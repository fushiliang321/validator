package utils

import "reflect"

func IsArrayOrSlice(v any) bool {
	var (
		rv   = reflect.TypeOf(v)
		kind = rv.Kind()
	)
	return kind == reflect.Array || kind == reflect.Slice
}

func IsObject(v any) bool {
	var (
		rv   = reflect.TypeOf(v)
		kind = rv.Kind()
	)
	return kind == reflect.Map || kind == reflect.Struct
}
