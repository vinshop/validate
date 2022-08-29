package validate

import (
	"math"
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

// Number create new Number validator
func Number(fns ...Rule) Rule {
	return &NumberValidate{fns: fns}
}

var floatType = reflect.TypeOf(float64(0))

// EQ check if number == n, if not return ErrEQ
func EQ(n float64) Rule {
	return RuleFn(func(data interface{}) error {
		return MustBeNumber(data, func(a float64) error {
			if cmp(a, n) != 0 {
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
			if cmp(a, n) == 0 {
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
			if cmp(a, n) != -1 {
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
			if cmp(a, n) == 1 {
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
			if cmp(a, n) != 1 {
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
			if cmp(a, n) == -1 {
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

// cmp compare 2 number, if a equal to b return 0, if a < b return -1, else return 1
func cmp(a, b float64) int {
	if math.Abs(a-b) < NumberEpsilon {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}
