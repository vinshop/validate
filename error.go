package validate

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
)

type ErrType string

// includeErrPath, if true (default) the error will return the full path to object triggered the error
var includeErrPath = true

var configMu sync.Mutex

// SetIncludeErrPath set true if you want the error message include the full path
func SetIncludeErrPath(value bool) {
	configMu.Lock()
	defer configMu.Unlock()
	includeErrPath = value
}

const (
	ErrField    ErrType = "field"
	ErrArrIndex ErrType = "arr_idx"
)

// Error common error type
type Error struct {
	Type ErrType
	Key  string
	Name string
	Err  error
}

// FieldError error for Field in Struct validator
func FieldError(key string, field reflect.StructField, err error) error {
	name := field.Name
	if json, ok := field.Tag.Lookup("json"); ok {
		parts := strings.Split(json, ",")
		if len(parts) > 0 {
			name = parts[0]
		}
	}
	return &Error{
		Type: ErrField,
		Key:  key,
		Name: name,
		Err:  err,
	}
}

// ArrayError error for Array
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
	if includeErrPath && path != "" {
		return fmt.Sprintf("%v: %v", path, err.Error())
	}
	return err.Error()
}

// GetLastErrorWithKey return the last Error has Key
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

// GetRootError return the root error
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
