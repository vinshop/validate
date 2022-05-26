package validate

import (
	"errors"
	"fmt"
	"regexp"
)

//StringValidate validator for String
type StringValidate struct {
	data interface{}
	fns  []Rule
}

func (s StringValidate) Validate() error {
	for _, fn := range s.fns {
		if err := fn.Do(s.data); err != nil {
			return err
		}
	}
	return nil
}

var (
	ErrNotString = errors.New("must be a string")
	ErrMaxLength = func(l int) error {
		return fmt.Errorf("string length must not be greater than %v", l)
	}
	ErrMinLength = func(l int) error {
		return fmt.Errorf("string length must be greater than %v", l)
	}
	ErrRegexNotMatch = func(regex string) error {
		return fmt.Errorf("string not match regex %v", regex)
	}
)

// String create new StringValidate
func String(fns ...Rule) Rule {
	return RuleFn(func(data interface{}) error {
		return MustBeString(data, func(s string) error {
			return StringValidate{
				data: data,
				fns:  fns,
			}.Validate()
		})
	})
}

// MaxLength check if string has max length of l, if not return ErrMaxLength
func MaxLength(l int) Rule {
	return RuleFn(func(data interface{}) error {
		return MustBeString(data, func(s string) error {
			if len(s) > l {
				return ErrMaxLength(l)
			}
			return nil
		})
	})
}

// MinLength check if string has min length of l, if not return ErrMinLength
func MinLength(l int) Rule {
	return RuleFn(func(data interface{}) error {
		return MustBeString(data, func(s string) error {
			if len(s) < l {
				return ErrMinLength(l)
			}
			return nil
		})
	})
}

// Match check if string match regex, if not return ErrRegexNotMatch
func Match(regex string) Rule {
	return RuleFn(func(data interface{}) error {
		return MustBeString(data, func(s string) error {
			return MustBeRegex(regex, func(r *regexp.Regexp) error {
				if !r.MatchString(s) {
					return ErrRegexNotMatch(regex)
				}
				return nil
			})
		})
	})
}

// StringCustom custom string validator
func StringCustom(fn func(s string) error) Rule {
	return RuleFn(func(data interface{}) error {
		return MustBeString(data, fn)
	})
}

// MustBeString check if data is String, if not return ErrNotString
func MustBeString(data interface{}, fn func(s string) error) error {
	s, ok := data.(string)
	if !ok {
		return ErrNotString
	}
	return fn(s)
}

// MustBeRegex check if data is regexp.Regexp, if not return err
func MustBeRegex(data string, fn func(r *regexp.Regexp) error) error {
	regex, err := regexp.Compile(data)
	if err != nil {
		return err
	}
	return fn(regex)
}
