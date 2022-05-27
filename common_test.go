package validate

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestRequire(t *testing.T) {
	t.Parallel()
	tests := []testCase{
		{"empty array", []string{}, ErrRequired},
		{"empty string", "", ErrRequired},
		{"empty number", 0, ErrRequired},
		{"nil", nil, ErrRequired},
		{"number", 1, nil},
		{"string", "abc", nil},
		{"arr", []interface{}{1, "2", 3}, nil},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expect, Use(test.value, Require).Validate())
		})
	}
}

func TestIn(t *testing.T) {
	t.Parallel()
	arr := []interface{}{1, "2", []interface{}{1, 2}}
	fns := In(arr)
	tests := []testCase{
		{"not include", "1", ErrMustIn(arr)},
		{"number", 1, nil},
		{"string", "2", nil},
		{"arr", []interface{}{1, 2}, nil},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expect, Use(test.value, fns).Validate())
		})
	}
}

func TestEqual(t *testing.T) {
	t.Parallel()
	value := "123"
	fns := Equal(value)
	tests := []testCase{
		{"nil", nil, ErrNotEqual(value)},
		{"not equal diff type", 123, ErrNotEqual(value)},
		{"not equal", "1234", ErrNotEqual(value)},
		{"equal", "123", nil},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expect, Use(test.value, fns).Validate())
		})
	}
}

func TestNotEqual(t *testing.T) {
	t.Parallel()
	value := "123"
	fns := NotEqual(value)
	tests := []testCase{
		{"nil", nil, nil},
		{"not equal diff type", 123, nil},
		{"not equal", "1234", nil},
		{"equal", "123", ErrEqual(value)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expect, Use(test.value, fns).Validate())
		})
	}
}

func TestIsKind(t *testing.T) {
	t.Parallel()
	value := reflect.String
	fns := IsKind(value)
	tests := []testCase{
		{"nil", nil, ErrNotKind(value)},
		{"number", 123, ErrNotKind(value)},
		{"string", "1234", nil},
		{"struct", mockStruct{}, ErrNotKind(value)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expect, Use(test.value, fns).Validate())
		})
	}
}
