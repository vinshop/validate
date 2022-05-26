package validate

import (
	"errors"
)

var (
	ErrRequired = errors.New("must not be empty")
)

// Require check if data is empty use IsZero method, if not return ErrRequired
var Require Rule = RuleFn(func(v interface{}) error {
	if IsZero(v) {
		return ErrRequired
	}
	return nil
})
