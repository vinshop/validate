package validate

import (
	"github.com/stretchr/testify/assert"
	"log"
	"math/rand"
	"testing"
)

func TestMustBeNumber(t *testing.T) {
	t.Parallel()
	tests := []testCase{
		{"string", "abc", ErrNotNumber},
		{"int", int(123), nil},
		{"int64", int64(123), nil},
		{"float32", float32(123), nil},
		{"float64", float64(123), nil},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expect, MustBeNumber(test.value, func(a float64) error {
				log.Println(a)
				return nil
			}))
		})
	}
}

func TestEQ(t *testing.T) {
	t.Parallel()
	num := rand.Float64()
	fns := Number(EQ(num))
	tests := []testCase{
		{"not number", "abc", ErrNotNumber},
		{"less", num - 1, ErrEQ(num)},
		{"more", num + 1, ErrEQ(num)},
		{"equal", num, nil},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expect, Use(test.value, fns).Validate())
		})
	}
}

func TestNEQ(t *testing.T) {
	t.Parallel()
	num := rand.Float64()
	fns := Number(NEQ(num))
	tests := []testCase{
		{"not number", "abc", ErrNotNumber},
		{"less", num - 1, nil},
		{"more", num + 1, nil},
		{"equal", num, ErrNEQ(num)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expect, Use(test.value, fns).Validate())
		})
	}
}

func TestLT(t *testing.T) {
	t.Parallel()
	num := rand.Float64()
	fns := Number(LT(num))
	tests := []testCase{
		{"not number", "abc", ErrNotNumber},
		{"less", num - 1, nil},
		{"more", num + 1, ErrLT(num)},
		{"equal", num, ErrLT(num)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expect, Use(test.value, fns).Validate())
		})
	}
}

func TestLTE(t *testing.T) {
	t.Parallel()
	num := rand.Float64()
	fns := Number(LTE(num))
	tests := []testCase{
		{"not number", "abc", ErrNotNumber},
		{"less", num - 1, nil},
		{"more", num + 1, ErrLTE(num)},
		{"equal", num, nil},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expect, Use(test.value, fns).Validate())
		})
	}
}

func TestGT(t *testing.T) {
	t.Parallel()
	num := rand.Float64()
	fns := Number(GT(num))
	tests := []testCase{
		{"not number", "abc", ErrNotNumber},
		{"less", num - 1, ErrGT(num)},
		{"more", num + 1, nil},
		{"equal", num, ErrGT(num)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expect, Use(test.value, fns).Validate())
		})
	}
}

func TestGTE(t *testing.T) {
	t.Parallel()
	num := rand.Float64()
	fns := Number(GTE(num))
	tests := []testCase{
		{"not number", "abc", ErrNotNumber},
		{"less", num - 1, ErrGTE(num)},
		{"more", num + 1, nil},
		{"equal", num, nil},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expect, Use(test.value, fns).Validate())
		})
	}
}

func TestDoMath(t *testing.T) {
	t.Parallel()
	num := rand.Float64()
	dif := rand.Float64()
	fns := Number(DoMath(func(n float64) float64 {
		return n + dif
	}, EQ(num)))
	tests := []testCase{
		{"not number", "abc", ErrNotNumber},
		{"-dif", num - dif, nil},
		{"+dif", num + dif, ErrEQ(num)},
		{"equal", num, ErrEQ(num)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expect, Use(test.value, fns).Validate())
		})
	}
}

func TestCustomNumber(t *testing.T) {
	t.Parallel()
	num := 10.0
	fns := Number(CustomNumber(func(n float64) error {
		if n != num {
			return ErrEQ(num)
		}
		return nil
	}))
	tests := []testCase{
		{"not number", "abc", ErrNotNumber},
		{"not equal", num + 1, ErrEQ(num)},
		{"equal", num, nil},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expect, Use(test.value, fns).Validate())
		})
	}
}
