package validate

import (
	"errors"
	"fmt"
	"reflect"
)

type StructValidator struct {
	key  string
	data reflect.Value
	fns  map[string][]Validate
	opts []Validate
}

type FieldError struct {
	Key  string
	Name string
	Err  error
}

func (f *FieldError) Error() string {
	return fmt.Sprintf("[%v]: %v", f.Name, f.Err.Error())
}

func (v StructValidator) Validate() error {
	for i := 0; i < v.data.NumField(); i++ {
		fName := v.data.Type().Field(i).Name
		fns := v.fns[fName]
		if len(fns) == 0 {
			continue
		}
		for _, fn := range fns {
			if err := fn(v.data.Field(i).Interface()); err != nil {
				return &FieldError{v.key, fName, err}
			}
		}
	}
	return nil
}

type StructFn func(v *StructValidator)

func Register(name string, fns ...Validate) StructFn {
	return func(v *StructValidator) {
		v.fns[name] = append(v.fns[name], fns...)
	}
}

func WithKey(key string) StructFn {
	return func(v *StructValidator) {
		v.key = key
	}
}

func Struct(fns ...StructFn) Validate {
	return func(data interface{}) error {
		return mustBeStruct(data, func(data reflect.Value) error {
			v := &StructValidator{
				data: data,
				fns:  make(map[string][]Validate),
			}
			for _, fn := range fns {
				fn(v)
			}
			return v.Validate()
		})
	}
}

var (
	ErrNotStruct = errors.New("must be an object")
)

func mustBeStruct(data interface{}, fn func(data reflect.Value) error) error {
	v := reflect.ValueOf(data)
	for v.Kind() == reflect.Pointer {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return ErrNotStruct
	}
	return fn(v)
}
