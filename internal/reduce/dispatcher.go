package reduce

import (
	"reflect"
)

func dispatch(items *reflect.Value, fn interface{}, info *reduceInfo) interface{} {
	switch info.fnIn1Type.Kind() {
	case reflect.Int:
		items := items.Interface().([]int)
		switch info.fnIn2Type.Kind() {
		case reflect.Int:
			return reduceIntToInt(items, fn.(func(int, int) int))
		}
	case reflect.String:
		items := items.Interface().([]string)
		switch info.fnIn2Type.Kind() {
		case reflect.Int:
			return reduceStringToInt(items, fn.(func(int, string) int))
		case reflect.String:
			return reduceStringToString(items, fn.(func(string, string) string))
		}

	}
	return nil
}
