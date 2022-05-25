package validate

import (
	"errors"
	"fmt"
	"reflect"
)

type ArrayValidate struct {
	data reflect.Value
	each []Validate
	all  []Validate
}

func (a *ArrayValidate) Validate() error {
	for _, fn := range a.all {
		if err := fn(a.data.Interface()); err != nil {
			return err
		}
	}
	for i := 0; i < a.data.Len(); i++ {
		for _, fn := range a.each {
			if err := fn(a.data.Index(i).Interface()); err != nil {
				return ArrayError(i, err)
			}
		}
	}
	return nil
}

var (
	ErrNotArray = errors.New("must be an array")
	ErrMinSize  = func(l int) error {
		return fmt.Errorf("size must be equal or greater than %v", l)
	}
	ErrMaxSize = func(l int) error {
		return fmt.Errorf("size must be equal or smaller than %v", l)
	}
)

type ArrayFn func(v *ArrayValidate)

func Array(fns ...ArrayFn) Validate {
	return func(data interface{}) error {
		return mustBeArray(data, func(data reflect.Value) error {
			v := &ArrayValidate{
				data: data,
				each: make([]Validate, 0),
				all:  make([]Validate, 0),
			}
			for _, fn := range fns {
				fn(v)
			}
			return v.Validate()
		})
	}
}

func Each(fns ...Validate) ArrayFn {
	return func(v *ArrayValidate) {
		v.each = append(v.each, fns...)
	}
}

func ArrayHas(fns ...Validate) ArrayFn {
	return func(v *ArrayValidate) {
		v.all = append(v.all, fns...)
	}
}

func MinSize(l int) Validate {
	return func(data interface{}) error {
		return mustBeArray(data, func(s reflect.Value) error {
			if s.Len() < l {
				return ErrMinSize(l)
			}
			return nil
		})
	}
}

func MaxSize(l int) Validate {
	return func(data interface{}) error {
		return mustBeArray(data, func(s reflect.Value) error {
			if s.Len() > l {
				return ErrMaxSize(l)
			}
			return nil
		})
	}
}

func mustBeArray(data interface{}, fn func(s reflect.Value) error) error {
	v := reflect.ValueOf(data)
	for v.Kind() == reflect.Pointer {
		v = v.Elem()
	}
	if v.Kind() != reflect.Array && v.Kind() != reflect.Slice {
		return ErrNotArray
	}
	return fn(v)
}
