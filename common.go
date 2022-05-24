package validate

import (
	"errors"
)

var (
	ErrRequired = errors.New("must not be empty")
)

var Require Validate = func(v interface{}) error {
	if IsZero(v) {
		return ErrRequired
	}
	return nil
}
