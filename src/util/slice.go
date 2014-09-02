package util

import (
	"reflect"
	"sort"
)

type sl struct {
	v reflect.Value
	f func(interface{}, interface{}) int
}

func (a sl) Len() int {
	return a.v.Len()
}

func (a sl) Index(i int) reflect.Value {
	return a.v.Index(i)
}

func (a sl) Less(i, j int) bool {
	v1, v2 := a.Index(i), a.Index(j)
	return a.f(v1.Interface(), v2.Interface()) < 0
}

func (a sl) Swap(i, j int) {
	v1, v2 := a.Index(i), a.Index(j)
	vt := reflect.ValueOf(v1.Interface())
	v1.Set(v2)
	v2.Set(vt)
}

func compareInt(i1, i2 int64) int {
	switch {
	case i1 < i2:
		return -1
	case i1 > i2:
		return 1
	default:
		return 0
	}
}

func compareUInt(i1, i2 uint64) int {
	switch {
	case i1 < i2:
		return -1
	case i1 > i2:
		return 1
	default:
		return 0
	}
}

func compareFloat(i1, i2 float64) int {
	switch {
	case i1 < i2:
		return -1
	case i1 > i2:
		return 1
	default:
		return 0
	}
}

func compareString(i1, i2 string) int {
	switch {
	case i1 < i2:
		return -1
	case i1 > i2:
		return 1
	default:
		return 0
	}
}

func isInt(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Int:
		fallthrough
	case reflect.Int8:
		fallthrough
	case reflect.Int16:
		fallthrough
	case reflect.Int32:
		fallthrough
	case reflect.Int64:
		return true
	default:
		return false
	}
}

func isUInt(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Uint:
		fallthrough
	case reflect.Uint8:
		fallthrough
	case reflect.Uint16:
		fallthrough
	case reflect.Uint32:
		fallthrough
	case reflect.Uint64:
		return true
	default:
		return false
	}
}

func isFloat(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Float32:
		fallthrough
	case reflect.Float64:
		return true
	default:
		return false
	}
}

func isString(v reflect.Value) bool {
	return v.Kind() == reflect.String
}

func getCompare(t reflect.Value) reflect.Value {
	return t.MethodByName("Compare")
}

func hasCompareMethod(v reflect.Value) bool {
	return getCompare(v).IsValid()
}

func canCompare(e interface{}) bool {
	v := reflect.ValueOf(e)
	return hasCompareMethod(v) || isInt(v) || isUInt(v) || isFloat(v) || isString(v)
}

func compare(i1, i2 interface{}) int {
	v1, v2 := reflect.ValueOf(i1), reflect.ValueOf(i2)
	switch {
	case v1.Kind() != v2.Kind():
		return 2
	case hasCompareMethod(v1):
		return int(getCompare(v1).Call([]reflect.Value{v2})[0].Int())
	case isInt(v1):
		return compareInt(v1.Int(), v2.Int())
	case isUInt(v1):
		return compareUInt(v1.Uint(), v2.Uint())
	case isFloat(v1):
		return compareFloat(v1.Float(), v2.Float())
	case isString(v1):
		return compareString(v1.String(), v2.String())
	default:
		return 2
	}
}

func SortBy(a interface{}, f func(interface{}, interface{}) int) interface{} {
	v := reflect.ValueOf(a)
	if v.Kind() != reflect.Slice {
		return nil
	}
	sort.Sort(sl{v, f})
	return v.Interface()
}

func Sort(a interface{}) interface{} {
	return SortBy(a, compare)
}

func Filter(a interface{}, f ...interface{}) interface{} {
	v := reflect.ValueOf(a)
	if v.Kind() != reflect.Slice {
		return nil
	}
	out := reflect.MakeSlice(v.Type(), 0, v.Cap())
	for i := 0; i < v.Len(); i++ {
		e := v.Index(i)
		if !Contains(f, e.Interface()) {
			out = reflect.Append(out, e)
		}
	}
	return out.Interface()
}

func Unic(a interface{}) interface{} {
	v := reflect.ValueOf(a)
	if v.Kind() != reflect.Slice {
		return nil
	}
	out := reflect.MakeSlice(v.Type(), 0, v.Cap())
	for i := 0; i < v.Len(); i++ {
		e := v.Index(i)
		if !Contains(out.Interface(), e.Interface()) {
			out = reflect.Append(out, e)
		}
	}
	return out.Interface()
}

func SortUnic(a interface{}) interface{} {
	return Sort(Unic(a))
}

func MkstringBy(a interface{}, beg, sep, end string, f func(interface{}) string) string {
	v := reflect.ValueOf(a)
	if v.Kind() != reflect.Slice {
		return ""
	}
	out := ""
	for i := 0; i < v.Len(); i++ {
		e := v.Index(i)
		Concat(&out, sep, f(e.Interface()))
	}
	return beg + out + end
}

func Mkstring(a interface{}, beg, sep, end string) string {
	return MkstringBy(a, beg, sep, end, StringOf)
}

func SliceStringBy(a interface{}, f func(interface{}) string) string {
	return MkstringBy(a, "[", "; ", "]", f)
}

func SliceString(a interface{}) string {
	return SliceStringBy(a, StringOf)
}
