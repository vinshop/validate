package validate

import (
	"github.com/stretchr/testify/assert"
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
