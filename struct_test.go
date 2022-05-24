package validate

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

type mockStruct struct {
	A interface{}
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
			assert.Equal(t, test.expect, mustBeStruct(test.value, func(data reflect.Value) error {
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
			"empty A", mockStruct{}, &FieldError{"abc", "A", ErrRequired},
		},
		{
			"A not string", mockStruct{124}, &FieldError{"abc", "A", ErrNotString},
		},
		{
			"A min length error", mockStruct{"1234"}, &FieldError{"abc", "A", ErrMinLength(5)},
		},
		{
			"A success", mockStruct{"12345"}, nil,
		},
		{
			"empty ptr A", &mockStruct{}, &FieldError{"abc", "A", ErrRequired},
		},
		{
			"ptr A not string", &mockStruct{124}, &FieldError{"abc", "A", ErrNotString},
		},
		{
			"ptr A min length error", &mockStruct{"1234"}, &FieldError{"abc", "A", ErrMinLength(5)},
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

func TestFieldError_Error(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	t.Parallel()
	for i := 0; i < 1000; i++ {
		t.Run(fmt.Sprint("test ", i), func(t *testing.T) {
			err := &FieldError{
				Key:  randomString(),
				Name: randomString(),
				Err:  errors.New(randomString()),
			}
			assert.Equal(t, err.Error(), fmt.Sprintf("[%v]: %v", err.Name, err.Err.Error()))
		})
	}
}

func randomString() string {
	length := rand.Intn(100)
	b := make([]byte, length)
	rand.Read(b)
	return string(b)
}
