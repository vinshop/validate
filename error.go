package validate

import (
	"fmt"
	"strings"
)

type ErrType string

const (
	ErrField    ErrType = "field"
	ErrArrIndex ErrType = "arr_idx"
)

type Error struct {
	Type ErrType
	Key  string
	Name string
	Err  error
}

func FieldError(key string, name string, err error) error {
	return &Error{
		Type: ErrField,
		Key:  key,
		Name: name,
		Err:  err,
	}
}

func ArrayError(index int, err error) error {
	return &Error{
		Type: ErrArrIndex,
		Key:  "",
		Name: fmt.Sprintf("[%v]", index),
		Err:  err,
	}
}

func (e *Error) Error() string {
	path, err := e.GetRootError()
	if path != "" {
		return fmt.Sprintf("%v: %v", path, err.Error())
	}
	return err.Error()
}

func (e *Error) GetLastErrorWithKey() error {
	var (
		err error
		res error
	)
	if e.Key != "" {
		res = e
	}
	err = e.Err
	for {
		cErr, ok := err.(*Error)
		if !ok {
			break
		}
		err = cErr.Err
		if cErr.Key != "" {
			res = cErr
		}
	}
	return res
}

func (e *Error) GetRootError() (string, error) {
	b := strings.Builder{}
	b.WriteString(e.Name)
	var (
		err error
	)
	err = e.Err

	for {
		cErr, ok := err.(*Error)
		if !ok {
			break
		}
		err = cErr.Err
		switch cErr.Type {
		case ErrArrIndex:
			b.WriteString(cErr.Name)
		case ErrField:
			b.WriteString(".")
			b.WriteString(cErr.Name)
		default:
			b.WriteString(".")
			b.WriteString(cErr.Name)
		}

	}
	return b.String(), err
}
