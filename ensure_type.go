package validate

import (
	"reflect"
	"regexp"
	"time"
)

// MustBeArray check if data is array, if not return ErrNotArray
func MustBeArray(data interface{}, fn func(s reflect.Value) error) error {
	w := Wrap(data)
	if w.Value.Kind() != reflect.Array && w.Value.Kind() != reflect.Slice {
		return ErrNotArray
	}
	return fn(w.Value)
}

// MustBeInt64 check if data is int64, if not return ErrNotInt64
func MustBeInt64(data interface{}, fn func(i int64) error) error {
	w := Wrap(data)
	if w.Value.Kind() != reflect.Int64 {
		return ErrNotInt64
	}
	return fn(w.Value.Int())
}

// MustBeNumber check if data is Number, if not return ErrNotNumber
func MustBeNumber(data interface{}, fn func(a float64) error) error {
	w := Wrap(data)
	if w.Value.Kind() == reflect.Float64 {
		return fn(w.Value.Float())
	}

	if !w.Value.Type().ConvertibleTo(floatType) {
		return ErrNotNumber
	}
	number := w.Value.Convert(floatType).Float()
	return fn(number)
}

// MustBeString check if data is String, if not return ErrNotString
func MustBeString(data interface{}, fn func(s string) error) error {
	w := Wrap(data)
	if w.Value.Kind() != reflect.String {
		return ErrNotString
	}
	s := w.Value.String()
	return fn(s)
}

// MustBeRegex check if data is regexp.Regexp, if not return err
func MustBeRegex(data string, fn func(r *regexp.Regexp) error) error {
	regex, err := regexp.Compile(data)
	if err != nil {
		return err
	}
	return fn(regex)
}

// MustBeStruct check if data is a struct, if not return ErrNotStruct
func MustBeStruct(data interface{}, fn func(data reflect.Value) error) error {
	w := Wrap(data)
	if w.Value.Kind() != reflect.Struct {
		return ErrNotStruct
	}
	return fn(w.Value)
}

// MustBeTime check if data is time.Time, if not return ErrTime
func MustBeTime(data interface{}, fn func(t time.Time) error) error {
	if data == nil {
		return ErrTime
	}
	w := Wrap(data)
	date, ok := w.Value.Interface().(time.Time)
	if !ok {
		return ErrTime
	}
	return fn(date)
}
