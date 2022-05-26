package validate

import (
	"errors"
	"reflect"
)

type StructValidator struct {
	key  string
	fns  map[string]Rules
	opts Rules
}

type Keyable interface {
	Key() string
}

func (v *StructValidator) Do(data interface{}) error {
	if v.key == "" {
		keyable, ok := data.(Keyable)
		if ok {
			v.key = keyable.Key()
		}
	}
	return MustBeStruct(data, func(data reflect.Value) error {
		for i := 0; i < data.NumField(); i++ {
			fName := data.Type().Field(i).Name
			fns := v.fns[fName]
			if len(fns) == 0 {
				continue
			}
			for _, fn := range fns {
				if err := fn.Do(data.Field(i).Interface()); err != nil {
					return FieldError(v.key, fName, err)
				}
			}
		}
		return nil
	})
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
	v := &StructValidator{
		key: "",
		fns: make(map[string]Rules),
	}
	for _, fn := range fns {
		fn(v)
	}

	return v
}

var (
	ErrNotStruct = errors.New("must be an object")
)

func MustBeStruct(data interface{}, fn func(data reflect.Value) error) error {
	v := reflect.ValueOf(data)
	for v.Kind() == reflect.Pointer {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return ErrNotStruct
	}
	return fn(v)
}
