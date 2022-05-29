package validate

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMustBeTime(t *testing.T) {
	t.Parallel()
	now := time.Now()
	tests := []testCase{
		{"nil", nil, ErrTime},
		{"string", "abc", ErrTime},
		{"time", now, nil},
		{"time ptr", &now, nil},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expect, MustBeTime(test.value, func(t time.Time) error {
				return nil
			}))
		})
	}
}

func TestParseTime(t *testing.T) {
	t.Parallel()
	dateString := "1998-12-22T00:00:00+07:00"
	tests := []testCase{
		{
			"no time source", Use(dateString, &TimeValidator{}).Validate(), ErrNoTimeSource,
		},
		{"wrong format", Use(dateString, Date("abc", time.Local)).Validate(), &time.ParseError{
			Layout:     "abc",
			Value:      dateString,
			LayoutElem: "abc",
			ValueElem:  dateString,
			Message:    "",
		}},
		{
			"right format", Use(dateString, Date(time.RFC3339, time.Local)).Validate(), nil,
		},
		{
			"right format default", Use(dateString, Date("", nil)).Validate(), nil,
		},
		{
			"second", Use(time.Now().Unix(), Second()).Validate(), nil,
		},
		{
			"nano second", Use(time.Now().UnixNano(), Nano()).Validate(), nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expect, test.value)
		})
	}
}

func TestBefore(t *testing.T) {
	t.Parallel()
	now := time.Now().Unix()
	nowDate := time.Unix(now, 0)
	fns := Second(Before(nowDate))
	tests := []testCase{
		{"before", now - int64(time.Second), nil},
		{"equal", now, ErrBefore(nowDate)},
		{"after", now + int64(time.Second), ErrBefore(nowDate)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expect, Use(test.value, fns).Validate())
		})
	}
}

func TestAfter(t *testing.T) {
	t.Parallel()
	now := time.Now().Unix()
	nowDate := time.Unix(now, 0)
	fns := Second(After(nowDate))
	tests := []testCase{
		{"before", now - int64(time.Second), ErrAfter(nowDate)},
		{"equal", now, ErrAfter(nowDate)},
		{"after", now + int64(time.Second), nil},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expect, Use(test.value, fns).Validate())
		})
	}
}

func TestChangeTime(t *testing.T) {
	t.Parallel()
	now := time.Now().Unix()
	nowDate := time.Unix(now, 0)
	fns := Second(ChangeTime(func(t time.Time) time.Time {
		return t.Add(time.Second)
	}, After(nowDate)))
	tests := []testCase{
		{"dec 1 sec", now - int64(time.Second), ErrAfter(nowDate)},
		{"no change", now, nil},
		{"add 1 sec", now + int64(time.Second), nil},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expect, Use(test.value, fns).Validate())
		})
	}
}
