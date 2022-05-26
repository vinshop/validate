package validate

import (
	"errors"
	"reflect"
)

type StructValidator struct {
	key  string
	data reflect.Value
	fns  map[string][]Rule
	opts []Rule
}

type Keyable interface {
	Key() string
}

func (v StructValidator) Validate() error {
	for i := 0; i < v.data.NumField(); i++ {
		fName := v.data.Type().Field(i).Name
		fns := v.fns[fName]
		if len(fns) == 0 {
			continue
		}
		for _, fn := range fns {
			if err := fn.Do(v.data.Field(i).Interface()); err != nil {
				return FieldError(v.key, fName, err)
			}
		}
	}
	return nil
}

type StructFn func(v *StructValidator)

func Field(name string, fns ...Rule) StructFn {
	return func(v *StructValidator) {
		v.fns[name] = append(v.fns[name], fns...)
	}
}

func WithKey(key string) StructFn {
	return func(v *StructValidator) {
		v.key = key
	}
}

func Struct(fns ...StructFn) Rule {
	return RuleFn(func(data interface{}) error {
		return mustBeStruct(data, func(data reflect.Value) error {
			v := &StructValidator{
				data: data,
				fns:  make(map[string][]Rule),
			}
			keyable, ok := data.Interface().(Keyable)
			if ok {
				v.key = keyable.Key()
			}
			for _, fn := range fns {
				fn(v)
			}
			return v.Validate()
		})
	})
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
