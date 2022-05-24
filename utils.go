package validate

import "reflect"

func IsZero(v interface{}) bool {
	if v == nil {
		return true
	}
	r := reflect.ValueOf(v)
	if r.Kind() == reflect.Slice || r.Kind() == reflect.Array {
		return r.Len() == 0
	}
	return reflect.DeepEqual(v, reflect.Zero(reflect.TypeOf(v)).Interface())
}
