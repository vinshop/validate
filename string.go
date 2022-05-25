package validate

import (
	"errors"
	"fmt"
	"regexp"
)

type StringValidate struct {
	data interface{}
	fns  []Validate
}

func (s StringValidate) Validate() error {
	for _, fn := range s.fns {
		if err := fn(s.data); err != nil {
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

func String(fns ...Validate) Validate {
	return func(data interface{}) error {
		return mustBeString(data, func(s string) error {
			return StringValidate{
				data: data,
				fns:  fns,
			}.Validate()
		})
	}
}

func MaxLength(l int) Validate {
	return func(data interface{}) error {
		return mustBeString(data, func(s string) error {
			if len(s) > l {
				return ErrMaxLength(l)
			}
			return nil
		})
	}
}

func MinLength(l int) Validate {
	return func(data interface{}) error {
		return mustBeString(data, func(s string) error {
			if len(s) < l {
				return ErrMinLength(l)
			}
			return nil
		})
	}
}

func Match(regex string) Validate {
	return func(data interface{}) error {
		return mustBeString(data, func(s string) error {
			return mustBeRegex(regex, func(r *regexp.Regexp) error {
				if !r.MatchString(s) {
					return ErrRegexNotMatch(regex)
				}
				return nil
			})
		})
	}
}

func StringCustom(fn func(s string) error) Validate {
	return func(data interface{}) error {
		return mustBeString(data, fn)
	}
}

func mustBeString(data interface{}, fn func(s string) error) error {
	s, ok := data.(string)
	if !ok {
		return ErrNotString
	}
	return fn(s)
}

func mustBeRegex(data string, fn func(r *regexp.Regexp) error) error {
	regex, err := regexp.Compile(data)
	if err != nil {
		return err
	}
	return fn(regex)
}
