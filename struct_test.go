package validate

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"reflect"
	"testing"
)

type mockStruct struct {
	A interface{}
}

func (m mockStruct) Key() string {
	return "key"
}

func TestMustBeStruct(t *testing.T) {
	t.Parallel()
	tests := []testCase{
		{"not struct", "abc", ErrNotStruct},
		{"is struct", mockStruct{}, nil},
		{"is struct ptr", &mockStruct{}, nil},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expect, MustBeStruct(test.value, func(data reflect.Value) error {
				return nil
			}))
		})
	}
}

func TestTructNoValidate(t *testing.T) {
	t.Parallel()
	fns := Struct(Field("A"))
	tests := []testCase{
		{"not struct", 123, ErrNotStruct},
		{"is struct", mockStruct{}, nil},
		{"is struct ptr", &mockStruct{}, nil},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expect, Use(test.value, fns).Validate())
		})
	}
}

func TestStructValidate(t *testing.T) {
	t.Parallel()
	fns := Struct(WithKey("abc"), Field("A", Require), Field("A", String(MinLength(5))))
	tests := []testCase{
		{
			"empty A", mockStruct{}, FieldError("abc", "A", ErrRequired),
		},
		{
			"A not string", mockStruct{124}, FieldError("abc", "A", ErrNotString),
		},
		{
			"A min length error", mockStruct{"1234"}, FieldError("abc", "A", ErrMinLength(5)),
		},
		{
			"A success", mockStruct{"12345"}, nil,
		},
		{
			"empty ptr A", &mockStruct{}, FieldError("abc", "A", ErrRequired),
		},
		{
			"ptr A not string", &mockStruct{124}, FieldError("abc", "A", ErrNotString),
		},
		{
			"ptr A min length error", &mockStruct{"1234"}, FieldError("abc", "A", ErrMinLength(5)),
		},
		{
			"ptr A success", &mockStruct{"12345"}, nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expect, Use(test.value, fns).Validate())
		})
	}
}

func randomString() string {
	length := rand.Intn(100)
	b := make([]byte, length)
	rand.Read(b)
	return string(b)
}
