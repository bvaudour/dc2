package util

import (
	"reflect"
)

func Contains(m interface{}, key interface{}) bool {
	r := reflect.ValueOf(m)
	switch r.Kind() {
	case reflect.Map:
		return r.MapIndex(reflect.ValueOf(key)).IsValid()
	case reflect.Slice:
		fallthrough
	case reflect.Array:
		for i := 0; i < r.Len(); i++ {
			if reflect.DeepEqual(r.Index(i).Interface(), key) {
				return true
			}
		}
		return false
	default:
		return false
	}
}
