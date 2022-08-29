package validate

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

// Common error
var (
	ErrRequired = errors.New("value required")
	ErrMustIn   = func(arr interface{}) error {
		return fmt.Errorf("value must be one of %v", arr)
	}
	ErrNotEqual = func(i interface{}) error {
		return fmt.Errorf("value must be equal to %v", i)
	}
	ErrEqual = func(i interface{}) error {
		return fmt.Errorf("value must not be %v", i)
	}
	ErrNotKind = func(k reflect.Kind) error {
		return fmt.Errorf("value kind must be %v", k.String())
	}
)

// Array error
var (
	ErrNotArray = errors.New("must be an array")
	ErrMinSize  = func(l int) error {
		return fmt.Errorf("size must be equal or greater than %v", l)
	}
	ErrMaxSize = func(l int) error {
		return fmt.Errorf("size must be equal or smaller than %v", l)
	}
)

// Number error

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

// String error

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
	ErrNotURL = errors.New("must be an url")

	ErrNotEmail = errors.New("must be an email address")

	ErrNotUUID = errors.New("string is not a valid UUID")

	ErrNotJSONString = errors.New("must be a json string")
)

// Struct error

var (
	ErrNotStruct = errors.New("must be an object")
)

// Time error

var (
	ErrTime         = errors.New("not a valid date")
	ErrNotInt64     = errors.New("must be int64")
	ErrNoTimeSource = errors.New("time source not chosen")
	ErrBefore       = func(t time.Time) error {
		return fmt.Errorf("time must before %v", t)
	}
	ErrAfter = func(t time.Time) error {
		return fmt.Errorf("time must after %v", t)
	}
)
