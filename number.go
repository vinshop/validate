package validate

import (
	"errors"
	"fmt"
	"reflect"
)

// NumberValidate Number validator
type NumberValidate struct {
	fns Rules
}

func (v *NumberValidate) Do(data interface{}) error {
	return MustBeNumber(data, func(a float64) error {
		return v.fns.Do(a)
	})
}

var (
	ErrNotNumber = errors.New("not a number")
	ErrNEQ       = func(n float64) error {
		return fmt.Errorf("must not be equal to %v", n)
	}
	ErrEQ = func(n float64) error {
		return fmt.Errorf("must be equal to %v", n)
	}
	ErrLT = func(n float64) error {
		return fmt.Errorf("must be less than %v", n)
	}
	ErrGT = func(n float64) error {
		return fmt.Errorf("must be greater than %v", n)
	}
	ErrLTE = func(n float64) error {
		return fmt.Errorf("must be equal or less than %v", n)
	}
	ErrGTE = func(n float64) error {
		return fmt.Errorf("must be equal or greater than %v", n)
	}
)

// Number create new Number validator
func Number(fns ...Rule) Rule {
	return &NumberValidate{fns: fns}
}

var floatType = reflect.TypeOf(float64(0))

// MustBeInt64 check if data is int64, if not return ErrNotInt64
func MustBeInt64(data interface{}, fn func(i int64) error) error {
	v := reflect.ValueOf(data)
	v = reflect.Indirect(v)
	if v.Kind() != reflect.Int64 {
		return ErrNotInt64
	}
	return fn(v.Int())
}

// MustBeNumber check if data is Number, if not return ErrNotNumber
func MustBeNumber(data interface{}, fn func(a float64) error) error {
	v := reflect.ValueOf(data)
	if v.Kind() == reflect.Float64 {
		return fn(v.Float())
	}
	v = reflect.Indirect(v)
	if !v.Type().ConvertibleTo(floatType) {
		return ErrNotNumber
	}
	number := v.Convert(floatType).Float()
	return fn(number)
}

// EQ check if number == n, if not return ErrEQ
func EQ(n float64) Rule {
	return RuleFn(func(data interface{}) error {
		return MustBeNumber(data, func(a float64) error {
			if a != n {
				return ErrEQ(n)
			}
			return nil
		})
	})
}

// NEQ check if number != n, if not return ErrNEQ
func NEQ(n float64) Rule {
	return RuleFn(func(data interface{}) error {
		return MustBeNumber(data, func(a float64) error {
			if a == n {
				return ErrNEQ(n)
			}
			return nil
		})
	})
}

// LT check if number < n, if not return ErrLT
func LT(n float64) Rule {
	return RuleFn(func(data interface{}) error {
		return MustBeNumber(data, func(a float64) error {
			if a >= n {
				return ErrLT(n)
			}
			return nil
		})
	})
}

// LTE check if number <= n, if not return ErrLTE
func LTE(n float64) Rule {
	return RuleFn(func(data interface{}) error {
		return MustBeNumber(data, func(a float64) error {
			if a > n {
				return ErrLTE(n)
			}
			return nil
		})
	})
}

// GT check if number > n, if not return ErrGT
func GT(n float64) Rule {
	return RuleFn(func(data interface{}) error {
		return MustBeNumber(data, func(a float64) error {
			if a <= n {
				return ErrGT(n)
			}
			return nil
		})
	})
}

// GTE check if number >= n, if not return ErrGTE
func GTE(n float64) Rule {
	return RuleFn(func(data interface{}) error {
		return MustBeNumber(data, func(a float64) error {
			if a < n {
				return ErrGTE(n)
			}
			return nil
		})
	})
}

// DoMath do some math to the number then validate
func DoMath(fn func(n float64) float64, fns ...Rule) Rule {
	return RuleFn(func(data interface{}) error {
		return MustBeNumber(data, func(a float64) error {
			a = fn(a)
			return Rules(fns).Do(a)
		})
	})
}

// CustomNumber custom validator for number
func CustomNumber(fn func(n float64) error) Rule {
	return RuleFn(func(data interface{}) error {
		return MustBeNumber(data, func(a float64) error {
			return fn(a)
		})
	})
}
