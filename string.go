package validate

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/mail"
	"net/url"
	"regexp"
)

// StringValidate validator for String
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

// URL check if string is valid URL using url.ParseRequestURI, if not return ErrNotURL
var URL Rule = RuleFn(func(data interface{}) error {
	return MustBeString(data, func(s string) error {
		if _, err := url.ParseRequestURI(s); err != nil {
			return ErrNotURL
		}
		return nil
	})
})

// Email check if string is valid Email using mail.ParseAddress, if not return ErrNotEmail
var Email Rule = RuleFn(func(data interface{}) error {
	return MustBeString(data, func(s string) error {
		if _, err := mail.ParseAddress(s); err != nil {
			return ErrNotEmail
		}
		return nil
	})
})

// UUID check if string is valid UUID using uuid.Parse, if not return ErrNotUUID
var UUID Rule = RuleFn(func(data interface{}) error {
	return MustBeString(data, func(s string) error {
		if _, err := uuid.Parse(s); err != nil {
			return ErrNotUUID
		}
		return nil
	})
})

// JSONString check if string is a valid json, if not return ErrNotJSONString
var JSONString Rule = RuleFn(func(data interface{}) error {
	return MustBeString(data, func(s string) error {
		var js interface{}
		if err := json.Unmarshal([]byte(s), &js); err != nil {
			return ErrNotJSONString
		}
		return nil
	})
})

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
