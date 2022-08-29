package validate

import (
	"reflect"
)

// ArrayValidate validator for array
type ArrayValidate struct {
	each []Rule
	all  []Rule
}

// ArrayFn Array function
type ArrayFn func(v *ArrayValidate)

// Array check if data is array
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

// Each check each element in array
func Each(fns ...Rule) ArrayFn {
	return func(v *ArrayValidate) {
		v.each = append(v.each, fns...)
	}
}

// ArrayHas check array info of array like MinSize, MaxSize, etc...
func ArrayHas(fns ...Rule) ArrayFn {
	return func(v *ArrayValidate) {
		v.all = append(v.all, fns...)
	}
}

// MinSize check array has length >= l, if not return ErrMinSize
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

// MaxSize check array has length <= l, if not return ErrMaxSize.
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
