package validate

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"regexp"
	"regexp/syntax"
	"testing"
)

type testCase struct {
	name   string
	value  interface{}
	expect error
}

func TestMustBeRegex(t *testing.T) {
	t.Parallel()
	tests := []testCase{
		{"not regex", `\l`, &syntax.Error{
			Code: syntax.ErrInvalidEscape,
			Expr: `\l`,
		}},
		{
			"regex", `\d+`, nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expect, mustBeRegex(test.value.(string), func(r *regexp.Regexp) error {
				return nil
			}))
		})
	}
}

func TestMustBeString(t *testing.T) {
	t.Parallel()
	tests := []testCase{
		{"not string", 123, ErrNotString},
		{"string", "abc", nil},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expect, mustBeString(test.value, func(s string) error {
				return nil
			}))
		})
	}
}

func TestMinLength(t *testing.T) {
	t.Parallel()
	l := 5
	fns := String(MinLength(l))
	tests := []testCase{
		{"empty", "", ErrMinLength(l)},
		{"smaller", "1234", ErrMinLength(l)},
		{"equal", "12345", nil},
		{"greater", "123456", nil},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expect, Use(test.value, fns).Validate())
		})
	}
}

func TestMaxLength(t *testing.T) {
	t.Parallel()
	l := 5
	fns := String(MaxLength(l))
	tests := []testCase{
		{"empty", "", nil},
		{"smaller", "1234", nil},
		{"equal", "12345", nil},
		{"greater", "1234567", ErrMaxLength(l)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expect, Use(test.value, fns).Validate())
		})
	}
}

func TestMatch(t *testing.T) {
	t.Parallel()
	regex := `^\d+$`
	fns := String(Match(regex))
	tests := []testCase{
		{"empty", "", ErrRegexNotMatch(regex)},
		{"not match", "abc123", ErrRegexNotMatch(regex)},
		{"match", "12345", nil},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expect, Use(test.value, fns).Validate())
		})
	}
}

func TestStringCustom(t *testing.T) {
	t.Parallel()
	err := errors.New("not abc")
	fns := String(StringCustom(func(s string) error {
		if s != "abc" {
			return err
		}
		return nil
	}))
	tests := []testCase{
		{"empty", "", err},
		{"not abc", "1234", err},
		{"abc", "abc", nil},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expect, Use(test.value, fns).Validate())
		})
	}
}
