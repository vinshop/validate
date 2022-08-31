package validate

import "reflect"

// IsZero check if data is zero or not
func IsZero(v interface{}) bool {
	w := Wrap(v)
	if w.Data == nil {
		return true
	}
	if w.Value.Kind() == reflect.Slice || w.Value.Kind() == reflect.Array {
		return w.Value.Len() == 0
	}
	return reflect.DeepEqual(w.Data, reflect.Zero(w.Type).Interface())
}

// BoolFunc function return boolean value, to make sure it's not panic
type BoolFunc func() bool
