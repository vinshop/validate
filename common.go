package validate

import (
	"errors"
	"fmt"
	"reflect"
)

var (
	ErrRequired = errors.New("value required")
	ErrMustIn   = func(arr interface{}) error {
		return fmt.Errorf("value must be one of %v", arr)
	}
	ErrNotEqual = func(i interface{}) error {
		return fmt.Errorf("value must be equal to %v", i)
	}
	ErrEqual = func(i interface{}) error {
		return fmt.Errorf("value must not be %v", i)
	}
	ErrNotKind = func(k reflect.Kind) error {
		return fmt.Errorf("value kind must be %v", k.String())
	}
)

// Require check data is empty use IsZero method, if not return ErrRequired
var Require Rule = RuleFn(func(v interface{}) error {
	if IsZero(v) {
		return ErrRequired
	}
	return nil
})

// In check data is an element in array, if not return ErrMustIn
func In(arr interface{}) Rule {
	return RuleFn(func(data interface{}) error {
		return MustBeArray(arr, func(s reflect.Value) error {
			for i := 0; i < s.Len(); i++ {
				if reflect.DeepEqual(s.Index(i).Interface(), data) {
					return nil
				}
			}
			return ErrMustIn(arr)
		})
	})
}

// Equal check data equal to interface, if not return ErrNotEqual
func Equal(i interface{}) Rule {
	return RuleFn(func(data interface{}) error {
		if !reflect.DeepEqual(data, i) {
			return ErrNotEqual(i)
		}
		return nil
	})
}

// NotEqual check data equal to interface, if equal return ErrNotEqual
func NotEqual(i interface{}) Rule {
	return RuleFn(func(data interface{}) error {
		if reflect.DeepEqual(data, i) {
			return ErrEqual(i)
		}
		return nil
	})
}

// IsKind check data kind is given kind,if not return ErrNotKind
func IsKind(kind reflect.Kind) Rule {
	return RuleFn(func(data interface{}) error {
		if reflect.ValueOf(data).Kind() != kind {
			return ErrNotKind(kind)
		}
		return nil
	})
}
