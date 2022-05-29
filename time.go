package validate

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

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

func MustBeTime(data interface{}, fn func(t time.Time) error) error {
	if data == nil {
		return ErrTime
	}
	v := reflect.Indirect(reflect.ValueOf(data))
	date, ok := v.Interface().(time.Time)
	if !ok {
		return ErrTime
	}
	return fn(date)
}

type timeSource int

const (
	fromString timeSource = iota + 1
	fromSecond
	fromNano
)

type TimeValidator struct {
	source timeSource
	layout string
	loc    *time.Location
	fns    Rules
}

func (t *TimeValidator) Do(data interface{}) error {
	switch t.source {
	case fromString:
		if t.layout == "" {
			t.layout = time.RFC3339
		}
		if t.loc == nil {
			t.loc = time.Local
		}
		return MustBeString(data, func(s string) error {
			date, err := time.ParseInLocation(t.layout, s, t.loc)
			if err != nil {
				return err
			}
			return t.fns.Do(date)
		})
	case fromSecond:
		return MustBeInt64(data, func(i int64) error {
			date := time.Unix(i, 0)
			return t.fns.Do(date)
		})
	case fromNano:
		return MustBeInt64(data, func(i int64) error {
			date := time.Unix(0, i)
			return t.fns.Do(date)
		})
	default:
		return ErrNoTimeSource
	}
}

// Date parse date from layout string and loc, default layout is time.RFC3339, default loc is time.Local
func Date(layout string, loc *time.Location, fns ...Rule) *TimeValidator {
	return &TimeValidator{
		source: fromString,
		layout: layout,
		loc:    loc,
		fns:    fns,
	}
}

// Second parse date from second
func Second(fns ...Rule) *TimeValidator {
	return &TimeValidator{
		source: fromSecond,
		layout: "",
		loc:    nil,
		fns:    fns,
	}
}

// Nano parse date from nano
func Nano(fns ...Rule) *TimeValidator {
	return &TimeValidator{
		source: fromNano,
		layout: "",
		loc:    nil,
		fns:    fns,
	}
}

// Before check if data is time instance before given time, if not return ErrBefore
func Before(before time.Time) Rule {
	return RuleFn(func(data interface{}) error {
		return MustBeTime(data, func(t time.Time) error {
			if t.Before(before) {
				return nil
			}
			return ErrBefore(before)
		})
	})
}

// After check if data is time instance after given time, if not return ErrAfter
func After(after time.Time) Rule {
	return RuleFn(func(data interface{}) error {
		return MustBeTime(data, func(t time.Time) error {
			if t.After(after) {
				return nil
			}
			return ErrAfter(after)
		})
	})
}

// ChangeTime modify the data time instance
func ChangeTime(fn func(t time.Time) time.Time, fns ...Rule) Rule {
	return RuleFn(func(data interface{}) error {
		return MustBeTime(data, func(t time.Time) error {
			t = fn(t)
			return Use(t, fns...).Validate()
		})
	})
}
