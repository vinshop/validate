package validate

import (
	"errors"
	"fmt"
	"reflect"
)

type ArrayValidate struct {
	each []Rule
	all  []Rule
}

func (a ArrayValidate) Do(data interface{}) error {
	for _, fn := range a.all {
		if err := fn.Do(data); err != nil {
			return err
		}
	}
	return MustBeArray(data, func(data reflect.Value) error {
		for i := 0; i < data.Len(); i++ {
			for _, fn := range a.each {
				if err := fn.Do(data.Index(i).Interface()); err != nil {
					return ArrayError(i, err)
				}
			}
		}
		return nil
	})

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

func Array(fns ...ArrayFn) Rule {
	v := &ArrayValidate{
		each: make(Rules, 0),
		all:  make(Rules, 0),
	}
	for _, fn := range fns {
		fn(v)
	}
	return v
}

func Each(fns ...Rule) ArrayFn {
	return func(v *ArrayValidate) {
		v.each = append(v.each, fns...)
	}
}

func ArrayHas(fns ...Rule) ArrayFn {
	return func(v *ArrayValidate) {
		v.all = append(v.all, fns...)
	}
}

func MinSize(l int) Rule {
	return RuleFn(func(data interface{}) error {
		return MustBeArray(data, func(s reflect.Value) error {
			if s.Len() < l {
				return ErrMinSize(l)
			}
			return nil
		})
	})
}

func MaxSize(l int) Rule {
	return RuleFn(func(data interface{}) error {
		return MustBeArray(data, func(s reflect.Value) error {
			if s.Len() > l {
				return ErrMaxSize(l)
			}
			return nil
		})
	})
}

func MustBeArray(data interface{}, fn func(s reflect.Value) error) error {
	v := reflect.ValueOf(data)
	for v.Kind() == reflect.Pointer {
		v = v.Elem()
	}
	if v.Kind() != reflect.Array && v.Kind() != reflect.Slice {
		return ErrNotArray
	}
	return fn(v)
}
