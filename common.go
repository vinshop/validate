package validate

import (
	"reflect"
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
				if reflect.DeepEqual(s.Index(i).Interface(), Wrap(data).Data) {
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
		if !reflect.DeepEqual(Wrap(data).Data, Wrap(i).Data) {
			return ErrNotEqual(i)
		}
		return nil
	})
}

// NotEqual check data equal to interface, if equal return ErrNotEqual
func NotEqual(i interface{}) Rule {
	return RuleFn(func(data interface{}) error {
		if reflect.DeepEqual(Wrap(data).Data, Wrap(i).Data) {
			return ErrEqual(i)
		}
		return nil
	})
}

// IsKind check data kind is given kind,if not return ErrNotKind
func IsKind(kind reflect.Kind) Rule {
	return RuleFn(func(data interface{}) error {
		if Wrap(data).Value.Kind() != kind {
			return ErrNotKind(kind)
		}
		return nil
	})
}

// Optional only validate if data is not empty
func Optional(fns ...Rule) Rule {
	return RuleFn(func(data interface{}) error {
		if IsZero(data) {
			return nil
		}
		return Rules(fns).Do(data)
	})
}
