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
