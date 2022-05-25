package validate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinSize(t *testing.T) {
	t.Parallel()
	l := 5
	fns := Array(ArrayHas(MinSize(l)))
	tests := []testCase{
		{"empty", []interface{}{}, ErrMinSize(l)},
		{"smaller", []interface{}{1, 2, 3, 4}, ErrMinSize(l)},
		{"equal", []interface{}{1, 2, 3, 4, 5}, nil},
		{"greater", []interface{}{1, 2, 3, 4, 5, 6}, nil},
		{"ptr empty", &[]interface{}{}, ErrMinSize(l)},
		{"ptr smaller", &[]interface{}{1, 2, 3, 4}, ErrMinSize(l)},
		{"ptr equal", &[]interface{}{1, 2, 3, 4, 5}, nil},
		{"ptr greater", &[]interface{}{1, 2, 3, 4, 5, 6}, nil},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expect, Use(test.value, fns).Validate())
		})
	}
}

func TestMaxSize(t *testing.T) {
	t.Parallel()
	l := 5
	fns := Array(ArrayHas(MaxSize(l)))
	tests := []testCase{
		{"empty", []interface{}{}, nil},
		{"smaller", []interface{}{1, 2, 3, 4}, nil},
		{"equal", []interface{}{1, 2, 3, 4, 5}, nil},
		{"greater", []interface{}{1, 2, 3, 4, 5, 6}, ErrMaxSize(l)},
		{"ptr empty", &[]interface{}{}, nil},
		{"ptr smaller", &[]interface{}{1, 2, 3, 4}, nil},
		{"ptr equal", &[]interface{}{1, 2, 3, 4, 5}, nil},
		{"ptr greater", &[]interface{}{1, 2, 3, 4, 5, 6}, ErrMaxSize(l)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expect, Use(test.value, fns).Validate())
		})
	}
}

func TestEach(t *testing.T) {
	t.Parallel()
	fns := Array(Each(String()))
	tests := []testCase{
		{"not arr", "abc", ErrNotArray},
		{"empty", []interface{}{}, nil},
		{"all string", []interface{}{"A", "B"}, nil},
		{"not all string", []interface{}{"A", 1, 2}, ArrayError(1, ErrNotString)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expect, Use(test.value, fns).Validate())
		})
	}
}
