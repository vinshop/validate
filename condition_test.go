package validate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConditional_Else(t *testing.T) {
	t.Parallel()
	tests := []testCase{
		{
			"true", true, ErrNotArray},
		{
			"false", false, ErrNotString,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expect, Use(test.value, If().Then(Array()).Else(String())).Validate())
		})
	}
}

func TestSwitchCase(t *testing.T) {
	t.Parallel()
	fns := Switch().Case("A", Array()).CaseMany([]interface{}{1, 2, 3}, String()).Default(Require)
	tests := []testCase{
		{"not array", "A", ErrNotArray},
		{"number", 1, ErrNotString},
		{"required", "", ErrRequired},
		{"default", "ABC", nil},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expect, Use(test.value, fns).Validate())
		})
	}
}
