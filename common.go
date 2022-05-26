package validate

import (
	"errors"
)

var (
	ErrRequired = errors.New("must not be empty")
)

var Require Rule = RuleFn(func(v interface{}) error {
	if IsZero(v) {
		return ErrRequired
	}
	return nil
})
