package validate

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

var (
	ErrNotDate      = errors.New("not a valid date")
	ErrNotInt64     = errors.New("must be int64")
	ErrNoTimeSource = errors.New("time source not chosen")
	ErrBefore       = func(t time.Time) error {
		return fmt.Errorf("time must before %v", t)
	}
	ErrAfter = func(t time.Time) error {
		return fmt.Errorf("time must after %v", t)
	}
)

var timeKind = reflect.TypeOf(time.Time{}).Kind()

func MustBeInt64(data interface{}, fn func(i int64) error) error {
	v := reflect.ValueOf(data)
	v = reflect.Indirect(v)
	if v.Kind() != reflect.Int64 {
		return ErrNotInt64
	}
	return fn(v.Int())
}

func MustBeTime(data interface{}, fn func(t time.Time) error) error {
	v := reflect.Indirect(reflect.ValueOf(data))
	date, ok := v.Interface().(time.Time)
	if !ok {
		return ErrNotDate
	}
	return fn(date)
}

type timeSource int

const (
	fromString timeSource = iota
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

func Date(layout string, loc *time.Location, fns ...Rule) *TimeValidator {
	return &TimeValidator{
		source: fromString,
		layout: layout,
		loc:    loc,
		fns:    fns,
	}
}

func Second(fns ...Rule) *TimeValidator {
	return &TimeValidator{
		source: fromSecond,
		layout: "",
		loc:    nil,
		fns:    fns,
	}
}

func Nano(fns ...Rule) *TimeValidator {
	return &TimeValidator{
		source: fromNano,
		layout: "",
		loc:    nil,
		fns:    fns,
	}
}

func Before(before time.Time) Rule {
	return RuleFn(func(data interface{}) error {
		return MustBeTime(data, func(t time.Time) error {
			if !t.Before(before) {
				return ErrBefore(before)
			}
			return nil
		})
	})
}

func After(after time.Time) Rule {
	return RuleFn(func(data interface{}) error {
		return MustBeTime(data, func(t time.Time) error {
			if !t.After(after) {
				return ErrAfter(after)
			}
			return nil
		})
	})
}
